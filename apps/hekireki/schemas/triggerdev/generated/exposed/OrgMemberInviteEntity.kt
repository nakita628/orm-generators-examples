package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class OrgMemberInviteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrgMemberInviteEntity>(OrgMemberInviteTable)

    var token by OrgMemberInviteTable.token
    var email by OrgMemberInviteTable.email
    var role by OrgMemberInviteTable.role

    /**
     * Optional RBAC role to assign on invite acceptance. When set, the
     * accept-invite flow calls the loaded RBAC plugin's setUserRole with
     * this id after creating the OrgMember. Null = legacy behaviour, the
     * runtime fallback derives the role from `role` above.
     *
     * Plain text (not an FK) — the RBAC plugin's RbacRole table lives on
     * a separate schema (Drizzle, not Prisma) so we can't model the FK
     * here. Validation happens at write time (action) and read time
     * (acceptInvite).
     */
    var rbacRoleId by OrgMemberInviteTable.rbacRoleId
    var organization by OrganizationEntity referencedOn OrgMemberInviteTable.organizationId
    var organizationId by OrgMemberInviteTable.organizationId
    var inviter by UserEntity referencedOn OrgMemberInviteTable.inviterId
    var inviterId by OrgMemberInviteTable.inviterId
    var createdAt by OrgMemberInviteTable.createdAt
    var updatedAt by OrgMemberInviteTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(OrgMemberInviteTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
