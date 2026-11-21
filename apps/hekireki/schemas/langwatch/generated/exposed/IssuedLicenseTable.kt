package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The registry of licenses LangWatch has issued (ADR-139). It lives on
 * LangWatch Cloud; a self-hosted install has the table and no rows.
 *
 * The row never holds the token an install presents. `tokenHash` is the
 * SHA-256 of that token, so read access to this table is not credential
 * access. Status is derived: revoked, superseded, past `expiresAt`, else active.
 */
object IssuedLicenseTable : IdTable<String>("\"IssuedLicense\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /** The `licenseId` inside the signed license (`lic-<uuid>`). */
    val licenseId = text("licenseId").uniqueIndex("\"IssuedLicense_licenseId_key\"")

    /** SHA-256 of the `lwl_` token. Looked up on every credential resolution. */
    val tokenHash = text("tokenHash").uniqueIndex("\"IssuedLicense_tokenHash_key\"")

    /**
     * The customer organization on LangWatch Cloud. Null for a license that is
     * recorded but not linked to a customer yet, such as a self-serve purchase:
     * it resolves to nothing until an operator links it.
     */
    val organizationId = text("organizationId").nullable()

    /** As signed into the license. */
    val organizationName = text("organizationName")
    val email = text("email")
    val planType = text("planType")
    val maxMembers = integer("maxMembers")

    /** The lite seats sold: null when not stated or when unlimited. Never a stand-in number. */
    val maxMembersLite = integer("maxMembersLite").nullable()

    /** True when the license sells unlimited lite seats; maxMembersLite is then null. */
    val liteSeatsUnlimited = bool("liteSeatsUnlimited").default(false)

    /** The Developer seats sold: null when not stated, which sells none. */
    val maxMembersDeveloper = integer("maxMembersDeveloper").nullable()
    val issuedAt = pgTimestamp("issuedAt", 3)
    val expiresAt = pgTimestamp("expiresAt", 3)
    val source1 = pgEnum("source", "\"IssuedLicenseSource\"", IssuedLicenseSource.entries, IssuedLicenseSource::dbName)

    /** The operator who issued or registered it. Null for the purchase flow and the script. */
    val issuedById = text("issuedById").nullable()
    val revokedAt = pgTimestamp("revokedAt", 3).nullable()
    val revokedById = text("revokedById").nullable()
    val revokedReason = text("revokedReason").nullable()

    /** Set when the install first presents the license that replaced this one. */
    val supersededAt = pgTimestamp("supersededAt", 3).nullable()

    /**
     * The id of the IssuedLicense this one replaces, when it was reissued. A
     * plain column ref (relationMode="prisma" cannot declare a self-relation on
     * Postgres); the registry service validates it. Unique, so a license is
     * replaced by one license at most.
     */
    val replacesId = text("replacesId").nullable().uniqueIndex("\"IssuedLicense_replacesId_key\"")

    /**
     * A reissued license waiting to be delivered over sync, encrypted. Erased
     * when the install first presents it. The only license text the registry holds.
     */
    val pendingDeliveryLicense = text("pendingDeliveryLicense").nullable()

    /** Hosted services this license may call, such as `instant_evals`. */
    val services = array("services", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats. */
    val seatRateCents = integer("seatRateCents").nullable()
    val seatCurrency = pgEnum("seatCurrency", "\"Currency\"", Currency.entries, Currency::dbName).nullable()

    /** The prepaid usage commit for the term, in USD cents. */
    val commitUsdCents = integer("commitUsdCents").default(0)

    /** Whether usage may continue past the commit, up to `overageMaxUsdCents`. */
    val overageEnabled = bool("overageEnabled").default(false)
    val overageMaxUsdCents = integer("overageMaxUsdCents").nullable()

    /** The install this license is bound to. Set on first use, cleared by an operator. */
    val instanceId = text("instanceId").nullable()
    val instanceBoundAt = pgTimestamp("instanceBoundAt", 3).nullable()

    /**
     * The last license sync this install reported: when it arrived, the
     * LangWatch version that sent it, and the seats in use at that moment.
     * Null until the install syncs for the first time.
     */
    val lastSyncAt = pgTimestamp("lastSyncAt", 3).nullable()
    val lastSyncVersion = text("lastSyncVersion").nullable()
    val reportedMembers = integer("reportedMembers").nullable()
    val reportedMembersLite = integer("reportedMembersLite").nullable()

    /**
     * The managed gateway key (purpose CONNECT) that hosted-service traffic of
     * this license runs under. One per license, so spend rows name the install.
     * Created when the license first resolves. A plain column, not a relation.
     */
    val virtualKeyId = text("virtualKeyId").nullable().uniqueIndex("\"IssuedLicense_virtualKeyId_key\"")

    /**
     * The seats of the license this one replaced, set only when a seat change
     * raised a linked license. Billing reads it through licensing to invoice
     * the added seats, so the fact commits with the license or not at all.
     */
    val seatsRaisedFrom = integer("seatsRaisedFrom").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"IssuedLicense_pkey\"")

    init {
        index("\"IssuedLicense_organizationId_idx\"", false, organizationId)
        index("\"IssuedLicense_expiresAt_idx\"", false, expiresAt)
        index("\"IssuedLicense_seatsRaisedFrom_idx\"", false, seatsRaisedFrom)
    }
}
