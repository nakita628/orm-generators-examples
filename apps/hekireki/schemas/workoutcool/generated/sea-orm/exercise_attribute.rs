use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "exercise_attributes")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "exerciseId", column_type = "Text")]
    pub exercise_id: String,
    #[sea_orm(column_name = "attributeNameId", column_type = "Text")]
    pub attribute_name_id: String,
    #[sea_orm(column_name = "attributeValueId", column_type = "Text")]
    pub attribute_value_id: String,
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
        belongs_to = "super::exercise::Entity",
        from = "Column::ExerciseId",
        to = "super::exercise::Column::Id",
        on_delete = "Cascade",
        fk_name = "exercise_attributes_exerciseId_fkey"
    )]
    Exercise,
    #[sea_orm(
        belongs_to = "super::exercise_attribute_name::Entity",
        from = "Column::AttributeNameId",
        to = "super::exercise_attribute_name::Column::Id",
        fk_name = "exercise_attributes_attributeNameId_fkey"
    )]
    AttributeName,
    #[sea_orm(
        belongs_to = "super::exercise_attribute_value::Entity",
        from = "Column::AttributeValueId",
        to = "super::exercise_attribute_value::Column::Id",
        fk_name = "exercise_attributes_attributeValueId_fkey"
    )]
    AttributeValue,
}

impl Related<super::exercise::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Exercise.def()
    }
}

impl Related<super::exercise_attribute_name::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttributeName.def()
    }
}

impl Related<super::exercise_attribute_value::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttributeValue.def()
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
