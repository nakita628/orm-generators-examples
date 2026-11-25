use super::processor_type::ProcessorType;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Payment")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "premiumId", column_type = "Text")]
    pub premium_id: Option<String>,
    #[sea_orm(column_name = "processorType", default_value = "LEMON_SQUEEZY")]
    pub processor_type: ProcessorType,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "processorId", column_type = "Text")]
    pub processor_id: Option<String>,
    #[sea_orm(column_name = "processorSubscriptionId", column_type = "Text")]
    pub processor_subscription_id: Option<String>,
    #[sea_orm(column_name = "processorCustomerId", column_type = "Text")]
    pub processor_customer_id: Option<String>,
    pub amount: i32,
    #[sea_orm(column_type = "Text")]
    pub currency: String,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    pub tax: i32,
    #[sea_orm(column_name = "taxInclusive")]
    pub tax_inclusive: bool,
    #[sea_orm(default_value = false)]
    pub refunded: bool,
    #[sea_orm(column_name = "refundedAt", column_type = "custom(\"timestamp(3)\")")]
    pub refunded_at: Option<DateTime>,
    #[sea_orm(column_name = "refundedAmount")]
    pub refunded_amount: Option<i32>,
    #[sea_orm(column_name = "billingReason", column_type = "Text")]
    pub billing_reason: Option<String>,
    #[sea_orm(
        column_name = "invoiceEmailSentAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub invoice_email_sent_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::premium::Entity",
        from = "Column::PremiumId",
        to = "super::premium::Column::Id",
        on_delete = "SetNull",
        fk_name = "Payment_premiumId_fkey"
    )]
    Premium,
}

impl Related<super::premium::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Premium.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
