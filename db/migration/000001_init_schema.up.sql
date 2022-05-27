CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" TIMESTAMP DEFAULT (now()) NOT NULL
);

CREATE TABLE "notes" (
  "id" SERIAL,
  "user_id" int NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar,
  "created_at" TIMESTAMP DEFAULT (now()) NOT NULL,
  PRIMARY KEY ("id", "user_id")
);

CREATE UNIQUE INDEX ON "notes" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
