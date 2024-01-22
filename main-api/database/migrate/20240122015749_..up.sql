CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "pictures" (
  "id" serial PRIMARY KEY,
  "description" varchar,
  "user_id" integer, -- NOT NULL,
  "url" varchar NOT NULL,
  "matches" integer NOT NULL,
  "rating" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "results" (
  "id" serial PRIMARY KEY,
  "picture_a" integer NOT NULL,
  "picture_b" integer NOT NULL,
  "complete" boolean NOT NULL DEFAULT false,
  "winner" integer,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "pictures" ("user_id");

CREATE INDEX ON "pictures" ("rating");

CREATE INDEX ON "results" ("winner");

CREATE INDEX ON "results" ("picture_a");

CREATE INDEX ON "results" ("picture_b");

CREATE INDEX ON "results" ("picture_a", "picture_b");

ALTER TABLE "pictures" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "results" ADD FOREIGN KEY ("picture_a") REFERENCES "pictures" ("id");

ALTER TABLE "results" ADD FOREIGN KEY ("picture_b") REFERENCES "pictures" ("id");

ALTER TABLE "results" ADD FOREIGN KEY ("winner") REFERENCES "pictures" ("id");
