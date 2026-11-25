use super::booking_status::BookingStatus;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BookingDenormalized")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: i32,
    #[sea_orm(column_type = "Text")]
    pub uid: String,
    #[sea_orm(column_name = "eventTypeId")]
    pub event_type_id: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "startTime", column_type = "custom(\"timestamp(3)\")")]
    pub start_time: DateTime,
    #[sea_orm(column_name = "endTime", column_type = "custom(\"timestamp(3)\")")]
    pub end_time: DateTime,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub location: Option<String>,
    pub paid: bool,
    pub status: BookingStatus,
    pub rescheduled: Option<bool>,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<i32>,
    #[sea_orm(column_name = "teamId")]
    pub team_id: Option<i32>,
    #[sea_orm(column_name = "eventLength")]
    pub event_length: Option<i32>,
    #[sea_orm(column_name = "eventParentId")]
    pub event_parent_id: Option<i32>,
    #[sea_orm(column_name = "userEmail", column_type = "Text")]
    pub user_email: Option<String>,
    #[sea_orm(column_name = "userName", column_type = "Text")]
    pub user_name: Option<String>,
    #[sea_orm(column_name = "userUsername", column_type = "Text")]
    pub user_username: Option<String>,
    #[sea_orm(column_name = "ratingFeedback", column_type = "Text")]
    pub rating_feedback: Option<String>,
    pub rating: Option<i32>,
    #[sea_orm(column_name = "noShowHost")]
    pub no_show_host: Option<bool>,
    #[sea_orm(column_name = "isTeamBooking")]
    pub is_team_booking: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
