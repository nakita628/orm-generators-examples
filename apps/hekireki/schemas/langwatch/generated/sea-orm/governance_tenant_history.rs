use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Every ClickHouse `TenantId` this organization has ever written governance
/// rows under, appended on first use and never pruned.
///
/// The live `resolveGovProjectId` lookup cannot be this translation: it filters
/// `archivedAt: null` and so returns null forever once someone archives the
/// governance project, while the write path re-reads the same project by slug
/// with no such filter and keeps landing rows under the old tenant. The
/// organization would map to zero tenants on read and one on write - a
/// permanent split-brain in which the cost screen reports "no governance data"
/// and an erasure job erases nothing and reports success.
///
/// Reads resolve against the whole history; erasure walks all of it, because
/// personal data does not stop existing in a tenant that stopped being current.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "GovernanceTenantHistory")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// A hidden governance project id this organization has written rows under.
    #[sea_orm(column_name = "tenantId", column_type = "Text")]
    pub tenant_id: String,
    #[sea_orm(column_name = "firstUsedAt", column_type = "custom(\"timestamp(3)\")")]
    pub first_used_at: DateTime,
    #[sea_orm(column_name = "lastUsedAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
