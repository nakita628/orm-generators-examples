use super::proration_status::ProrationStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "MonthlyProration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "teamId")]
    pub team_id: i32,
    #[sea_orm(column_name = "monthKey", column_type = "Text")]
    pub month_key: String,
    #[sea_orm(column_name = "periodStart", column_type = "custom(\"timestamp(3)\")")]
    pub period_start: DateTime,
    #[sea_orm(column_name = "periodEnd", column_type = "custom(\"timestamp(3)\")")]
    pub period_end: DateTime,
    #[sea_orm(column_name = "seatsAtStart")]
    pub seats_at_start: i32,
    #[sea_orm(column_name = "seatsAdded")]
    pub seats_added: i32,
    #[sea_orm(column_name = "seatsRemoved")]
    pub seats_removed: i32,
    #[sea_orm(column_name = "netSeatIncrease")]
    pub net_seat_increase: i32,
    #[sea_orm(column_name = "seatsAtEnd")]
    pub seats_at_end: i32,
    #[sea_orm(column_name = "subscriptionId", column_type = "Text")]
    pub subscription_id: String,
    #[sea_orm(column_name = "subscriptionItemId", column_type = "Text")]
    pub subscription_item_id: String,
    #[sea_orm(column_name = "customerId", column_type = "Text")]
    pub customer_id: String,
    #[sea_orm(
        column_name = "subscriptionStart",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub subscription_start: DateTime,
    #[sea_orm(
        column_name = "subscriptionEnd",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub subscription_end: DateTime,
    #[sea_orm(column_name = "remainingDays")]
    pub remaining_days: i32,
    #[sea_orm(column_name = "pricePerSeat")]
    pub price_per_seat: i32,
    #[sea_orm(column_name = "proratedAmount")]
    pub prorated_amount: i32,
    #[sea_orm(column_name = "invoiceItemId", column_type = "Text")]
    pub invoice_item_id: Option<String>,
    #[sea_orm(column_name = "invoiceId", column_type = "Text")]
    pub invoice_id: Option<String>,
    #[sea_orm(column_name = "invoiceUrl", column_type = "Text")]
    pub invoice_url: Option<String>,
    #[sea_orm(default_value = "PENDING")]
    pub status: ProrationStatus,
    #[sea_orm(column_name = "chargedAt", column_type = "custom(\"timestamp(3)\")")]
    pub charged_at: Option<DateTime>,
    #[sea_orm(column_name = "failedAt", column_type = "custom(\"timestamp(3)\")")]
    pub failed_at: Option<DateTime>,
    #[sea_orm(column_name = "failureReason", column_type = "Text")]
    pub failure_reason: Option<String>,
    #[sea_orm(column_name = "retryCount", default_value = 0)]
    pub retry_count: i32,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "teamBillingId", column_type = "Text")]
    pub team_billing_id: Option<String>,
    #[sea_orm(column_name = "organizationBillingId", column_type = "Text")]
    pub organization_billing_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "MonthlyProration_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::team_billing::Entity",
        from = "Column::TeamBillingId",
        to = "super::team_billing::Column::Id",
        fk_name = "MonthlyProration_teamBillingId_fkey"
    )]
    TeamBilling,
    #[sea_orm(
        belongs_to = "super::organization_billing::Entity",
        from = "Column::OrganizationBillingId",
        to = "super::organization_billing::Column::Id",
        fk_name = "MonthlyProration_organizationBillingId_fkey"
    )]
    OrganizationBilling,
    #[sea_orm(has_many = "super::seat_change_log::Entity")]
    SeatChanges,
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::team_billing::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TeamBilling.def()
    }
}

impl Related<super::organization_billing::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrganizationBilling.def()
    }
}

impl Related<super::seat_change_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SeatChanges.def()
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
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
