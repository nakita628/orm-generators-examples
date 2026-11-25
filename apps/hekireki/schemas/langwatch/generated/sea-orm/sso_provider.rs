use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SsoProvider")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub issuer: String,
    #[sea_orm(column_name = "oidcConfig", column_type = "Text")]
    pub oidc_config: Option<String>,
    #[sea_orm(column_name = "samlConfig", column_type = "Text")]
    pub saml_config: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "providerId", column_type = "Text")]
    pub provider_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub domain: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
