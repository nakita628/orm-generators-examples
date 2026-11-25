use super::approved_submit_for_order::APPROVED_SUBMIT_FOR_ORDER;
use super::creation_method::CreationMethod;
use super::state::State;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Post")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(default_value = "QUEUE")]
    pub state: State,
    #[sea_orm(column_name = "publishDate", column_type = "custom(\"timestamp(3)\")")]
    pub publish_date: DateTime,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "integrationId", column_type = "Text")]
    pub integration_id: String,
    #[sea_orm(column_type = "Text")]
    pub content: String,
    #[sea_orm(default_value = 0)]
    pub delay: i32,
    #[sea_orm(column_type = "Text")]
    pub group: String,
    #[sea_orm(column_name = "batchId", column_type = "Text")]
    pub batch_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub title: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "parentPostId", column_type = "Text")]
    pub parent_post_id: Option<String>,
    #[sea_orm(column_name = "releaseId", column_type = "Text")]
    pub release_id: Option<String>,
    #[sea_orm(column_name = "releaseURL", column_type = "Text")]
    pub release_url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub settings: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub image: Option<String>,
    #[sea_orm(column_name = "submittedForOrderId", column_type = "Text")]
    pub submitted_for_order_id: Option<String>,
    #[sea_orm(column_name = "submittedForOrganizationId", column_type = "Text")]
    pub submitted_for_organization_id: Option<String>,
    #[sea_orm(column_name = "approvedSubmitForOrder", default_value = "NO")]
    pub approved_submit_for_order: APPROVED_SUBMIT_FOR_ORDER,
    #[sea_orm(column_name = "creationMethod", default_value = "UNKNOWN")]
    pub creation_method: CreationMethod,
    #[sea_orm(column_name = "lastMessageId", column_type = "Text")]
    pub last_message_id: Option<String>,
    #[sea_orm(column_name = "intervalInDays")]
    pub interval_in_days: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
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
        belongs_to = "super::integration::Entity",
        from = "Column::IntegrationId",
        to = "super::integration::Column::Id",
        fk_name = "Post_integrationId_fkey"
    )]
    Integration,
    #[sea_orm(
        belongs_to = "super::messages::Entity",
        from = "Column::LastMessageId",
        to = "super::messages::Column::Id",
        fk_name = "Post_lastMessageId_fkey"
    )]
    LastMessage,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "Post_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::post::Entity",
        from = "Column::ParentPostId",
        to = "super::post::Column::Id",
        fk_name = "Post_parentPostId_fkey"
    )]
    ParentPost,
    #[sea_orm(
        belongs_to = "super::orders::Entity",
        from = "Column::SubmittedForOrderId",
        to = "super::orders::Column::Id",
        fk_name = "Post_submittedForOrderId_fkey"
    )]
    SubmittedForOrder,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::SubmittedForOrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "Post_submittedForOrganizationId_fkey"
    )]
    SubmittedForOrganization,
    #[sea_orm(has_many = "super::comments::Entity")]
    Comments,
    #[sea_orm(has_many = "super::errors::Entity")]
    Errors,
    #[sea_orm(has_many = "super::payout_problems::Entity")]
    PayoutProblems,
    #[sea_orm(
        has_many = "super::post::Entity",
        from = "Column::Id",
        to = "super::post::Column::ParentPostId"
    )]
    ChildrenPost,
    #[sea_orm(has_many = "super::tags_posts::Entity")]
    Tags,
}

impl Related<super::integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Integration.def()
    }
}

impl Related<super::messages::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LastMessage.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::post::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ParentPost.def()
    }
}

impl Related<super::orders::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SubmittedForOrder.def()
    }
}

impl Related<super::comments::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Comments.def()
    }
}

impl Related<super::errors::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Errors.def()
    }
}

impl Related<super::payout_problems::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PayoutProblems.def()
    }
}

impl Related<super::tags_posts::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tags.def()
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
