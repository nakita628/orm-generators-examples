use super::activity_action::ActivityAction;
use super::activity_entity::ActivityEntity;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ActivityEvent")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "occurredAt", column_type = "TimestampWithTimeZone")]
    pub occurred_at: DateTimeWithTimeZone,
    #[sea_orm(column_name = "actorUserId")]
    pub actor_user_id: Option<String>,
    #[sea_orm(column_name = "actorSnapshot")]
    pub actor_snapshot: Option<Json>,
    pub action: ActivityAction,
    #[sea_orm(column_name = "entityType")]
    pub entity_type: ActivityEntity,
    #[sea_orm(column_name = "entityId")]
    pub entity_id: String,
    #[sea_orm(column_name = "assetId")]
    pub asset_id: Option<String>,
    #[sea_orm(column_name = "bookingId")]
    pub booking_id: Option<String>,
    #[sea_orm(column_name = "auditSessionId")]
    pub audit_session_id: Option<String>,
    #[sea_orm(column_name = "auditAssetId")]
    pub audit_asset_id: Option<String>,
    #[sea_orm(column_name = "kitId")]
    pub kit_id: Option<String>,
    #[sea_orm(column_name = "locationId")]
    pub location_id: Option<String>,
    #[sea_orm(column_name = "teamMemberId")]
    pub team_member_id: Option<String>,
    #[sea_orm(column_name = "targetUserId")]
    pub target_user_id: Option<String>,
    pub field: Option<String>,
    #[sea_orm(column_name = "fromValue")]
    pub from_value: Option<Json>,
    #[sea_orm(column_name = "toValue")]
    pub to_value: Option<Json>,
    pub meta: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.occurred_at.is_not_set() {
            self.occurred_at = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
