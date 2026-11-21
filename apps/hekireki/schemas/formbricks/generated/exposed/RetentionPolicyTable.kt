package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One policy per organisation and kind of data: a notice `warnDays` before the action, and the action
 * `periodDays` after the target's clock. The action follows from the kind of data: responses are
 * deleted, surveys archived (the archive purge deletes them 30 days later, as it does every archived
 * survey), members deactivated. The dates themselves are never stored; they are computed by the shared
 * helper in `apps/web/modules/ee/data-retention/lib/schedule.ts`, so the UI, the emails and the sweep
 * always agree.
 *
 * The allowed ranges are validated by the API. The migration repeats the ones that guard against
 * acting early or re-notifying (a notice of at least 14 days and shorter than the period, an enabled
 * policy always has `enabledAt`) and the structural ones (conditions on the surveys policy only) as
 * CHECK constraints, which Prisma cannot express.
 */
object RetentionPolicyTable : IdTable<String>("\"RetentionPolicy\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionPolicy_organizationId_fkey\"",
        )
    val entity = pgEnum("entity", "\"RetentionEntity\"", RetentionEntity.entries, RetentionEntity::dbName)
    val enabled = bool("enabled").default(false)

    /**
     * When the current configuration took effect: the policy was switched on or unpaused, or changed
     * in a way that could bring an action earlier (a shorter period or warning, different
     * conditions). Lengthening does not move it. The warning always runs in full from here, so no
     * change can delete data on the night it is made (ENG-3614). Also the clock for members with no
     * recorded sign-in.
     */
    val enabledAt = pgTimestamp("enabledAt", 3).nullable()

    /**
     * How long before the action the notice goes out. The action never runs before the notice has
     * been out for this long, so the warning always runs in full.
     */
    val warnDays = integer("warnDays").default(60)

    /**
     * The period, counted from the target's clock, after which the policy acts: responses are
     * deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
     */
    val periodDays = integer("periodDays")

    /** Surveys policy only. */
    val conditions =
        array(
            "conditions",
            PgEnumColumnType("\"RetentionSurveyCondition\"", RetentionSurveyCondition.entries, RetentionSurveyCondition::dbName),
        ).nullable().databaseDefault("'{}'") { emptyList() }

    /**
     * The last night the sweep's time ran out before it reached this policy's organisation, licensed.
     * The organisation was due and the sweep was running, so it counts as reached for the gap that
     * restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
     */
    val deferredAt = pgTimestamp("deferredAt", 3).nullable()
    val updatedById =
        optReference(
            "updatedById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionPolicy_updatedById_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"RetentionPolicy_pkey\"")

    init {
        index("\"RetentionPolicy_updatedById_idx\"", false, updatedById)
        uniqueIndex("\"RetentionPolicy_organizationId_entity_key\"", organizationId, entity)
    }
}
