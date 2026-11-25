use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LinkPreset")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "pId", column_type = "Text")]
    pub p_id: Option<String>,
    #[sea_orm(column_name = "enableCustomMetaTag", default_value = false)]
    pub enable_custom_meta_tag: Option<bool>,
    #[sea_orm(column_name = "metaTitle", column_type = "Text")]
    pub meta_title: Option<String>,
    #[sea_orm(column_name = "metaDescription", column_type = "Text")]
    pub meta_description: Option<String>,
    #[sea_orm(column_name = "metaImage", column_type = "Text")]
    pub meta_image: Option<String>,
    #[sea_orm(column_name = "metaFavicon", column_type = "Text")]
    pub meta_favicon: Option<String>,
    #[sea_orm(column_name = "enableNotification", default_value = false)]
    pub enable_notification: Option<bool>,
    #[sea_orm(column_name = "emailProtected", default_value = true)]
    pub email_protected: Option<bool>,
    #[sea_orm(column_name = "emailAuthenticated", default_value = false)]
    pub email_authenticated: Option<bool>,
    #[sea_orm(column_name = "allowDownload", default_value = false)]
    pub allow_download: Option<bool>,
    #[sea_orm(column_name = "enableAllowList", default_value = false)]
    pub enable_allow_list: Option<bool>,
    #[sea_orm(
        column_name = "allowList",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub allow_list: Vec<String>,
    #[sea_orm(column_name = "enableDenyList", default_value = false)]
    pub enable_deny_list: Option<bool>,
    #[sea_orm(
        column_name = "denyList",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub deny_list: Vec<String>,
    #[sea_orm(column_name = "expiresIn")]
    pub expires_in: Option<i32>,
    #[sea_orm(column_name = "enableScreenshotProtection", default_value = false)]
    pub enable_screenshot_protection: Option<bool>,
    #[sea_orm(column_name = "enableConfidentialView", default_value = false)]
    pub enable_confidential_view: Option<bool>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_name = "enablePassword", default_value = false)]
    pub enable_password: Option<bool>,
    #[sea_orm(column_type = "Text")]
    pub password: Option<String>,
    #[sea_orm(column_name = "enableWatermark", default_value = false)]
    pub enable_watermark: Option<bool>,
    #[sea_orm(column_name = "watermarkConfig", column_type = "JsonBinary")]
    pub watermark_config: Option<Json>,
    #[sea_orm(column_name = "enableAgreement", default_value = false)]
    pub enable_agreement: Option<bool>,
    #[sea_orm(column_name = "agreementId", column_type = "Text")]
    pub agreement_id: Option<String>,
    #[sea_orm(column_name = "enableCustomFields", default_value = false)]
    pub enable_custom_fields: Option<bool>,
    #[sea_orm(column_name = "customFields", column_type = "JsonBinary")]
    pub custom_fields: Option<Json>,
    #[sea_orm(column_name = "showBanner", default_value = false)]
    pub show_banner: Option<bool>,
    #[sea_orm(column_name = "welcomeMessage", column_type = "Text")]
    pub welcome_message: Option<String>,
    #[sea_orm(column_name = "isDefault", default_value = false)]
    pub is_default: bool,
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
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "LinkPreset_teamId_fkey"
    )]
    Team,
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
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
