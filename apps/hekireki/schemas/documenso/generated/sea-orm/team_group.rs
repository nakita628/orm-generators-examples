use super::team_member_role::TeamMemberRole;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TeamGroup")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organisationGroupId", column_type = "Text")]
    pub organisation_group_id: String,
    #[sea_orm(column_name = "teamRole")]
    pub team_role: TeamMemberRole,
    #[sea_orm(column_name = "teamId")]
    pub team_id: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organisation_group::Entity",
        from = "Column::OrganisationGroupId",
        to = "super::organisation_group::Column::Id",
        on_delete = "Cascade",
        fk_name = "TeamGroup_organisationGroupId_fkey"
    )]
    OrganisationGroup,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "TeamGroup_teamId_fkey"
    )]
    Team,
}

impl Related<super::organisation_group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrganisationGroup.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
