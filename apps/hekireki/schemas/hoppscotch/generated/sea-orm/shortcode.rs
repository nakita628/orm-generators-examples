use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Shortcode")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub request: Json,
    #[sea_orm(column_name = "creatorUid", column_type = "Text")]
    pub creator_uid: Option<String>,
    #[sea_orm(column_name = "embedProperties", column_type = "JsonBinary")]
    pub embed_properties: Option<Json>,
    #[sea_orm(
        column_name = "createdOn",
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_on: DateTimeWithTimeZone,
    #[sea_orm(
        column_name = "updatedOn",
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_on: DateTimeWithTimeZone,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatorUid",
        to = "super::user::Column::Uid",
        fk_name = "Shortcode_creatorUid_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_on.is_not_set() {
            self.created_on = Set(now.fixed_offset());
        }
        if !self.updated_on.is_set() {
            self.updated_on = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
