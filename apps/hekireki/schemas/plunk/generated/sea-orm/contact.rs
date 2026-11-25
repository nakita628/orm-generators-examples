use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "contacts")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub email: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub data: Option<Json>,
    #[sea_orm(default_value = true)]
    pub subscribed: bool,
    /// When set, `subscribed = false` is temporary rather than permanent: the snooze
    /// sweep resubscribes the contact once this passes. Suppression itself is still
    /// carried entirely by `subscribed`, so every send path filters snoozed contacts
    /// out without knowing this column exists.
    ///
    /// Cleared by every other write to `subscribed` — including the bounce and
    /// complaint paths, so a suppressed address can never be woken back up.
    #[sea_orm(column_name = "snoozedUntil", column_type = "custom(\"timestamp(3)\")")]
    pub snoozed_until: Option<DateTime>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
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
        fk_name = "contacts_projectId_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::email::Entity")]
    Emails,
    #[sea_orm(has_many = "super::workflow_execution::Entity")]
    WorkflowExecutions,
    #[sea_orm(has_many = "super::event::Entity")]
    Events,
    #[sea_orm(has_many = "super::segment_membership::Entity")]
    SegmentMemberships,
    #[sea_orm(has_many = "super::throttle_claim::Entity")]
    ThrottleClaims,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::email::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Emails.def()
    }
}

impl Related<super::workflow_execution::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkflowExecutions.def()
    }
}

impl Related<super::event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Events.def()
    }
}

impl Related<super::segment_membership::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SegmentMemberships.def()
    }
}

impl Related<super::throttle_claim::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ThrottleClaims.def()
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
