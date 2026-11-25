use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "UserBusinessIntel")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "howDidYouHearAboutUs", column_type = "Text")]
    pub how_did_you_hear_about_us: Option<String>,
    #[sea_orm(column_name = "jobTitle", column_type = "Text")]
    pub job_title: Option<String>,
    #[sea_orm(column_name = "teamSize", column_type = "Text")]
    pub team_size: Option<String>,
    #[sea_orm(column_name = "companyName", column_type = "Text")]
    pub company_name: Option<String>,
    #[sea_orm(column_name = "primaryUseCase", column_type = "Text")]
    pub primary_use_case: Option<String>,
    #[sea_orm(column_name = "currentSolution", column_type = "Text")]
    pub current_solution: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub timeline: Option<String>,
    #[sea_orm(column_name = "signupPlan", column_type = "Text")]
    pub signup_plan: Option<String>,
    #[sea_orm(column_name = "signupTrial")]
    pub signup_trial: Option<bool>,
    #[sea_orm(column_name = "utmSource", column_type = "Text")]
    pub utm_source: Option<String>,
    #[sea_orm(column_name = "utmMedium", column_type = "Text")]
    pub utm_medium: Option<String>,
    #[sea_orm(column_name = "utmCampaign", column_type = "Text")]
    pub utm_campaign: Option<String>,
    #[sea_orm(column_name = "utmContent", column_type = "Text")]
    pub utm_content: Option<String>,
    #[sea_orm(unique)]
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
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "UserBusinessIntel_userId_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
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
