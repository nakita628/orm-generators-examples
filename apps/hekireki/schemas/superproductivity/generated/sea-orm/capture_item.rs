use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// One sealed capture waiting for a device. No relation to its grant: items
/// outlive a grant that is revoked without discarding them.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "capture_items")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    pub user_id: i32,
    #[sea_orm(column_type = "Text")]
    pub grant_id: String,
    #[sea_orm(column_type = "Text")]
    pub ciphertext: String,
    /// Task id the claiming device must use; replaced when another device takes over.
    #[sea_orm(column_type = "Text")]
    pub task_id: String,
    pub created_at: i64,
    #[sea_orm(column_type = "Text")]
    pub lease_device_id: Option<String>,
    pub lease_until: Option<i64>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "capture_items_user_id_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
