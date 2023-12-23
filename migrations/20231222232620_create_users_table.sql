-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "google_id" bigint NULL,
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
