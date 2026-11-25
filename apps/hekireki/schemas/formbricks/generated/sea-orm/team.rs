use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a team within an organization.
/// Enables group-based access control and collaboration.
///
/// @property id - Unique identifier for the team
/// @property name - Display name of the team
/// @property organization - The parent organization
/// @property teamUsers - Users who are part of this team
/// @property workspaceTeams - Workspaces this team has access to
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Team")]
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
        fk_name = "Team_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::team_user::Entity")]
    TeamUsers,
    #[sea_orm(has_many = "super::workspace_team::Entity")]
    WorkspaceTeams,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::team_user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TeamUsers.def()
    }
}

impl Related<super::workspace_team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkspaceTeams.def()
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
