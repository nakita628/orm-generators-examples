use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A custom field value an AssetModel gives the assets that join it.
///
/// `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
/// asset is a plain copy. Rows cascade away with their model or field.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AssetModelCustomFieldDefault")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub value: Json,
    #[sea_orm(column_name = "assetModelId", column_type = "Text")]
    pub asset_model_id: String,
    #[sea_orm(column_name = "customFieldId", column_type = "Text")]
    pub custom_field_id: String,
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
        belongs_to = "super::asset_model::Entity",
        from = "Column::AssetModelId",
        to = "super::asset_model::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "AssetModelCustomFieldDefault_assetModelId_fkey"
    )]
    AssetModel,
}

impl Related<super::asset_model::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetModel.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
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
