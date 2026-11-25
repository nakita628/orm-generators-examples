use super::currency::Currency;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// How a connected customer is invoiced (ADR-139, section 7): one payment
/// provider customer, one quarterly usage subscription in USD, and the
/// commercial terms the annual invoice was raised from.
///
/// One row per organization. Every write through it is addressed by a stored
/// provider id, so running onboarding or a renewal twice changes nothing.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ConnectedBillingAccount")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// The customer at the payment provider. Invoiced, never charged.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeCustomerId", column_type = "Text")]
    pub stripe_customer_id: String,
    /// The one metered subscription, and the item hosted usage is reported on.
    /// Null between the customer being created and the subscription existing,
    /// which is where a half-finished onboarding resumes from.
    #[sea_orm(column_name = "usageSubscriptionId", column_type = "Text")]
    pub usage_subscription_id: Option<String>,
    #[sea_orm(column_name = "usageSubscriptionItemId", column_type = "Text")]
    pub usage_subscription_item_id: Option<String>,
    #[sea_orm(column_name = "termStartsAt", column_type = "custom(\"timestamp(3)\")")]
    pub term_starts_at: DateTime,
    #[sea_orm(column_name = "termEndsAt", column_type = "custom(\"timestamp(3)\")")]
    pub term_ends_at: DateTime,
    /// Prepaid usage for the term, in USD cents. The credit grants add up to it.
    #[sea_orm(column_name = "commitUsdCents", default_value = 0)]
    pub commit_usd_cents: i32,
    /// The seat contract, which may be in another currency than hosted usage.
    #[sea_orm(column_name = "seatCurrency")]
    pub seat_currency: Currency,
    #[sea_orm(column_name = "seatRateCents")]
    pub seat_rate_cents: i32,
    pub seats: i32,
    /// `us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
    /// virtual bank account at the provider. Null when it wires to LangWatch's
    /// own account and finance marks the invoice paid.
    #[sea_orm(column_name = "bankTransferType", column_type = "Text")]
    pub bank_transfer_type: Option<String>,
    #[sea_orm(column_name = "bankTransferCountry", column_type = "Text")]
    pub bank_transfer_country: Option<String>,
    #[sea_orm(column_name = "billingEmail", column_type = "Text")]
    pub billing_email: String,
    /// A renewal whose credit is waiting for the last usage invoice of the old
    /// term to be finalized: the new term's end and commit, and who asked.
    #[sea_orm(column_name = "pendingRenewal", column_type = "JsonBinary")]
    pub pending_renewal: Option<Json>,
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: Option<String>,
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
        on_delete = "Restrict",
        fk_name = "ConnectedBillingAccount_organizationId_fkey"
    )]
    Organization,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
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
