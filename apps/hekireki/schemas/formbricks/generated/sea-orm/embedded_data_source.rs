use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Where an Embedded Data field's value comes from.
/// `computed` is set by survey logic, `ingested` arrives from a URL param or the SDK, and
/// `reserved` is auto-captured system metadata. Reserved fields are never stored as rows —
/// they are a static catalog in code — but the value exists so every layer shares one type.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "EmbeddedDataSource")]
pub enum EmbeddedDataSource {
    #[sea_orm(string_value = "computed")]
    Computed,
    #[sea_orm(string_value = "ingested")]
    Ingested,
    #[sea_orm(string_value = "reserved")]
    Reserved,
}
