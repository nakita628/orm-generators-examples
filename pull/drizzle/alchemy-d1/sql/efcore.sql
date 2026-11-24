CREATE TABLE "users" (
    "id" INTEGER NOT NULL CONSTRAINT "users_pkey" PRIMARY KEY AUTOINCREMENT,
    "email" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "created_at" INTEGER NOT NULL DEFAULT ((unixepoch()))
);

CREATE TABLE "posts" (
    "id" INTEGER NOT NULL CONSTRAINT "posts_pkey" PRIMARY KEY AUTOINCREMENT,
    "user_id" INTEGER NOT NULL,
    "title" TEXT NOT NULL,
    "body" TEXT NOT NULL,
    "created_at" INTEGER NOT NULL DEFAULT ((unixepoch())),
    CONSTRAINT "posts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
