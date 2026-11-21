package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a complete survey configuration including questions, styling, and display rules.
 * Core model for the survey functionality in Formbricks.
 *
 * @property id - Unique identifier for the survey
 * @property name - Display name of the survey
 * @property type - Survey delivery method (link, web, website, app)
 * @property status - Current state of the survey (draft, active, completed, etc)
 * @property workspace - The workspace this survey belongs to
 * @property questions - JSON array containing survey questions configuration
 * @property displayOption - Rules for how often the survey can be shown
 * @property triggers - Actions that can trigger this survey
 * @property attributeFilters - Rules for targeting specific contacts
 * @property customHeadScripts - Survey-specific custom HTML scripts (self-hosted only)
 * @property customHeadScriptsMode - "add" (merge with workspace) or "replace" (override workspace)
 */
class SurveyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SurveyEntity>(SurveyTable)

    /** [SurveyHiddenFields] */
    var hiddenFields by SurveyTable.hiddenFields

    /** [SurveyVariables] */
    var variables by SurveyTable.variables
    var ownershipInitialized by SurveyTable.ownershipInitialized
    var createdAt by SurveyTable.createdAt
    var updatedAt by SurveyTable.updatedAt
    var name by SurveyTable.name
    var redirectUrl by SurveyTable.redirectUrl
    var type by SurveyTable.type
    var workspace by WorkspaceEntity referencedOn SurveyTable.workspaceId
    var workspaceId by SurveyTable.workspaceId
    var creator by UserEntity optionalReferencedOn SurveyTable.createdBy
    var createdBy by SurveyTable.createdBy

    /**
     * ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
     * the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
     */
    var visibility by SurveyTable.visibility
    var owner by UserEntity optionalReferencedOn SurveyTable.ownerId
    var ownerId by SurveyTable.ownerId

    /**
     * Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
     * with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
     * projection, decided from PostgreSQL facts until the projector acknowledges it.
     */
    var visibilityVersion by SurveyTable.visibilityVersion

    /** The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was. */
    var visibilityProjectedVersion by SurveyTable.visibilityProjectedVersion

    /**
     * "visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
     * trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
     * in create/update data.
     */
    var visibilityPending by SurveyTable.visibilityPending
    var visibilityChangedAt by SurveyTable.visibilityChangedAt
    var visibilityChangedById by SurveyTable.visibilityChangedById
    var status by SurveyTable.status

    /** [SurveyWelcomeCard] */
    var welcomeCard by SurveyTable.welcomeCard

    /** [SurveyQuestions] */
    var questions by SurveyTable.questions

    /** [SurveyBlocks] */
    var blocks by SurveyTable.blocks

    /** [SurveyEnding] */
    var endings by SurveyTable.endings
    val responses by ResponseEntity referrersOn ResponseTable.surveyId
    val quotas by SurveyQuotaEntity referrersOn SurveyQuotaTable.surveyId
    var displayOption by SurveyTable.displayOption
    var recontactDays by SurveyTable.recontactDays
    var displayLimit by SurveyTable.displayLimit
    val triggers by SurveyTriggerEntity referrersOn SurveyTriggerTable.surveyId

    /** [SurveyInlineTriggers] */
    var inlineTriggers by SurveyTable.inlineTriggers
    val attributeFilters by SurveyAttributeFilterEntity referrersOn SurveyAttributeFilterTable.surveyId
    val displays by DisplayEntity referrersOn DisplayTable.surveyId
    var autoClose by SurveyTable.autoClose
    var keepEndingCardOpen by SurveyTable.keepEndingCardOpen
    var autoComplete by SurveyTable.autoComplete
    var delay by SurveyTable.delay
    var publishOn by SurveyTable.publishOn
    var closeOn by SurveyTable.closeOn
    var archivedAt by SurveyTable.archivedAt

    /** [SurveyClosedMessage] */
    var surveyClosedMessage by SurveyTable.surveyClosedMessage
    var segmentId by SurveyTable.segmentId
    var segment by SegmentEntity optionalReferencedOn SurveyTable.segmentId

    /** [SurveyWorkspaceOverwrites] */
    var workspaceOverwrites by SurveyTable.workspaceOverwrites

    /** [SurveyStyling] */
    var styling by SurveyTable.styling

    /** [SurveySingleUse] */
    var singleUse by SurveyTable.singleUse
    var isVerifyEmailEnabled by SurveyTable.isVerifyEmailEnabled
    var isBackButtonHidden by SurveyTable.isBackButtonHidden
    var isAutoProgressingEnabled by SurveyTable.isAutoProgressingEnabled
    var isCaptureIpEnabled by SurveyTable.isCaptureIpEnabled
    var isAnonymizeResponsesEnabled by SurveyTable.isAnonymizeResponsesEnabled
    var pin by SurveyTable.pin
    var displayPercentage by SurveyTable.displayPercentage
    val languages by SurveyLanguageEntity referrersOn SurveyLanguageTable.surveyId
    var showLanguageSwitch by SurveyTable.showLanguageSwitch

    /**
     * Open the survey in the respondent's browser language when it matches an enabled survey
     * language. Null on surveys that predate the setting, which reads as off.
     */
    var autoSelectLanguage by SurveyTable.autoSelectLanguage
    val followUps by SurveyFollowUpEntity referrersOn SurveyFollowUpTable.surveyId

    /** [SurveyRecaptcha] */
    var recaptcha by SurveyTable.recaptcha

    /** [SurveyLinkMetadata] */
    var metadata by SurveyTable.metadata
    val retentionExemptions by RetentionExemptionEntity optionalReferrersOn RetentionExemptionTable.surveyId
    val retentionNotices by RetentionNoticeEntity optionalReferrersOn RetentionNoticeTable.surveyId
    val retentionRestore by RetentionSurveyRestoreEntity optionalBackReferencedOn RetentionSurveyRestoreTable.id
    var slug by SurveyTable.slug
    var customHeadScripts by SurveyTable.customHeadScripts
    var customHeadScriptsMode by SurveyTable.customHeadScriptsMode

    /**
     * [CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
     * by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
     */
    var customCss by SurveyTable.customCss

    /**
     * [SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
     * (ENG-3873). null = no conditions, which is what every existing survey keeps.
     */
    var showConditions by SurveyTable.showConditions

    /**
     * [SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
     * survey on load. null = none, which is what every existing survey keeps.
     */
    var startLogic by SurveyTable.startLogic

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SurveyTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
