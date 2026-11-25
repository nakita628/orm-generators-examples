use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Plugs")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "plugFunction", column_type = "Text")]
    pub plug_function: String,
    #[sea_orm(column_type = "Text")]
    pub data: String,
    #[sea_orm(column_name = "integrationId", column_type = "Text")]
    pub integration_id: String,
    #[sea_orm(default_value = true)]
    pub activated: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::integration::Entity",
        from = "Column::IntegrationId",
        to = "super::integration::Column::Id",
        fk_name = "Plugs_integrationId_fkey"
    )]
    Integration,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "Plugs_organizationId_fkey"
    )]
    Organization,
}

impl Related<super::integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Integration.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }
}
