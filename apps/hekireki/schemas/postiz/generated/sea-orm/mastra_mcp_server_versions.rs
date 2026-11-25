use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_mcp_server_versions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "mcpServerId", column_type = "Text")]
    pub mcp_server_id: String,
    #[sea_orm(column_name = "versionNumber")]
    pub version_number: i32,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub version: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub instructions: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub repository: Option<Json>,
    #[sea_orm(column_name = "releaseDate", column_type = "Text")]
    pub release_date: Option<String>,
    #[sea_orm(column_name = "isLatest")]
    pub is_latest: Option<bool>,
    #[sea_orm(column_name = "packageCanonical", column_type = "Text")]
    pub package_canonical: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub tools: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub agents: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub workflows: Option<Json>,
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
