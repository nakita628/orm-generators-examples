use super::document_visibility::DocumentVisibility;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganisationGlobalSettings")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "documentVisibility", default_value = "EVERYONE")]
    pub document_visibility: DocumentVisibility,
    #[sea_orm(
        column_name = "documentLanguage",
        column_type = "Text",
        default_value = "en"
    )]
    pub document_language: String,
    #[sea_orm(column_name = "includeSenderDetails", default_value = true)]
    pub include_sender_details: bool,
    #[sea_orm(column_name = "includeSigningCertificate", default_value = true)]
    pub include_signing_certificate: bool,
    #[sea_orm(column_name = "includeAuditLog", default_value = false)]
    pub include_audit_log: bool,
    #[sea_orm(column_name = "documentTimezone", column_type = "Text")]
    pub document_timezone: Option<String>,
    #[sea_orm(
        column_name = "documentDateFormat",
        column_type = "Text",
        default_value = "yyyy-MM-dd hh:mm a"
    )]
    pub document_date_format: String,
    #[sea_orm(column_name = "delegateDocumentOwnership", default_value = false)]
    pub delegate_document_ownership: bool,
    #[sea_orm(column_name = "typedSignatureEnabled", default_value = true)]
    pub typed_signature_enabled: bool,
    #[sea_orm(column_name = "uploadSignatureEnabled", default_value = true)]
    pub upload_signature_enabled: bool,
    #[sea_orm(column_name = "drawSignatureEnabled", default_value = true)]
    pub draw_signature_enabled: bool,
    /// [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
    #[sea_orm(column_name = "defaultRecipients", column_type = "JsonBinary")]
    pub default_recipients: Option<Json>,
    #[sea_orm(column_name = "emailId", column_type = "Text")]
    pub email_id: Option<String>,
    #[sea_orm(column_name = "emailReplyTo", column_type = "Text")]
    pub email_reply_to: Option<String>,
    /// [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema)
    #[sea_orm(column_name = "emailDocumentSettings", column_type = "JsonBinary")]
    pub email_document_settings: Json,
    #[sea_orm(column_name = "brandingEnabled", default_value = false)]
    pub branding_enabled: bool,
    #[sea_orm(column_name = "brandingLogo", column_type = "Text", default_value = "")]
    pub branding_logo: String,
    #[sea_orm(column_name = "brandingUrl", column_type = "Text", default_value = "")]
    pub branding_url: String,
    #[sea_orm(
        column_name = "brandingCompanyDetails",
        column_type = "Text",
        default_value = ""
    )]
    pub branding_company_details: String,
    /// [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
    #[sea_orm(column_name = "brandingColors", column_type = "JsonBinary")]
    pub branding_colors: Option<Json>,
    #[sea_orm(column_name = "brandingCss", column_type = "Text", default_value = "")]
    pub branding_css: String,
    /// [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
    #[sea_orm(column_name = "envelopeExpirationPeriod", column_type = "JsonBinary")]
    pub envelope_expiration_period: Option<Json>,
    /// [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
    #[sea_orm(column_name = "reminderSettings", column_type = "JsonBinary")]
    pub reminder_settings: Option<Json>,
    #[sea_orm(column_name = "aiFeaturesEnabled", default_value = false)]
    pub ai_features_enabled: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organisation_email::Entity",
        from = "Column::EmailId",
        to = "super::organisation_email::Column::Id",
        fk_name = "OrganisationGlobalSettings_emailId_fkey"
    )]
    Email,
    #[sea_orm(has_one = "super::organisation::Entity")]
    Organisation,
}

impl Related<super::organisation_email::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Email.def()
    }
}

impl Related<super::organisation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organisation.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
