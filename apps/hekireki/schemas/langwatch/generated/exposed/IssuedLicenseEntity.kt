package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class IssuedLicenseEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, IssuedLicenseEntity>(IssuedLicenseTable)

    /** The `licenseId` inside the signed license (`lic-<uuid>`). */
    var licenseId by IssuedLicenseTable.licenseId

    /** SHA-256 of the `lwl_` token. Looked up on every credential resolution. */
    var tokenHash by IssuedLicenseTable.tokenHash

    /**
     * The customer organization on LangWatch Cloud. Null for a license that is
     * recorded but not linked to a customer yet, such as a self-serve purchase:
     * it resolves to nothing until an operator links it.
     */
    var organizationId by IssuedLicenseTable.organizationId

    /** As signed into the license. */
    var organizationName by IssuedLicenseTable.organizationName
    var email by IssuedLicenseTable.email
    var planType by IssuedLicenseTable.planType
    var maxMembers by IssuedLicenseTable.maxMembers

    /** The lite seats sold: null when not stated or when unlimited. Never a stand-in number. */
    var maxMembersLite by IssuedLicenseTable.maxMembersLite

    /** True when the license sells unlimited lite seats; maxMembersLite is then null. */
    var liteSeatsUnlimited by IssuedLicenseTable.liteSeatsUnlimited

    /** The Developer seats sold: null when not stated, which sells none. */
    var maxMembersDeveloper by IssuedLicenseTable.maxMembersDeveloper
    var issuedAt by IssuedLicenseTable.issuedAt
    var expiresAt by IssuedLicenseTable.expiresAt
    var source by IssuedLicenseTable.source1

    /** The operator who issued or registered it. Null for the purchase flow and the script. */
    var issuedById by IssuedLicenseTable.issuedById
    var revokedAt by IssuedLicenseTable.revokedAt
    var revokedById by IssuedLicenseTable.revokedById
    var revokedReason by IssuedLicenseTable.revokedReason

    /** Set when the install first presents the license that replaced this one. */
    var supersededAt by IssuedLicenseTable.supersededAt

    /**
     * The id of the IssuedLicense this one replaces, when it was reissued. A
     * plain column ref (relationMode="prisma" cannot declare a self-relation on
     * Postgres); the registry service validates it. Unique, so a license is
     * replaced by one license at most.
     */
    var replacesId by IssuedLicenseTable.replacesId

    /**
     * A reissued license waiting to be delivered over sync, encrypted. Erased
     * when the install first presents it. The only license text the registry holds.
     */
    var pendingDeliveryLicense by IssuedLicenseTable.pendingDeliveryLicense

    /** Hosted services this license may call, such as `instant_evals`. */
    var services by IssuedLicenseTable.services

    /** The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats. */
    var seatRateCents by IssuedLicenseTable.seatRateCents
    var seatCurrency by IssuedLicenseTable.seatCurrency

    /** The prepaid usage commit for the term, in USD cents. */
    var commitUsdCents by IssuedLicenseTable.commitUsdCents

    /** Whether usage may continue past the commit, up to `overageMaxUsdCents`. */
    var overageEnabled by IssuedLicenseTable.overageEnabled
    var overageMaxUsdCents by IssuedLicenseTable.overageMaxUsdCents

    /** The install this license is bound to. Set on first use, cleared by an operator. */
    var instanceId by IssuedLicenseTable.instanceId
    var instanceBoundAt by IssuedLicenseTable.instanceBoundAt

    /**
     * The last license sync this install reported: when it arrived, the
     * LangWatch version that sent it, and the seats in use at that moment.
     * Null until the install syncs for the first time.
     */
    var lastSyncAt by IssuedLicenseTable.lastSyncAt
    var lastSyncVersion by IssuedLicenseTable.lastSyncVersion
    var reportedMembers by IssuedLicenseTable.reportedMembers
    var reportedMembersLite by IssuedLicenseTable.reportedMembersLite

    /**
     * The managed gateway key (purpose CONNECT) that hosted-service traffic of
     * this license runs under. One per license, so spend rows name the install.
     * Created when the license first resolves. A plain column, not a relation.
     */
    var virtualKeyId by IssuedLicenseTable.virtualKeyId

    /**
     * The seats of the license this one replaced, set only when a seat change
     * raised a linked license. Billing reads it through licensing to invoice
     * the added seats, so the fact commits with the license or not at all.
     */
    var seatsRaisedFrom by IssuedLicenseTable.seatsRaisedFrom
    var createdAt by IssuedLicenseTable.createdAt
    var updatedAt by IssuedLicenseTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(IssuedLicenseTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
