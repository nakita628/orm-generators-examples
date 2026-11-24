CREATE TABLE "users" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "email" text NOT NULL UNIQUE, "name" text NOT NULL, "created_at" integer DEFAULT (((unixepoch()))) NOT NULL);

CREATE TABLE "posts" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "title" text NOT NULL, "body" text NOT NULL, "created_at" integer DEFAULT (((unixepoch()))) NOT NULL, "user_id" integer NOT NULL REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED);
