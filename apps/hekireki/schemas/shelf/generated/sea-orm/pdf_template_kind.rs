use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The kind of PDF a PdfTemplate lays out. Only the booking PDF has saved
/// layouts today; the check-in receipt and the audit receipt can join as new
/// kinds.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "PdfTemplateKind")]
pub enum PdfTemplateKind {
    #[sea_orm(string_value = "BOOKING")]
    Booking,
}
