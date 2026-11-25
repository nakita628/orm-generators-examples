use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Capture tokens and OAuth grants store only hashes of bearer credentials.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "capture_grants")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    pub user_id: i32,
    #[sea_orm(column_type = "Text")]
    pub kind: String,
    #[sea_orm(column_type = "Text")]
    pub client_name: String,
    #[sea_orm(column_type = "Text")]
    pub secret_hash: String,
    pub created_at: i64,
    pub expires_at: i64,
    pub last_used_at: Option<i64>,
    pub revoked_at: Option<i64>,
    #[sea_orm(column_type = "Text")]
    pub client_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub resource: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub refresh_hash: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub prev_refresh_hash: Option<String>,
    pub rotated_at: Option<i64>,
    pub refresh_expires_at: Option<i64>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "capture_grants_user_id_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
