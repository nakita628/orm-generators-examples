package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class ActivationCodeEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ActivationCodeEntity>(ActivationCodeTable)

    /** SHA-256 of the normalised code. Looked up on every redemption. */
    var codeHash by ActivationCodeTable.codeHash

    /**
     * The last four characters, so an operator tells two codes apart in a list
     * without the list holding either one.
     */
    var codeHint by ActivationCodeTable.codeHint

    /** The customer this code belongs to, and the terms it mints. */
    var organizationId by ActivationCodeTable.organizationId
    var organizationName by ActivationCodeTable.organizationName
    var email by ActivationCodeTable.email
    var planType by ActivationCodeTable.planType
    var maxMembers by ActivationCodeTable.maxMembers

    /** The lite seats the minted license carries: null when unlimited. Never a stand-in number. */
    var maxMembersLite by ActivationCodeTable.maxMembersLite

    /** True when the minted license carries unlimited lite seats; maxMembersLite is then null. */
    var liteSeatsUnlimited by ActivationCodeTable.liteSeatsUnlimited

    /** The Developer seats the minted license carries: null when not stated, which sells none. */
    var maxMembersDeveloper by ActivationCodeTable.maxMembersDeveloper

    /** How long the license this code mints runs for. */
    var licenseTermDays by ActivationCodeTable.licenseTermDays

    /** Hosted services the minted license may call. */
    var services by ActivationCodeTable.services

    /** When the code stops working, which is not the term of what it mints. */
    var expiresAt by ActivationCodeTable.expiresAt

    /** A code a customer may redeem on more than one install, for a rollout. */
    var reusable by ActivationCodeTable.reusable

    /**
     * Set by the redemption that won, which is what makes a single-use code
     * single use: the write itself carries the condition.
     */
    var redeemedAt by ActivationCodeTable.redeemedAt
    var redeemedByInstanceId by ActivationCodeTable.redeemedByInstanceId

    /** The license the winning redemption minted. */
    var issuedLicenseId by ActivationCodeTable.issuedLicenseId

    /** How many installs have redeemed it, which only moves on a reusable code. */
    var redemptionCount by ActivationCodeTable.redemptionCount
    var revokedAt by ActivationCodeTable.revokedAt
    var revokedById by ActivationCodeTable.revokedById
    var createdById by ActivationCodeTable.createdById
    var createdAt by ActivationCodeTable.createdAt
    var updatedAt by ActivationCodeTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ActivationCodeTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
