use super::wrong_assignment_report_status::WrongAssignmentReportStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WrongAssignmentReport")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: Uuid,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "bookingUid", column_type = "Text")]
    pub booking_uid: String,
    #[sea_orm(column_name = "reportedById")]
    pub reported_by_id: Option<i32>,
    #[sea_orm(column_name = "correctAssignee", column_type = "Text")]
    pub correct_assignee: Option<String>,
    #[sea_orm(column_name = "additionalNotes", column_type = "Text")]
    pub additional_notes: String,
    #[sea_orm(column_name = "teamId")]
    pub team_id: Option<i32>,
    #[sea_orm(default_value = "PENDING")]
    pub status: WrongAssignmentReportStatus,
    #[sea_orm(column_name = "reviewedById")]
    pub reviewed_by_id: Option<i32>,
    #[sea_orm(column_name = "reviewedAt", column_type = "custom(\"timestamp(3)\")")]
    pub reviewed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::booking::Entity",
        from = "Column::BookingUid",
        to = "super::booking::Column::Uid",
        on_delete = "Cascade",
        fk_name = "WrongAssignmentReport_bookingUid_fkey"
    )]
    Booking,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ReportedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "WrongAssignmentReport_reportedById_fkey"
    )]
    ReportedBy,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "SetNull",
        fk_name = "WrongAssignmentReport_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ReviewedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "WrongAssignmentReport_reviewedById_fkey"
    )]
    ReviewedBy,
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Booking.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ReportedBy.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4()),
            ..ActiveModelTrait::default()
        }
    }

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
