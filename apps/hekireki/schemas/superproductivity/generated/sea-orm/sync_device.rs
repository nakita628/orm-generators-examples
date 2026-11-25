use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "sync_devices")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub user_id: i32,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_type = "Text")]
    pub device_name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub user_agent: Option<String>,
    /// Bare semver the client last reported on download (`18.22.0`), NULL until
    /// it does. Read by the checkpoint gate (#9962); never shown to users.
    #[sea_orm(column_type = "Text")]
    pub app_version: Option<String>,
    pub last_seen_at: i64,
    /// Reserved: written once at device registration, never read or updated.
    #[sea_orm(default_value = 0)]
    pub last_acked_seq: i32,
    pub created_at: i64,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "sync_devices_user_id_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
