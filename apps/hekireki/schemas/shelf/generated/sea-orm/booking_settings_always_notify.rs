use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "_BookingSettingsAlwaysNotify")]
pub struct Model {
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "A",
        column_type = "Text"
    )]
    pub booking_settings_id: String,
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
        belongs_to = "super::booking_settings::Entity",
        from = "Column::BookingSettingsId",
        to = "super::booking_settings::Column::Id",
        fk_name = "_BookingSettingsAlwaysNotify_A_fkey"
    )]
    BookingSettings,
    #[sea_orm(
        belongs_to = "super::team_member::Entity",
        from = "Column::TeamMemberId",
        to = "super::team_member::Column::Id",
        fk_name = "_BookingSettingsAlwaysNotify_B_fkey"
    )]
    TeamMember,
}

impl ActiveModelBehavior for ActiveModel {}
