use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "automations")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub trigger_id: String,
    #[sea_orm(column_type = "Text")]
    pub action_id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::trigger::Entity",
        from = "Column::TriggerId",
        to = "super::trigger::Column::Id",
        on_delete = "Cascade",
        fk_name = "automations_trigger_id_fkey"
    )]
    Trigger,
    #[sea_orm(
        belongs_to = "super::action::Entity",
        from = "Column::ActionId",
        to = "super::action::Column::Id",
        on_delete = "Cascade",
        fk_name = "automations_action_id_fkey"
    )]
    Action,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "automations_project_id_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::automation_execution::Entity")]
    AutomationExecution,
}

impl Related<super::trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Trigger.def()
    }
}

impl Related<super::action::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Action.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::automation_execution::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AutomationExecution.def()
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
