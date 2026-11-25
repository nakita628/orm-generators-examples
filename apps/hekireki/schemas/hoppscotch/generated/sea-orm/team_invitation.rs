use super::team_access_role::TeamAccessRole;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TeamInvitation")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "teamID", column_type = "Text")]
    pub team_id: String,
    #[sea_orm(column_name = "creatorUid", column_type = "Text")]
    pub creator_uid: String,
    #[sea_orm(column_name = "inviteeEmail", column_type = "Text")]
    pub invitee_email: String,
    #[sea_orm(column_name = "inviteeRole")]
    pub invitee_role: TeamAccessRole,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "TeamInvitation_teamID_fkey"
    )]
    Team,
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
