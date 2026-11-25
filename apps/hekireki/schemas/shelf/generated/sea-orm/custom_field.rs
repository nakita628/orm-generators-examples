use super::custom_field_format::CustomFieldFormat;
use super::custom_field_type::CustomFieldType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "CustomField")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "helpText", column_type = "Text")]
    pub help_text: Option<String>,
    #[sea_orm(default_value = false)]
    pub required: bool,
    #[sea_orm(default_value = true)]
    pub active: bool,
    #[sea_orm(default_value = "TEXT")]
    pub r#type: CustomFieldType,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub options: Vec<String>,
    /// How a TEXT field's values are read. LINK requires a web link in every
    /// value; any other type stays PLAIN (enforced by a CHECK constraint).
    #[sea_orm(default_value = "PLAIN")]
    pub format: CustomFieldFormat,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "deletedAt", column_type = "custom(\"timestamp(3)\")")]
    pub deleted_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "CustomField_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        fk_name = "CustomField_userId_fkey"
    )]
    CreatedBy,
    #[sea_orm(has_many = "super::asset_custom_field_value::Entity")]
    AssetCustomFieldsValues,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::asset_custom_field_value::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetCustomFieldsValues.def()
    }
}

impl Related<super::category::Entity> for Entity {
    fn to() -> RelationDef {
        super::category_to_custom_field::Relation::Category.def()
    }
    fn via() -> Option<RelationDef> {
        Some(
            super::category_to_custom_field::Relation::CustomField
                .def()
                .rev(),
        )
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
