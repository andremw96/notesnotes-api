CREATE TABLE "users" (
  "id" SERIAL,
  "full_name" varchar,
  "first_name" varchar,
  "last_name" varchar,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" TIMESTAMP DEFAULT (now()) NOT NULL,
  "updated_at" TIMESTAMP DEFAULT (now()) NOT NULL,
  "is_deleted" BOOLEAN DEFAULT FALSE NOT NULL,
  PRIMARY KEY ("id", "username")
);

CREATE TABLE "notes" (
  "id" SERIAL,
  "user_id" int NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar,
  "created_at" TIMESTAMP DEFAULT (now()) NOT NULL,
  "updated_at" TIMESTAMP DEFAULT (now()) NOT NULL,
  "is_deleted" BOOLEAN DEFAULT FALSE NOT NULL,
  PRIMARY KEY ("id", "user_id")
);

CREATE UNIQUE INDEX ON "notes" ("id");

CREATE UNIQUE INDEX ON "users" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users" ADD CONSTRAINT unique_username UNIQUE (username);
