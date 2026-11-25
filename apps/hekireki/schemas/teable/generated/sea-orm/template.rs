use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "template")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub base_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub cover: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub markdown_description: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub category_id: Vec<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_time: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_modified_time: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub last_modified_by: Option<String>,
    pub is_system: Option<bool>,
    pub is_published: Option<bool>,
    pub featured: Option<bool>,
    #[sea_orm(column_type = "Text")]
    pub snapshot: Option<String>,
    pub order: f64,
    #[sea_orm(default_value = 0)]
    pub usage_count: i32,
    #[sea_orm(column_type = "JsonBinary")]
    pub publish_info: Option<Json>,
    #[sea_orm(default_value = 0)]
    pub visit_count: i32,
    #[sea_orm(column_type = "Text")]
    pub kind: Option<String>,
    #[sea_orm(default_value = 1)]
    pub version: i32,
    #[sea_orm(column_type = "JsonBinary")]
    pub manifest: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub review: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub history: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub parent_template_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub slug: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_time.is_not_set() {
            self.created_time = Set(now.naive_utc());
        }
        if !self.last_modified_time.is_set() {
            self.last_modified_time = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
