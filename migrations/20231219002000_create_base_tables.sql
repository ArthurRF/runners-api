-- Create "events" table
CREATE TABLE "public"."events" (
	"id" bigserial NOT NULL,
	"name" text NOT NULL,
	"description" text NULL,
	"runners" bigint NULL DEFAULT 0,
	"created_at" timestamptz NULL,
	"updated_at" timestamptz NULL,
	"deleted_at" timestamptz NULL,
	PRIMARY KEY ("id")
);

-- Create index "idx_events_deleted_at" to table: "events"
CREATE INDEX "idx_events_deleted_at" ON "public"."events" ("deleted_at");

-- Create "states" table
CREATE TABLE "public"."states" (
	"id" bigserial NOT NULL,
	"uf" text NOT NULL,
	"name" text NOT NULL,
	"ibge" bigint NOT NULL,
	"ddd" text NULL,
	PRIMARY KEY ("id")
);

-- Create index "states_name_key" to table: "states"
CREATE UNIQUE INDEX "states_name_key" ON "public"."states" ("name");

-- Create index "states_uf_key" to table: "states"
CREATE UNIQUE INDEX "states_uf_key" ON "public"."states" ("uf");

-- Create "cities" table
CREATE TABLE "public"."cities" (
	"id" bigserial NOT NULL,
	"name" text NOT NULL,
	"state_id" bigint NOT NULL,
	"ibge" text NOT NULL,
	"lat_lon" text NULL,
	"cod_tom" bigint NULL,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_cities_state" FOREIGN KEY ("state_id") REFERENCES "public"."states" ("id") ON UPDATE CASCADE ON DELETE
	SET
		NULL
);

-- Create "users" table
CREATE TABLE "public"."users" (
	"id" uuid NOT NULL DEFAULT gen_random_uuid(),
	"google_id" text NULL,
	"name" text NOT NULL,
	"email" text NULL,
	"avatar_url" text NULL,
	"created_at" timestamptz NULL,
	"updated_at" timestamptz NULL,
	"deleted_at" timestamptz NULL,
	PRIMARY KEY ("id")
);

-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");

-- Create index "users_google_id_key" to table: "users"
CREATE UNIQUE INDEX "users_google_id_key" ON "public"."users" ("google_id");