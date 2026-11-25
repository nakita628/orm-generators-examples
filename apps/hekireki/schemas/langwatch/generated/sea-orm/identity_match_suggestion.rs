use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Match candidates a background job computed and a human has yet to confirm
/// (ADR-128 §12).
///
/// Stored rather than computed when the review screen asks, and the reason is
/// measured rather than stylistic: fuzzy matching is quadratic, there is no
/// database route to it here, and at this document's own example size - 2,000
/// discovered people against 500 accounts - plain edit distance over the
/// 1,000,000 pairs blocked the event loop for 2.9 seconds. Per page load,
/// uncached, stalling every other request on the instance. A stored row also
/// makes a pending count answerable without paying for the whole sweep.
///
/// Rows are written only by the job, which prefilters before it scores and
/// recomputes when its inputs change. Confirming one opens an `IdentityMatch`
/// and deletes the suggestion.
///
/// Suspension deliberately does NOT live here - it is on `DiscoveredPerson`,
/// because a halt on automatic linking that the next recompute clears is not a
/// halt.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "IdentityMatchSuggestion")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "discoveredPersonId", column_type = "Text")]
    pub discovered_person_id: String,
    /// The platform user this might be. Not nullable: a suggestion with no
    /// candidate is not a suggestion.
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    /// Similarity of the two name texts, after the prefilter, in [0, 1]. A
    /// display hint and an ordering key - never a threshold anything auto-links
    /// on, which is the whole distinction this table exists to hold.
    pub score: f64,
    #[sea_orm(column_name = "computedAt", column_type = "custom(\"timestamp(3)\")")]
    pub computed_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
