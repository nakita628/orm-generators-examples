use super::billing_period::BillingPeriod;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganizationOnboarding")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: i32,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "orgOwnerEmail", column_type = "Text")]
    pub org_owner_email: String,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: Option<i32>,
    #[sea_orm(column_name = "billingPeriod")]
    pub billing_period: BillingPeriod,
    #[sea_orm(column_name = "pricePerSeat")]
    pub price_per_seat: f64,
    pub seats: i32,
    #[sea_orm(column_name = "isPlatform", default_value = false)]
    pub is_platform: bool,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_type = "Text")]
    pub logo: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub bio: Option<String>,
    #[sea_orm(column_name = "brandColor", column_type = "Text")]
    pub brand_color: Option<String>,
    #[sea_orm(column_name = "bannerUrl", column_type = "Text")]
    pub banner_url: Option<String>,
    #[sea_orm(column_name = "isDomainConfigured", default_value = false)]
    pub is_domain_configured: bool,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeCustomerId", column_type = "Text")]
    pub stripe_customer_id: Option<String>,
    #[sea_orm(column_name = "stripeSubscriptionId", column_type = "Text")]
    pub stripe_subscription_id: Option<String>,
    #[sea_orm(column_name = "stripeSubscriptionItemId", column_type = "Text")]
    pub stripe_subscription_item_id: Option<String>,
    /// @zod.import(["import { orgOnboardingInvitedMembersSchema } from '../../zod-utils'"]).custom.use(orgOnboardingInvitedMembersSchema)
    #[sea_orm(
        column_name = "invitedMembers",
        column_type = "JsonBinary",
        default_value = "[]"
    )]
    pub invited_members: Json,
    /// @zod.import(["import { orgOnboardingTeamsSchema } from '../../zod-utils'"]).custom.use(orgOnboardingTeamsSchema)
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub teams: Json,
    #[sea_orm(column_name = "isComplete", default_value = false)]
    pub is_complete: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "OrganizationOnboarding_createdById_fkey"
    )]
    CreatedBy,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::OrganizationId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "OrganizationOnboarding_organizationId_fkey"
    )]
    Organization,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
