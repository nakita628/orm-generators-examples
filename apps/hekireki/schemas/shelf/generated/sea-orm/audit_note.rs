use super::note_type::NoteType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuditNote")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub content: String,
    #[sea_orm(default_value = "COMMENT")]
    pub r#type: NoteType,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "auditSessionId")]
    pub audit_session_id: String,
    #[sea_orm(column_name = "auditAssetId")]
    pub audit_asset_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::audit_session::Entity",
        from = "Column::AuditSessionId",
        to = "super::audit_session::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    AuditSession,
    #[sea_orm(
        belongs_to = "super::audit_asset::Entity",
        from = "Column::AuditAssetId",
        to = "super::audit_asset::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull"
    )]
    AuditAsset,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::audit_session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditSession.def()
    }
}

impl Related<super::audit_asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditAsset.def()
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
