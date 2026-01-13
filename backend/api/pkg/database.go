package pkg

import (
	// "database/sql"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type Zone struct {
	Id          int32           `json:"id"`
	Name        string          `json:"name"`
	Region      string          `json:"region"`
	Description string          `json:"description"`
	Latitude    decimal.Decimal `json:"latitude"`
	Longitude   decimal.Decimal `json:"longitude"`
	Metadata    map[string]any  `json:"metadata"`
}

// -- eg Land of Confusion
type Crag struct {
	Id          int32           `json:"id"`
	ZoneId      int32           `json:"zone_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Latitude    decimal.Decimal `json:"latitude"`
	Longitude   decimal.Decimal `json:"longitude"`
	Metadata    map[string]any  `json:"metadata"`
}

// -- eg corn and Bung
type Area struct {
	Id          int32           `json:"id"`
	CragId      int32           `json:"crag_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Latitude    decimal.Decimal `json:"latitude"`
	Longitude   decimal.Decimal `json:"longitude"`
	Metadata    map[string]any  `json:"metadata"`
}

// -- Upper Boulder West
type Boulder struct {
	Id          int32           `json:"id"`
	AreaId      int32           `json:"area_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Latitude    decimal.Decimal `json:"latitude"`
	Longitude   decimal.Decimal `json:"longitude"`
	Metadata    map[string]any  `json:"metadata"`
}

// CREATE TYPE DIRECTION AS ENUM ('north', 'south', 'east', 'west');
type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
	None  Direction = ""
)

func (d Direction) IsValid() bool {
	switch d {
	case North, South, East, West, None:
		return true
	}
	return false
}

// Scan implements the sql.Scanner interface to read from the DB
func (d *Direction) Scan(value any) error {
	if value == nil {
		*d = None
		return nil
	}

	switch v := value.(type) {
	case []byte:
		*d = Direction(v)
	case string:
		*d = Direction(v)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into Direction", value)
	}

	if !d.IsValid() {
		return fmt.Errorf("invalid Direction value: %s", string(*d))
	}

	return nil
}

// Value implements the driver.Valuer interface to write to the DB
func (d Direction) Value() (driver.Value, error) {
	return string(d), nil
}

type LineItem struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type Route struct {
	Id    int32      `json:"id"`
	Lines []LineItem `json:"lines"`
}

type Climb struct {
	Id          int32          `json:"id"`
	BoulderId   int32          `json:"boulder_id"`
	Face        Direction      `json:"face"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Grade       string         `json:"grade"`
	Line        Route          `json:"line"`
	Metadata    map[string]any `json:"metadata"`
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
	Id      int32 `json:"Id"`
	TagId   int32 `json:"TagId"`
	ClimbId int32 `json:"ClimbId"`
}

func GetClimb(ctx context.Context, db *pgxpool.Pool, id int32) (*Climb, error) {
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
		FROM climbs where id = $1`

	var c Climb
	var lineBytes, metaBytes []byte
	err := db.QueryRow(ctx, query, id).Scan(
		&c.Id,
		&c.BoulderId,
		&c.Face,
		&c.Name,
		&c.Description,
		&c.Grade,
		&lineBytes,
		&metaBytes,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning climb: %w", err)
	}

	if len(lineBytes) > 0 {
		if err := json.Unmarshal(lineBytes, &c.Line); err != nil {
			log.Printf("failed marshalling Lines: %v \n bytes: %#v", err, lineBytes)
			return nil, err
		}
	}
	if len(metaBytes) > 0 {
		if err := json.Unmarshal(metaBytes, &c.Metadata); err != nil {
			log.Printf("failed marshalling metadata: %v \n bytes: %#v", err, lineBytes)
			return nil, err
		}
	}

	return &c, nil
}

func GetAllClimbs(ctx context.Context, db *pgxpool.Pool) ([]*Climb, error) {
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

	var climbs []*Climb

	// 2. Iterate through rows
	for rows.Next() {
		var c Climb
		var lineBytes []byte // Temporary holder for the JSONB raw bytes
		var metaBytes []byte // Temporary holder for the JSONB raw bytes

		// 3. Scan into variables (use pointers for basic types, byte slices for JSON)
		err := rows.Scan(
			&c.Id,
			&c.BoulderId,
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

		climbs = append(climbs, &c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return climbs, nil
}

func CreateClimb(ctx context.Context, db *pgxpool.Pool, c *Climb) error {
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
        INSERT INTO climbs (boulder_id, face, name, description, grade, line, metadata)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id`

	// QueryRow lets us scan the new ID back into the pointer
	err = db.QueryRow(ctx, query,
		c.BoulderId, c.Face, c.Name, c.Description, c.Grade, lineJSON, metadataJSON,
	).Scan(&c.Id)

	return err
}

