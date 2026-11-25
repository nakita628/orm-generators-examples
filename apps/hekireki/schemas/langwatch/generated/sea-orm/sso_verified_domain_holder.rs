use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// A connection currently entitled to use an organization-owned verified
/// domain. Usually there is one holder. During an Auth0 migration there are
/// exactly two: the grandfathered predecessor and its direct replacement.
///
/// The projection repository enforces that relationship before inserting a
/// second holder. The flattened migration also installs partial unique live
/// legacy/direct slots on `SsoConnection`, closing concurrent third-writer
/// races that Prisma cannot describe declaratively.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SsoVerifiedDomainHolder")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub domain: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "connectionId", column_type = "Text")]
    pub connection_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::sso_verified_domain::Entity",
        from = "(Column::Domain, Column::OrganizationId)",
        to = "(super::sso_verified_domain::Column::Domain, super::sso_verified_domain::Column::OrganizationId)",
        on_delete = "Cascade",
        fk_name = "SsoVerifiedDomainHolder_domain_organizationId_fkey"
    )]
    Ownership,
}

impl Related<super::sso_verified_domain::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Ownership.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
