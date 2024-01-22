CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "pictures" (
  "id" serial PRIMARY KEY,
  "description" varchar,
  "user_id" integer NOT NULL,
  "url" varchar NOT NULL,
  "matches" integer NOT NULL,
  "rating" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "results" (
  "id" serial PRIMARY KEY,
  "cat_a" integer NOT NULL,
  "cat_b" integer NOT NULL,
  "winner" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "pictures" ("user_id");

CREATE INDEX ON "pictures" ("rating");

CREATE INDEX ON "results" ("winner");

CREATE INDEX ON "results" ("cat_a");

CREATE INDEX ON "results" ("cat_b");

CREATE INDEX ON "results" ("cat_a", "cat_b");

ALTER TABLE "pictures" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "results" ADD FOREIGN KEY ("cat_a") REFERENCES "pictures" ("id");

ALTER TABLE "results" ADD FOREIGN KEY ("cat_b") REFERENCES "pictures" ("id");

ALTER TABLE "results" ADD FOREIGN KEY ("winner") REFERENCES "pictures" ("id");
