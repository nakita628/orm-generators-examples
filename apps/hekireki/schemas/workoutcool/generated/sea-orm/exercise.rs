use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "exercises")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "nameEn", column_type = "Text")]
    pub name_en: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "descriptionEn", column_type = "Text")]
    pub description_en: Option<String>,
    #[sea_orm(column_name = "fullVideoUrl", column_type = "Text")]
    pub full_video_url: Option<String>,
    #[sea_orm(column_name = "fullVideoImageUrl", column_type = "Text")]
    pub full_video_image_url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub introduction: Option<String>,
    #[sea_orm(column_name = "introductionEn", column_type = "Text")]
    pub introduction_en: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub slug: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "slugEn", column_type = "Text")]
    pub slug_en: Option<String>,
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
    #[sea_orm(has_many = "super::exercise_attribute::Entity")]
    Attributes,
    #[sea_orm(has_many = "super::workout_session_exercise::Entity")]
    WorkoutSessionExercise,
    #[sea_orm(has_many = "super::program_session_exercise::Entity")]
    ProgramSessionExercise,
    #[sea_orm(has_many = "super::user_favorite_exercise::Entity")]
    FavoritesByUsers,
}

impl Related<super::exercise_attribute::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Attributes.def()
    }
}

impl Related<super::workout_session_exercise::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkoutSessionExercise.def()
    }
}

impl Related<super::program_session_exercise::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ProgramSessionExercise.def()
    }
}

impl Related<super::user_favorite_exercise::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FavoritesByUsers.def()
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
