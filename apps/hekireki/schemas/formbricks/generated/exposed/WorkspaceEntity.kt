package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Main grouping mechanism for resources in Formbricks.
 * Each organization can have multiple workspaces to separate different applications or products.
 *
 * @property id - Unique identifier for the workspace
 * @property name - Display name of the workspace
 * @property organization - Reference to parent organization
 * @property surveys - Collection of surveys in this workspace
 * @property styling - Workspace-wide styling configuration
 * @property config - Workspace-specific configuration
 * @property recontactDays - Default recontact delay for surveys
 * @property placement - Default widget placement for in-app surveys
 */
class WorkspaceEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WorkspaceEntity>(WorkspaceTable)

    var createdAt by WorkspaceTable.createdAt
    var updatedAt by WorkspaceTable.updatedAt
    var name by WorkspaceTable.name
    var legacyEnvironmentId by WorkspaceTable.legacyEnvironmentId
    var organization by OrganizationEntity referencedOn WorkspaceTable.organizationId
    var organizationId by WorkspaceTable.organizationId

    /** [Styling] */
    var styling by WorkspaceTable.styling

    /** [WorkspaceConfig] */
    var config by WorkspaceTable.config
    var recontactDays by WorkspaceTable.recontactDays
    var linkSurveyBranding by WorkspaceTable.linkSurveyBranding
    var inAppSurveyBranding by WorkspaceTable.inAppSurveyBranding
    var placement by WorkspaceTable.placement
    var clickOutsideClose by WorkspaceTable.clickOutsideClose
    var overlay by WorkspaceTable.overlay

    /** Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`. */
    var overlayColor by WorkspaceTable.overlayColor
    var overlayOpacity by WorkspaceTable.overlayOpacity
    val languages by LanguageEntity referrersOn LanguageTable.workspaceId
    val embeddedData by EmbeddedDataEntity referrersOn EmbeddedDataTable.workspaceId

    /** [Logo] */
    var logo by WorkspaceTable.logo
    val workspaceTeams by WorkspaceTeamEntity referrersOn WorkspaceTeamTable.workspaceId
    var appSetupCompleted by WorkspaceTable.appSetupCompleted
    var customHeadScripts by WorkspaceTable.customHeadScripts

    /** [CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service. */
    var customCss by WorkspaceTable.customCss

    /**
     * [CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
     * submitting its source again through the current processor and authorization checks.
     */
    var customCssPrevious by WorkspaceTable.customCssPrevious
    val feedbackDirectoryWorkspaces by FeedbackDirectoryWorkspaceEntity referrersOn FeedbackDirectoryWorkspaceTable.workspaceId
    val charts by ChartEntity referrersOn ChartTable.workspaceId
    val dashboards by DashboardEntity referrersOn DashboardTable.workspaceId
    val creatives by CreativeEntity referrersOn CreativeTable.workspaceId
    val attachmentExports by AttachmentExportEntity referrersOn AttachmentExportTable.workspaceId
    val surveys by SurveyEntity referrersOn SurveyTable.workspaceId
    val contacts by ContactEntity referrersOn ContactTable.workspaceId
    val actionClasses by ActionClassEntity referrersOn ActionClassTable.workspaceId
    val contactAttributeKeys by ContactAttributeKeyEntity referrersOn ContactAttributeKeyTable.workspaceId
    val webhooks by WebhookEntity referrersOn WebhookTable.workspaceId
    val tags by TagEntity referrersOn TagTable.workspaceId
    val segments by SegmentEntity referrersOn SegmentTable.workspaceId
    val integrations by IntegrationEntity referrersOn IntegrationTable.workspaceId
    val apiKeyWorkspaces by ApiKeyWorkspaceEntity referrersOn ApiKeyWorkspaceTable.workspaceId
    val feedbackSources by FeedbackSourceEntity referrersOn FeedbackSourceTable.workspaceId
    val workflows by WorkflowEntity referrersOn WorkflowTable.workspaceId
    val workflowRuns by WorkflowRunEntity referrersOn WorkflowRunTable.workspaceId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WorkspaceTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
