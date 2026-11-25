use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "prices")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub model_id: String,
    #[sea_orm(column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub pricing_tier_id: String,
    #[sea_orm(column_type = "Text")]
    pub usage_type: String,
    #[sea_orm(column_type = "Decimal(Some((65, 30)))")]
    pub price: Decimal,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::model::Entity",
        from = "Column::ModelId",
        to = "super::model::Column::Id",
        on_delete = "Cascade",
        fk_name = "prices_model_id_fkey"
    )]
    Model,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "prices_project_id_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::pricing_tier::Entity",
        from = "Column::PricingTierId",
        to = "super::pricing_tier::Column::Id",
        on_delete = "Cascade",
        fk_name = "prices_pricing_tier_id_fkey"
    )]
    PricingTier,
}

impl Related<super::model::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Model.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::pricing_tier::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PricingTier.def()
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
