package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
 * assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
 * table, so omitting it risks an adapter error on a code path we do not control.
 */
class OauthClientAssertionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OauthClientAssertionEntity>(OauthClientAssertionTable)

    var expiresAt by OauthClientAssertionTable.expiresAt
}
