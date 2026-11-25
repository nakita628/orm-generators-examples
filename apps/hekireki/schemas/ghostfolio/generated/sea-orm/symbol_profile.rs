use super::asset_class::AssetClass;
use super::asset_sub_class::AssetSubClass;
use super::data_gathering_frequency::DataGatheringFrequency;
use super::data_source::DataSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SymbolProfile")]
pub struct Model {
    #[sea_orm(column_name = "assetClass")]
    pub asset_class: Option<AssetClass>,
    #[sea_orm(column_name = "assetSubClass")]
    pub asset_sub_class: Option<AssetSubClass>,
    #[sea_orm(column_type = "Text")]
    pub comment: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub countries: Option<Json>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub currency: String,
    #[sea_orm(column_type = "Text")]
    pub cusip: Option<String>,
    #[sea_orm(column_name = "dataGatheringFrequency", default_value = "DAILY")]
    pub data_gathering_frequency: DataGatheringFrequency,
    #[sea_orm(column_name = "dataSource")]
    pub data_source: DataSource,
    #[sea_orm(column_type = "Text")]
    pub figi: Option<String>,
    #[sea_orm(column_name = "figiComposite", column_type = "Text")]
    pub figi_composite: Option<String>,
    #[sea_orm(column_name = "figiShareClass", column_type = "Text")]
    pub figi_share_class: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub holdings: Option<Json>,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "isActive", default_value = true)]
    pub is_active: bool,
    #[sea_orm(column_type = "Text")]
    pub isin: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_name = "scraperConfiguration", column_type = "JsonBinary")]
    pub scraper_configuration: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub sectors: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub symbol: String,
    #[sea_orm(column_name = "symbolMapping", column_type = "JsonBinary")]
    pub symbol_mapping: Option<Json>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub url: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "SymbolProfile_userId_fkey"
    )]
    User,
    #[sea_orm(has_many = "super::order::Entity")]
    Activities,
    #[sea_orm(has_many = "super::asset_profile_split::Entity")]
    AssetProfileSplits,
    #[sea_orm(has_one = "super::asset_profile_overrides::Entity")]
    AssetProfileOverrides,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::order::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Activities.def()
    }
}

impl Related<super::asset_profile_split::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetProfileSplits.def()
    }
}

impl Related<super::asset_profile_overrides::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetProfileOverrides.def()
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
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
