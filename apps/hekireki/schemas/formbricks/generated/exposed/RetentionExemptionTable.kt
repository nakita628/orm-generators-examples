package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Holds one survey back from one policy until `until`. Keyed by organisation + entity rather than a
 * policy id, so removing and re-adding a policy doesn't orphan its exemptions.
 *
 * Active means `revokedAt IS NULL AND until > now()`, so expiry needs no job. "One active exemption
 * per survey and policy" is the partial unique index `RetentionExemption_surveyId_entity_active_key`
 * (`WHERE "revokedAt" IS NULL`) in
 * `migration/20261007120000_add_data_retention_tables/migration.sql`; Prisma's `@@unique` has no
 * `where`, so it is deliberately not declared here (see "Indexes Prisma cannot express" in the
 * README). An expired row still holds that slot until it is revoked, so creating a new exemption
 * must close an expired one in the same transaction.
 *
 * `surveyId` is nullable and a CHECK requires exactly one target, so a later `workflowId` is one more
 * column in that CHECK. The survey's organisation can't be tied to `organizationId` in the database
 * (a survey carries its workspace, not its organisation): writers derive it from the survey.
 */
object RetentionExemptionTable : IdTable<String>("\"RetentionExemption\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionExemption_organizationId_fkey\"",
        )
    val entity = pgEnum("entity", "\"RetentionEntity\"", RetentionEntity.entries, RetentionEntity::dbName)
    val surveyId =
        optReference(
            "surveyId",
            SurveyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionExemption_surveyId_fkey\"",
        )
    val until = pgTimestamp("until", 3)
    val source1 =
        pgEnum("source", "\"RetentionExemptionSource\"", RetentionExemptionSource.entries, RetentionExemptionSource::dbName).default(
            RetentionExemptionSource.MANUAL,
        )

    /** Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`). */
    val reason = text("reason").nullable()
    val createdById =
        optReference(
            "createdById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionExemption_createdById_fkey\"",
        )
    val revokedAt = pgTimestamp("revokedAt", 3).nullable()
    val revokedById =
        optReference(
            "revokedById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionExemption_revokedById_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"RetentionExemption_pkey\"")

    init {
        index("\"RetentionExemption_organizationId_created_at_id_idx\"", false, organizationId, createdAt, id)
        index("\"RetentionExemption_surveyId_idx\"", false, surveyId)
        index("\"RetentionExemption_createdById_idx\"", false, createdById)
        index("\"RetentionExemption_revokedById_idx\"", false, revokedById)
    }
}
