use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "revenuecat_webhook_events")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "eventType", column_type = "Text")]
    pub event_type: String,
    #[sea_orm(
        column_name = "eventTimestamp",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub event_timestamp: DateTime,
    #[sea_orm(column_name = "appUserId", column_type = "Text")]
    pub app_user_id: String,
    #[sea_orm(column_type = "Text")]
    pub environment: String,
    #[sea_orm(column_name = "productId", column_type = "Text")]
    pub product_id: Option<String>,
    #[sea_orm(column_name = "transactionId", column_type = "Text")]
    pub transaction_id: Option<String>,
    #[sea_orm(column_name = "originalTransactionId", column_type = "Text")]
    pub original_transaction_id: Option<String>,
    #[sea_orm(column_name = "entitlementIds", column_type = "Text")]
    pub entitlement_ids: Option<String>,
    #[sea_orm(default_value = false)]
    pub processed: bool,
    #[sea_orm(column_name = "processingError", column_type = "Text")]
    pub processing_error: Option<String>,
    #[sea_orm(column_name = "retryCount", default_value = 0)]
    pub retry_count: i32,
    #[sea_orm(column_name = "rawEventData", column_type = "JsonBinary")]
    pub raw_event_data: Json,
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
pub enum Relation {}

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
