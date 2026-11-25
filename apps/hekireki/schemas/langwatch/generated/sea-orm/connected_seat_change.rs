use super::connected_seat_change_state::ConnectedSeatChangeState;
use super::currency::Currency;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// What a mid-term seat change of one connected license owes (ADR-141,
/// section 7): the seats added on the reissued license, invoiced prorated to
/// the end of the term.
///
/// The row is written before the payment provider is called and completed
/// after it, so a call that failed halfway is retried by the next seat
/// invoicing pass without inventing a second invoice. One row per reissued license.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ConnectedSeatChange")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The reissued `IssuedLicense` row that carries the new seat count. A
    /// plain column ref, like `replacesId`: relationMode="prisma" carries no
    /// foreign keys.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "licenseId", column_type = "Text")]
    pub license_id: String,
    /// The `ConnectedBillingAccount` row the invoice is raised against. Null
    /// when the customer had no billing account (`not_onboarded`).
    #[sea_orm(column_name = "accountId", column_type = "Text")]
    pub account_id: Option<String>,
    /// The customer organization the license belongs to.
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    /// When the seats changed, which is the day the proration counts from.
    #[sea_orm(column_name = "changedAt", column_type = "custom(\"timestamp(3)\")")]
    pub changed_at: DateTime,
    /// Seats on the new license above the seats on the one it replaced.
    #[sea_orm(column_name = "addedSeats")]
    pub added_seats: i32,
    /// What one added seat costs for the rest of the term, in cents of `currency`.
    #[sea_orm(column_name = "unitAmountCents")]
    pub unit_amount_cents: i32,
    /// `unitAmountCents` times `addedSeats`. Zero when nothing is owed.
    #[sea_orm(column_name = "amountCents")]
    pub amount_cents: i32,
    /// Null when no billing account named one (`not_onboarded`).
    pub currency: Option<Currency>,
    /// The one-off invoice the added seats were billed on.
    #[sea_orm(column_name = "stripeInvoiceId", column_type = "Text")]
    pub stripe_invoice_id: Option<String>,
    pub state: ConnectedSeatChangeState,
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
pub enum Relation {}

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
