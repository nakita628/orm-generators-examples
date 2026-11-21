use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// SCIM's own copy of an organization's SSO connections: a peer fold over
/// identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
/// reducer, one row per connection. SCIM reads this instead of asking identity.
/// Rebuilt from identity's event log by a projection replay, so no relation:
/// a foreign key would make a replay depend on state the log does not carry.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ScimSsoConnectionView")]
pub struct Model {
    /// The aggregate: identity's connection id.
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt).
    pub folded: Json,
    /// The event ids of the last delivery folded in, so a redelivery folds once.
    #[sea_orm(column_name = "appliedEventIds")]
    pub applied_event_ids: Vec<String>,
    /// A row of another version reads as undecodable and is re-folded from the log.
    #[sea_orm(column_name = "projectionVersion")]
    pub projection_version: String,
    /// Event-derived, so a replayed row equals the row it rebuilds.
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
