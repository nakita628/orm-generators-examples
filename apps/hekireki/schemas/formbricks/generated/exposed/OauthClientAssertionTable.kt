package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
 * assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
 * table, so omitting it risks an adapter error on a code path we do not control.
 */
object OauthClientAssertionTable : IdTable<String>("\"oauthClientAssertion\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val expiresAt = pgTimestamp("expiresAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"oauthClientAssertion_pkey\"")

    init {
        index("\"oauthClientAssertion_expiresAt_idx\"", false, expiresAt)
    }
}
