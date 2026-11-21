use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// One member's or pending invitation's seat as the seat peak last applied it (ADR-002).
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganizationSeatPeakHolder")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// `member`, `invitation` or `hold` (a seat held after somebody left).
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "holderKind")]
    pub holder_kind: String,
    /// The member's or the hold's userId, or the invitation's id.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "holderId")]
    pub holder_id: String,
    /// `full`, `lite`, `developer` or `none`.
    pub seat: String,
    #[sea_orm(column_name = "lastChangedAt")]
    pub last_changed_at: f64,
    /// Absent for a seat the opening snapshot recorded.
    #[sea_orm(column_name = "lastChangeId")]
    pub last_change_id: Option<String>,
    /// An invitation's expiry; null for a member or an invitation that never expires.
    #[sea_orm(column_name = "expiresAt")]
    pub expires_at: Option<f64>,
    /// The fold gave the invitation's seat back at `expiresAt`.
    #[sea_orm(column_name = "expiredByFold", default_value = false)]
    pub expired_by_fold: bool,
    /// A hold whose seat a live full invitation to its address took: that invitation's id.
    #[sea_orm(column_name = "suppressedBy")]
    pub suppressed_by: Option<String>,
    /// The fold gave a suppressed hold its seat back when that invitation expired.
    #[sea_orm(column_name = "restoredByFold", default_value = false)]
    pub restored_by_fold: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
