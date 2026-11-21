use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Each purpose is a separate key set with its own issuer, rotation and publication.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "OidcSigningKeyPurpose"
)]
pub enum OidcSigningKeyPurpose {
    #[sea_orm(string_value = "EXTERNAL_OIDC")]
    ExternalOidc,
    #[sea_orm(string_value = "INTERNAL_RUN")]
    InternalRun,
}
