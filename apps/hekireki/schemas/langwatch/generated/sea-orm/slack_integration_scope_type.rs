use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Who can use a Slack connection (ADR-093 §5a): every project in the
/// organization, or one project. Chosen per connection.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "SlackIntegrationScopeType"
)]
pub enum SlackIntegrationScopeType {
    #[sea_orm(string_value = "ORGANIZATION")]
    Organization,
    #[sea_orm(string_value = "PROJECT")]
    Project,
}
