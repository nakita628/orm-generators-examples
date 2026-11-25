use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a feedback directory (Hub tenant) owned by an organization.
/// Directories group feedback data and are assigned to workspaces for access control.
///
/// @property id - Unique identifier for the directory
/// @property name - Display name of the directory
/// @property isArchived - Soft delete flag
/// @property organization - The parent organization
/// @property workspaces - Workspaces assigned to this directory
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "FeedbackDirectory")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "isArchived", default_value = false)]
    pub is_archived: bool,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "FeedbackDirectory_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::feedback_directory_workspace::Entity")]
    Workspaces,
    #[sea_orm(has_many = "super::feedback_source::Entity")]
    FeedbackSources,
    #[sea_orm(has_many = "super::chart::Entity")]
    Charts,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::feedback_directory_workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspaces.def()
    }
}

impl Related<super::feedback_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackSources.def()
    }
}

impl Related<super::chart::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Charts.def()
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
