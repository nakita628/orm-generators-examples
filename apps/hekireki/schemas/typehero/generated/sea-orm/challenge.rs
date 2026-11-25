use super::challenge_status::ChallengeStatus;
use super::difficulty::Difficulty;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Challenge")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"datetime(3)\")")]
    pub updated_at: DateTime,
    pub difficulty: Difficulty,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub slug: String,
    /// this is the description to describe the challenge on description panel
    #[sea_orm(column_type = "custom(\"mediumtext\")")]
    pub description: String,
    /// this is a short blue thats shown on the explore card
    #[sea_orm(
        column_name = "shortDescription",
        column_type = "String(StringLen::N(191))"
    )]
    pub short_description: String,
    /// Starter code for the challenge
    #[sea_orm(column_type = "custom(\"mediumtext\")")]
    pub code: String,
    /// Test cases that need to be passed in order have the submission be accepted.
    #[sea_orm(column_type = "custom(\"mediumtext\")")]
    pub tests: String,
    #[sea_orm(default_value = "PENDING")]
    pub status: ChallengeStatus,
    pub tsconfig: Option<Json>,
    #[sea_orm(column_name = "userId", column_type = "String(StringLen::N(191))")]
    pub user_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        fk_name = "Challenge_userId_fkey"
    )]
    User,
    #[sea_orm(has_many = "super::vote::Entity")]
    Vote,
    #[sea_orm(has_many = "super::submission::Entity")]
    Submission,
    #[sea_orm(has_many = "super::comment::Entity")]
    Comment,
    #[sea_orm(has_many = "super::bookmark::Entity")]
    Bookmark,
    #[sea_orm(has_many = "super::shared_solution::Entity")]
    SharedSolution,
    #[sea_orm(has_many = "super::report::Entity")]
    Report,
    #[sea_orm(has_many = "super::track_challenge::Entity")]
    TrackChallenge,
    #[sea_orm(has_many = "super::notification::Entity")]
    Notification,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::vote::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Vote.def()
    }
}

impl Related<super::submission::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Submission.def()
    }
}

impl Related<super::comment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Comment.def()
    }
}

impl Related<super::bookmark::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Bookmark.def()
    }
}

impl Related<super::shared_solution::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SharedSolution.def()
    }
}

impl Related<super::report::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Report.def()
    }
}

impl Related<super::track_challenge::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TrackChallenge.def()
    }
}

impl Related<super::notification::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notification.def()
    }
}

impl Related<super::tag::Entity> for Entity {
    fn to() -> RelationDef {
        super::challenge_to_tag::Relation::Tag.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::challenge_to_tag::Relation::Challenge.def().rev())
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
