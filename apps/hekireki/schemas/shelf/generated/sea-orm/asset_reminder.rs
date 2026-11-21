use super::reminder_recurrence_unit::ReminderRecurrenceUnit;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AssetReminder")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub name: String,
    pub message: String,
    #[sea_orm(column_name = "alertDateTime")]
    pub alert_date_time: DateTime,
    #[sea_orm(column_name = "activeSchedulerReference")]
    pub active_scheduler_reference: Option<String>,
    #[sea_orm(column_name = "recurrenceUnit")]
    pub recurrence_unit: Option<ReminderRecurrenceUnit>,
    #[sea_orm(column_name = "recurrenceInterval")]
    pub recurrence_interval: Option<i32>,
    #[sea_orm(column_name = "recurrenceTimezone")]
    pub recurrence_timezone: Option<String>,
    #[sea_orm(column_name = "recurrenceEndsAt")]
    pub recurrence_ends_at: Option<DateTime>,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "assetId")]
    pub asset_id: String,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::asset::Entity",
        from = "Column::AssetId",
        to = "super::asset::Column::Id",
        on_delete = "Cascade"
    )]
    Asset,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id"
    )]
    CreatedBy,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Asset.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        super::asset_reminder_to_team_member::Relation::TeamMember.def()
    }
    fn via() -> Option<RelationDef> {
        Some(
            super::asset_reminder_to_team_member::Relation::AssetReminder
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
