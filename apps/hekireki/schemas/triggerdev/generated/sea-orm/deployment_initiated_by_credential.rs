use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "DeploymentInitiatedByCredential"
)]
pub enum DeploymentInitiatedByCredential {
    #[sea_orm(string_value = "ROOT_API_KEY")]
    RootApiKey,
    #[sea_orm(string_value = "SCOPED_API_KEY")]
    ScopedApiKey,
    #[sea_orm(string_value = "PERSONAL_ACCESS_TOKEN")]
    PersonalAccessToken,
    #[sea_orm(string_value = "ORGANIZATION_ACCESS_TOKEN")]
    OrganizationAccessToken,
    #[sea_orm(string_value = "TRUSTED_PUBLISHING")]
    TrustedPublishing,
}
