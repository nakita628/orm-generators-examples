use super::connected_invoice_kind::ConnectedInvoiceKind;
use super::currency::Currency;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// An invoice LangWatch raised for a connected customer, as the backoffice
/// reads it.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ConnectedInvoice")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The `ConnectedBillingAccount` row, as a plain column ref.
    #[sea_orm(column_name = "accountId", column_type = "Text")]
    pub account_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeInvoiceId", column_type = "Text")]
    pub stripe_invoice_id: String,
    pub kind: ConnectedInvoiceKind,
    pub currency: Currency,
    #[sea_orm(column_name = "amountCents")]
    pub amount_cents: i32,
    /// The provider's own status: `draft`, `open`, `paid`, `void` or `uncollectible`.
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(
        column_name = "paidOutOfBandAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub paid_out_of_band_at: Option<DateTime>,
    #[sea_orm(column_name = "termStartsAt", column_type = "custom(\"timestamp(3)\")")]
    pub term_starts_at: Option<DateTime>,
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
