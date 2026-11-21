package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object RoleTable : IdTable<String>("\"Role\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = pgEnum("name", "\"Roles\"", Roles.entries, Roles::dbName).default(Roles.USER).uniqueIndex("\"Role_name_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Role_pkey\"")
}
