use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Each gateway virtual key, folded from gateway's vk_lifecycle facts: nurturing's own copy. The
/// creator comes from the created fact alone; the holder and purpose are each the latest non-null
/// value by the fact's time (stamped beside them); the state times keep their latest.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingVirtualKey")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "virtualKeyId", column_type = "Text")]
    pub virtual_key_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "occurredAt", column_type = "custom(\"timestamp(3)\")")]
    pub occurred_at: DateTime,
    #[sea_orm(column_name = "keyCreatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub key_created_at: Option<DateTime>,
    #[sea_orm(column_name = "creatorUserId", column_type = "Text")]
    pub creator_user_id: Option<String>,
    #[sea_orm(column_name = "principalUserId", column_type = "Text")]
    pub principal_user_id: Option<String>,
    #[sea_orm(
        column_name = "principalUserIdAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub principal_user_id_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub purpose: Option<String>,
    #[sea_orm(column_name = "purposeAt", column_type = "custom(\"timestamp(3)\")")]
    pub purpose_at: Option<DateTime>,
    #[sea_orm(column_name = "disabledAt", column_type = "custom(\"timestamp(3)\")")]
    pub disabled_at: Option<DateTime>,
    #[sea_orm(column_name = "enabledAt", column_type = "custom(\"timestamp(3)\")")]
    pub enabled_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedAt", column_type = "custom(\"timestamp(3)\")")]
    pub revoked_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
