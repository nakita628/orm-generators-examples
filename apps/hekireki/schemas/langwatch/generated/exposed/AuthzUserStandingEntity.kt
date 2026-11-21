package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant"). */
class AuthzUserStandingEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuthzUserStandingEntity>(AuthzUserStandingTable)

    var deactivatedAt by AuthzUserStandingTable.deactivatedAt
    var erasedAt by AuthzUserStandingTable.erasedAt
    var standingChangedAt by AuthzUserStandingTable.standingChangedAt
}
