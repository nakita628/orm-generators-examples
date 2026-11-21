package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

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
class RetentionExemptionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, RetentionExemptionEntity>(RetentionExemptionTable)

    var createdAt by RetentionExemptionTable.createdAt
    var organization by OrganizationEntity referencedOn RetentionExemptionTable.organizationId
    var organizationId by RetentionExemptionTable.organizationId
    var entity by RetentionExemptionTable.entity
    var survey by SurveyEntity optionalReferencedOn RetentionExemptionTable.surveyId
    var surveyId by RetentionExemptionTable.surveyId
    var until by RetentionExemptionTable.until
    var source by RetentionExemptionTable.source1

    /** Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`). */
    var reason by RetentionExemptionTable.reason
    var createdBy by UserEntity optionalReferencedOn RetentionExemptionTable.createdById
    var createdById by RetentionExemptionTable.createdById
    var revokedAt by RetentionExemptionTable.revokedAt
    var revokedBy by UserEntity optionalReferencedOn RetentionExemptionTable.revokedById
    var revokedById by RetentionExemptionTable.revokedById
}
