use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Per-user tracking of notification views and dismissals
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PlatformNotificationInteraction")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "notificationId", column_type = "Text")]
    pub notification_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    /// Set by beacon/CLI GET on first impression
    #[sea_orm(
        column_name = "firstSeenAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub first_seen_at: DateTime,
    /// Times shown (incremented per beacon/CLI view)
    #[sea_orm(column_name = "showCount", default_value = 1)]
    pub show_count: i32,
    /// User dismissed in webapp
    #[sea_orm(
        column_name = "webappDismissedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub webapp_dismissed_at: Option<DateTime>,
    /// User clicked a link in webapp
    #[sea_orm(
        column_name = "webappClickedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub webapp_clicked_at: Option<DateTime>,
    /// Auto-dismissed or expired in CLI
    #[sea_orm(
        column_name = "cliDismissedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub cli_dismissed_at: Option<DateTime>,
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
        belongs_to = "super::platform_notification::Entity",
        from = "Column::NotificationId",
        to = "super::platform_notification::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "PlatformNotificationInteraction_notificationId_fkey"
    )]
    Notification,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "PlatformNotificationInteraction_userId_fkey"
    )]
    User,
}

impl Related<super::platform_notification::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notification.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.first_seen_at.is_not_set() {
            self.first_seen_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
