use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ClippingClip")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "clippingId", column_type = "Text")]
    pub clipping_id: String,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_type = "Text")]
    pub content: String,
    pub start: f64,
    pub end: f64,
    #[sea_orm(column_name = "trimStart")]
    pub trim_start: Option<i32>,
    #[sea_orm(column_type = "Text", default_value = "pending")]
    pub status: String,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "mediaId", column_type = "Text")]
    pub media_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub path: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub thumbnail: Option<String>,
    #[sea_orm(column_name = "draftedAt", column_type = "custom(\"timestamp(3)\")")]
    pub drafted_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::clipping::Entity",
        from = "Column::ClippingId",
        to = "super::clipping::Column::Id",
        fk_name = "ClippingClip_clippingId_fkey"
    )]
    Clipping,
}

impl Related<super::clipping::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Clipping.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
