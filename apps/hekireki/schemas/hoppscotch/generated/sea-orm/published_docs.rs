use super::workspace_type::WorkspaceType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PublishedDocs")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_name = "collectionID", column_type = "Text")]
    pub collection_id: String,
    #[sea_orm(column_name = "creatorUid", column_type = "Text")]
    pub creator_uid: String,
    #[sea_orm(column_type = "Text")]
    pub version: String,
    #[sea_orm(column_name = "autoSync")]
    pub auto_sync: bool,
    #[sea_orm(column_name = "documentTree", column_type = "JsonBinary")]
    pub document_tree: Option<Json>,
    #[sea_orm(column_name = "workspaceType")]
    pub workspace_type: WorkspaceType,
    #[sea_orm(column_name = "workspaceID", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "environmentID", column_type = "Text")]
    pub environment_id: Option<String>,
    #[sea_orm(column_name = "environmentName", column_type = "Text")]
    pub environment_name: Option<String>,
    #[sea_orm(column_name = "environmentVariables", column_type = "JsonBinary")]
    pub environment_variables: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(
        column_name = "createdOn",
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_on: DateTimeWithTimeZone,
    #[sea_orm(column_name = "updatedOn", column_type = "custom(\"timestamptz(3)\")")]
    pub updated_on: DateTimeWithTimeZone,
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
        if insert && self.created_on.is_not_set() {
            self.created_on = Set(now.fixed_offset());
        }
        if !self.updated_on.is_set() {
            self.updated_on = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
