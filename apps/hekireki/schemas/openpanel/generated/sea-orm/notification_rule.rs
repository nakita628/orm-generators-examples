use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "notification_rules")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(default_expr = "Expr::cust(\"gen_random_uuid()\")")]
    pub id: Uuid,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "sendToApp", default_value = false)]
    pub send_to_app: bool,
    #[sea_orm(column_name = "sendToEmail", default_value = false)]
    pub send_to_email: bool,
    /// [IPrismaNotificationRuleConfig]
    #[sea_orm(column_type = "JsonBinary")]
    pub config: Json,
    #[sea_orm(column_type = "Text")]
    pub template: Option<String>,
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
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "notification_rules_projectId_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::notification::Entity")]
    Notifications,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::notification::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notifications.def()
    }
}

impl Related<super::integration::Entity> for Entity {
    fn to() -> RelationDef {
        super::integration_to_notification_rule::Relation::Integration.def()
    }
    fn via() -> Option<RelationDef> {
        Some(
            super::integration_to_notification_rule::Relation::NotificationRule
                .def()
                .rev(),
        )
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
