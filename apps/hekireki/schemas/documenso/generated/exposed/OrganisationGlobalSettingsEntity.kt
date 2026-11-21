package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"]) */
class OrganisationGlobalSettingsEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrganisationGlobalSettingsEntity>(OrganisationGlobalSettingsTable)

    val organisation by OrganisationEntity optionalBackReferencedOn OrganisationTable.organisationGlobalSettingsId
    var documentVisibility by OrganisationGlobalSettingsTable.documentVisibility
    var documentLanguage by OrganisationGlobalSettingsTable.documentLanguage
    var includeSenderDetails by OrganisationGlobalSettingsTable.includeSenderDetails
    var includeSigningCertificate by OrganisationGlobalSettingsTable.includeSigningCertificate
    var includeAuditLog by OrganisationGlobalSettingsTable.includeAuditLog
    var documentTimezone by OrganisationGlobalSettingsTable.documentTimezone
    var documentDateFormat by OrganisationGlobalSettingsTable.documentDateFormat
    var delegateDocumentOwnership by OrganisationGlobalSettingsTable.delegateDocumentOwnership
    var typedSignatureEnabled by OrganisationGlobalSettingsTable.typedSignatureEnabled
    var uploadSignatureEnabled by OrganisationGlobalSettingsTable.uploadSignatureEnabled
    var drawSignatureEnabled by OrganisationGlobalSettingsTable.drawSignatureEnabled

    /** [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema) */
    var defaultRecipients by OrganisationGlobalSettingsTable.defaultRecipients
    var emailId by OrganisationGlobalSettingsTable.emailId
    var email by OrganisationEmailEntity optionalReferencedOn OrganisationGlobalSettingsTable.emailId
    var emailReplyTo by OrganisationGlobalSettingsTable.emailReplyTo

    /** [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema) */
    var emailDocumentSettings by OrganisationGlobalSettingsTable.emailDocumentSettings
    var brandingEnabled by OrganisationGlobalSettingsTable.brandingEnabled
    var brandingLogo by OrganisationGlobalSettingsTable.brandingLogo
    var brandingUrl by OrganisationGlobalSettingsTable.brandingUrl
    var brandingCompanyDetails by OrganisationGlobalSettingsTable.brandingCompanyDetails

    /** [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema) */
    var brandingColors by OrganisationGlobalSettingsTable.brandingColors
    var brandingCss by OrganisationGlobalSettingsTable.brandingCss

    /** [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema) */
    var envelopeExpirationPeriod by OrganisationGlobalSettingsTable.envelopeExpirationPeriod

    /** [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema) */
    var reminderSettings by OrganisationGlobalSettingsTable.reminderSettings
    var aiFeaturesEnabled by OrganisationGlobalSettingsTable.aiFeaturesEnabled
}
