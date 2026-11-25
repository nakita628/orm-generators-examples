use super::webhook_delivery_channel::WebhookDeliveryChannel;
use super::webhook_delivery_outcome::WebhookDeliveryOutcome;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The per-ATTEMPT delivery log for BOTH webhook channels: the endpoints
/// platform's batches and the automations channel's trigger fires. One row per
/// HTTP attempt; `dispatchId` groups every attempt of one logical send (the
/// value of the X-LangWatch-Delivery-Id or X-LangWatch-Event-Id header). A slim
/// facts table: outcome, status, latency, a capped error message, and a
/// truncated copy of a failure response. Our request content (URL, headers,
/// body) is never stored. Pruned after 30 days.
///
/// The two channels used to keep separate tables with the same columns and two
/// prune implementations. They are one table with a discriminator now, so the
/// tenancy columns of the channel that did not write a row are null.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WebhookEndpointDelivery")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(default_value = "platform")]
    pub channel: WebhookDeliveryChannel,
    /// Platform tenancy. Null on automations rows.
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "endpointId", column_type = "Text")]
    pub endpoint_id: Option<String>,
    /// Automations tenancy. Null on platform rows.
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_name = "triggerId", column_type = "Text")]
    pub trigger_id: Option<String>,
    #[sea_orm(column_name = "dispatchId", column_type = "Text")]
    pub dispatch_id: String,
    /// 1-based attempt and the batch's envelope count. Both null on automations
    /// rows, which record neither: one trigger fire is one event, and the
    /// channel's sender does not thread its attempt number into the log.
    pub attempt: Option<i32>,
    #[sea_orm(column_name = "eventCount")]
    pub event_count: Option<i32>,
    #[sea_orm(column_name = "responseStatus")]
    pub response_status: Option<i32>,
    #[sea_orm(column_name = "latencyMs")]
    pub latency_ms: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub response: Option<Json>,
    pub outcome: WebhookDeliveryOutcome,
    #[sea_orm(
        column_name = "firedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub fired_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::webhook_endpoint::Entity",
        from = "Column::EndpointId",
        to = "super::webhook_endpoint::Column::Id",
        on_delete = "Cascade",
        fk_name = "WebhookEndpointDelivery_endpointId_fkey"
    )]
    Endpoint,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "WebhookEndpointDelivery_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::trigger::Entity",
        from = "Column::TriggerId",
        to = "super::trigger::Column::Id",
        on_delete = "Cascade",
        fk_name = "WebhookEndpointDelivery_triggerId_fkey"
    )]
    Trigger,
}

impl Related<super::webhook_endpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Endpoint.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Trigger.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.fired_at.is_not_set() {
            self.fired_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
