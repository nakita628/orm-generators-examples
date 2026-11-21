use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A name-text seen on a provider's cost or audit rows: "the provider told us
/// someone called m.silva exists". One row per provider-scoped identity.
///
/// This is the unit of GDPR erasure (ADR-128 §9), and deliberately so: most of
/// these people have no LangWatch login at all - contractors, seat holders who
/// never signed in, anyone whose email a provider put on a cost row - so there
/// is no platform user to key on and a user-deletion event can never fire for
/// them.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DiscoveredPerson")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// "anthropic" | "openai" | "databricks" | ... - the provider that named them.
    pub provider: String,
    /// What the provider calls them, verbatim. The join key onto the money rows'
    /// RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
    /// the row addressable for spend attribution without keeping the identifier.
    #[sea_orm(column_name = "rawActorId")]
    pub raw_actor_id: String,
    /// Name or email text as seen. Pseudonymized alongside rawActorId on erasure.
    #[sea_orm(column_name = "displayText")]
    pub display_text: String,
    /// person | service_account - determined deterministically per ADR-128 §10
    /// (a bare UUID in Databricks `executed_by` is a service account, never an
    /// agent), never guessed from the shape of a name.
    pub kind: String,
    /// The department the provider's directory filed this person under, verbatim.
    /// An OBSERVATION, not the `Department` entity: it is free text a tenant owns
    /// on their side, it exists for people who hold no LangWatch account, and
    /// nothing rolls spend up by it. Spend attribution stays on `Department`, via
    /// the membership the match engine proves (directoryDepartmentSync.service).
    /// Null means no directory has named one - never "no department", because a
    /// blank field in the tenant's directory and a person no directory lists are
    /// the same null here.
    pub department: Option<String>,
    #[sea_orm(column_name = "firstSeenAt")]
    pub first_seen_at: DateTime,
    #[sea_orm(column_name = "lastSeenAt")]
    pub last_seen_at: DateTime,
    /// Set when this row's identity has been erased. The row survives so its
    /// spend stays attributable to *someone*; the identifier does not.
    #[sea_orm(column_name = "erasedAt")]
    pub erased_at: Option<DateTime>,
    /// Set alongside `erasedAt` and cleared once the daily cost rows have been
    /// removed and their rebuild requested. Non-null means the ClickHouse half
    /// of the erasure did not finish — a crash between the delete and the
    /// replay leaves days short by the erased amount, and without this marker a
    /// re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
    /// did no work. The re-run resumes at that step instead.
    #[sea_orm(column_name = "moneyRowsPendingAt")]
    pub money_rows_pending_at: Option<DateTime>,
    /// The earliest day to rebuild from, recorded BEFORE the rows are deleted.
    /// Once they are gone there is nothing left to ask which days they were on,
    /// so a resume with no record of this could only replay everything or
    /// nothing. Null means nothing was rebuildable — either no day carried the
    /// identifier, or every one of them is older than the event log keeps.
    #[sea_orm(column_name = "moneyRebuildSince")]
    pub money_rebuild_since: Option<String>,
    /// ADR-128 §12 conflict rule: automatic linking has stopped for this
    /// identity because evidence pointed at two candidates, or contradicted an
    /// existing link. Stored HERE rather than on the suggestion rows, because a
    /// halt on auto-linking that a recompute or a restart clears is not a halt.
    #[sea_orm(column_name = "suspendedAt")]
    pub suspended_at: Option<DateTime>,
    /// What tripped the suspension, for the human who reviews it.
    #[sea_orm(column_name = "suspendedReason")]
    pub suspended_reason: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
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
