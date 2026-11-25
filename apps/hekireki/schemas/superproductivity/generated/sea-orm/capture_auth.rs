use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "capture_auth")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    pub user_id: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_type = "Text")]
    pub client_name: String,
    #[sea_orm(column_type = "Text")]
    pub redirect_uri: String,
    #[sea_orm(column_type = "Text")]
    pub code_challenge: String,
    #[sea_orm(column_type = "Text")]
    pub resource: String,
    #[sea_orm(column_type = "Text")]
    pub state: String,
    #[sea_orm(column_type = "Text")]
    pub scope: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub code_hash: Option<String>,
    pub token_version: Option<i32>,
    pub expires_at: i64,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "capture_auth_user_id_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
