use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NotificationState")]
pub struct Model {
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "channelId", column_type = "Text")]
    pub channel_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "actorId", column_type = "Text")]
    pub actor_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "tableName", column_type = "Text")]
    pub table_name: String,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(
        column_name = "flappingSince",
        column_type = "custom(\"timestamptz(3)\")"
    )]
    pub flapping_since: Option<DateTimeWithTimeZone>,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(
        column_name = "lastNotification",
        column_type = "custom(\"timestamptz(3)\")"
    )]
    pub last_notification: DateTimeWithTimeZone,
    #[sea_orm(column_name = "statusChangeId")]
    pub status_change_id: i64,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
