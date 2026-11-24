CREATE TABLE "users" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "email" varchar NOT NULL UNIQUE, "name" varchar NOT NULL, "created_at" integer NOT NULL );

CREATE TABLE "posts" ( "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL, "title" varchar NOT NULL, "body" varchar NOT NULL, "created_at" integer NOT NULL, FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION );
