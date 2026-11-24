use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Encrypted value of an HTTP request secret header (ENG-7168, ADR 0043). The definition and every version
/// store only `{name, secret: true, secretId}`; Replace overwrites this row in place, so the enabled
/// version uses the new value at once. Rows nothing references are deleted on Enable; all are deleted
/// with the workflow. A value is only sent to the origin it was set for (ADR 0051).
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WorkflowSecret")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    #[sea_orm(column_name = "workflowId")]
    pub workflow_id: String,
    #[sea_orm(column_name = "workspaceId")]
    pub workspace_id: String,
    /// `symmetricEncrypt` ciphertext (AES-256-GCM). Never logged, never returned.
    pub value: String,
    /// Origin (scheme, host, port) of the step's URL when the value was set or replaced; the value is only
    /// sent there (ADR 0051). Null for a value set before host binding: held to the step's saved URL
    /// until it is set again.
    #[sea_orm(column_name = "boundOrigin")]
    pub bound_origin: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "(Column::WorkflowId, Column::WorkspaceId)",
        to = "(super::workflow::Column::Id, super::workflow::Column::WorkspaceId)",
        on_delete = "Cascade"
    )]
    Workflow,
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflow.def()
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
