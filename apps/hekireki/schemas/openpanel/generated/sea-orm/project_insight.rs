use super::insight_state::InsightState;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "project_insights")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(default_expr = "Expr::cust(\"gen_random_uuid()\")")]
    pub id: Uuid,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "moduleKey", column_type = "Text")]
    pub module_key: String,
    #[sea_orm(column_name = "dimensionKey", column_type = "Text")]
    pub dimension_key: String,
    #[sea_orm(column_name = "windowKind", column_type = "Text")]
    pub window_kind: String,
    #[sea_orm(default_value = "active")]
    pub state: InsightState,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_type = "Text")]
    pub summary: Option<String>,
    #[sea_orm(column_name = "displayName", column_type = "Text", default_value = "")]
    pub display_name: String,
    /// [IPrismaProjectInsightPayload]
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub payload: Json,
    #[sea_orm(column_type = "Text")]
    pub direction: Option<String>,
    #[sea_orm(column_name = "impactScore", default_value = 0)]
    pub impact_score: f64,
    #[sea_orm(column_name = "severityBand", column_type = "Text")]
    pub severity_band: Option<String>,
    #[sea_orm(column_name = "relevanceScore")]
    pub relevance_score: Option<f64>,
    #[sea_orm(column_name = "aiSummary", column_type = "Text")]
    pub ai_summary: Option<String>,
    #[sea_orm(column_name = "aiCategory", column_type = "Text")]
    pub ai_category: Option<String>,
    #[sea_orm(column_name = "emailWorthy", default_value = false)]
    pub email_worthy: bool,
    #[sea_orm(column_name = "referenceWorthy", default_value = false)]
    pub reference_worthy: bool,
    #[sea_orm(column_name = "enrichedAt", column_type = "custom(\"timestamp(3)\")")]
    pub enriched_at: Option<DateTime>,
    #[sea_orm(column_name = "enrichVersion")]
    pub enrich_version: Option<i32>,
    #[sea_orm(default_value = 1)]
    pub version: i32,
    #[sea_orm(
        column_name = "threadId",
        default_expr = "Expr::cust(\"gen_random_uuid()\")"
    )]
    pub thread_id: Uuid,
    #[sea_orm(column_name = "windowStart", column_type = "custom(\"timestamp(3)\")")]
    pub window_start: Option<DateTime>,
    #[sea_orm(column_name = "windowEnd", column_type = "custom(\"timestamp(3)\")")]
    pub window_end: Option<DateTime>,
    #[sea_orm(
        column_name = "firstDetectedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub first_detected_at: DateTime,
    #[sea_orm(
        column_name = "lastUpdatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub last_updated_at: DateTime,
    #[sea_orm(
        column_name = "lastSeenAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub last_seen_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::insight_event::Entity")]
    Events,
}

impl Related<super::insight_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Events.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.first_detected_at.is_not_set() {
            self.first_detected_at = Set(now.naive_utc());
        }
        if !self.last_updated_at.is_set() {
            self.last_updated_at = Set(now.naive_utc());
        }
        if insert && self.last_seen_at.is_not_set() {
            self.last_seen_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
