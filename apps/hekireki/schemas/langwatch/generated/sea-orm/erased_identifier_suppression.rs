use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Hashes of erased identifiers - the do-not-reimport list.
///
/// Without it the erasure is undone by the pipeline that produced it: the
/// pullers look 30 days back, so the next day's pull re-ingests the same email
/// and re-writes what we erased. Every write path carrying an identifier
/// consults this list and drops a suppressed one (ADR-128 §9 step 1).
///
/// It stores hashes and never identifiers, so it is not itself a copy of the
/// data it exists to keep out. The stored `identifierHash` and the pseudonym
/// written in place of the erased value are the same function of the same
/// input - HMAC-SHA256(secret, identifier) - so a write path computes that digest
/// once and uses it twice: as the membership test, and, on a hit, as the
/// replacement value. That is why no original-to-pseudonym mapping table
/// exists; one would hold the erased identifier in plaintext forever.
///
/// Scoped to the provider as well as the organization: the same string can be a
/// different person under a different provider.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ErasedIdentifierSuppression")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_type = "Text")]
    pub provider: String,
    #[sea_orm(column_name = "identifierHash", column_type = "Text")]
    pub identifier_hash: String,
    #[sea_orm(column_name = "erasedAt", column_type = "custom(\"timestamp(3)\")")]
    pub erased_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
