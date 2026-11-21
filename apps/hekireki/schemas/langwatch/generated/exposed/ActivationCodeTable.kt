package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A short code a fresh install pastes instead of a license blob (ADR-139,
 * section 5). Redeeming it mints the license the code describes and returns
 * it signed, which is what identifies a new customer without anybody emailing
 * a license key around.
 *
 * The row holds a hash of the code, never the code. Every lookup hashes what
 * the caller presented first, so the stored value is a step away from anything
 * that could be presented: read access to this table is not credential access.
 */
object ActivationCodeTable : IdTable<String>("\"ActivationCode\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /** SHA-256 of the normalised code. Looked up on every redemption. */
    val codeHash = text("codeHash").uniqueIndex("\"ActivationCode_codeHash_key\"")

    /**
     * The last four characters, so an operator tells two codes apart in a list
     * without the list holding either one.
     */
    val codeHint = text("codeHint")

    /** The customer this code belongs to, and the terms it mints. */
    val organizationId = text("organizationId")
    val organizationName = text("organizationName")
    val email = text("email")
    val planType = text("planType")
    val maxMembers = integer("maxMembers")

    /** The lite seats the minted license carries: null when unlimited. Never a stand-in number. */
    val maxMembersLite = integer("maxMembersLite").nullable().default(0)

    /** True when the minted license carries unlimited lite seats; maxMembersLite is then null. */
    val liteSeatsUnlimited = bool("liteSeatsUnlimited").default(false)

    /** The Developer seats the minted license carries: null when not stated, which sells none. */
    val maxMembersDeveloper = integer("maxMembersDeveloper").nullable()

    /** How long the license this code mints runs for. */
    val licenseTermDays = integer("licenseTermDays")

    /** Hosted services the minted license may call. */
    val services = array("services", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** When the code stops working, which is not the term of what it mints. */
    val expiresAt = pgTimestamp("expiresAt", 3)

    /** A code a customer may redeem on more than one install, for a rollout. */
    val reusable = bool("reusable").default(false)

    /**
     * Set by the redemption that won, which is what makes a single-use code
     * single use: the write itself carries the condition.
     */
    val redeemedAt = pgTimestamp("redeemedAt", 3).nullable()
    val redeemedByInstanceId = text("redeemedByInstanceId").nullable()

    /** The license the winning redemption minted. */
    val issuedLicenseId = text("issuedLicenseId").nullable()

    /** How many installs have redeemed it, which only moves on a reusable code. */
    val redemptionCount = integer("redemptionCount").default(0)
    val revokedAt = pgTimestamp("revokedAt", 3).nullable()
    val revokedById = text("revokedById").nullable()
    val createdById = text("createdById").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ActivationCode_pkey\"")

    init {
        index("\"ActivationCode_organizationId_idx\"", false, organizationId)
        index("\"ActivationCode_expiresAt_idx\"", false, expiresAt)
    }
}
