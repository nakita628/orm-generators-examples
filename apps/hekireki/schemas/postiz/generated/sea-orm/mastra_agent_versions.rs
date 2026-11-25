use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_agent_versions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "agentId", column_type = "Text")]
    pub agent_id: String,
    #[sea_orm(column_name = "versionNumber")]
    pub version_number: i32,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub instructions: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub model: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub tools: Option<Json>,
    #[sea_orm(column_name = "defaultOptions", column_type = "JsonBinary")]
    pub default_options: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub workflows: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub agents: Option<Json>,
    #[sea_orm(column_name = "integrationTools", column_type = "JsonBinary")]
    pub integration_tools: Option<Json>,
    #[sea_orm(column_name = "inputProcessors", column_type = "JsonBinary")]
    pub input_processors: Option<Json>,
    #[sea_orm(column_name = "outputProcessors", column_type = "JsonBinary")]
    pub output_processors: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub memory: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub scorers: Option<Json>,
    #[sea_orm(column_name = "mcpClients", column_type = "JsonBinary")]
    pub mcp_clients: Option<Json>,
    #[sea_orm(column_name = "requestContextSchema", column_type = "JsonBinary")]
    pub request_context_schema: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub workspace: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub skills: Option<Json>,
    #[sea_orm(column_name = "skillsFormat", column_type = "Text")]
    pub skills_format: Option<String>,
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
    pub durable: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub browser: Option<Json>,
    #[sea_orm(column_name = "toolProviders", column_type = "JsonBinary")]
    pub tool_providers: Option<Json>,
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
