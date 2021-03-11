CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar(32) UNIQUE NOT NULL,
  "name" varchar(32) NOT NULL
);

CREATE TABLE "api_keys" (
  "id" SERIAL PRIMARY KEY,
  "key" varchar(32) NOT NULL,
  "owner" int NOT NULL,
  "enabled" boolean DEFAULT 'true',
  "hits" int DEFAULT 0,
  "errors" int DEFAULT 0
);

CREATE TABLE "storage" (
  "id" SERIAL PRIMARY KEY,
  "key" varchar(32) UNIQUE NOT NULL,
  "value" varchar NOT NULL DEFAULT 'empty',
  "available" boolean DEFAULT 'true',
  "by" int NOT NULL,
  "downloads" int DEFAULT 0,
  "errors" int DEFAULT 0
);

ALTER TABLE "api_keys" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "storage" ADD FOREIGN KEY ("by") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

CREATE INDEX "user name" ON "users" ("username");

CREATE INDEX "id for users" ON "users" ("id");

CREATE INDEX "api key" ON "api_keys" ("key");

CREATE INDEX "key owner" ON "api_keys" ("owner");

CREATE INDEX "by key" ON "storage" ("key");

CREATE INDEX "by creator" ON "storage" ("by");

COMMENT ON COLUMN "users"."username" IS 'username';

COMMENT ON COLUMN "users"."name" IS 'Name of the User';
