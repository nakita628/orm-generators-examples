use super::chat_provider::ChatProvider;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RuntimeMediaIdCache")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "publicTypebotId", column_type = "Text")]
    pub public_typebot_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    pub provider: ChatProvider,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub url: String,
    #[sea_orm(column_name = "mediaId", column_type = "Text")]
    pub media_id: String,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::public_typebot::Entity",
        from = "Column::PublicTypebotId",
        to = "super::public_typebot::Column::Id",
        on_delete = "Cascade",
        fk_name = "RuntimeMediaIdCache_publicTypebotId_fkey"
    )]
    PublicTypebot,
}

impl Related<super::public_typebot::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PublicTypebot.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
