use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Public half of the user's capture key. The row outlives switching the
/// feature off: `disabledAt` stops a stale device re-registering that key.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "capture_keys")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub user_id: i32,
    /// Base64 SPKI DER of an ECDH P-256 public key.
    #[sea_orm(column_type = "Text")]
    pub public_key: String,
    #[sea_orm(column_type = "Text")]
    pub fingerprint: String,
    pub created_at: i64,
    pub disabled_at: Option<i64>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "capture_keys_user_id_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
