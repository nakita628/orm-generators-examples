use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A known LLM model or model pattern for cost tracking
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "llm_models")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub friendly_id: String,
    #[sea_orm(column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub model_name: String,
    #[sea_orm(column_type = "Text")]
    pub match_pattern: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub start_date: Option<DateTime>,
    #[sea_orm(column_type = "Text", default_value = "default")]
    pub source: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub provider: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    pub context_window: Option<i32>,
    pub max_output_tokens: Option<i32>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub capabilities: Vec<String>,
    #[sea_orm(default_value = false)]
    pub is_hidden: bool,
    #[sea_orm(column_type = "Text")]
    pub base_model_name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub pricing_unit: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "llm_models_project_id_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::llm_pricing_tier::Entity")]
    PricingTiers,
    #[sea_orm(has_many = "super::llm_price::Entity")]
    Prices,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::llm_pricing_tier::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PricingTiers.def()
    }
}

impl Related<super::llm_price::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Prices.def()
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
