use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ScimToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// D08: the SSO connection this token was issued for, and its ENTIRE write
    /// authority — a token minted for one connection can never touch another
    /// connection's people. No relation, matching `SsoConnection` itself: that
    /// projection carries no foreign keys, so neither may a row pointing at it.
    ///
    /// Nullable only for the tokens that predate the column. They were minted
    /// when "which organization" was the whole answer, so they keep the
    /// organization-wide authority they were sold with; the backfill points
    /// each at its organization's first connection where there is one. Minting
    /// a NEW token without a connection is refused (`scim_connection_required`).
    #[sea_orm(column_name = "connectionId")]
    pub connection_id: Option<String>,
    /// UNIQUE, and that is a tenancy boundary rather than tidiness.
    ///
    /// The lookup that turns a bearer token into an organization is keyed on
    /// this column alone — it has to be, since the token is the only thing the
    /// caller presents. While every token was 32 random bytes a collision was
    /// impossible and the missing constraint cost nothing. An administrator may
    /// now supply a value they already hold, so two organizations CAN choose the
    /// same string, and without this the lookup would answer with whichever row
    /// the planner reached first: one customer's directory provisioning and
    /// deleting another customer's people.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "hashedToken")]
    pub hashed_token: String,
    /// How `hashedToken` was derived, so the two schemes can coexist while
    /// tokens minted under the old one are still in use.
    ///
    /// `sha256` is a bare digest, which is correct for 32 random bytes and wrong
    /// for a value a person chose — a database dump of those is a wordlist away
    /// from live credentials. `hmac-sha256` is keyed on the deployment's own
    /// secret, so a dump on its own is inert. Every new row is the latter.
    #[sea_orm(column_name = "hashScheme", default_value = "sha256")]
    pub hash_scheme: String,
    pub description: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "lastUsedAt")]
    pub last_used_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade"
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
        Ok(self)
    }
}
