use super::seat_change_type::SeatChangeType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SeatChangeLog")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "teamId")]
    pub team_id: i32,
    #[sea_orm(column_name = "changeType")]
    pub change_type: SeatChangeType,
    #[sea_orm(column_name = "seatCount")]
    pub seat_count: i32,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<i32>,
    #[sea_orm(column_name = "triggeredBy")]
    pub triggered_by: Option<i32>,
    #[sea_orm(
        column_name = "changeDate",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub change_date: DateTime,
    #[sea_orm(column_name = "monthKey", column_type = "Text")]
    pub month_key: String,
    #[sea_orm(column_name = "operationId", column_type = "Text")]
    pub operation_id: Option<String>,
    #[sea_orm(column_name = "processedInProrationId", column_type = "Text")]
    pub processed_in_proration_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "teamBillingId", column_type = "Text")]
    pub team_billing_id: Option<String>,
    #[sea_orm(column_name = "organizationBillingId", column_type = "Text")]
    pub organization_billing_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "SeatChangeLog_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::monthly_proration::Entity",
        from = "Column::ProcessedInProrationId",
        to = "super::monthly_proration::Column::Id",
        fk_name = "SeatChangeLog_processedInProrationId_fkey"
    )]
    Proration,
    #[sea_orm(
        belongs_to = "super::team_billing::Entity",
        from = "Column::TeamBillingId",
        to = "super::team_billing::Column::Id",
        fk_name = "SeatChangeLog_teamBillingId_fkey"
    )]
    TeamBilling,
    #[sea_orm(
        belongs_to = "super::organization_billing::Entity",
        from = "Column::OrganizationBillingId",
        to = "super::organization_billing::Column::Id",
        fk_name = "SeatChangeLog_organizationBillingId_fkey"
    )]
    OrganizationBilling,
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::monthly_proration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Proration.def()
    }
}

impl Related<super::team_billing::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TeamBilling.def()
    }
}

impl Related<super::organization_billing::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrganizationBilling.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.change_date.is_not_set() {
            self.change_date = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
