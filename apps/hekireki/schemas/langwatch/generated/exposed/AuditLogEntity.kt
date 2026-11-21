package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class AuditLogEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuditLogEntity>(AuditLogTable)

    var createdAt by AuditLogTable.createdAt
    var userId by AuditLogTable.userId
    var projectId by AuditLogTable.projectId
    var organizationId by AuditLogTable.organizationId
    var action by AuditLogTable.action
    var args by AuditLogTable.args
    var error by AuditLogTable.error
    var ipAddress by AuditLogTable.ipAddress
    var userAgent by AuditLogTable.userAgent
    var metadata by AuditLogTable.metadata

    /**
     * WHO REALLY DID IT, when that is not `userId`.
     *
     * Under an impersonation the session's `user.id` is the SUBJECT — the
     * customer whose access an operator is borrowing — so a row naming only
     * `userId` filed the operator's act against the customer. An operator
     * granting a permanent way into somebody's organization appeared, in the
     * only durable record of it, to be that somebody.
     *
     * Null on an ordinary request, where the two would be the same person and
     * saying so twice adds nothing. ADR-118 / D06: every decision names both.
     */
    var actorUserId by AuditLogTable.actorUserId
    var targetKind by AuditLogTable.targetKind
    var targetId by AuditLogTable.targetId
    var before by AuditLogTable.before
    var after by AuditLogTable.after

    /** A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2). */
    var idempotencyKey by AuditLogTable.idempotencyKey
}