func DeleteClimb(ctx context.Context, db *pgxpool.Pool, id int32) error {

	// 2. Define the update query
	query := `
		DELETE from climbs
		WHERE id = $1`

	// 3. Execute the query
	result, err := db.Exec(ctx, query,
		id,
	)

	if err != nil {
		return fmt.Errorf("error executing update: %w", err)
	}
	// 4. Verify a row was actually updated
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no climb found with id %d", id)
	}
	return err
}

// UpdateClimb updates an existing climb record based on the struct's ID.
func UpdateClimb(ctx context.Context, db *pgxpool.Pool, c *Climb) error {
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
		c.BoulderId,
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

// Wrapper structs to hold the hierarchy
type ZoneNode struct {
	Zone
	Crags []*CragNode `json:"crags"`
}

type CragNode struct {
	Crag
	Areas []*AreaNode `json:"areas"`
}

type AreaNode struct {
	Area
	Boulders []*BoulderNode `json:"boulders"`
}

type BoulderNode struct {
	Boulder
	Climbs []Climb `json:"climbs"`
	// Climbs are the leaves, so they use the base struct
}

func GetFullHierarchy(ctx context.Context, db *pgxpool.Pool) ([]*ZoneNode, error) {
	// 1. Join all tables from top (Zones) to bottom (Climbs)
	query := `
SELECT
        -- Zone (Always present, no Coalesce needed)
        z.id, z.name, z.region, z.description, z.latitude, z.longitude, z.metadata,

        -- Crag (Left Join -> Might be NULL)
        COALESCE(c.id, 0) as id,
        COALESCE(c.zone_id, 0) as zone_id,
        COALESCE(c.name, '') as name,
        COALESCE(c.description, '') as description,
        COALESCE(c.latitude, 0) as latitude,
        COALESCE(c.longitude, 0) as longitude,
        c.metadata, -- Metadata handles NULLs automatically in Go (reads as nil byte slice)

        -- Area
        COALESCE(a.id, 0) as id,
        COALESCE(a.crag_id, 0) as crag_id,
        COALESCE(a.name, '') as name,
        COALESCE(a.description, '') as description,
        COALESCE(a.latitude, 0) as latitude,
        COALESCE(a.longitude, 0) as longitude,
        a.metadata,

        -- Boulder
        COALESCE(b.id, 0) as id,
        COALESCE(b.area_id, 0) as area_id,
        COALESCE(b.name, '') as name,
        COALESCE(b.description, '') as description,
        COALESCE(b.latitude, 0) as latitude,
        COALESCE(b.longitude, 0) as longitude,
        b.metadata,

        -- Climb
        COALESCE(cl.id, 0) as id,
        COALESCE(cl.boulder_id, 0) as boulder_id,
        COALESCE(cl.face, '') as face,
        COALESCE(cl.name, '') as name,
        COALESCE(cl.description, '') as description,
        COALESCE(cl.grade, '') as grade,
        cl.line,
        cl.metadata
    FROM zones z
    LEFT JOIN crags c ON c.zone_id = z.id
    LEFT JOIN areas a ON a.crag_id = c.id
    LEFT JOIN boulders b ON b.area_id = a.id
    LEFT JOIN climbs cl ON cl.boulder_id = b.id
    ORDER BY z.id, c.id, a.id, b.id, cl.id;
	`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	// 2. Maps to keep track of unique pointers to avoid duplication
	zoneMap := make(map[int32]*ZoneNode)
	cragMap := make(map[int32]*CragNode)
	areaMap := make(map[int32]*AreaNode)
	boulderMap := make(map[int32]*BoulderNode)

	// Result slice (roots)
	var rootZones []*ZoneNode

	for rows.Next() {
		// Temporary holders for nullable IDs (to handle LEFT JOIN nulls)
		var zID int32
		var cID, aID, bID, clID sql.NullInt32

		// Temporary structs to scan data into
		var z Zone
		var c Crag
		var a Area
		var b Boulder
		var cl Climb

		// Byte slices for JSON fields
		var zMeta, cMeta, aMeta, bMeta, clMeta, clLine []byte

		// 3. Scan all columns
		// Note: We scan IDs into temp variables first to check for NULLs
		err := rows.Scan(
			// Zone
			&zID, &z.Name, &z.Region, &z.Description, &z.Latitude, &z.Longitude, &zMeta,
			// Crag
			&cID, &c.ZoneId, &c.Name, &c.Description, &c.Latitude, &c.Longitude, &cMeta,
			// Area
			&aID, &a.CragId, &a.Name, &a.Description, &a.Latitude, &a.Longitude, &aMeta,
			// Boulder
			&bID, &b.AreaId, &b.Name, &b.Description, &b.Latitude, &b.Longitude, &bMeta,
			// Climb
			&clID, &cl.BoulderId, &cl.Face, &cl.Name, &cl.Description, &cl.Grade, &clLine, &clMeta,
		)
		if err != nil {
			return nil, err
		}

		// 4. Process Zone
		z.Id = zID
		if len(zMeta) > 0 {
			_ = json.Unmarshal(zMeta, &z.Metadata)
		}

		zoneNode, exists := zoneMap[z.Id]
		if !exists {
			zoneNode = &ZoneNode{Zone: z, Crags: []*CragNode{}}
			zoneMap[z.Id] = zoneNode
			rootZones = append(rootZones, zoneNode)
		}

		// 5. Process Crag (Check if exists, because LEFT JOIN might return NULL)
		if !cID.Valid {
			continue
		}
		c.Id = int32(cID.Int32)
		if len(cMeta) > 0 {
			_ = json.Unmarshal(cMeta, &c.Metadata)
		}

		if c.Id == 0 {
			continue
		}
		cragNode, exists := cragMap[c.Id]
		if !exists {
			cragNode = &CragNode{Crag: c, Areas: []*AreaNode{}}
			cragMap[c.Id] = cragNode
			// Link to parent
			zoneNode.Crags = append(zoneNode.Crags, cragNode)
		}

		// 6. Process Area
		if !aID.Valid {
			continue
		}
		a.Id = int32(aID.Int32)
		if len(aMeta) > 0 {
			_ = json.Unmarshal(aMeta, &a.Metadata)
		}
		if a.Id == 0 {
			continue
		}
		areaNode, exists := areaMap[a.Id]
		if !exists {
			areaNode = &AreaNode{Area: a, Boulders: []*BoulderNode{}}
			areaMap[a.Id] = areaNode
			// Link to parent
			cragNode.Areas = append(cragNode.Areas, areaNode)
		}

		// 7. Process Boulder
		if !bID.Valid {
			continue
		}
		b.Id = int32(bID.Int32)
		if len(bMeta) > 0 {
			_ = json.Unmarshal(bMeta, &b.Metadata)
		}
		if b.Id == 0 {
			continue
		}
		boulderNode, exists := boulderMap[b.Id]
		if !exists {
			boulderNode = &BoulderNode{Boulder: b, Climbs: []Climb{}}
			boulderMap[b.Id] = boulderNode
			// Link to parent
			areaNode.Boulders = append(areaNode.Boulders, boulderNode)
		}

		// 8. Process Climb
		if !clID.Valid {
			continue
		}
		cl.Id = int32(clID.Int32)
		if len(clMeta) > 0 {
			_ = json.Unmarshal(clMeta, &cl.Metadata)
		}
		if len(clLine) > 0 {
			_ = json.Unmarshal(clLine, &cl.Line)
		}

		// Add climb to boulder (Climbs are leaves, no map needed unless deduping specific climbs)
		boulderNode.Climbs = append(boulderNode.Climbs, cl)
	}

	return rootZones, nil
}
