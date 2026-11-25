use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A short code a fresh install pastes instead of a license blob (ADR-139,
/// section 5). Redeeming it mints the license the code describes and returns
/// it signed, which is what identifies a new customer without anybody emailing
/// a license key around.
///
/// The row holds a hash of the code, never the code. Every lookup hashes what
/// the caller presented first, so the stored value is a step away from anything
/// that could be presented: read access to this table is not credential access.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ActivationCode")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// SHA-256 of the normalised code. Looked up on every redemption.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "codeHash", column_type = "Text")]
    pub code_hash: String,
    /// The last four characters, so an operator tells two codes apart in a list
    /// without the list holding either one.
    #[sea_orm(column_name = "codeHint", column_type = "Text")]
    pub code_hint: String,
    /// The customer this code belongs to, and the terms it mints.
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "organizationName", column_type = "Text")]
    pub organization_name: String,
    #[sea_orm(column_type = "Text")]
    pub email: String,
    #[sea_orm(column_name = "planType", column_type = "Text")]
    pub plan_type: String,
    #[sea_orm(column_name = "maxMembers")]
    pub max_members: i32,
    /// The lite seats the minted license carries: null when unlimited. Never a stand-in number.
    #[sea_orm(column_name = "maxMembersLite", default_value = 0)]
    pub max_members_lite: Option<i32>,
    /// True when the minted license carries unlimited lite seats; maxMembersLite is then null.
    #[sea_orm(column_name = "liteSeatsUnlimited", default_value = false)]
    pub lite_seats_unlimited: bool,
    /// The Developer seats the minted license carries: null when not stated, which sells none.
    #[sea_orm(column_name = "maxMembersDeveloper")]
    pub max_members_developer: Option<i32>,
    /// How long the license this code mints runs for.
    #[sea_orm(column_name = "licenseTermDays")]
    pub license_term_days: i32,
    /// Hosted services the minted license may call.
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub services: Vec<String>,
    /// When the code stops working, which is not the term of what it mints.
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: DateTime,
    /// A code a customer may redeem on more than one install, for a rollout.
    #[sea_orm(default_value = false)]
    pub reusable: bool,
    /// Set by the redemption that won, which is what makes a single-use code
    /// single use: the write itself carries the condition.
    #[sea_orm(column_name = "redeemedAt", column_type = "custom(\"timestamp(3)\")")]
    pub redeemed_at: Option<DateTime>,
    #[sea_orm(column_name = "redeemedByInstanceId", column_type = "Text")]
    pub redeemed_by_instance_id: Option<String>,
    /// The license the winning redemption minted.
    #[sea_orm(column_name = "issuedLicenseId", column_type = "Text")]
    pub issued_license_id: Option<String>,
    /// How many installs have redeemed it, which only moves on a reusable code.
    #[sea_orm(column_name = "redemptionCount", default_value = 0)]
    pub redemption_count: i32,
    #[sea_orm(column_name = "revokedAt", column_type = "custom(\"timestamp(3)\")")]
    pub revoked_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedById", column_type = "Text")]
    pub revoked_by_id: Option<String>,
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: Option<String>,
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
