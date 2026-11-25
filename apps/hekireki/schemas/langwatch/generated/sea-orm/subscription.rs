use super::plan_types::PlanTypes;
use super::subscription_status::SubscriptionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Subscription")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    pub plan: PlanTypes,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeSubscriptionId", column_type = "Text")]
    pub stripe_subscription_id: Option<String>,
    pub status: SubscriptionStatus,
    #[sea_orm(column_name = "startDate", column_type = "custom(\"timestamp(3)\")")]
    pub start_date: Option<DateTime>,
    #[sea_orm(column_name = "endDate", column_type = "custom(\"timestamp(3)\")")]
    pub end_date: Option<DateTime>,
    #[sea_orm(
        column_name = "lastPaymentFailedDate",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_payment_failed_date: Option<DateTime>,
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
    #[sea_orm(column_name = "maxMembers")]
    pub max_members: Option<i32>,
    #[sea_orm(column_name = "maxProjects")]
    pub max_projects: Option<i32>,
    #[sea_orm(column_name = "maxMessagesPerMonth")]
    pub max_messages_per_month: Option<i32>,
    #[sea_orm(column_name = "maxWorkflows")]
    pub max_workflows: Option<i32>,
    #[sea_orm(column_name = "maxRetentionDays")]
    pub max_retention_days: Option<i32>,
    #[sea_orm(column_name = "maxMembersLite")]
    pub max_members_lite: Option<i32>,
    /// Developer seats the organization pays for; null falls to the plan's own cap.
    #[sea_orm(column_name = "maxMembersDeveloper")]
    pub max_members_developer: Option<i32>,
    #[sea_orm(column_name = "maxTeams")]
    pub max_teams: Option<i32>,
    #[sea_orm(column_name = "maxPrompts")]
    pub max_prompts: Option<i32>,
    #[sea_orm(column_name = "maxEvaluators")]
    pub max_evaluators: Option<i32>,
    #[sea_orm(column_name = "maxScenarios")]
    pub max_scenarios: Option<i32>,
    #[sea_orm(column_name = "maxAgents")]
    pub max_agents: Option<i32>,
    #[sea_orm(column_name = "maxExperiments")]
    pub max_experiments: Option<i32>,
    #[sea_orm(column_name = "maxOnlineEvaluations")]
    pub max_online_evaluations: Option<i32>,
    #[sea_orm(column_name = "maxDatasets")]
    pub max_datasets: Option<i32>,
    #[sea_orm(column_name = "maxDashboards")]
    pub max_dashboards: Option<i32>,
    #[sea_orm(column_name = "maxCustomGraphs")]
    pub max_custom_graphs: Option<i32>,
    #[sea_orm(column_name = "maxAutomations")]
    pub max_automations: Option<i32>,
    /// Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil.
    #[sea_orm(column_name = "seatsKept")]
    pub seats_kept: Option<i32>,
    #[sea_orm(
        column_name = "seatsKeptUntil",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub seats_kept_until: Option<DateTime>,
    /// The seat quantity the last finished seat change set the seat line to; provider events never write it.
    #[sea_orm(column_name = "seatLineQuantity")]
    pub seat_line_quantity: Option<i32>,
    #[sea_orm(
        column_name = "seatLineQuantityAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub seat_line_quantity_at: Option<DateTime>,
    /// A Stripe seat line write recorded before it is sent and cleared once it is known to have
    /// finished under the lease; one left behind is settled by the seat line repair.
    /// Every seat quantity written while this record was pending.
    #[sea_orm(column_name = "seatLineWriteQuantities")]
    pub seat_line_write_quantities: Vec<i32>,
    /// The first write still pending; a later holder's write keeps it.
    #[sea_orm(
        column_name = "seatLineWriteSince",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub seat_line_write_since: Option<DateTime>,
    /// The latest write sent; the repair waits a grace window after it.
    #[sea_orm(
        column_name = "seatLineWriteAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub seat_line_write_at: Option<DateTime>,
    #[sea_orm(column_name = "seatLineWriteAttempts")]
    pub seat_line_write_attempts: Option<i32>,
    #[sea_orm(
        column_name = "seatLineWriteFlaggedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub seat_line_write_flagged_at: Option<DateTime>,
    /// The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled.
    #[sea_orm(column_name = "seatRepairQuantity")]
    pub seat_repair_quantity: Option<i32>,
    /// The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it.
    #[sea_orm(column_name = "seatRepairLiftedQuantity")]
    pub seat_repair_lifted_quantity: Option<i32>,
    #[sea_orm(
        column_name = "seatRepairSince",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub seat_repair_since: Option<DateTime>,
    #[sea_orm(column_name = "seatRepairAttempts")]
    pub seat_repair_attempts: Option<i32>,
    /// Set once the seat line repair has failed often enough that Ops should look at it.
    #[sea_orm(
        column_name = "seatRepairFlaggedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub seat_repair_flagged_at: Option<DateTime>,
    /// The lease a seat change or seat line repair holds while it reads and moves the seat line.
    #[sea_orm(column_name = "seatLineLeaseToken", column_type = "Text")]
    pub seat_line_lease_token: Option<String>,
    #[sea_orm(
        column_name = "seatLineLeaseUntil",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub seat_line_lease_until: Option<DateTime>,
    /// An open failed payment recovery case: the unpaid renewal invoice, day 0 (its first failed
    /// charge) and the last notice sent (3, 7, 14 or 21). Cleared when the case closes.
    #[sea_orm(column_name = "paymentRecoveryInvoiceId", column_type = "Text")]
    pub payment_recovery_invoice_id: Option<String>,
    #[sea_orm(
        column_name = "paymentRecoveryStartedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub payment_recovery_started_at: Option<DateTime>,
    #[sea_orm(column_name = "paymentRecoveryNoticeDay")]
    pub payment_recovery_notice_day: Option<i32>,
    /// The over-plan mark: traces recorded from overPlanSince are withheld while it is open and the
    /// organization is on Free with more members than Free allows; ended once that no longer holds.
    #[sea_orm(
        column_name = "overPlanSince",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub over_plan_since: Option<DateTime>,
    #[sea_orm(
        column_name = "overPlanEndedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub over_plan_ended_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "Subscription_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::invoice::Entity")]
    Invoices,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::invoice::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Invoices.def()
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
