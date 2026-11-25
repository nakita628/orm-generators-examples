use super::platform_notification_scope::PlatformNotificationScope;
use super::platform_notification_surface::PlatformNotificationSurface;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Admin-created notification definitions
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PlatformNotification")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    /// Admin-facing title for identification in admin tools
    #[sea_orm(column_type = "Text")]
    pub title: String,
    /// Versioned JSON rendering content (see payload schema in spec)
    #[sea_orm(column_type = "JsonBinary")]
    pub payload: Json,
    pub surface: PlatformNotificationSurface,
    pub scope: PlatformNotificationScope,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    /// When notification becomes active
    #[sea_orm(column_name = "startsAt", column_type = "custom(\"timestamp(3)\")")]
    pub starts_at: DateTime,
    /// When notification expires
    #[sea_orm(column_name = "endsAt", column_type = "custom(\"timestamp(3)\")")]
    pub ends_at: DateTime,
    /// CLI: auto-expire N days after user first saw it
    #[sea_orm(column_name = "cliMaxDaysAfterFirstSeen")]
    pub cli_max_days_after_first_seen: Option<i32>,
    /// CLI: auto-expire after shown N times to user
    #[sea_orm(column_name = "cliMaxShowCount")]
    pub cli_max_show_count: Option<i32>,
    /// CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…)
    #[sea_orm(column_name = "cliShowEvery")]
    pub cli_show_every: Option<i32>,
    /// Ordering within same scope level (higher = more important)
    #[sea_orm(default_value = 0)]
    pub priority: i32,
    /// Staged-but-unpublished. While true the notification is hidden from all
    /// user-facing reads regardless of startsAt/endsAt; publishing sets real
    /// dates and flips this to false.
    #[sea_orm(column_name = "isDraft", default_value = false)]
    pub is_draft: bool,
    /// Soft delete
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
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
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "PlatformNotification_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "PlatformNotification_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "PlatformNotification_projectId_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::platform_notification_interaction::Entity")]
    Interactions,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::platform_notification_interaction::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Interactions.def()
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
