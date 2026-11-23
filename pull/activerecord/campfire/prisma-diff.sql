-- RedefineTables
PRAGMA defer_foreign_keys=ON;
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_active_storage_attachments" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "blob_id" INTEGER NOT NULL,
    "created_at" datetime(6) NOT NULL,
    "name" TEXT NOT NULL,
    "record_id" BIGINT NOT NULL,
    "record_type" TEXT NOT NULL,
    CONSTRAINT "active_storage_attachments_blob_id_fkey" FOREIGN KEY ("blob_id") REFERENCES "active_storage_blobs" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);
INSERT INTO "new_active_storage_attachments" ("blob_id", "created_at", "id", "name", "record_id", "record_type") SELECT "blob_id", "created_at", "id", "name", "record_id", "record_type" FROM "active_storage_attachments";
DROP TABLE "active_storage_attachments";
ALTER TABLE "new_active_storage_attachments" RENAME TO "active_storage_attachments";
CREATE INDEX "index_active_storage_attachments_on_blob_id" ON "active_storage_attachments"("blob_id");
CREATE UNIQUE INDEX "index_active_storage_attachments_uniqueness" ON "active_storage_attachments"("record_type", "record_id", "name", "blob_id");
CREATE TABLE "new_active_storage_variant_records" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "blob_id" INTEGER NOT NULL,
    "variation_digest" TEXT NOT NULL,
    CONSTRAINT "active_storage_variant_records_blob_id_fkey" FOREIGN KEY ("blob_id") REFERENCES "active_storage_blobs" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);
INSERT INTO "new_active_storage_variant_records" ("blob_id", "id", "variation_digest") SELECT "blob_id", "id", "variation_digest" FROM "active_storage_variant_records";
DROP TABLE "active_storage_variant_records";
ALTER TABLE "new_active_storage_variant_records" RENAME TO "active_storage_variant_records";
CREATE UNIQUE INDEX "index_active_storage_variant_records_uniqueness" ON "active_storage_variant_records"("blob_id", "variation_digest");
PRAGMA foreign_keys=ON;
PRAGMA defer_foreign_keys=OFF;
