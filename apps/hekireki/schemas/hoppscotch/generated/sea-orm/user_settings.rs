use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "UserSettings")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "userUid", column_type = "Text")]
    pub user_uid: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub properties: Json,
    #[sea_orm(column_name = "updatedOn", column_type = "custom(\"timestamptz(3)\")")]
    pub updated_on: DateTimeWithTimeZone,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserUid",
        to = "super::user::Column::Uid",
        on_delete = "Cascade",
        fk_name = "UserSettings_userUid_fkey"
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
    async fn before_save<C>(mut self, _db: &C, _insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if !self.updated_on.is_set() {
            self.updated_on = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
