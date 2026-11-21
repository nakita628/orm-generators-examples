package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** @zod.import(["import { ZRecipientAuthOptionsSchema } from '@documenso/lib/types/document-auth';"]) */
class RecipientEntity(
    id: EntityID<Int>,
) : Entity<Int>(id) {
    companion object : EntityClass<Int, RecipientEntity>(RecipientTable)

    var envelopeId by RecipientTable.envelopeId
    var email by RecipientTable.email
    var name by RecipientTable.name
    var token by RecipientTable.token
    var documentDeletedAt by RecipientTable.documentDeletedAt
    var expired by RecipientTable.expired
    var expiresAt by RecipientTable.expiresAt
    var expirationNotifiedAt by RecipientTable.expirationNotifiedAt
    var sentAt by RecipientTable.sentAt
    var signedAt by RecipientTable.signedAt
    var lastReminderSentAt by RecipientTable.lastReminderSentAt
    var nextReminderAt by RecipientTable.nextReminderAt
    var reminderCount by RecipientTable.reminderCount

    /** [RecipientAuthOptions] @zod.custom.use(ZRecipientAuthOptionsSchema) */
    var authOptions by RecipientTable.authOptions

    /** @zod.number.describe("The order in which the recipient should sign the document. Only works if the document is set to sequential signing.") */
    var signingOrder by RecipientTable.signingOrder
    var rejectionReason by RecipientTable.rejectionReason
    var role by RecipientTable.role
    var readStatus by RecipientTable.readStatus
    var signingStatus by RecipientTable.signingStatus
    var sendStatus by RecipientTable.sendStatus
    var envelope by EnvelopeEntity referencedOn RecipientTable.envelopeId
    val fields by FieldEntity referrersOn FieldTable.recipientId
    val signatures by SignatureEntity referrersOn SignatureTable.recipientId
    val cscCredential by CscCredentialEntity optionalBackReferencedOn CscCredentialTable.recipientId
    val cscSession by CscSessionEntity optionalBackReferencedOn CscSessionTable.recipientId
}
