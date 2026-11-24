package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** @zod.import(["import { ZAdminTagTypeSchema } from '@documenso/lib/types/admin-tags';"]) */
class AdminTagEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AdminTagEntity>(AdminTagTable)

    var createdAt by AdminTagTable.createdAt
    var updatedAt by AdminTagTable.updatedAt

    /** [AdminTagType] @zod.custom.use(ZAdminTagTypeSchema) */
    var type by AdminTagTable.type
    var description by AdminTagTable.description
    var userId by AdminTagTable.userId
    var user by UserEntity optionalReferencedOn AdminTagTable.userId
    var organisationId by AdminTagTable.organisationId
    var organisation by OrganisationEntity optionalReferencedOn AdminTagTable.organisationId
    var createdByUserId by AdminTagTable.createdByUserId
    var createdByUser by UserEntity optionalReferencedOn AdminTagTable.createdByUserId
    var adminRuleId by AdminTagTable.adminRuleId
    var adminRule by AdminRuleEntity optionalReferencedOn AdminTagTable.adminRuleId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AdminTagTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
