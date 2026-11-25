use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Team")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(column_name = "isPersonal", default_value = false)]
    pub is_personal: bool,
    #[sea_orm(column_name = "ownerUserId", column_type = "String(StringLen::N(255))")]
    pub owner_user_id: Option<String>,
    /// Department assignment inherited as the default for members and
    /// projects that have none of their own. Accounting-only.
    #[sea_orm(column_name = "departmentId", column_type = "Text")]
    pub department_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "Team_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::OwnerUserId",
        to = "super::user::Column::Id",
        fk_name = "Team_ownerUserId_fkey"
    )]
    OwnerUser,
    #[sea_orm(has_many = "super::team_user::Entity")]
    Members,
    #[sea_orm(has_many = "super::project::Entity")]
    Projects,
    #[sea_orm(has_many = "super::ingestion_source::Entity")]
    IngestionSources,
    #[sea_orm(has_many = "super::ai_tool_entry_team::Entity")]
    AiToolEntryTeams,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OwnerUser.def()
    }
}

impl Related<super::team_user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Members.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Projects.def()
    }
}

impl Related<super::ingestion_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::IngestionSources.def()
    }
}

impl Related<super::ai_tool_entry_team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AiToolEntryTeams.def()
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
