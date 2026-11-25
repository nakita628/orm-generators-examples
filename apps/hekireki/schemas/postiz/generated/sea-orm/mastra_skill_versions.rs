use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_skill_versions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "skillId", column_type = "Text")]
    pub skill_id: String,
    #[sea_orm(column_name = "versionNumber")]
    pub version_number: i32,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: String,
    #[sea_orm(column_type = "Text")]
    pub instructions: String,
    #[sea_orm(column_type = "Text")]
    pub license: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub compatibility: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub source: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub references: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub scripts: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub assets: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub tree: Option<Json>,
    #[sea_orm(column_name = "changedFields", column_type = "JsonBinary")]
    pub changed_fields: Option<Json>,
    #[sea_orm(column_name = "changeMessage", column_type = "Text")]
    pub change_message: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "createdAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_type = "JsonBinary")]
    pub files: Option<Json>,
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
        if insert && self.created_at_z.is_not_set() {
            self.created_at_z = Set(Some(now.fixed_offset()));
        }
        Ok(self)
    }
}
