use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SelectedCalendar")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(column_type = "Text")]
    pub integration: String,
    #[sea_orm(column_name = "externalId", column_type = "Text")]
    pub external_id: String,
    #[sea_orm(column_name = "credentialId")]
    pub credential_id: Option<i32>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: Option<DateTime>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
    #[sea_orm(column_name = "googleChannelId", column_type = "Text")]
    pub google_channel_id: Option<String>,
    #[sea_orm(column_name = "googleChannelKind", column_type = "Text")]
    pub google_channel_kind: Option<String>,
    #[sea_orm(column_name = "googleChannelResourceId", column_type = "Text")]
    pub google_channel_resource_id: Option<String>,
    #[sea_orm(column_name = "googleChannelResourceUri", column_type = "Text")]
    pub google_channel_resource_uri: Option<String>,
    #[sea_orm(column_name = "googleChannelExpiration", column_type = "Text")]
    pub google_channel_expiration: Option<String>,
    #[sea_orm(column_name = "channelId", column_type = "Text")]
    pub channel_id: Option<String>,
    #[sea_orm(column_name = "channelKind", column_type = "Text")]
    pub channel_kind: Option<String>,
    #[sea_orm(column_name = "channelResourceId", column_type = "Text")]
    pub channel_resource_id: Option<String>,
    #[sea_orm(column_name = "channelResourceUri", column_type = "Text")]
    pub channel_resource_uri: Option<String>,
    #[sea_orm(
        column_name = "channelExpiration",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub channel_expiration: Option<DateTime>,
    #[sea_orm(
        column_name = "syncSubscribedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub sync_subscribed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "syncSubscribedErrorAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub sync_subscribed_error_at: Option<DateTime>,
    #[sea_orm(column_name = "syncSubscribedErrorCount", default_value = 0)]
    pub sync_subscribed_error_count: i32,
    #[sea_orm(column_name = "syncToken", column_type = "Text")]
    pub sync_token: Option<String>,
    #[sea_orm(column_name = "syncedAt", column_type = "custom(\"timestamp(3)\")")]
    pub synced_at: Option<DateTime>,
    #[sea_orm(column_name = "syncErrorAt", column_type = "custom(\"timestamp(3)\")")]
    pub sync_error_at: Option<DateTime>,
    #[sea_orm(column_name = "syncErrorCount", default_value = 0)]
    pub sync_error_count: Option<i32>,
    #[sea_orm(column_name = "delegationCredentialId", column_type = "Text")]
    pub delegation_credential_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "lastErrorAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_error_at: Option<DateTime>,
    #[sea_orm(column_name = "watchAttempts", default_value = 0)]
    pub watch_attempts: i32,
    #[sea_orm(column_name = "unwatchAttempts", default_value = 0)]
    pub unwatch_attempts: i32,
    #[sea_orm(column_name = "maxAttempts", default_value = 3)]
    pub max_attempts: i32,
    #[sea_orm(column_name = "eventTypeId")]
    pub event_type_id: Option<i32>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "SelectedCalendar_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::credential::Entity",
        from = "Column::CredentialId",
        to = "super::credential::Column::Id",
        on_delete = "Cascade",
        fk_name = "SelectedCalendar_credentialId_fkey"
    )]
    Credential,
    #[sea_orm(
        belongs_to = "super::delegation_credential::Entity",
        from = "Column::DelegationCredentialId",
        to = "super::delegation_credential::Column::Id",
        on_delete = "Cascade",
        fk_name = "SelectedCalendar_delegationCredentialId_fkey"
    )]
    DelegationCredential,
    #[sea_orm(
        belongs_to = "super::event_type::Entity",
        from = "Column::EventTypeId",
        to = "super::event_type::Column::Id",
        fk_name = "SelectedCalendar_eventTypeId_fkey"
    )]
    EventType,
    #[sea_orm(has_many = "super::calendar_cache_event::Entity")]
    CalendarCacheEvents,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::credential::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Credential.def()
    }
}

impl Related<super::delegation_credential::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DelegationCredential.def()
    }
}

impl Related<super::event_type::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EventType.def()
    }
}

impl Related<super::calendar_cache_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CalendarCacheEvents.def()
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
            self.created_at = Set(Some(now.naive_utc()));
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
