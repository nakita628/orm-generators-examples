package documenso.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/** @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';"]) */
object DocumentMetaTable : IdTable<String>("\"DocumentMeta\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val subject = text("subject").nullable()
    val message = text("message").nullable()
    val timezone = text("timezone").nullable().default("Etc/UTC")
    val dateFormat = text("dateFormat").nullable().default("yyyy-MM-dd hh:mm a")
    val redirectUrl = text("redirectUrl").nullable()
    val signingOrder =
        pgEnum("signingOrder", "\"DocumentSigningOrder\"", DocumentSigningOrder.entries, DocumentSigningOrder::dbName).default(
            DocumentSigningOrder.PARALLEL,
        )
    val allowDictateNextSigner = bool("allowDictateNextSigner").default(false)
    val typedSignatureEnabled = bool("typedSignatureEnabled").default(true)
    val uploadSignatureEnabled = bool("uploadSignatureEnabled").default(true)
    val drawSignatureEnabled = bool("drawSignatureEnabled").default(true)
    val language = text("language").default("en")
    val distributionMethod =
        pgEnum(
            "distributionMethod",
            "\"DocumentDistributionMethod\"",
            DocumentDistributionMethod.entries,
            DocumentDistributionMethod::dbName,
        ).default(DocumentDistributionMethod.EMAIL)

    /** [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema) */
    val emailSettings = jsonb<String>("emailSettings", { it }, { it }).nullable()
    val emailReplyTo = text("emailReplyTo").nullable()
    val emailId = text("emailId").nullable()

    /** [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema) */
    val envelopeExpirationPeriod = jsonb<String>("envelopeExpirationPeriod", { it }, { it }).nullable()

    /** [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema) */
    val reminderSettings = jsonb<String>("reminderSettings", { it }, { it }).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"DocumentMeta_pkey\"")
}
