use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "notifications")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(default_expr = "Expr::cust(\"gen_random_uuid()\")")]
    pub id: Uuid,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_type = "Text")]
    pub message: String,
    #[sea_orm(column_name = "isReadAt", column_type = "custom(\"timestamp(3)\")")]
    pub is_read_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "sendToApp", default_value = false)]
    pub send_to_app: bool,
    #[sea_orm(column_name = "sendToEmail", default_value = false)]
    pub send_to_email: bool,
    #[sea_orm(column_name = "integrationId")]
    pub integration_id: Option<Uuid>,
    #[sea_orm(column_name = "notificationRuleId")]
    pub notification_rule_id: Option<Uuid>,
    /// [IPrismaNotificationPayload]
    #[sea_orm(column_type = "JsonBinary")]
    pub payload: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "notifications_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::integration::Entity",
        from = "Column::IntegrationId",
        to = "super::integration::Column::Id",
        on_delete = "Cascade",
        fk_name = "notifications_integrationId_fkey"
    )]
    Integration,
    #[sea_orm(
        belongs_to = "super::notification_rule::Entity",
        from = "Column::NotificationRuleId",
        to = "super::notification_rule::Column::Id",
        on_delete = "Cascade",
        fk_name = "notifications_notificationRuleId_fkey"
    )]
    NotificationRule,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Integration.def()
    }
}

impl Related<super::notification_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::NotificationRule.def()
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
