use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The format of a TEXT custom field. Unlike the type, it can change after the
/// field exists.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "CustomFieldFormat")]
pub enum CustomFieldFormat {
    #[sea_orm(string_value = "PLAIN")]
    Plain,
    #[sea_orm(string_value = "LINK")]
    Link,
}
