package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
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
object WorkspaceTable : IdTable<String>("\"Workspace\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val name = text("name")
    val legacyEnvironmentId = text("legacyEnvironmentId").nullable().uniqueIndex("\"Workspace_legacyEnvironmentId_key\"")
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Workspace_organizationId_fkey\"",
        )

    /** [Styling] */
    val styling = jsonb<String>("styling", { it }, { it }).default("{\"allowStyleOverwrite\":true}")

    /** [WorkspaceConfig] */
    val config = jsonb<String>("config", { it }, { it }).default("{}")
    val recontactDays = integer("recontactDays").default(7)
    val linkSurveyBranding = bool("linkSurveyBranding").default(true)
    val inAppSurveyBranding = bool("inAppSurveyBranding").default(true)
    val placement =
        pgEnum("placement", "\"WidgetPlacement\"", WidgetPlacement.entries, WidgetPlacement::dbName).default(WidgetPlacement.BOTTOM_RIGHT)
    val clickOutsideClose = bool("clickOutsideClose").default(true)
    val overlay = pgEnum("overlay", "\"SurveyOverlay\"", SurveyOverlay.entries, SurveyOverlay::dbName).default(SurveyOverlay.NONE)

    /** Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`. */
    val overlayColor = text("overlayColor").nullable()
    val overlayOpacity = integer("overlayOpacity").nullable()

    /** [Logo] */
    val logo = jsonb<String>("logo", { it }, { it }).nullable()
    val appSetupCompleted = bool("appSetupCompleted").default(false)
    val customHeadScripts = text("customHeadScripts").nullable()

    /** [CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service. */
    val customCss = jsonb<String>("customCss", { it }, { it }).nullable()

    /**
     * [CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
     * submitting its source again through the current processor and authorization checks.
     */
    val customCssPrevious = jsonb<String>("customCssPrevious", { it }, { it }).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Workspace_pkey\"")

    init {
        uniqueIndex("\"Workspace_organizationId_name_key\"", organizationId, name)
    }
}
