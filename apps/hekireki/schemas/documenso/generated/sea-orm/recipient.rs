use super::read_status::ReadStatus;
use super::recipient_role::RecipientRole;
use super::send_status::SendStatus;
use super::signing_status::SigningStatus;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZRecipientAuthOptionsSchema } from '@documenso/lib/types/document-auth';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Recipient")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(column_name = "envelopeId", column_type = "Text")]
    pub envelope_id: String,
    #[sea_orm(column_type = "String(StringLen::N(255))")]
    pub email: String,
    #[sea_orm(column_type = "String(StringLen::N(255))", default_value = "")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub token: String,
    #[sea_orm(
        column_name = "documentDeletedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub document_deleted_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub expired: Option<DateTime>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(
        column_name = "expirationNotifiedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub expiration_notified_at: Option<DateTime>,
    #[sea_orm(column_name = "sentAt", column_type = "custom(\"timestamp(3)\")")]
    pub sent_at: Option<DateTime>,
    #[sea_orm(column_name = "signedAt", column_type = "custom(\"timestamp(3)\")")]
    pub signed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "lastReminderSentAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_reminder_sent_at: Option<DateTime>,
    #[sea_orm(
        column_name = "nextReminderAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub next_reminder_at: Option<DateTime>,
    #[sea_orm(column_name = "reminderCount", default_value = 0)]
    pub reminder_count: i32,
    /// [RecipientAuthOptions] @zod.custom.use(ZRecipientAuthOptionsSchema)
    #[sea_orm(column_name = "authOptions", column_type = "JsonBinary")]
    pub auth_options: Option<Json>,
    /// @zod.number.describe("The order in which the recipient should sign the document. Only works if the document is set to sequential signing.")
    #[sea_orm(column_name = "signingOrder")]
    pub signing_order: Option<i32>,
    #[sea_orm(column_name = "rejectionReason", column_type = "Text")]
    pub rejection_reason: Option<String>,
    #[sea_orm(default_value = "SIGNER")]
    pub role: RecipientRole,
    #[sea_orm(column_name = "readStatus", default_value = "NOT_OPENED")]
    pub read_status: ReadStatus,
    #[sea_orm(column_name = "signingStatus", default_value = "NOT_SIGNED")]
    pub signing_status: SigningStatus,
    #[sea_orm(column_name = "sendStatus", default_value = "NOT_SENT")]
    pub send_status: SendStatus,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::envelope::Entity",
        from = "Column::EnvelopeId",
        to = "super::envelope::Column::Id",
        on_delete = "Cascade",
        fk_name = "Recipient_envelopeId_fkey"
    )]
    Envelope,
    #[sea_orm(has_many = "super::field::Entity")]
    Fields,
    #[sea_orm(has_many = "super::signature::Entity")]
    Signatures,
    #[sea_orm(has_one = "super::csc_credential::Entity")]
    CscCredential,
    #[sea_orm(has_one = "super::csc_session::Entity")]
    CscSession,
}

impl Related<super::envelope::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Envelope.def()
    }
}

impl Related<super::field::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Fields.def()
    }
}

impl Related<super::signature::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Signatures.def()
    }
}

impl Related<super::csc_credential::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CscCredential.def()
    }
}

impl Related<super::csc_session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CscSession.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
