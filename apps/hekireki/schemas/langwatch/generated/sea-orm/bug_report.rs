use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BugReport")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTimeWithTimeZone,
    #[sea_orm(column_type = "Text")]
    pub source: String,
    #[sea_orm(column_type = "Text")]
    pub kind: String,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_type = "Text")]
    pub summary: Option<String>,
    #[sea_orm(column_name = "sessionData", column_type = "Text")]
    pub session_data: Option<String>,
    #[sea_orm(column_name = "sessionTruncated", default_value = false)]
    pub session_truncated: bool,
    #[sea_orm(column_type = "Text")]
    pub agent: Option<String>,
    #[sea_orm(column_name = "contactEmail", column_type = "Text")]
    pub contact_email: Option<String>,
    #[sea_orm(column_name = "cliVersion", column_type = "Text")]
    pub cli_version: Option<String>,
    #[sea_orm(column_name = "linkedProjectId", column_type = "Text")]
    pub linked_project_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
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
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
