use super::gateway_budget_breach_action::GatewayBudgetBreachAction;
use super::gateway_budget_scope_type::GatewayBudgetScopeType;
use super::gateway_budget_window::GatewayBudgetWindow;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "GatewayBudget")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "scopeType")]
    pub scope_type: GatewayBudgetScopeType,
    #[sea_orm(column_name = "scopeId", column_type = "Text")]
    pub scope_id: String,
    #[sea_orm(column_name = "providerKey", column_type = "Text")]
    pub provider_key: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    pub window: GatewayBudgetWindow,
    #[sea_orm(column_name = "limitUsd", column_type = "Decimal(Some((18, 6)))")]
    pub limit_usd: Decimal,
    #[sea_orm(column_name = "onBreach", default_value = "BLOCK")]
    pub on_breach: GatewayBudgetBreachAction,
    #[sea_orm(column_type = "Text")]
    pub timezone: Option<String>,
    #[sea_orm(column_name = "externalId", column_type = "Text")]
    pub external_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub metadata: Json,
    #[sea_orm(
        column_name = "spentUsd",
        column_type = "Decimal(Some((18, 6)))",
        default_value = 0
    )]
    pub spent_usd: Decimal,
    #[sea_orm(
        column_name = "currentPeriodStartedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub current_period_started_at: DateTime,
    #[sea_orm(column_name = "resetsAt", column_type = "custom(\"timestamp(3)\")")]
    pub resets_at: DateTime,
    #[sea_orm(column_name = "lastResetAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_reset_at: Option<DateTime>,
    #[sea_orm(
        column_name = "cycleAnchorAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub cycle_anchor_at: Option<DateTime>,
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
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: String,
    #[sea_orm(column_name = "managedByVirtualKeyId", column_type = "Text")]
    pub managed_by_virtual_key_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        fk_name = "GatewayBudget_createdById_fkey"
    )]
    CreatedBy,
    #[sea_orm(has_many = "super::gateway_budget_ledger::Entity")]
    LedgerEntries,
    #[sea_orm(has_many = "super::gateway_budget_bucket_boundary::Entity")]
    BucketBoundaries,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::gateway_budget_ledger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LedgerEntries.def()
    }
}

impl Related<super::gateway_budget_bucket_boundary::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BucketBoundaries.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.current_period_started_at.is_not_set() {
            self.current_period_started_at = Set(now.naive_utc());
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
