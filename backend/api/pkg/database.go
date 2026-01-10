package database

import (
	// "database/sql"
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type Zone struct {
	Id          int32           `json:'id'`
	Name        string          `json:'name'`
	Region      string          `json:'region'`
	Description string          `json:'description'`
	Latitude    decimal.Decimal `json:'latitude'`
	Longitude   decimal.Decimal `json:'longitude'`
	Metadata    map[string]any  `json:'metadata'`
}

// -- eg Land of Confusion
type Crag struct {
	Id          int32
	Zone_id     int32
	Name        string
	Description string
	Latitude    decimal.Decimal
	Longitude   decimal.Decimal
	Metadata    map[string]any
}

// -- eg corn and Bung
type Area struct {
	Id          int32
	Crag_id     int32
	Name        string
	Description string
	Latitude    decimal.Decimal
	Longitude   decimal.Decimal
	Metadata    map[string]any
}

// -- Upper Boulder West
type Boulder struct {
	Id          int32
	Area_id     int32
	Name        string
	Description string
	Latitude    decimal.Decimal
	Longitude   decimal.Decimal
	Metadata    map[string]any
}

// CREATE TYPE DIRECTION AS ENUM ('north', 'south', 'east', 'west');
type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
)

func (d Direction) IsValid() bool {
	switch d {
	case North, South, East, West:
		return true
	}
	return false
}

// Scan implements the sql.Scanner interface to read from the DB
func (d *Direction) Scan(value any) error {
	if value == nil {
		*d = ""
		return nil
	}

	val, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into Direction", value)
	}

	*d = Direction(val)
	return nil
}

// Value implements the driver.Valuer interface to write to the DB
func (d Direction) Value() (driver.Value, error) {
	return string(d), nil
}

type LineItem struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Route struct {
	ID    int
	Lines []LineItem // This maps to JSONB DEFAULT '[]'
}

type Climb struct {
	Id          int32
	Boulder_id  int32
	Face        Direction
	Name        string
	Description string
	Grade       string
	Line        Route
	Metadata    map[string]any
}

type LineItems []LineItem

// Value converts our Go slice to JSON for Postgres
func (l LineItems) Value() (driver.Value, error) {
	if l == nil {
		return json.Marshal([]LineItem{}) // Ensures '[]' instead of 'null'
	}
	return json.Marshal(l)
}

// Scan converts Postgres JSONB to Go slice
func (l *LineItems) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, l)
}

type Tag struct {
	Id   int32
	Name string
}

// -- tags associated to the specific climb.
type ClimbingTag struct {
	Id       int32
	Tag_id   int32
	Climb_id int32
}

// GetAllClimbs retrieves all records from the climbs table.
func GetAllClimbs(ctx context.Context, db *pgxpool.Pool) ([]Climb, error) {
	// 1. Define the query
	query := `
		SELECT
			id,
			boulder_id,
			face,
			name,
			description,
			grade,
			line,
			metadata
		FROM climbs`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying climbs: %w", err)
	}
	defer rows.Close()

	var climbs []Climb

	// 2. Iterate through rows
	for rows.Next() {
		var c Climb
		var lineBytes []byte // Temporary holder for the JSONB raw bytes
		var metaBytes []byte // Temporary holder for the JSONB raw bytes

		// 3. Scan into variables (use pointers for basic types, byte slices for JSON)
		err := rows.Scan(
			&c.Id,
			&c.Boulder_id,
			&c.Face,
			&c.Name,
			&c.Description,
			&c.Grade,
			&lineBytes,
			&metaBytes,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning climb row: %w", err)
		}

		// 4. Unmarshal the JSON bytes into the struct fields
		if lineBytes != nil {
			if err := json.Unmarshal(lineBytes, &c.Line); err != nil {
				return nil, fmt.Errorf("error unmarshaling line json for climb %d: %w", c.Id, err)
			}
		}

		if metaBytes != nil {
			if err := json.Unmarshal(metaBytes, &c.Metadata); err != nil {
				return nil, fmt.Errorf("error unmarshaling metadata json for climb %d: %w", c.Id, err)
			}
		}

		climbs = append(climbs, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return climbs, nil
}

// UpdateClimb updates an existing climb record based on the struct's ID.
func UpdateClimb(ctx context.Context, db *pgxpool.Pool, c Climb) error {
	// 1. Marshal complex types to JSON
	lineJSON, err := json.Marshal(c.Line)
	if err != nil {
		return fmt.Errorf("error marshaling line: %w", err)
	}

	metadataJSON, err := json.Marshal(c.Metadata)
	if err != nil {
		return fmt.Errorf("error marshaling metadata: %w", err)
	}

	// 2. Define the update query
	query := `
		UPDATE climbs
		SET
			boulder_id = $1,
			face = $2,
			name = $3,
			description = $4,
			grade = $5,
			line = $6,
			metadata = $7
		WHERE id = $8`

	// 3. Execute the query
	result, err := db.Exec(ctx, query,
		c.Boulder_id,
		c.Face,
		c.Name,
		c.Description,
		c.Grade,
		lineJSON,
		metadataJSON,
		c.Id,
	)
	if err != nil {
		return fmt.Errorf("error executing update: %w", err)
	}

	// 4. Verify a row was actually updated
	rowsAffected := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no climb found with id %d", c.Id)
	}

	return nil
}

func GetAllTags(ctx context.Context, db *pgxpool.Pool) (tags []Tag, err error) {

	query := "select * from tags"

	res, err := db.Query(ctx, query)
	if err != nil {
		return tags, fmt.Errorf("error querying results: %w", err)
	}
	for res.Next() {
		var tag Tag
		err = res.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return tags, fmt.Errorf("error scanning rows: %w", err)
		}
		tags = append(tags, tag)
	}
	return
}
