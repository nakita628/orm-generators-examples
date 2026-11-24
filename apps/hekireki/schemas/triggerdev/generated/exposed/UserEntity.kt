package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class UserEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserEntity>(UserTable)

    var email by UserTable.email

    /** Address awaiting confirmation; becomes `email` only after a link sent to it is opened */
    var pendingEmail by UserTable.pendingEmail
    var authenticationMethod by UserTable.authenticationMethod
    var authenticationProfile by UserTable.authenticationProfile
    var authenticationExtraParams by UserTable.authenticationExtraParams
    var authIdentifier by UserTable.authIdentifier
    var displayName by UserTable.displayName
    var name by UserTable.name
    var avatarUrl by UserTable.avatarUrl
    var admin by UserTable.admin

    /** Preferences for the dashboard */
    var dashboardPreferences by UserTable.dashboardPreferences
    var createdAt by UserTable.createdAt
    var updatedAt by UserTable.updatedAt

    /** @deprecated */
    var isOnCloudWaitlist by UserTable.isOnCloudWaitlist

    /** @deprecated */
    var featureCloud by UserTable.featureCloud

    /** @deprecated */
    var isOnHostedRepoWaitlist by UserTable.isOnHostedRepoWaitlist
    var marketingEmails by UserTable.marketingEmails
    var confirmedBasicDetails by UserTable.confirmedBasicDetails
    var referralSource by UserTable.referralSource
    var onboardingData by UserTable.onboardingData
    val orgMemberships by OrgMemberEntity referrersOn OrgMemberTable.userId
    val sentInvites by OrgMemberInviteEntity referrersOn OrgMemberInviteTable.inviterId
    var mfaEnabledAt by UserTable.mfaEnabledAt
    var mfaSecretReference by SecretReferenceEntity optionalReferencedOn UserTable.mfaSecretReferenceId
    var mfaSecretReferenceId by UserTable.mfaSecretReferenceId

    /** Hash of the last used code to prevent replay attacks */
    var mfaLastUsedCode by UserTable.mfaLastUsedCode

    /**
     * Maximum session lifetime in seconds for this user. Default = 1 year (Gregorian).
     * May be further restricted by Organization.maxSessionDuration on any of the user's orgs.
     */
    var sessionDuration by UserTable.sessionDuration

    /**
     * Absolute timestamp at which the user's current session must end. Written
     * at login (and any time the effective duration is recomputed) and shortened
     * when an admin lowers an org cap. Read on requireUser / getUser; null is
     * treated as "no enforced deadline" (legacy sessions from before this
     * shipped). Read-on-demand has zero per-request DB cost — the value
     * piggybacks on the User fetch that requireUser already does.
     */
    var nextSessionEnd by UserTable.nextSessionEnd
    var invitationCode by InvitationCodeEntity optionalReferencedOn UserTable.invitationCodeId
    var invitationCodeId by UserTable.invitationCodeId
    val personalAccessTokens by PersonalAccessTokenEntity referrersOn PersonalAccessTokenTable.userId
    val createdApiKeys by ApiKeyEntity optionalReferrersOn ApiKeyTable.createdByUserId
    val deployments by WorkerDeploymentEntity optionalReferrersOn WorkerDeploymentTable.triggeredById
    val backupCodes by MfaBackupCodeEntity referrersOn MfaBackupCodeTable.userId
    val bulkActions by BulkActionGroupEntity optionalReferrersOn BulkActionGroupTable.userId
    val impersonationsPerformed by ImpersonationAuditLogEntity referrersOn ImpersonationAuditLogTable.adminId
    val impersonationsReceived by ImpersonationAuditLogEntity referrersOn ImpersonationAuditLogTable.targetId
    val supportAccessRequestsMade by SupportAccessRequestEntity optionalReferrersOn SupportAccessRequestTable.requestedById
    val supportAccessRequestsApproved by SupportAccessRequestEntity optionalReferrersOn SupportAccessRequestTable.approvedById
    val customerQueries by CustomerQueryEntity optionalReferrersOn CustomerQueryTable.userId
    val metricsDashboards by MetricsDashboardEntity optionalReferrersOn MetricsDashboardTable.ownerId
    val platformNotifications by PlatformNotificationEntity optionalReferrersOn PlatformNotificationTable.userId
    val platformNotificationInteractions by PlatformNotificationInteractionEntity referrersOn PlatformNotificationInteractionTable.userId
    val environmentVariableIgnores by EnvironmentVariableIgnoreEntity optionalReferrersOn EnvironmentVariableIgnoreTable.createdById

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(UserTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
