use super::organization_roles::OrganizationRoles;
use super::role_change_source::RoleChangeSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RoleChangeLog")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "previousRole")]
    pub previous_role: OrganizationRoles,
    #[sea_orm(column_name = "newRole")]
    pub new_role: OrganizationRoles,
    /// How the change was made. SSO rows name the member themself as `changedBy`.
    #[sea_orm(default_value = "MANUAL")]
    pub source: RoleChangeSource,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    #[sea_orm(column_name = "changedById")]
    pub changed_by_id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ChangedById",
        to = "super::user::Column::Id"
    )]
    ChangedBy,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id"
    )]
    Organization,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
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
        Ok(self)
    }
}
