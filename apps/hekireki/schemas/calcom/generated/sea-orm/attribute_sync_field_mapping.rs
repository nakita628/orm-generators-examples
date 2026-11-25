use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AttributeSyncFieldMapping")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "integrationFieldName", column_type = "Text")]
    pub integration_field_name: String,
    #[sea_orm(column_name = "attributeId", column_type = "Text")]
    pub attribute_id: String,
    pub enabled: bool,
    #[sea_orm(column_name = "integrationAttributeSyncId", column_type = "Text")]
    pub integration_attribute_sync_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::attribute::Entity",
        from = "Column::AttributeId",
        to = "super::attribute::Column::Id",
        on_delete = "Cascade",
        fk_name = "AttributeSyncFieldMapping_attributeId_fkey"
    )]
    Attribute,
    #[sea_orm(
        belongs_to = "super::integration_attribute_sync::Entity",
        from = "Column::IntegrationAttributeSyncId",
        to = "super::integration_attribute_sync::Column::Id",
        on_delete = "Cascade",
        fk_name = "AttributeSyncFieldMapping_integrationAttributeSyncId_fkey"
    )]
    IntegrationAttributeSync,
}

impl Related<super::attribute::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Attribute.def()
    }
}

impl Related<super::integration_attribute_sync::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::IntegrationAttributeSync.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
