use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One transactional notification held for one recipient (or for the whole project) by a
/// send with `delivery.throttle`. While it is held, the same notification is not sent again.
///
/// The claim is taken by inserting (or taking over an expired) row, so the unique constraint
/// decides the race between concurrent sends, like `IdempotencyKey`. It is held from the
/// moment the send is accepted, re-anchored to the real send time once SES has the message,
/// and released when the email never gets there (cancelled, or failed before sending).
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "throttle_claims")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub key: String,
    #[sea_orm(column_name = "recipientKey", column_type = "Text")]
    pub recipient_key: String,
    #[sea_orm(column_name = "contactId", column_type = "Text")]
    pub contact_id: Option<String>,
    #[sea_orm(column_name = "emailId", column_type = "Text")]
    pub email_id: Option<String>,
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: Option<String>,
    #[sea_orm(column_name = "windowSeconds")]
    pub window_seconds: Option<i32>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
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
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "throttle_claims_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::contact::Entity",
        from = "Column::ContactId",
        to = "super::contact::Column::Id",
        on_delete = "Cascade",
        fk_name = "throttle_claims_contactId_fkey"
    )]
    Contact,
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "Column::WorkflowId",
        to = "super::workflow::Column::Id",
        on_delete = "Cascade",
        fk_name = "throttle_claims_workflowId_fkey"
    )]
    Workflow,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::contact::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Contact.def()
    }
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflow.def()
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
