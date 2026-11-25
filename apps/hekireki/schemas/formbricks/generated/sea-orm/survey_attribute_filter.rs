use super::survey_attribute_filter_condition::SurveyAttributeFilterCondition;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Defines targeting rules for surveys based on contact attributes.
/// Used to show surveys only to contacts matching specific criteria.
///
/// @property id - Unique identifier for the filter
/// @property attributeKey - The contact attribute to filter on
/// @property survey - The survey being filtered
/// @property condition - The comparison operator to use
/// @property value - The value to compare against
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SurveyAttributeFilter")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "attributeKeyId", column_type = "Text")]
    pub attribute_key_id: String,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: String,
    pub condition: SurveyAttributeFilterCondition,
    #[sea_orm(column_type = "Text")]
    pub value: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::contact_attribute_key::Entity",
        from = "Column::AttributeKeyId",
        to = "super::contact_attribute_key::Column::Id",
        on_delete = "Cascade",
        fk_name = "SurveyAttributeFilter_attributeKeyId_fkey"
    )]
    AttributeKey,
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "Column::SurveyId",
        to = "super::survey::Column::Id",
        on_delete = "Cascade",
        fk_name = "SurveyAttributeFilter_surveyId_fkey"
    )]
    Survey,
}

impl Related<super::contact_attribute_key::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttributeKey.def()
    }
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
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
