
DROP TABLE IF EXISTS CLIMB_TAGS CASCADE;
DROP TABLE IF EXISTS TAGS CASCADE;
DROP TABLE IF EXISTS CLIMBS CASCADE;
DROP TABLE IF EXISTS BOULDERS CASCADE;
DROP TABLE IF EXISTS AREAS CASCADE;
DROP TABLE IF EXISTS CRAGS CASCADE;
DROP TABLE IF EXISTS ZONES CASCADE;
DROP TYPE IF EXISTS DIRECTION;


-- eg peggy's cove
CREATE TABLE ZONES (
    id SERIAL PRIMARY KEY,
    name TEXT,
    region TEXT,
    description TEXT,
    latitude DECIMAL(9,6) DEFAULT 0,
    longitude DECIMAL(9,6) DEFAULT 0,
    metadata JSONB
);

-- eg Land of Confusion
CREATE TABLE CRAGS (
    id SERIAL PRIMARY KEY,
    zone_id INTEGER REFERENCES zones(id) ON DELETE CASCADE,
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    latitude DECIMAL(9,6) DEFAULT 0,
    longitude DECIMAL(9,6) DEFAULT 0,
    metadata JSONB default '{}'
);

-- eg corn and Bung
CREATE TABLE AREAS (
    id SERIAL PRIMARY KEY,
    crag_id INTEGER REFERENCES crags(id) ON DELETE CASCADE,
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    latitude DECIMAL(9,6) DEFAULT 0,
    longitude DECIMAL(9,6) DEFAULT 0,
    metadata JSONB default '{}'
);

-- Upper Boulder West
CREATE TABLE BOULDERS (
    id SERIAL PRIMARY KEY,
    area_id INTEGER REFERENCES areas(id) ON DELETE CASCADE,
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    latitude DECIMAL(9,6) DEFAULT 0,
    longitude DECIMAL(9,6) DEFAULT 0,
    metadata JSONB default '{}'
);

CREATE TYPE DIRECTION AS ENUM ('north', 'south', 'east', 'west');

CREATE TABLE CLIMBS (
    id SERIAL PRIMARY KEY,
    boulder_id INTEGER REFERENCES boulders(id) ON DELETE CASCADE,
    face DIRECTION,
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    grade TEXT DEFAULT 'V?', -- e.g., 'V5' or '7A'
    line JSONB DEFAULT '[]',
    metadata JSONB default '{}'
);

CREATE TABLE TAGS (
    id SERIAL PRIMARY KEY,
    name TEXT
);

-- tags associated to the specific climb.
CREATE TABLE CLIMB_TAGS (
    id SERIAL PRIMARY KEY,
    tag_id INTEGER REFERENCES TAGS(id) ON DELETE CASCADE,
    climb_id INTEGER REFERENCES CLIMBS(id) ON DELETE CASCADE
);


-- Index the hierarchy foreign keys for fast joins
CREATE INDEX idx_crags_zone_id ON crags(zone_id);
CREATE INDEX idx_areas_crag_id ON areas(crag_id);
CREATE INDEX idx_boulders_area_id ON boulders(area_id);
CREATE INDEX idx_climbs_boulder_id ON climbs(boulder_id);
CREATE INDEX idx_climbs_tags_tag_id ON climb_tags(tag_id);
CREATE INDEX idx_climbs_tags_climb_id ON climb_tags(climb_id);
CREATE INDEX idx_climbs_boulder_grade ON climbs(grade);


CREATE INDEX idx_climbs_metadata ON climbs USING GIN (metadata);
CREATE INDEX idx_crags_metadata ON crags USING GIN (metadata);