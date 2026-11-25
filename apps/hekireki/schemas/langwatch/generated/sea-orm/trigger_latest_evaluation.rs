use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TriggerLatestEvaluation")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "triggerId", column_type = "Text")]
    pub trigger_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(
        column_name = "evaluatedAt",
        column_type = "custom(\"timestamptz(3)\")"
    )]
    pub evaluated_at: DateTimeWithTimeZone,
    #[sea_orm(column_type = "Text")]
    pub verdict: String,
    #[sea_orm(column_name = "observedValue")]
    pub observed_value: Option<f64>,
    pub threshold: Option<f64>,
    #[sea_orm(column_type = "Text")]
    pub operator: Option<String>,
    #[sea_orm(column_name = "timePeriodMinutes")]
    pub time_period_minutes: Option<i32>,
    #[sea_orm(column_name = "skipCode", column_type = "Text")]
    pub skip_code: Option<String>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamptz(3)\")")]
    pub updated_at: DateTimeWithTimeZone,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::trigger::Entity",
        from = "Column::TriggerId",
        to = "super::trigger::Column::Id",
        on_delete = "Cascade",
        fk_name = "TriggerLatestEvaluation_triggerId_fkey"
    )]
    Trigger,
}

impl Related<super::trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Trigger.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, _insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
