use super::document_storage_type::DocumentStorageType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AgreementResponse")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "agreementId", column_type = "Text")]
    pub agreement_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "viewId", column_type = "Text")]
    pub view_id: Option<String>,
    #[sea_orm(column_name = "linkId", column_type = "Text")]
    pub link_id: Option<String>,
    #[sea_orm(column_name = "signerEmail", column_type = "Text")]
    pub signer_email: Option<String>,
    #[sea_orm(column_name = "signerName", column_type = "Text")]
    pub signer_name: Option<String>,
    #[sea_orm(
        column_name = "signingStatus",
        column_type = "Text",
        default_value = "PENDING"
    )]
    pub signing_status: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "signingExternalId", column_type = "Text")]
    pub signing_external_id: Option<String>,
    #[sea_orm(column_name = "signingEnvelopeId", column_type = "Text")]
    pub signing_envelope_id: Option<String>,
    #[sea_orm(column_name = "signingDocumentId")]
    pub signing_document_id: Option<i32>,
    #[sea_orm(column_name = "signedAt", column_type = "custom(\"timestamp(3)\")")]
    pub signed_at: Option<DateTime>,
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(column_name = "signedFileKey", column_type = "Text")]
    pub signed_file_key: Option<String>,
    #[sea_orm(column_name = "signedFileName", column_type = "Text")]
    pub signed_file_name: Option<String>,
    #[sea_orm(column_name = "signedFileStorageType")]
    pub signed_file_storage_type: Option<DocumentStorageType>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::agreement::Entity",
        from = "Column::AgreementId",
        to = "super::agreement::Column::Id",
        on_delete = "Cascade",
        fk_name = "AgreementResponse_agreementId_fkey"
    )]
    Agreement,
    #[sea_orm(
        belongs_to = "super::view::Entity",
        from = "Column::ViewId",
        to = "super::view::Column::Id",
        on_delete = "Cascade",
        fk_name = "AgreementResponse_viewId_fkey"
    )]
    View,
}

impl Related<super::agreement::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Agreement.def()
    }
}

impl Related<super::view::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::View.def()
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
