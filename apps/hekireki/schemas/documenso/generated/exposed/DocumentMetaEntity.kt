package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';"]) */
class DocumentMetaEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DocumentMetaEntity>(DocumentMetaTable)

    var subject by DocumentMetaTable.subject
    var message by DocumentMetaTable.message
    var timezone by DocumentMetaTable.timezone
    var dateFormat by DocumentMetaTable.dateFormat
    var redirectUrl by DocumentMetaTable.redirectUrl
    var signingOrder by DocumentMetaTable.signingOrder
    var allowDictateNextSigner by DocumentMetaTable.allowDictateNextSigner
    var typedSignatureEnabled by DocumentMetaTable.typedSignatureEnabled
    var uploadSignatureEnabled by DocumentMetaTable.uploadSignatureEnabled
    var drawSignatureEnabled by DocumentMetaTable.drawSignatureEnabled
    var language by DocumentMetaTable.language
    var distributionMethod by DocumentMetaTable.distributionMethod

    /** [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema) */
    var emailSettings by DocumentMetaTable.emailSettings
    var emailReplyTo by DocumentMetaTable.emailReplyTo
    var emailId by DocumentMetaTable.emailId

    /** [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema) */
    var envelopeExpirationPeriod by DocumentMetaTable.envelopeExpirationPeriod

    /** [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema) */
    var reminderSettings by DocumentMetaTable.reminderSettings
    val envelope by EnvelopeEntity optionalBackReferencedOn EnvelopeTable.documentMetaId
}
