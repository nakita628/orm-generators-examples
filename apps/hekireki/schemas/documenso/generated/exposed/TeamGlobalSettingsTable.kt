package documenso.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/** @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"]) */
object TeamGlobalSettingsTable : IdTable<String>("\"TeamGlobalSettings\"") {
    override val id = text("id").entityId()
    val documentVisibility =
        pgEnum("documentVisibility", "\"DocumentVisibility\"", DocumentVisibility.entries, DocumentVisibility::dbName).nullable()
    val documentLanguage = text("documentLanguage").nullable()
    val documentTimezone = text("documentTimezone").nullable()
    val documentDateFormat = text("documentDateFormat").nullable()
    val delegateDocumentOwnership = bool("delegateDocumentOwnership").nullable()
    val includeSenderDetails = bool("includeSenderDetails").nullable()
    val includeSigningCertificate = bool("includeSigningCertificate").nullable()
    val includeAuditLog = bool("includeAuditLog").nullable()
    val typedSignatureEnabled = bool("typedSignatureEnabled").nullable()
    val uploadSignatureEnabled = bool("uploadSignatureEnabled").nullable()
    val drawSignatureEnabled = bool("drawSignatureEnabled").nullable()

    /** [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema) */
    val defaultRecipients = jsonb<String>("defaultRecipients", { it }, { it }).nullable()
    val emailId =
        optReference(
            "emailId",
            OrganisationEmailTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TeamGlobalSettings_emailId_fkey\"",
        )
    val emailReplyTo = text("emailReplyTo").nullable()

    /** [DocumentEmailSettingsNullable] @zod.custom.use(ZDocumentEmailSettingsSchema) */
    val emailDocumentSettings = jsonb<String>("emailDocumentSettings", { it }, { it }).nullable()
    val brandingEnabled = bool("brandingEnabled").nullable()
    val brandingLogo = text("brandingLogo").nullable()
    val brandingUrl = text("brandingUrl").nullable()
    val brandingCompanyDetails = text("brandingCompanyDetails").nullable()

    /** [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema) */
    val brandingColors = jsonb<String>("brandingColors", { it }, { it }).nullable()
    val brandingCss = text("brandingCss").nullable()

    /** [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema) */
    val envelopeExpirationPeriod = jsonb<String>("envelopeExpirationPeriod", { it }, { it }).nullable()

    /** [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema) */
    val reminderSettings = jsonb<String>("reminderSettings", { it }, { it }).nullable()
    val aiFeaturesEnabled = bool("aiFeaturesEnabled").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"TeamGlobalSettings_pkey\"")
}
