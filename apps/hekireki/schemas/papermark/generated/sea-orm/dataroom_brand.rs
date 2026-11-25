use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DataroomBrand")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub logo: Option<String>,
    #[sea_orm(column_name = "hideLogo")]
    pub hide_logo: Option<bool>,
    #[sea_orm(column_type = "Text")]
    pub banner: Option<String>,
    #[sea_orm(column_name = "brandColor", column_type = "Text")]
    pub brand_color: Option<String>,
    #[sea_orm(column_name = "accentColor", column_type = "Text")]
    pub accent_color: Option<String>,
    #[sea_orm(column_name = "accentButtonColor", column_type = "Text")]
    pub accent_button_color: Option<String>,
    #[sea_orm(column_name = "applyAccentColorToDataroomView", default_value = false)]
    pub apply_accent_color_to_dataroom_view: bool,
    #[sea_orm(column_name = "welcomeMessage", column_type = "Text")]
    pub welcome_message: Option<String>,
    #[sea_orm(
        column_name = "cardLayout",
        column_type = "Text",
        default_value = "LIST"
    )]
    pub card_layout: String,
    #[sea_orm(column_name = "showFolderTree", default_value = true)]
    pub show_folder_tree: bool,
    #[sea_orm(
        column_name = "viewerLayoutPreset",
        column_type = "Text",
        default_value = "STANDARD"
    )]
    pub viewer_layout_preset: String,
    #[sea_orm(
        column_name = "viewerHeaderStyle",
        column_type = "Text",
        default_value = "DEFAULT"
    )]
    pub viewer_header_style: String,
    #[sea_orm(column_name = "hideFolderIconsInMain", default_value = false)]
    pub hide_folder_icons_in_main: bool,
    #[sea_orm(column_name = "ctaLabel", column_type = "Text")]
    pub cta_label: Option<String>,
    #[sea_orm(column_name = "ctaUrl", column_type = "Text")]
    pub cta_url: Option<String>,
    #[sea_orm(column_name = "customLinkPreviewEnabled", default_value = false)]
    pub custom_link_preview_enabled: bool,
    #[sea_orm(column_name = "linkPreviewTitle", column_type = "Text")]
    pub link_preview_title: Option<String>,
    #[sea_orm(column_name = "linkPreviewDescription", column_type = "Text")]
    pub link_preview_description: Option<String>,
    #[sea_orm(column_name = "linkPreviewImage", column_type = "Text")]
    pub link_preview_image: Option<String>,
    #[sea_orm(column_name = "linkPreviewFavicon", column_type = "Text")]
    pub link_preview_favicon: Option<String>,
    #[sea_orm(
        column_name = "defaultLanguage",
        column_type = "Text",
        default_value = "en"
    )]
    pub default_language: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "dataroomId", column_type = "Text")]
    pub dataroom_id: String,
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
        belongs_to = "super::dataroom::Entity",
        from = "Column::DataroomId",
        to = "super::dataroom::Column::Id",
        on_delete = "Cascade",
        fk_name = "DataroomBrand_dataroomId_fkey"
    )]
    Dataroom,
}

impl Related<super::dataroom::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dataroom.def()
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
