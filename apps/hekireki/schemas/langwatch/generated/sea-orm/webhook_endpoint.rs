use super::webhook_destination_kind::WebhookDestinationKind;
use super::webhook_endpoint_status::WebhookEndpointStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Org-anchored outbound webhook endpoint (the webhook platform, enterprise
/// only). The signing secret is AES-256-GCM encrypted at rest (shared
/// encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
/// once, at create or roll time. `enabledEvents` holds registry event types
/// or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
/// 72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WebhookEndpoint")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// Where this endpoint delivers. Everything above the last hop is shared:
    /// same batching, same retry ladder, same delivery log, same signature over
    /// the same bytes.
    #[sea_orm(column_name = "destinationKind", default_value = "http")]
    pub destination_kind: WebhookDestinationKind,
    /// The receiver URL. Present on `http` endpoints and null on the others;
    /// the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
    #[sea_orm(column_type = "Text")]
    pub url: Option<String>,
    /// The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
    /// region and the owning account are read off it rather than configured
    /// separately, so they can never disagree with the queue.
    #[sea_orm(column_name = "sqsQueueUrl", column_type = "Text")]
    pub sqs_queue_url: Option<String>,
    /// The role we assume to write to the queue, and the ExternalId that proves
    /// the request came from us rather than from anyone who learned the role's
    /// name. The recommended credential mode.
    #[sea_orm(column_name = "sqsRoleArn", column_type = "Text")]
    pub sqs_role_arn: Option<String>,
    #[sea_orm(column_name = "sqsExternalId", column_type = "Text")]
    pub sqs_external_id: Option<String>,
    /// A static key pair for the queue, encrypted at rest exactly like the
    /// signing secret. Both halves are present or both are absent: a half-filled
    /// pair is what the AWS SDK reads as "we have credentials" and stops looking
    /// with, so the CHECK constraint refuses it.
    #[sea_orm(column_name = "sqsAccessKeyId", column_type = "Text")]
    pub sqs_access_key_id: Option<String>,
    #[sea_orm(column_name = "sqsSecretAccessKeyEncrypted", column_type = "Text")]
    pub sqs_secret_access_key_encrypted: Option<String>,
    #[sea_orm(column_name = "secretEncrypted", column_type = "Text")]
    pub secret_encrypted: String,
    /// How deliveries are formatted and signed. Null: the batch envelope signed
    /// `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
    /// only for endpoints migrated from governance anomaly destinations.
    #[sea_orm(column_name = "signatureScheme", column_type = "Text")]
    pub signature_scheme: Option<String>,
    /// Set only by governance's anomaly destination migration: a create that
    /// repeats it in the same organization answers the endpoint it made. An
    /// archived endpoint gives its key up, so a retry after cleanup creates anew.
    #[sea_orm(column_name = "idempotencyKey", column_type = "Text")]
    pub idempotency_key: Option<String>,
    /// The secret this endpoint signed with before the last roll. Deliveries
    /// carry a second `v1=` computed from it until it expires, so a receiver
    /// swaps secrets on its own schedule instead of dropping deliveries in the
    /// gap. Null once the window closes or if the secret was never rolled.
    #[sea_orm(column_name = "previousSecretEncrypted", column_type = "Text")]
    pub previous_secret_encrypted: Option<String>,
    /// When the previous secret stops being signed with and accepted.
    #[sea_orm(
        column_name = "previousSecretExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub previous_secret_expires_at: Option<DateTime>,
    #[sea_orm(
        column_name = "enabledEvents",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub enabled_events: Vec<String>,
    #[sea_orm(default_value = "ACTIVE")]
    pub status: WebhookEndpointStatus,
    /// Why the endpoint is disabled: "auto_failures_72h" or "manual".
    #[sea_orm(column_name = "disabledReason", column_type = "Text")]
    pub disabled_reason: Option<String>,
    #[sea_orm(column_name = "disabledAt", column_type = "custom(\"timestamp(3)\")")]
    pub disabled_at: Option<DateTime>,
    /// Start of the current consecutive-failure streak; cleared on any 2xx.
    /// 72 hours of unbroken failures from this instant auto-disables.
    #[sea_orm(column_name = "failingSince", column_type = "custom(\"timestamp(3)\")")]
    pub failing_since: Option<DateTime>,
    /// Delivery controls, each server-bounded (see webhookDeliveryControls
    /// constants): envelopes per POST, coalescing window before a partial
    /// batch ships, and concurrent POSTs to this receiver.
    #[sea_orm(column_name = "maxBatchSize", default_value = 100)]
    pub max_batch_size: i32,
    #[sea_orm(column_name = "maxBatchDelayMs", default_value = 250)]
    pub max_batch_delay_ms: i32,
    #[sea_orm(column_name = "maxInFlight", default_value = 4)]
    pub max_in_flight: i32,
    /// Self-hosted only: deliveries accept the receiver's self-signed TLS
    /// certificate. Off by default; the hosted product ignores it.
    #[sea_orm(column_name = "allowSelfSignedCertificate", default_value = false)]
    pub allow_self_signed_certificate: bool,
    #[sea_orm(
        column_name = "lastSuccessAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_success_at: Option<DateTime>,
    #[sea_orm(
        column_name = "lastFailureAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_failure_at: Option<DateTime>,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "WebhookEndpoint_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::webhook_endpoint_delivery::Entity")]
    Deliveries,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::webhook_endpoint_delivery::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Deliveries.def()
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
