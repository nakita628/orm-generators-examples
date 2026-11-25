use super::booking_audit_action::BookingAuditAction;
use super::booking_audit_source::BookingAuditSource;
use super::booking_audit_type::BookingAuditType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BookingAudit")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: Uuid,
    #[sea_orm(column_name = "bookingUid", column_type = "Text")]
    pub booking_uid: String,
    #[sea_orm(column_name = "actorId")]
    pub actor_id: Uuid,
    pub r#type: BookingAuditType,
    pub action: BookingAuditAction,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub timestamp: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    pub source: BookingAuditSource,
    #[sea_orm(column_name = "operationId", column_type = "Text")]
    pub operation_id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub data: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub context: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::audit_actor::Entity",
        from = "Column::ActorId",
        to = "super::audit_actor::Column::Id",
        on_delete = "Restrict",
        fk_name = "BookingAudit_actorId_fkey"
    )]
    Actor,
}

impl Related<super::audit_actor::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Actor.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::now_v7()),
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
