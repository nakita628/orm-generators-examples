use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "VerificationToken")]
pub struct Model {
    #[sea_orm(column_type = "Text")]
    pub identifier: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub token: String,
    #[sea_orm(column_type = "Text")]
    pub value: Option<String>,
    #[sea_orm(column_name = "failedAttempts", default_value = 0)]
    pub failed_attempts: i32,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub expires: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
