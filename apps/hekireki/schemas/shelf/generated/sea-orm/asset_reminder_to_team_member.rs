use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "_AssetReminderToTeamMember")]
pub struct Model {
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "A",
        column_type = "Text"
    )]
    pub asset_reminder_id: String,
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "B",
        column_type = "Text"
    )]
    pub team_member_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::asset_reminder::Entity",
        from = "Column::AssetReminderId",
        to = "super::asset_reminder::Column::Id",
        fk_name = "_AssetReminderToTeamMember_A_fkey"
    )]
    AssetReminder,
    #[sea_orm(
        belongs_to = "super::team_member::Entity",
        from = "Column::TeamMemberId",
        to = "super::team_member::Column::Id",
        fk_name = "_AssetReminderToTeamMember_B_fkey"
    )]
    TeamMember,
}

impl ActiveModelBehavior for ActiveModel {}
