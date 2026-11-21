package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"]) */
class TeamGlobalSettingsEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TeamGlobalSettingsEntity>(TeamGlobalSettingsTable)

    val team by TeamEntity optionalBackReferencedOn TeamTable.teamGlobalSettingsId
    var documentVisibility by TeamGlobalSettingsTable.documentVisibility
    var documentLanguage by TeamGlobalSettingsTable.documentLanguage
    var documentTimezone by TeamGlobalSettingsTable.documentTimezone
    var documentDateFormat by TeamGlobalSettingsTable.documentDateFormat
    var delegateDocumentOwnership by TeamGlobalSettingsTable.delegateDocumentOwnership
    var includeSenderDetails by TeamGlobalSettingsTable.includeSenderDetails
    var includeSigningCertificate by TeamGlobalSettingsTable.includeSigningCertificate
    var includeAuditLog by TeamGlobalSettingsTable.includeAuditLog
    var typedSignatureEnabled by TeamGlobalSettingsTable.typedSignatureEnabled
    var uploadSignatureEnabled by TeamGlobalSettingsTable.uploadSignatureEnabled
    var drawSignatureEnabled by TeamGlobalSettingsTable.drawSignatureEnabled

    /** [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema) */
    var defaultRecipients by TeamGlobalSettingsTable.defaultRecipients
    var emailId by TeamGlobalSettingsTable.emailId
    var email by OrganisationEmailEntity optionalReferencedOn TeamGlobalSettingsTable.emailId
    var emailReplyTo by TeamGlobalSettingsTable.emailReplyTo

    /** [DocumentEmailSettingsNullable] @zod.custom.use(ZDocumentEmailSettingsSchema) */
    var emailDocumentSettings by TeamGlobalSettingsTable.emailDocumentSettings
    var brandingEnabled by TeamGlobalSettingsTable.brandingEnabled
    var brandingLogo by TeamGlobalSettingsTable.brandingLogo
    var brandingUrl by TeamGlobalSettingsTable.brandingUrl
    var brandingCompanyDetails by TeamGlobalSettingsTable.brandingCompanyDetails

    /** [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema) */
    var brandingColors by TeamGlobalSettingsTable.brandingColors
    var brandingCss by TeamGlobalSettingsTable.brandingCss

    /** [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema) */
    var envelopeExpirationPeriod by TeamGlobalSettingsTable.envelopeExpirationPeriod

    /** [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema) */
    var reminderSettings by TeamGlobalSettingsTable.reminderSettings
    var aiFeaturesEnabled by TeamGlobalSettingsTable.aiFeaturesEnabled
}
