use super::workflow_status::WorkflowStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Workflow")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    pub name: String,
    pub description: Option<String>,
    #[sea_orm(default_value = "draft")]
    pub status: WorkflowStatus,
    #[sea_orm(column_name = "workspaceId")]
    pub workspace_id: String,
    #[sea_orm(column_name = "createdBy")]
    pub created_by: Option<String>,
    /// [WorkflowDefinition]
    pub definition: Json,
    /// Webhook Secret URL token (ADRs 0048, 0057): SHA-256 hash to check calls, and an encrypted copy so
    /// the editor can show the URL again. Never part of the definition, its versions or its copies.
    #[sea_orm(column_name = "webhookTokenHash")]
    pub webhook_token_hash: Option<String>,
    #[sea_orm(column_name = "webhookTokenEncrypted")]
    pub webhook_token_encrypted: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull"
    )]
    Creator,
    #[sea_orm(has_many = "super::workflow_version::Entity")]
    Versions,
    #[sea_orm(has_many = "super::workflow_run::Entity")]
    Runs,
    #[sea_orm(has_many = "super::workflow_secret::Entity")]
    Secrets,
    #[sea_orm(has_many = "super::workflow_sample::Entity")]
    Samples,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creator.def()
    }
}

impl Related<super::workflow_version::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Versions.def()
    }
}

impl Related<super::workflow_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Runs.def()
    }
}

impl Related<super::workflow_secret::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Secrets.def()
    }
}

impl Related<super::workflow_sample::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Samples.def()
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
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
