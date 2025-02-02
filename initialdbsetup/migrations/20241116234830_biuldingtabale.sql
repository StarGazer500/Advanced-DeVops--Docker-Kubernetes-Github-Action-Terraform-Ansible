-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
	    CREATE TABLE "building" (gid serial,
"shape__are" numeric,
"shape__len" numeric,
"globalid" varchar(38),
"creationda" date,
"creator" varchar(128),
"editdate" date,
"editor" varchar(128),
"num_storey" float8,
"building_t" varchar(254),
"ghanapost_" varchar(254),
"plot_numbe" varchar(254),
"developmen" varchar(254),
"name" varchar(254),
"parcel_id" varchar(254),
"exact_use" varchar(254),
"building_u" varchar(254),
"mixed_use" varchar(254),
"remarks" varchar(254),
"other_info" varchar(254),
"other_in_1" varchar(254));
ALTER TABLE "building" ADD PRIMARY KEY (gid);
SELECT AddGeometryColumn('','building','geom','0','MULTIPOLYGON',4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE building

-- +goose StatementEnd
