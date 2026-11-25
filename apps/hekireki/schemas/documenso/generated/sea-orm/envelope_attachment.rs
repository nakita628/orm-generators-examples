use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZEnvelopeAttachmentTypeSchema } from '@documenso/lib/types/envelope-attachment';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "EnvelopeAttachment")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// [EnvelopeAttachmentType] @zod.custom.use(ZEnvelopeAttachmentTypeSchema)
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_type = "Text")]
    pub label: String,
    #[sea_orm(column_type = "Text")]
    pub data: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "envelopeId", column_type = "Text")]
    pub envelope_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::envelope::Entity",
        from = "Column::EnvelopeId",
        to = "super::envelope::Column::Id",
        on_delete = "Cascade",
        fk_name = "EnvelopeAttachment_envelopeId_fkey"
    )]
    Envelope,
}

impl Related<super::envelope::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Envelope.def()
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
