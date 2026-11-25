use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "oauth_app")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub logo: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub homepage: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_type = "Text")]
    pub redirect_uris: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub scopes: Option<String>,
    #[sea_orm(default_value = false)]
    pub allow_device_flow: bool,
    #[sea_orm(default_value = false)]
    pub dynamic_registration: bool,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_time: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_modified_time: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub created_by: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::oauth_app_secret::Entity")]
    Secrets,
    #[sea_orm(has_many = "super::oauth_app_authorized::Entity")]
    Authorized,
    #[sea_orm(has_many = "super::oauth_app_token::Entity")]
    Tokens,
}

impl Related<super::oauth_app_secret::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Secrets.def()
    }
}

impl Related<super::oauth_app_authorized::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Authorized.def()
    }
}

impl Related<super::oauth_app_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tokens.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_time.is_not_set() {
            self.created_time = Set(now.naive_utc());
        }
        if !self.last_modified_time.is_set() {
            self.last_modified_time = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
