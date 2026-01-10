
DELETE FROM ZONES;
INSERT INTO ZONES (
    id,
    name ,
    region,
    latitude,
    longitude
    -- metadata
) values
(1,'Herring Cove', 'Southern Shore', 44.577007, -63.550592),
(2,'Chebucto Head', 'Southern Shore', 44.501763, -63.524395),
(3,'Crystal Crescent', 'Southern Shore', 44.459990, -63.622073),
(4,'Prospect', 'Southern Shore', 44.471228, -63.790843),
(5,'West Dover', 'Southern Shore', 44.491659, -63.869845),
(6,'Peggy''s Cove', 'Southern Shore', 44.497615, -63.915399),
(7,'West Pennant', 'Southern Shore', 0,0),
(8,'Terrence Bay', 'Southern Shore', 0,0)
;

DELETE FROM CRAGS;
INSERT INTO CRAGS (
    id,
    zone_id,
    name,
    latitude,
    longitude
    -- metadata JSONB default '{}'
) values
(1, 6,'Land of Confusion', 0,0),
(2, 6,'Dover Island', 0,0),
(3, 6,'Polly''s Cove', 0,0)
;

-- (2,,'West Pennant', 'Southern Shore', 0,0, ''),
-- (3,'Prospect', 'Southern Shore', 0,0, ''),
-- (4,'Chebucto Head', 'Southern Shore', 0,0, ''),
-- (5,'Terrence Bay', 'Southern Shore', 0,0, ''),
-- (6,'Herring Cove', 'Southern Shore', 0,0, '')
-- (7,'Crystal Crescent', 'Southern Shore', 0,0, '')


-- eg corn and Bung
DELETE FROM AREAS;
INSERT INTO AREAS (
    id,
    crag_id,
    name,
    -- description,
    latitude ,
    longitude
    -- metadata JSONB default '{}'
) values
(1, 1,'Corn & Bung', 44.498149, -63.902713),
(2, 1,'Gros Poisson & Jungle', 44.496239, -63.897941),
(3, 1,'The Scoop', 44.494133, -63.892400);



DELETE FROM BOULDERS;
INSERT INTO BOULDERS (
    id,
    area_id,
    name,
    description,
    latitude,
    longitude
    -- metadata,
) values
(1, 1, 'Corn', 'Upper Boulder',  0,0),
(2, 1, 'Bung', 'Lower Boulder', 0,0)
;



-- CREATE TYPE DIRECTION AS ENUM ('north', 'south', 'east', 'west');

DELETE FROM CLIMBS;
INSERT INTO CLIMBS (
    id,
    boulder_id,
    face,
    name,
    -- description,
    grade
    -- line,
    -- metadata,
) values
(1,1,'west','Jackalopes Can Kill','V4'),
(2,1,'west','Jackpot','V0'),
(3,1,'west','Cornholio','V4'),
(4,1,'west','It''s A Disease','V8'),
(5,1,'north','Resurrection','V8'),
(6,1,'north','Carbosaurus','V10'),
(7,1,'north','Insurrection','V10'),
(8,1,'north','The Whales Back','V3'),
(9,1,'north','Can''t Trust Skinny People','V6'),
(10,1,'east','Captain Hook','V5'),
(11,2,'west','Taco Dish','V3'),
(12,2,'west','Overdrive','V10'),
(13,2,'north','Piranha','V6'),
(14,2,'north','Joe Boxer','V6'),
(15,2,'east','Memory Block','V2'),
(16,2,'east','White Cracker','V0'),
(17,2,'east','Razors Edge','V2'),
(18,2,'east','Dancing to the New Bolaro','V9'),
(19,2,'south','Milking The Cow','V7'),
(20,2,'south','Milking The Cow (Stand)','V5');



-- DELETE FROM TAGS CASCADE where id = 0;
DELETE FROM TAGS;
INSERT INTO TAGS (id,name)
values
 (0,'TEST'),
 (1,'SDS'),
 (2,'Sit Start'),
 (3,'Standing Start'),
 (4,'Crimpy'),
 (5,'Slopers'),
 (6,'Juggy'),
 (7,'Pinchy'),
 (8,'Pockets'),
 (9,'Volumes'),
 (10,'Slab'),
 (11,'Vertical'),
 (12,'Overhang'),
 (13,'Roof'),
 (14,'Arete'),
 (15,'Dihedral'),
 (16,'Dynamic'),
 (17,'Static'),
 (18,'Technical'),
 (19,'Powerful'),
 (20,'Compression'),
 (21,'Balancey'),
 (22,'Dyno'),
 (23,'Heel Hook'),
 (24,'Toe Hook'),
 (25,'Knee Bar'),
 (26,'Mantle'),
 (27,'Match'),
 (28,'Highball'),
 (29,'Traverse');

-- tags associated to the specific climb.
DELETE FROM CLIMB_TAGS;
INSERT INTO CLIMB_TAGS (
    id ,
    tag_id ,
    climb_id
) values
(1,0,1),
(2,0,2),
(3,0,3),
(4,0,4),
(5,0,5),
(6,0,6),
(7,0,7),
(8,0,8),
(9,0,9),
(10,0,10),
(11,0,11),
(12,0,12),
(13,0,13),
(14,0,14),
(15,0,15),
(16,0,16),
(17,0,17),
(18,0,18),
(19,0,19),
(20,0,20);
