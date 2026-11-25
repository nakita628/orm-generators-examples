use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RuleHistory")]
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
    #[sea_orm(column_name = "ruleId", column_type = "Text")]
    pub rule_id: String,
    pub version: i32,
    #[sea_orm(column_name = "triggerType", column_type = "Text")]
    pub trigger_type: String,
    #[sea_orm(column_name = "promptText", column_type = "Text")]
    pub prompt_text: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub instructions: Option<String>,
    pub enabled: bool,
    pub automate: bool,
    #[sea_orm(column_name = "runOnThreads")]
    pub run_on_threads: bool,
    #[sea_orm(column_name = "conditionalOperator", column_type = "Text")]
    pub conditional_operator: String,
    #[sea_orm(column_type = "Text")]
    pub from: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub to: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub subject: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub body: Option<String>,
    #[sea_orm(column_name = "categoryFilterType", column_type = "Text")]
    pub category_filter_type: Option<String>,
    #[sea_orm(column_name = "systemType", column_type = "Text")]
    pub system_type: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub actions: Json,
    #[sea_orm(column_name = "categoryFilters", column_type = "JsonBinary")]
    pub category_filters: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::rule::Entity",
        from = "Column::RuleId",
        to = "super::rule::Column::Id",
        on_delete = "Cascade",
        fk_name = "RuleHistory_ruleId_fkey"
    )]
    Rule,
}

impl Related<super::rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Rule.def()
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
