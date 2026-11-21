package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/** Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant"). */
object AuthzUserStandingTable : IdTable<String>("\"AuthzUserStanding\"") {
    override val id = text("userId").entityId()
    val deactivatedAt = pgTimestamp("deactivatedAt", 3).nullable()
    val erasedAt = pgTimestamp("erasedAt", 3).nullable()
    val standingChangedAt = pgTimestamp("standingChangedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"AuthzUserStanding_pkey\"")
}
