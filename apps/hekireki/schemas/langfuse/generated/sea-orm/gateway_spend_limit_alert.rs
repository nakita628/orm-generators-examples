use super::gateway_spend_limit_alert_type::GatewaySpendLimitAlertType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One notification per limit, subject, alert type, threshold, counting start and window, so the gateway may report a limit any number of times.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "gateway_spend_limit_alerts")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub spend_limit_id: String,
    #[sea_orm(default_value = "")]
    pub subject_id: String,
    pub r#type: GatewaySpendLimitAlertType,
    #[sea_orm(default_value = 100)]
    pub threshold_percent: i32,
    pub counting_started_at: DateTime,
    pub window_start: DateTime,
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::gateway_spend_limit::Entity",
        from = "Column::SpendLimitId",
        to = "super::gateway_spend_limit::Column::Id",
        on_delete = "Cascade"
    )]
    SpendLimit,
}

impl Related<super::gateway_spend_limit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SpendLimit.def()
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
        Ok(self)
    }
}
