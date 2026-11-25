use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DigestItem")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "messageId", column_type = "Text")]
    pub message_id: String,
    #[sea_orm(column_name = "threadId", column_type = "Text")]
    pub thread_id: String,
    #[sea_orm(column_type = "Text")]
    pub content: String,
    #[sea_orm(column_name = "digestId", column_type = "Text")]
    pub digest_id: String,
    #[sea_orm(column_name = "actionId", column_type = "Text")]
    pub action_id: Option<String>,
    #[sea_orm(column_name = "coldEmailId", column_type = "Text")]
    pub cold_email_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::digest::Entity",
        from = "Column::DigestId",
        to = "super::digest::Column::Id",
        on_delete = "Cascade",
        fk_name = "DigestItem_digestId_fkey"
    )]
    Digest,
    #[sea_orm(
        belongs_to = "super::executed_action::Entity",
        from = "Column::ActionId",
        to = "super::executed_action::Column::Id",
        on_delete = "Cascade",
        fk_name = "DigestItem_actionId_fkey"
    )]
    Action,
    #[sea_orm(
        belongs_to = "super::cold_email::Entity",
        from = "Column::ColdEmailId",
        to = "super::cold_email::Column::Id",
        fk_name = "DigestItem_coldEmailId_fkey"
    )]
    ColdEmail,
}

impl Related<super::digest::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Digest.def()
    }
}

impl Related<super::executed_action::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Action.def()
    }
}

impl Related<super::cold_email::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ColdEmail.def()
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
