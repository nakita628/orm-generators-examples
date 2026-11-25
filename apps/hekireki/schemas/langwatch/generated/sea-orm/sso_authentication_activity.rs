use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// A successful callback through one SSO connection. Identifier.lastUsedAt
/// also advances during routing and resolution, so it cannot prove that a
/// member actually authenticated through the replacement before cutover.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SsoAuthenticationActivity")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "connectionId", column_type = "Text")]
    pub connection_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(
        column_name = "authenticatedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub authenticated_at: DateTime,
    /// The provider-side account the assertion named; null before #7633's arrival door recorded it.
    #[sea_orm(column_name = "providerAccountId", column_type = "Text")]
    pub provider_account_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
