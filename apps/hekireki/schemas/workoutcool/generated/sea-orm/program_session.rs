use super::exercise_attribute_value_enum::ExerciseAttributeValueEnum;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "program_sessions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "weekId", column_type = "Text")]
    pub week_id: String,
    #[sea_orm(column_name = "sessionNumber")]
    pub session_number: i32,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_name = "titleEn", column_type = "Text")]
    pub title_en: String,
    #[sea_orm(column_name = "titleEs", column_type = "Text")]
    pub title_es: String,
    #[sea_orm(column_name = "titlePt", column_type = "Text")]
    pub title_pt: String,
    #[sea_orm(column_name = "titleRu", column_type = "Text")]
    pub title_ru: String,
    #[sea_orm(column_name = "titleZhCn", column_type = "Text")]
    pub title_zh_cn: String,
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_name = "slugEn", column_type = "Text")]
    pub slug_en: String,
    #[sea_orm(column_name = "slugEs", column_type = "Text")]
    pub slug_es: String,
    #[sea_orm(column_name = "slugPt", column_type = "Text")]
    pub slug_pt: String,
    #[sea_orm(column_name = "slugRu", column_type = "Text")]
    pub slug_ru: String,
    #[sea_orm(column_name = "slugZhCn", column_type = "Text")]
    pub slug_zh_cn: String,
    #[sea_orm(column_type = "Text")]
    pub description: String,
    #[sea_orm(column_name = "descriptionEn", column_type = "Text")]
    pub description_en: String,
    #[sea_orm(column_name = "descriptionEs", column_type = "Text")]
    pub description_es: String,
    #[sea_orm(column_name = "descriptionPt", column_type = "Text")]
    pub description_pt: String,
    #[sea_orm(column_name = "descriptionRu", column_type = "Text")]
    pub description_ru: String,
    #[sea_orm(column_name = "descriptionZhCn", column_type = "Text")]
    pub description_zh_cn: String,
    pub equipment: Vec<ExerciseAttributeValueEnum>,
    #[sea_orm(column_name = "estimatedMinutes")]
    pub estimated_minutes: i32,
    #[sea_orm(column_name = "isPremium", default_value = true)]
    pub is_premium: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::program_week::Entity",
        from = "Column::WeekId",
        to = "super::program_week::Column::Id",
        on_delete = "Cascade",
        fk_name = "program_sessions_weekId_fkey"
    )]
    Week,
    #[sea_orm(has_many = "super::program_session_exercise::Entity")]
    Exercises,
    #[sea_orm(has_many = "super::user_session_progress::Entity")]
    UserProgress,
}

impl Related<super::program_week::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Week.def()
    }
}

impl Related<super::program_session_exercise::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Exercises.def()
    }
}

impl Related<super::user_session_progress::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UserProgress.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
