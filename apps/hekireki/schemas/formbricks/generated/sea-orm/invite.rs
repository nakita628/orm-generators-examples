use super::organization_role::OrganizationRole;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents pending invitations to join an organization.
/// Used to manage the process of adding new users to an organization.
/// Once accepted, invites are converted into memberships.
///
/// @property id - Unique identifier for the invite
/// @property email - Email address of the invited user
/// @property name - Optional display name for the invited user
/// @property organization - The organization sending the invite
/// @property creator - The user who created the invite
/// @property acceptor - The user who accepted the invite (if accepted)
/// @property expiresAt - When the invite becomes invalid
/// @property role - Intended role for the invited user
/// @property teamIds - Teams the user will be added to upon acceptance
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Invite")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub email: String,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "creatorId", column_type = "Text")]
    pub creator_id: String,
    #[sea_orm(column_name = "acceptorId", column_type = "Text")]
    pub acceptor_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(default_value = "member")]
    pub role: OrganizationRole,
    #[sea_orm(
        column_name = "teamIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub team_ids: Vec<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "Invite_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatorId",
        to = "super::user::Column::Id",
        fk_name = "Invite_creatorId_fkey"
    )]
    Creator,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::AcceptorId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Invite_acceptorId_fkey"
    )]
    Acceptor,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creator.def()
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
        Ok(self)
    }
}
