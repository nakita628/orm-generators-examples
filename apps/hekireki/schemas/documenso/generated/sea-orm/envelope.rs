use super::document_source::DocumentSource;
use super::document_status::DocumentStatus;
use super::document_visibility::DocumentVisibility;
use super::envelope_type::EnvelopeType;
use super::template_type::TemplateType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZDocumentAuthOptionsSchema } from '@documenso/lib/types/document-auth';", "import { ZDocumentFormValuesSchema } from '@documenso/lib/types/document-form-values';", "import { ZSignatureLevelSchema } from '@documenso/lib/types/signature-level';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Envelope")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "secondaryId", column_type = "Text")]
    pub secondary_id: String,
    /// @zod.string.describe("A custom external ID you can use to identify the document.")
    #[sea_orm(column_name = "externalId", column_type = "Text")]
    pub external_id: Option<String>,
    pub r#type: EnvelopeType,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(column_name = "deletedAt", column_type = "custom(\"timestamp(3)\")")]
    pub deleted_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(default_value = "DRAFT")]
    pub status: DocumentStatus,
    pub source: DocumentSource,
    /// @zod.string.describe("The token for viewing the document using the QR code on the certificate.")
    #[sea_orm(column_name = "qrToken", column_type = "Text")]
    pub qr_token: Option<String>,
    /// [SignatureLevel] @zod.custom.use(ZSignatureLevelSchema)
    #[sea_orm(column_name = "signatureLevel", column_type = "Text")]
    pub signature_level: String,
    #[sea_orm(column_name = "internalVersion")]
    pub internal_version: i32,
    #[sea_orm(column_name = "useLegacyFieldInsertion", default_value = false)]
    pub use_legacy_field_insertion: bool,
    /// [DocumentAuthOptions] @zod.custom.use(ZDocumentAuthOptionsSchema)
    #[sea_orm(column_name = "authOptions", column_type = "JsonBinary")]
    pub auth_options: Option<Json>,
    /// [DocumentFormValues] @zod.custom.use(ZDocumentFormValuesSchema)
    #[sea_orm(column_name = "formValues", column_type = "JsonBinary")]
    pub form_values: Option<Json>,
    #[sea_orm(default_value = "EVERYONE")]
    pub visibility: DocumentVisibility,
    #[sea_orm(column_name = "templateType", default_value = "PRIVATE")]
    pub template_type: TemplateType,
    #[sea_orm(column_name = "publicTitle", column_type = "Text", default_value = "")]
    pub public_title: String,
    #[sea_orm(
        column_name = "publicDescription",
        column_type = "Text",
        default_value = ""
    )]
    pub public_description: String,
    #[sea_orm(column_name = "templateId")]
    pub template_id: Option<i32>,
    /// @zod.number.describe("The ID of the user that created this document.")
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(column_name = "teamId")]
    pub team_id: i32,
    #[sea_orm(column_name = "folderId", column_type = "Text")]
    pub folder_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "documentMetaId", column_type = "Text")]
    pub document_meta_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Envelope_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "Envelope_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::folder::Entity",
        from = "Column::FolderId",
        to = "super::folder::Column::Id",
        on_delete = "SetNull",
        fk_name = "Envelope_folderId_fkey"
    )]
    Folder,
    #[sea_orm(
        belongs_to = "super::document_meta::Entity",
        from = "Column::DocumentMetaId",
        to = "super::document_meta::Column::Id",
        fk_name = "Envelope_documentMetaId_fkey"
    )]
    DocumentMeta,
    #[sea_orm(has_many = "super::envelope_item::Entity")]
    EnvelopeItems,
    #[sea_orm(has_many = "super::recipient::Entity")]
    Recipients,
    #[sea_orm(has_many = "super::field::Entity")]
    Fields,
    #[sea_orm(has_many = "super::envelope_content::Entity")]
    Contents,
    #[sea_orm(has_many = "super::document_share_link::Entity")]
    ShareLinks,
    #[sea_orm(has_many = "super::document_audit_log::Entity")]
    AuditLogs,
    #[sea_orm(has_many = "super::envelope_attachment::Entity")]
    EnvelopeAttachments,
    #[sea_orm(has_one = "super::template_direct_link::Entity")]
    DirectLink,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::folder::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Folder.def()
    }
}

impl Related<super::document_meta::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DocumentMeta.def()
    }
}

impl Related<super::envelope_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EnvelopeItems.def()
    }
}

impl Related<super::recipient::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Recipients.def()
    }
}

impl Related<super::field::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Fields.def()
    }
}

impl Related<super::envelope_content::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Contents.def()
    }
}

impl Related<super::document_share_link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ShareLinks.def()
    }
}

impl Related<super::document_audit_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditLogs.def()
    }
}

impl Related<super::envelope_attachment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EnvelopeAttachments.def()
    }
}

impl Related<super::template_direct_link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DirectLink.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
