use super::currency::Currency;
use super::issued_license_source::IssuedLicenseSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The registry of licenses LangWatch has issued (ADR-139). It lives on
/// LangWatch Cloud; a self-hosted install has the table and no rows.
///
/// The row never holds the token an install presents. `tokenHash` is the
/// SHA-256 of that token, so read access to this table is not credential
/// access. Status is derived: revoked, superseded, past `expiresAt`, else active.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "IssuedLicense")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// The `licenseId` inside the signed license (`lic-<uuid>`).
    #[sea_orm(unique)]
    #[sea_orm(column_name = "licenseId")]
    pub license_id: String,
    /// SHA-256 of the `lwl_` token. Looked up on every credential resolution.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "tokenHash")]
    pub token_hash: String,
    /// The customer organization on LangWatch Cloud. Null for a license that is
    /// recorded but not linked to a customer yet, such as a self-serve purchase:
    /// it resolves to nothing until an operator links it.
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: Option<String>,
    /// As signed into the license.
    #[sea_orm(column_name = "organizationName")]
    pub organization_name: String,
    pub email: String,
    #[sea_orm(column_name = "planType")]
    pub plan_type: String,
    #[sea_orm(column_name = "maxMembers")]
    pub max_members: i32,
    /// The lite seats sold: null when not stated or when unlimited. Never a stand-in number.
    #[sea_orm(column_name = "maxMembersLite")]
    pub max_members_lite: Option<i32>,
    /// True when the license sells unlimited lite seats; maxMembersLite is then null.
    #[sea_orm(column_name = "liteSeatsUnlimited", default_value = false)]
    pub lite_seats_unlimited: bool,
    /// The Developer seats sold: null when not stated, which sells none.
    #[sea_orm(column_name = "maxMembersDeveloper")]
    pub max_members_developer: Option<i32>,
    #[sea_orm(column_name = "issuedAt")]
    pub issued_at: DateTime,
    #[sea_orm(column_name = "expiresAt")]
    pub expires_at: DateTime,
    pub source: IssuedLicenseSource,
    /// The operator who issued or registered it. Null for the purchase flow and the script.
    #[sea_orm(column_name = "issuedById")]
    pub issued_by_id: Option<String>,
    #[sea_orm(column_name = "revokedAt")]
    pub revoked_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedById")]
    pub revoked_by_id: Option<String>,
    #[sea_orm(column_name = "revokedReason")]
    pub revoked_reason: Option<String>,
    /// Set when the install first presents the license that replaced this one.
    #[sea_orm(column_name = "supersededAt")]
    pub superseded_at: Option<DateTime>,
    /// The id of the IssuedLicense this one replaces, when it was reissued. A
    /// plain column ref (relationMode="prisma" cannot declare a self-relation on
    /// Postgres); the registry service validates it. Unique, so a license is
    /// replaced by one license at most.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "replacesId")]
    pub replaces_id: Option<String>,
    /// A reissued license waiting to be delivered over sync, encrypted. Erased
    /// when the install first presents it. The only license text the registry holds.
    #[sea_orm(column_name = "pendingDeliveryLicense", column_type = "Text")]
    pub pending_delivery_license: Option<String>,
    /// Hosted services this license may call, such as `instant_evals`.
    pub services: Vec<String>,
    /// The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats.
    #[sea_orm(column_name = "seatRateCents")]
    pub seat_rate_cents: Option<i32>,
    #[sea_orm(column_name = "seatCurrency")]
    pub seat_currency: Option<Currency>,
    /// The prepaid usage commit for the term, in USD cents.
    #[sea_orm(column_name = "commitUsdCents", default_value = 0)]
    pub commit_usd_cents: i32,
    /// Whether usage may continue past the commit, up to `overageMaxUsdCents`.
    #[sea_orm(column_name = "overageEnabled", default_value = false)]
    pub overage_enabled: bool,
    #[sea_orm(column_name = "overageMaxUsdCents")]
    pub overage_max_usd_cents: Option<i32>,
    /// The install this license is bound to. Set on first use, cleared by an operator.
    #[sea_orm(column_name = "instanceId")]
    pub instance_id: Option<String>,
    #[sea_orm(column_name = "instanceBoundAt")]
    pub instance_bound_at: Option<DateTime>,
    /// The last license sync this install reported: when it arrived, the
    /// LangWatch version that sent it, and the seats in use at that moment.
    /// Null until the install syncs for the first time.
    #[sea_orm(column_name = "lastSyncAt")]
    pub last_sync_at: Option<DateTime>,
    #[sea_orm(column_name = "lastSyncVersion")]
    pub last_sync_version: Option<String>,
    #[sea_orm(column_name = "reportedMembers")]
    pub reported_members: Option<i32>,
    #[sea_orm(column_name = "reportedMembersLite")]
    pub reported_members_lite: Option<i32>,
    /// The managed gateway key (purpose CONNECT) that hosted-service traffic of
    /// this license runs under. One per license, so spend rows name the install.
    /// Created when the license first resolves. A plain column, not a relation.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "virtualKeyId")]
    pub virtual_key_id: Option<String>,
    /// The seats of the license this one replaced, set only when a seat change
    /// raised a linked license. Billing reads it through licensing to invoice
    /// the added seats, so the fact commits with the license or not at all.
    #[sea_orm(column_name = "seatsRaisedFrom")]
    pub seats_raised_from: Option<i32>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Restrict"
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
