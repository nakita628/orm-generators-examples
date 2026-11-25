use super::connected_credit_grant_kind::ConnectedCreditGrantKind;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A paid credit at the payment provider, which only metered usage draws down.
/// Kept here so a grant is created once per reason and the backoffice can show
/// what the customer holds without asking the provider.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ConnectedCreditGrant")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The `ConnectedBillingAccount` row. A plain column ref, like
    /// `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
    #[sea_orm(column_name = "accountId", column_type = "Text")]
    pub account_id: String,
    /// The grant at the payment provider. Unique, so a retry records one row.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeCreditGrantId", column_type = "Text")]
    pub stripe_credit_grant_id: String,
    #[sea_orm(column_name = "amountUsdCents")]
    pub amount_usd_cents: i32,
    pub kind: ConnectedCreditGrantKind,
    /// The term this credit was granted for, and when the provider lets it lapse:
    /// after the last invoice of that term has been finalized, never on the day
    /// the term ends.
    #[sea_orm(column_name = "termEndsAt", column_type = "custom(\"timestamp(3)\")")]
    pub term_ends_at: DateTime,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
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
        Ok(self)
    }
}
