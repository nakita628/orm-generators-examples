package calcom.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object WorkspacePlatformTable : IdTable<Int>("\"WorkspacePlatform\"") {
    override val id = integer("id").autoIncrement().entityId()

    /** @zod.string.min(1) */
    val slug = text("slug").uniqueIndex("\"WorkspacePlatform_slug_key\"")

    /** @zod.string.min(1) */
    val name = text("name")
    val description = text("description")

    /** @zod.import(["import { serviceAccountKeySchema } from '../../zod-utils'"]).custom.use(serviceAccountKeySchema) */
    val defaultServiceAccountKey = jsonb<String>("defaultServiceAccountKey", { it }, { it })
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val enabled = bool("enabled").default(false)
    override val primaryKey = PrimaryKey(id, name = "\"WorkspacePlatform_pkey\"")
}
