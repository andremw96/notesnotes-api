CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar,
  "first_name" varchar,
  "last_name" varchar,
  "username" varchar,
  "email" varchar,
  "password" varchar,
  "created_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "notes" (
  "id" SERIAL,
  "user_id" int,
  "title" varchar,
  "description" varchar,
  "created_at" TIMESTAMP DEFAULT (now()),
  PRIMARY KEY ("id", "user_id")
);

CREATE UNIQUE INDEX ON "notes" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
