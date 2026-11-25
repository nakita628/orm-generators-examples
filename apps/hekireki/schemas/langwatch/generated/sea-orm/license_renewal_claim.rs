use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// One renewed licence per paid Stripe invoice of a self-serve licence
/// subscription (specs/licensing/self-serving-license-purchase.feature). A
/// redelivered invoice event finds the row and mints nothing. A claim never
/// delivered is taken over once `claimedAt` is older than the claim lease.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LicenseRenewalClaim")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "stripeInvoiceId", column_type = "Text")]
    pub stripe_invoice_id: String,
    #[sea_orm(column_name = "stripeSubscriptionId", column_type = "Text")]
    pub stripe_subscription_id: String,
    /// Fences completion and release to the delivery that holds the claim.
    #[sea_orm(column_name = "claimToken", column_type = "Text")]
    pub claim_token: String,
    #[sea_orm(column_name = "claimedAt", column_type = "custom(\"timestamp(3)\")")]
    pub claimed_at: DateTime,
    /// The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered.
    #[sea_orm(column_name = "issuedLicense", column_type = "JsonBinary")]
    pub issued_license: Option<Json>,
    #[sea_orm(column_name = "deliveredAt", column_type = "custom(\"timestamp(3)\")")]
    pub delivered_at: Option<DateTime>,
    /// The `licenseId` inside the delivered licence.
    #[sea_orm(column_name = "licenseId", column_type = "Text")]
    pub license_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
