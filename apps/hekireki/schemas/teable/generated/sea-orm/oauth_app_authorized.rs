use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "oauth_app_authorized")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub authorized_time: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::oauth_app::Entity",
        from = "Column::ClientId",
        to = "super::oauth_app::Column::ClientId",
        on_delete = "Cascade",
        fk_name = "oauth_app_authorized_client_id_fkey"
    )]
    OAuthApp,
}

impl Related<super::oauth_app::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OAuthApp.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
