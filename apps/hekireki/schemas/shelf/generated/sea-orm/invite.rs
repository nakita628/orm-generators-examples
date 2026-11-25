use super::invite_statuses::InviteStatuses;
use super::organization_roles::OrganizationRoles;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Invite")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "inviterId", column_type = "Text")]
    pub inviter_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "inviteeUserId", column_type = "Text")]
    pub invitee_user_id: Option<String>,
    #[sea_orm(column_name = "teamMemberId", column_type = "Text")]
    pub team_member_id: String,
    #[sea_orm(column_name = "inviteeEmail", column_type = "Text")]
    pub invitee_email: String,
    #[sea_orm(default_value = "PENDING")]
    pub status: InviteStatuses,
    #[sea_orm(column_name = "inviteCode", column_type = "Text")]
    pub invite_code: String,
    pub roles: Vec<OrganizationRoles>,
    #[sea_orm(
        column_name = "inviteMessage",
        column_type = "String(StringLen::N(1000))"
    )]
    pub invite_message: Option<String>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
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
        belongs_to = "super::user::Entity",
        from = "Column::InviterId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Invite_inviterId_fkey"
    )]
    Inviter,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Invite_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::InviteeUserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Invite_inviteeUserId_fkey"
    )]
    InviteeUser,
    #[sea_orm(
        belongs_to = "super::team_member::Entity",
        from = "Column::TeamMemberId",
        to = "super::team_member::Column::Id",
        fk_name = "Invite_teamMemberId_fkey"
    )]
    InviteeTeamMember,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Inviter.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InviteeTeamMember.def()
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
