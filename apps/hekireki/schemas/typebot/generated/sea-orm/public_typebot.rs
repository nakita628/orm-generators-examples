use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PublicTypebot")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub version: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(
        column_name = "lastActivityAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_activity_at: Option<DateTime>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "typebotId", column_type = "Text")]
    pub typebot_id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub groups: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub events: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub variables: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub edges: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub theme: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub settings: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::typebot::Entity",
        from = "Column::TypebotId",
        to = "super::typebot::Column::Id",
        on_delete = "Cascade",
        fk_name = "PublicTypebot_typebotId_fkey"
    )]
    Typebot,
    #[sea_orm(has_many = "super::runtime_media_id_cache::Entity")]
    MediaIdsCache,
}

impl Related<super::typebot::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Typebot.def()
    }
}

impl Related<super::runtime_media_id_cache::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::MediaIdsCache.def()
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
        if insert && self.updated_at.is_not_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
