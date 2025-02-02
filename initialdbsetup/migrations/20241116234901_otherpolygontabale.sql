-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
 		CREATE TABLE "other_polygon_structure" (gid serial,
"shape__are" numeric,
"shape__len" numeric,
"globalid" varchar(38),
"creationda" date,
"creator" varchar(128),
"editdate" date,
"editor" varchar(128),
"exact_use" varchar(254),
"usage1" varchar(254),
"developmen" varchar(254),
"structure_" varchar(254),
"ghanapostg" varchar(254),
"street_nam" varchar(254),
"name" varchar(254),
"parcel_id" varchar(254),
"remarks" varchar(254),
"other_info" varchar(254),
"other_in_1" varchar(254),
"mixed_usag" varchar(254));
ALTER TABLE "other_polygon_structure" ADD PRIMARY KEY (gid);
SELECT AddGeometryColumn('','other_polygon_structure','geom','0','MULTIPOLYGON',4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE other_polygon_structure
-- +goose StatementEnd
