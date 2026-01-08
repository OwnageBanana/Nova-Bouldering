
CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    region TEXT,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    metadata JSONB,
);

CREATE TABLE areas (
    id SERIAL PRIMARY KEY,
    location_id INTEGER REFERENCES locations(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,0
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
);

CREATE TABLE boulders (
    id SERIAL PRIMARY KEY,
    area_id INTEGER REFERENCES areas(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6)
);

CREATE TABLE climbs (
    id SERIAL PRIMARY KEY,
    boulder_id INTEGER REFERENCES boulders(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    grade TEXT NOT NULL, -- e.g., 'V5' or '7A'
    description TEXT,
    tags TEXT[] DEFAULT '{}' -- The tagging system
);

-- Create the index for fast tag lookups
CREATE INDEX idx_climbs_tags ON climbs USING GIN (tags);

-- Index the hierarchy foreign keys for fast joins
CREATE INDEX idx_areas_location_id ON areas(location_id);
CREATE INDEX idx_boulders_area_id ON boulders(area_id);
CREATE INDEX idx_climbs_boulder_id ON climbs(boulder_id);


CREATE INDEX idx_climbs_metadata ON climbs USING GIN (metadata);