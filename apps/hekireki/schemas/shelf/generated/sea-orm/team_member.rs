use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TeamMember")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub name: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "deletedAt")]
    pub deleted_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull"
    )]
    User,
    #[sea_orm(has_many = "super::custody::Entity")]
    Custodies,
    #[sea_orm(has_many = "super::invite::Entity")]
    ReceivedInvites,
    #[sea_orm(has_many = "super::kit_custody::Entity")]
    KitCustodies,
    #[sea_orm(has_many = "super::booking::Entity")]
    Bookings,
    #[sea_orm(has_many = "super::consumption_log::Entity")]
    ConsumptionLogs,
    #[sea_orm(has_many = "super::team_member_note::Entity")]
    TeamMemberNotes,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::custody::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Custodies.def()
    }
}

impl Related<super::invite::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ReceivedInvites.def()
    }
}

impl Related<super::kit_custody::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::KitCustodies.def()
    }
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Bookings.def()
    }
}

impl Related<super::consumption_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ConsumptionLogs.def()
    }
}

impl Related<super::team_member_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TeamMemberNotes.def()
    }
}

impl Related<super::asset_reminder::Entity> for Entity {
    fn to() -> RelationDef {
        super::asset_reminder_to_team_member::Relation::AssetReminder.def()
    }
    fn via() -> Option<RelationDef> {
        Some(
            super::asset_reminder_to_team_member::Relation::TeamMember
                .def()
                .rev(),
        )
    }
}

impl Related<super::booking_settings::Entity> for Entity {
    fn to() -> RelationDef {
        super::booking_settings_always_notify::Relation::BookingSettings.def()
    }
    fn via() -> Option<RelationDef> {
        Some(
            super::booking_settings_always_notify::Relation::TeamMember
                .def()
                .rev(),
        )
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
