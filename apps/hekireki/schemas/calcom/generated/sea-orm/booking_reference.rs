use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BookingReference")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    /// @zod.string.min(1)
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    /// @zod.string.min(1)
    #[sea_orm(column_type = "Text")]
    pub uid: String,
    #[sea_orm(column_name = "meetingId", column_type = "Text")]
    pub meeting_id: Option<String>,
    #[sea_orm(column_name = "thirdPartyRecurringEventId", column_type = "Text")]
    pub third_party_recurring_event_id: Option<String>,
    #[sea_orm(column_name = "meetingPassword", column_type = "Text")]
    pub meeting_password: Option<String>,
    #[sea_orm(column_name = "meetingUrl", column_type = "Text")]
    pub meeting_url: Option<String>,
    #[sea_orm(column_name = "bookingId")]
    pub booking_id: Option<i32>,
    #[sea_orm(column_name = "externalCalendarId", column_type = "Text")]
    pub external_calendar_id: Option<String>,
    pub deleted: Option<bool>,
    #[sea_orm(column_name = "credentialId")]
    pub credential_id: Option<i32>,
    #[sea_orm(column_name = "delegationCredentialId", column_type = "Text")]
    pub delegation_credential_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::booking::Entity",
        from = "Column::BookingId",
        to = "super::booking::Column::Id",
        on_delete = "Cascade",
        fk_name = "BookingReference_bookingId_fkey"
    )]
    Booking,
    #[sea_orm(
        belongs_to = "super::credential::Entity",
        from = "Column::CredentialId",
        to = "super::credential::Column::Id",
        on_delete = "SetNull",
        fk_name = "BookingReference_credentialId_fkey"
    )]
    Credential,
    #[sea_orm(
        belongs_to = "super::delegation_credential::Entity",
        from = "Column::DelegationCredentialId",
        to = "super::delegation_credential::Column::Id",
        on_delete = "SetNull",
        fk_name = "BookingReference_delegationCredentialId_fkey"
    )]
    DelegationCredential,
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Booking.def()
    }
}

impl Related<super::credential::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Credential.def()
    }
}

impl Related<super::delegation_credential::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DelegationCredential.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
