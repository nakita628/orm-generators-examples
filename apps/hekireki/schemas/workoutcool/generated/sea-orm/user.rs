use super::user_role::UserRole;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "user")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "firstName", column_type = "Text", default_value = "")]
    pub first_name: String,
    #[sea_orm(column_name = "lastName", column_type = "Text", default_value = "")]
    pub last_name: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub email: String,
    #[sea_orm(column_name = "emailVerified")]
    pub email_verified: bool,
    #[sea_orm(column_type = "Text")]
    pub image: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "fr")]
    pub locale: Option<String>,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(default_value = "user")]
    pub role: Option<UserRole>,
    #[sea_orm(default_value = false)]
    pub banned: Option<bool>,
    #[sea_orm(column_name = "banReason", column_type = "Text")]
    pub ban_reason: Option<String>,
    #[sea_orm(column_name = "banExpires", column_type = "custom(\"timestamp(3)\")")]
    pub ban_expires: Option<DateTime>,
    #[sea_orm(column_name = "isPremium", default_value = false)]
    pub is_premium: Option<bool>,
    #[sea_orm(column_name = "onboardingPreferences", column_type = "JsonBinary")]
    pub onboarding_preferences: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::session::Entity")]
    Sessions,
    #[sea_orm(has_many = "super::account::Entity")]
    Accounts,
    #[sea_orm(has_many = "super::feedbacks::Entity")]
    Feedbacks,
    #[sea_orm(has_many = "super::workout_session::Entity")]
    WorkoutSession,
    #[sea_orm(has_many = "super::user_favorite_exercise::Entity")]
    FavoriteExercises,
    #[sea_orm(has_many = "super::subscription::Entity")]
    Subscriptions,
    #[sea_orm(has_many = "super::license::Entity")]
    Licenses,
    #[sea_orm(has_many = "super::user_program_enrollment::Entity")]
    ProgramEnrollments,
}

impl Related<super::session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Sessions.def()
    }
}

impl Related<super::account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Accounts.def()
    }
}

impl Related<super::feedbacks::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Feedbacks.def()
    }
}

impl Related<super::workout_session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkoutSession.def()
    }
}

impl Related<super::user_favorite_exercise::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FavoriteExercises.def()
    }
}

impl Related<super::subscription::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Subscriptions.def()
    }
}

impl Related<super::license::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Licenses.def()
    }
}

impl Related<super::user_program_enrollment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ProgramEnrollments.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
