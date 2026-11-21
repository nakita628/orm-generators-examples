use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Idempotency marker for a source event consumed by a process definition.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ProcessManagerInbox")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "processName")]
    pub process_name: String,
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    #[sea_orm(column_name = "processKey")]
    pub process_key: String,
    #[sea_orm(column_name = "tenantId")]
    pub tenant_id: String,
    /// The raw source event id, kept UNINDEXED for diagnostics. A caller composes
    /// this and its length is that domain's business, so it must never carry a
    /// btree constraint — see `sourceEventKey`.
    #[sea_orm(column_name = "sourceEventId")]
    pub source_event_id: String,
    /// The fixed-width digest of `sourceEventId` the uniqueness is enforced on,
    /// derived by the store (`stores/inboxKey.ts`). Postgres cannot index a btree
    /// row past ~2704 bytes; keying on a digest makes the constraint the same
    /// width no matter what any pipeline concatenates into an idempotency key.
    ///
    /// Nullable ONLY for the rolling deploy that introduces it — pods running the
    /// previous build write rows without it, and a NOT NULL would fail their
    /// commits outright. The store always supplies it.
    ///
    /// The trade, stated plainly: a NULL is not constrained by the unique index
    /// (btree treats NULLs as distinct), so while an old writer is serving, the
    /// one case that loses its guard is the same source event reaching two
    /// different process keys at once — the store's own read-check plus its
    /// per-processKey advisory lock cover everything else. A follow-up migration
    /// adds NOT NULL once the fleet has cycled and closes it.
    #[sea_orm(column_name = "sourceEventKey")]
    pub source_event_key: Option<String>,
    #[sea_orm(column_name = "consumedAt", column_type = "TimestampWithTimeZone")]
    pub consumed_at: DateTimeWithTimeZone,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
