use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Which webhook pipeline wrote a delivery row, and therefore which columns
/// carry its tenancy: `platform` rows anchor on organizationId + endpointId,
/// `automations` rows on projectId + triggerId.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "WebhookDeliveryChannel"
)]
pub enum WebhookDeliveryChannel {
    #[sea_orm(string_value = "platform")]
    Platform,
    #[sea_orm(string_value = "automations")]
    Automations,
}
