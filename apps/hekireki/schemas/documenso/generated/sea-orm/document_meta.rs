use super::document_distribution_method::DocumentDistributionMethod;
use super::document_signing_order::DocumentSigningOrder;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DocumentMeta")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub subject: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub message: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "Etc/UTC")]
    pub timezone: Option<String>,
    #[sea_orm(
        column_name = "dateFormat",
        column_type = "Text",
        default_value = "yyyy-MM-dd hh:mm a"
    )]
    pub date_format: Option<String>,
    #[sea_orm(column_name = "redirectUrl", column_type = "Text")]
    pub redirect_url: Option<String>,
    #[sea_orm(column_name = "signingOrder", default_value = "PARALLEL")]
    pub signing_order: DocumentSigningOrder,
    #[sea_orm(column_name = "allowDictateNextSigner", default_value = false)]
    pub allow_dictate_next_signer: bool,
    #[sea_orm(column_name = "typedSignatureEnabled", default_value = true)]
    pub typed_signature_enabled: bool,
    #[sea_orm(column_name = "uploadSignatureEnabled", default_value = true)]
    pub upload_signature_enabled: bool,
    #[sea_orm(column_name = "drawSignatureEnabled", default_value = true)]
    pub draw_signature_enabled: bool,
    #[sea_orm(column_type = "Text", default_value = "en")]
    pub language: String,
    #[sea_orm(column_name = "distributionMethod", default_value = "EMAIL")]
    pub distribution_method: DocumentDistributionMethod,
    /// [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema)
    #[sea_orm(column_name = "emailSettings", column_type = "JsonBinary")]
    pub email_settings: Option<Json>,
    #[sea_orm(column_name = "emailReplyTo", column_type = "Text")]
    pub email_reply_to: Option<String>,
    #[sea_orm(column_name = "emailId", column_type = "Text")]
    pub email_id: Option<String>,
    /// [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
    #[sea_orm(column_name = "envelopeExpirationPeriod", column_type = "JsonBinary")]
    pub envelope_expiration_period: Option<Json>,
    /// [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
    #[sea_orm(column_name = "reminderSettings", column_type = "JsonBinary")]
    pub reminder_settings: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_one = "super::envelope::Entity")]
    Envelope,
}

impl Related<super::envelope::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Envelope.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
