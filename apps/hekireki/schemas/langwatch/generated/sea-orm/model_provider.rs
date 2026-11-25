use super::gateway_provider_health_status::GatewayProviderHealthStatus;
use super::gateway_provider_rotation_policy::GatewayProviderRotationPolicy;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ModelProvider")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub provider: String,
    #[sea_orm(column_name = "routingHandle", column_type = "Text")]
    pub routing_handle: Option<String>,
    pub enabled: bool,
    #[sea_orm(column_name = "customKeys", column_type = "JsonBinary")]
    pub custom_keys: Option<Json>,
    #[sea_orm(column_name = "extraHeaders", column_type = "JsonBinary")]
    pub extra_headers: Option<Json>,
    #[sea_orm(column_name = "customModels", column_type = "JsonBinary")]
    pub custom_models: Option<Json>,
    #[sea_orm(column_name = "customEmbeddingsModels", column_type = "JsonBinary")]
    pub custom_embeddings_models: Option<Json>,
    #[sea_orm(column_name = "deploymentMapping", column_type = "JsonBinary")]
    pub deployment_mapping: Option<Json>,
    #[sea_orm(column_name = "rateLimitRpm")]
    pub rate_limit_rpm: Option<i32>,
    #[sea_orm(column_name = "rateLimitTpm")]
    pub rate_limit_tpm: Option<i32>,
    #[sea_orm(column_name = "rateLimitRpd")]
    pub rate_limit_rpd: Option<i32>,
    #[sea_orm(column_name = "rotationPolicy", default_value = "MANUAL")]
    pub rotation_policy: GatewayProviderRotationPolicy,
    #[sea_orm(column_name = "providerConfig", column_type = "JsonBinary")]
    pub provider_config: Option<Json>,
    #[sea_orm(column_name = "fallbackPriorityGlobal")]
    pub fallback_priority_global: Option<i32>,
    #[sea_orm(column_name = "langySkipPermissionsModels", column_type = "JsonBinary")]
    pub langy_skip_permissions_models: Option<Json>,
    #[sea_orm(column_name = "healthStatus", default_value = "UNKNOWN")]
    pub health_status: GatewayProviderHealthStatus,
    #[sea_orm(
        column_name = "circuitOpenedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub circuit_opened_at: Option<DateTime>,
    #[sea_orm(
        column_name = "lastHealthCheckAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_health_check_at: Option<DateTime>,
    #[sea_orm(column_name = "disabledAt", column_type = "custom(\"timestamp(3)\")")]
    pub disabled_at: Option<DateTime>,
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
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::model_provider_scope::Entity")]
    Scopes,
}

impl Related<super::model_provider_scope::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scopes.def()
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
