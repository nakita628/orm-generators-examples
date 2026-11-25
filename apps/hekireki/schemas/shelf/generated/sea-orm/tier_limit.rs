use super::tier_id::TierId;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TierLimit")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: TierId,
    #[sea_orm(column_name = "canImportAssets", default_value = false)]
    pub can_import_assets: bool,
    #[sea_orm(column_name = "canExportAssets", default_value = false)]
    pub can_export_assets: bool,
    #[sea_orm(column_name = "canImportNRM", default_value = false)]
    pub can_import_nrm: bool,
    #[sea_orm(column_name = "canHideShelfBranding", default_value = false)]
    pub can_hide_shelf_branding: bool,
    #[sea_orm(column_name = "canUseReports", default_value = false)]
    pub can_use_reports: bool,
    #[sea_orm(column_name = "canUseRecurringReminders", default_value = false)]
    pub can_use_recurring_reminders: bool,
    #[sea_orm(column_name = "canUsePushNotifications", default_value = false)]
    pub can_use_push_notifications: bool,
    #[sea_orm(column_name = "maxCustomFields", default_value = 0)]
    pub max_custom_fields: i32,
    #[sea_orm(column_name = "maxOrganizations", default_value = 1)]
    pub max_organizations: i32,
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
    #[sea_orm(has_one = "super::tier::Entity")]
    Tier,
}

impl Related<super::tier::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tier.def()
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
