package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.math.BigDecimal
import java.time.Instant
import java.time.temporal.ChronoUnit

object RuntimeEnvironmentTable : IdTable<String>("\"RuntimeEnvironment\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val slug = text("slug")
    val apiKey = text("apiKey").uniqueIndex("\"RuntimeEnvironment_apiKey_key\"")

    /** Non-null prevents the root API key from being returned to users. */
    val rootApiKeyHiddenAt = pgTimestamp("rootApiKeyHiddenAt", 3).nullable()

    /** @deprecated was for v2 */
    val pkApiKey = text("pkApiKey").uniqueIndex("\"RuntimeEnvironment_pkApiKey_key\"")
    val type =
        pgEnum("type", "\"RuntimeEnvironmentType\"", RuntimeEnvironmentType.entries, RuntimeEnvironmentType::dbName).default(
            RuntimeEnvironmentType.DEVELOPMENT,
        )

    /**
     * If true, this environment has branches and is treated differently in the dashboard/API
     * NB: this flag is NOT used for Development branches, instead (type, parentEnvironmentId) = (DEVELOPMENT, NULL) is used
     */
    val isBranchableEnvironment = bool("isBranchableEnvironment").default(false)
    val branchName = text("branchName").nullable()
    val parentEnvironmentId =
        optReference(
            "parentEnvironmentId",
            RuntimeEnvironmentTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RuntimeEnvironment_parentEnvironmentId_fkey\"",
        )

    /** Only configured on the PREVIEW root. Null disables automatic archiving. */
    val previewAutoArchiveAfterDays = integer("previewAutoArchiveAfterDays").nullable()
    val previewAutoArchiveExcludedBranches =
        array("previewAutoArchiveExcludedBranches", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** Durable per-project cleanup progress. Null next-check disables scheduling. */
    val previewAutoArchiveNextCheckAt = pgTimestamp("previewAutoArchiveNextCheckAt", 3).nullable()
    val previewAutoArchiveCursorCreatedAt = pgTimestamp("previewAutoArchiveCursorCreatedAt", 3).nullable()
    val previewAutoArchiveCursorId = text("previewAutoArchiveCursorId").nullable()
    val git = jsonb<String>("git", { it }, { it }).nullable()

    /** When set API calls will fail */
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()

    /** A memorable code for the environment */
    val shortcode = text("shortcode")
    val maximumConcurrencyLimit = integer("maximumConcurrencyLimit").default(5)
    val concurrencyLimitBurstFactor = decimal("concurrencyLimitBurstFactor", 4, 2).default(BigDecimal("2"))
    val paused = bool("paused").default(false)
    val pauseSource =
        pgEnum("pauseSource", "\"EnvironmentPauseSource\"", EnvironmentPauseSource.entries, EnvironmentPauseSource::dbName).nullable()
    val autoEnableInternalSources = bool("autoEnableInternalSources").default(true)
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RuntimeEnvironment_organizationId_fkey\"",
        )
    val projectId =
        reference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RuntimeEnvironment_projectId_fkey\"",
        )
    val orgMemberId =
        optReference(
            "orgMemberId",
            OrgMemberTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RuntimeEnvironment_orgMemberId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** Allows us to customize the built-in environment variables for a specific environment, like TRIGGER_OTEL_SPAN_ATTRIBUTE_VALUE_LENGTH_LIMIT */
    val builtInEnvironmentVariableOverrides = jsonb<String>("builtInEnvironmentVariableOverrides", { it }, { it }).nullable()
    val tunnelId = text("tunnelId").nullable()
    val currentSessionId =
        optReference(
            "currentSessionId",
            RuntimeEnvironmentSessionTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RuntimeEnvironment_currentSessionId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"RuntimeEnvironment_pkey\"")

    init {
        index("\"RuntimeEnvironment_parentEnvironmentId_idx\"", false, parentEnvironmentId)
        index("\"RuntimeEnvironment_projectId_idx\"", false, projectId)
        index("\"RuntimeEnvironment_organizationId_idx\"", false, organizationId)
        uniqueIndex("\"RuntimeEnvironment_projectId_slug_orgMemberId_key\"", projectId, slug, orgMemberId)
        uniqueIndex("\"RuntimeEnvironment_projectId_shortcode_key\"", projectId, shortcode)
    }
}
