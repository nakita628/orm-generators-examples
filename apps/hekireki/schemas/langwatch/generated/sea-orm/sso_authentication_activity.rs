use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// A successful callback through one SSO connection. Identifier.lastUsedAt
/// also advances during routing and resolution, so it cannot prove that a
/// member actually authenticated through the replacement before cutover.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SsoAuthenticationActivity")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "connectionId")]
    pub connection_id: String,
    #[sea_orm(column_name = "userId")]
    pub user_id: String,
    #[sea_orm(column_name = "authenticatedAt")]
    pub authenticated_at: DateTime,
    /// The provider-side account the assertion named; null before #7633's arrival door recorded it.
    #[sea_orm(column_name = "providerAccountId")]
    pub provider_account_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
