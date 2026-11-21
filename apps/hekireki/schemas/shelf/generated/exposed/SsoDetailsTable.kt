package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object SsoDetailsTable : IdTable<String>("\"SsoDetails\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val domain = text("domain")
    val baseUserGroupId = text("baseUserGroupId").nullable()
    val selfServiceGroupId = text("selfServiceGroupId").nullable()
    val adminGroupId = text("adminGroupId").nullable()
    val managerGroupId = text("managerGroupId").nullable()
    val requireSsoLogin = bool("requireSsoLogin").default(true)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SsoDetails_pkey\"")
}
