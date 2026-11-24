use super::booking_status::BookingStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Booking")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub name: String,
    #[sea_orm(default_value = "DRAFT")]
    pub status: BookingStatus,
    #[sea_orm(default_value = "")]
    pub description: Option<String>,
    #[sea_orm(column_name = "activeSchedulerReference")]
    pub active_scheduler_reference: Option<String>,
    #[sea_orm(column_name = "creatorId")]
    pub creator_id: String,
    #[sea_orm(column_name = "custodianUserId")]
    pub custodian_user_id: Option<String>,
    #[sea_orm(column_name = "custodianTeamMemberId")]
    pub custodian_team_member_id: Option<String>,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "createdAt", column_type = "TimestampWithTimeZone")]
    pub created_at: DateTimeWithTimeZone,
    #[sea_orm(column_name = "updatedAt", column_type = "TimestampWithTimeZone")]
    pub updated_at: DateTimeWithTimeZone,
    #[sea_orm(column_type = "TimestampWithTimeZone")]
    pub from: DateTimeWithTimeZone,
    #[sea_orm(column_type = "TimestampWithTimeZone")]
    pub to: DateTimeWithTimeZone,
    #[sea_orm(column_name = "originalFrom", column_type = "TimestampWithTimeZone")]
    pub original_from: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "originalTo", column_type = "TimestampWithTimeZone")]
    pub original_to: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "autoArchivedAt", column_type = "TimestampWithTimeZone")]
    pub auto_archived_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "archivedWithoutCheckin", default_value = false)]
    pub archived_without_checkin: bool,
    #[sea_orm(column_name = "cancellationReason")]
    pub cancellation_reason: Option<String>,
    #[sea_orm(column_name = "approvedAt", column_type = "TimestampWithTimeZone")]
    pub approved_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "approvedById")]
    pub approved_by_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatorId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Creator,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CustodianUserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    CustodianUser,
    #[sea_orm(
        belongs_to = "super::team_member::Entity",
        from = "Column::CustodianTeamMemberId",
        to = "super::team_member::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    CustodianTeamMember,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ApprovedById",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull"
    )]
    ApprovedBy,
    #[sea_orm(has_many = "super::booking_asset::Entity")]
    BookingAssets,
    #[sea_orm(has_many = "super::consumption_log::Entity")]
    ConsumptionLogs,
    #[sea_orm(has_many = "super::booking_note::Entity")]
    Notes,
    #[sea_orm(has_many = "super::booking_model_request::Entity")]
    ModelRequests,
    #[sea_orm(has_many = "super::partial_booking_checkin::Entity")]
    PartialCheckins,
    #[sea_orm(has_many = "super::partial_booking_checkout::Entity")]
    PartialCheckouts,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creator.def()
    }
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustodianTeamMember.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::booking_asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BookingAssets.def()
    }
}

impl Related<super::consumption_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ConsumptionLogs.def()
    }
}

impl Related<super::booking_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notes.def()
    }
}

impl Related<super::booking_model_request::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ModelRequests.def()
    }
}

impl Related<super::partial_booking_checkin::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PartialCheckins.def()
    }
}

impl Related<super::partial_booking_checkout::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PartialCheckouts.def()
    }
}

impl Related<super::tag::Entity> for Entity {
    fn to() -> RelationDef {
        super::booking_to_tag::Relation::Tag.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::booking_to_tag::Relation::Booking.def().rev())
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
            self.created_at = Set(now.fixed_offset());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
