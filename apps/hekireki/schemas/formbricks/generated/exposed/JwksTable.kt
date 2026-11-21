package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/** Better Auth JWT signing keys for OAuth/OIDC access and ID tokens. */
object JwksTable : IdTable<String>("jwks") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val publicKey = text("publicKey")
    val privateKey = text("privateKey")
    val createdAt = pgTimestamp("createdAt", 3)
    val expiresAt = pgTimestamp("expiresAt", 3).nullable()

    /**
     * Better Auth 1.7 (ENG-2343): the JWS algorithm this key signs with, so a keyring can hold more
     * than one alg and `getLatestKeyByAlg` can pick the right key per resource.
     *
     * NOT optional upstream in practice, despite being declared `required: false`: `createJwk` always
     * puts `alg` in the row it writes (defaulting to EdDSA), so without this column every key mint is
     * rejected by Prisma as an unknown argument — which breaks JWT signing, and with it the whole MCP
     * OAuth flow, on any deployment that has yet to mint a key. Nullable because rows written before
     * 1.7 have no value; upstream reads treat `alg: null` as the configured default alg, so existing
     * keys keep working untouched.
     */
    val alg = text("alg").nullable()

    /**
     * Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
     * the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
     * same always-written path as `alg` above.
     */
    val crv = text("crv").nullable()
    override val primaryKey = PrimaryKey(id, name = "jwks_pkey")
}
