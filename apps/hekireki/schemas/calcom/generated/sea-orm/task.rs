use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Task")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(
        column_name = "scheduledAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub scheduled_at: DateTime,
    #[sea_orm(column_name = "succeededAt", column_type = "custom(\"timestamp(3)\")")]
    pub succeeded_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_type = "Text")]
    pub payload: String,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    #[sea_orm(column_name = "maxAttempts", default_value = 3)]
    pub max_attempts: i32,
    #[sea_orm(column_name = "lastError", column_type = "Text")]
    pub last_error: Option<String>,
    #[sea_orm(
        column_name = "lastFailedAttemptAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_failed_attempt_at: Option<DateTime>,
    #[sea_orm(column_name = "referenceUid", column_type = "Text")]
    pub reference_uid: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
        if insert && self.scheduled_at.is_not_set() {
            self.scheduled_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
