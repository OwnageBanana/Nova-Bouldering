
DROP TABLE IF EXISTS ZONE_IMAGES CASCADE;
DROP TABLE IF EXISTS CRAG_IMAGES CASCADE;
DROP TABLE IF EXISTS AREA_IMAGES CASCADE;
DROP TABLE IF EXISTS BOULDER_IMAGES CASCADE;
DROP TABLE IF EXISTS CLIMB_IMAGES CASCADE;
DROP TABLE IF EXISTS IMAGES CASCADE;
DROP TABLE IF EXISTS CLIMB_TAGS CASCADE;
DROP TABLE IF EXISTS TAGS CASCADE;
DROP TABLE IF EXISTS CLIMBS CASCADE;
DROP TABLE IF EXISTS BOULDERS CASCADE;
DROP TABLE IF EXISTS AREAS CASCADE;
DROP TABLE IF EXISTS CRAGS CASCADE;
DROP TABLE IF EXISTS ZONES CASCADE;
DROP TYPE IF EXISTS DIRECTION;


-- Images stored in R2
CREATE TABLE IMAGES (
    id SERIAL PRIMARY KEY,
    r2_key TEXT NOT NULL,       -- R2 object key/path
    filename TEXT DEFAULT '',
    content_type TEXT DEFAULT '',
    size_bytes INTEGER DEFAULT 0,
    alt_text TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    metadata JSONB NOT NULL DEFAULT '{}'
);


-- eg peggy's cove
CREATE TABLE ZONES (
    id SERIAL PRIMARY KEY,
    image_id INTEGER REFERENCES images(id) ON DELETE SET NULL,
    name TEXT DEFAULT '',
    region TEXT DEFAULT '',
    description TEXT DEFAULT '',
    latitude DECIMAL(9,6) DEFAULT 0,
    longitude DECIMAL(9,6) DEFAULT 0,
    metadata JSONB NOT NULL default '{}'
);

-- eg Land of Confusion
CREATE TABLE CRAGS (
    id SERIAL PRIMARY KEY,
    zone_id INTEGER REFERENCES zones(id) ON DELETE CASCADE,
    image_id INTEGER REFERENCES images(id) ON DELETE SET NULL,
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    latitude DECIMAL(9,6) DEFAULT 0,
    longitude DECIMAL(9,6) DEFAULT 0,
    metadata JSONB NOT NULL default '{}'
);

-- eg corn and Bung
CREATE TABLE AREAS (
    id SERIAL PRIMARY KEY,
    crag_id INTEGER REFERENCES crags(id) ON DELETE CASCADE,
    image_id INTEGER REFERENCES images(id) ON DELETE SET NULL,
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    latitude DECIMAL(9,6) DEFAULT 0,
    longitude DECIMAL(9,6) DEFAULT 0,
    metadata JSONB NOT NULL default '{}'
);

-- Upper Boulder
CREATE TABLE BOULDERS (
    id SERIAL PRIMARY KEY,
    area_id INTEGER REFERENCES areas(id) ON DELETE CASCADE,
    image_id INTEGER REFERENCES images(id) ON DELETE SET NULL,
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    latitude DECIMAL(9,6) DEFAULT 0,
    longitude DECIMAL(9,6) DEFAULT 0,
    metadata JSONB NOT NULL default '{}'
);

CREATE TYPE DIRECTION AS ENUM ('north', 'south', 'east', 'west', '');

CREATE TABLE CLIMBS (
    id SERIAL PRIMARY KEY,
    boulder_id INTEGER REFERENCES boulders(id) ON DELETE CASCADE,
    image_id INTEGER REFERENCES images(id) ON DELETE SET NULL,
    face DIRECTION DEFAULT '',
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    grade TEXT DEFAULT 'V?', -- e.g., 'V5' or '7A'
    line JSONB NOT NULL DEFAULT '{}',
    metadata JSONB NOT NULL default '{}'
);

CREATE TABLE TAGS (
    id SERIAL PRIMARY KEY,
    name TEXT default ''
);

-- tags associated to the specific climb.
CREATE TABLE CLIMB_TAGS (
    id SERIAL PRIMARY KEY,
    tag_id INTEGER REFERENCES TAGS(id) ON DELETE CASCADE,
    climb_id INTEGER REFERENCES CLIMBS(id) ON DELETE CASCADE
);

-- Join tables for many-to-one image relationships
CREATE TABLE ZONE_IMAGES (
    id SERIAL PRIMARY KEY,
    zone_id INTEGER NOT NULL REFERENCES zones(id) ON DELETE CASCADE,
    image_id INTEGER NOT NULL REFERENCES images(id) ON DELETE CASCADE,
    sort_order INTEGER DEFAULT 0,
    UNIQUE(zone_id, image_id)
);

CREATE TABLE CRAG_IMAGES (
    id SERIAL PRIMARY KEY,
    crag_id INTEGER NOT NULL REFERENCES crags(id) ON DELETE CASCADE,
    image_id INTEGER NOT NULL REFERENCES images(id) ON DELETE CASCADE,
    sort_order INTEGER DEFAULT 0,
    UNIQUE(crag_id, image_id)
);

CREATE TABLE AREA_IMAGES (
    id SERIAL PRIMARY KEY,
    area_id INTEGER NOT NULL REFERENCES areas(id) ON DELETE CASCADE,
    image_id INTEGER NOT NULL REFERENCES images(id) ON DELETE CASCADE,
    sort_order INTEGER DEFAULT 0,
    UNIQUE(area_id, image_id)
);

CREATE TABLE BOULDER_IMAGES (
    id SERIAL PRIMARY KEY,
    boulder_id INTEGER NOT NULL REFERENCES boulders(id) ON DELETE CASCADE,
    image_id INTEGER NOT NULL REFERENCES images(id) ON DELETE CASCADE,
    sort_order INTEGER DEFAULT 0,
    UNIQUE(boulder_id, image_id)
);

CREATE TABLE CLIMB_IMAGES (
    id SERIAL PRIMARY KEY,
    climb_id INTEGER NOT NULL REFERENCES climbs(id) ON DELETE CASCADE,
    image_id INTEGER NOT NULL REFERENCES images(id) ON DELETE CASCADE,
    sort_order INTEGER DEFAULT 0,
    UNIQUE(climb_id, image_id)
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

-- Image indexes
CREATE UNIQUE INDEX idx_images_r2_key ON images(r2_key);
CREATE INDEX idx_zone_images_zone_id ON zone_images(zone_id);
CREATE INDEX idx_crag_images_crag_id ON crag_images(crag_id);
CREATE INDEX idx_area_images_area_id ON area_images(area_id);
CREATE INDEX idx_boulder_images_boulder_id ON boulder_images(boulder_id);
CREATE INDEX idx_climb_images_climb_id ON climb_images(climb_id);
