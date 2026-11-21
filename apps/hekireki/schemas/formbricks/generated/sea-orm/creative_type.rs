use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Creatives (ENG-3993): page elements the web SDK renders next to surveys. Every type lives in
/// `Creative`; `content` and `styling` are JSON validated per type by the API v3 schemas, so a new type is
/// a new payload variant, not new columns.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "CreativeType")]
pub enum CreativeType {
    #[sea_orm(string_value = "feedbackButton")]
    Feedbackbutton,
    #[sea_orm(string_value = "banner")]
    Banner,
}
