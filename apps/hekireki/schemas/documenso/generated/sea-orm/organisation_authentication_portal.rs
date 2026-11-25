use super::organisation_member_role::OrganisationMemberRole;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganisationAuthenticationPortal")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(default_value = false)]
    pub enabled: bool,
    #[sea_orm(column_name = "clientId", column_type = "Text", default_value = "")]
    pub client_id: String,
    #[sea_orm(column_name = "clientSecret", column_type = "Text", default_value = "")]
    pub client_secret: String,
    #[sea_orm(column_name = "wellKnownUrl", column_type = "Text", default_value = "")]
    pub well_known_url: String,
    #[sea_orm(column_name = "defaultOrganisationRole", default_value = "MEMBER")]
    pub default_organisation_role: OrganisationMemberRole,
    #[sea_orm(column_name = "autoProvisionUsers", default_value = true)]
    pub auto_provision_users: bool,
    #[sea_orm(
        column_name = "allowedDomains",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub allowed_domains: Vec<String>,
    #[sea_orm(column_name = "allowPersonalOrganisations", default_value = false)]
    pub allow_personal_organisations: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_one = "super::organisation::Entity")]
    Organisation,
}

impl Related<super::organisation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organisation.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
