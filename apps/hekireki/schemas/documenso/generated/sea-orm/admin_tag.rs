use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZAdminTagTypeSchema } from '@documenso/lib/types/admin-tags';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AdminTag")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    /// [AdminTagType] @zod.custom.use(ZAdminTagTypeSchema)
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_type = "Text")]
    pub description: String,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<i32>,
    #[sea_orm(column_name = "organisationId", column_type = "Text")]
    pub organisation_id: Option<String>,
    #[sea_orm(column_name = "createdByUserId")]
    pub created_by_user_id: Option<i32>,
    #[sea_orm(column_name = "adminRuleId", column_type = "Text")]
    pub admin_rule_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "AdminTag_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::organisation::Entity",
        from = "Column::OrganisationId",
        to = "super::organisation::Column::Id",
        on_delete = "Cascade",
        fk_name = "AdminTag_organisationId_fkey"
    )]
    Organisation,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedByUserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "AdminTag_createdByUserId_fkey"
    )]
    CreatedByUser,
    #[sea_orm(
        belongs_to = "super::admin_rule::Entity",
        from = "Column::AdminRuleId",
        to = "super::admin_rule::Column::Id",
        on_delete = "SetNull",
        fk_name = "AdminTag_adminRuleId_fkey"
    )]
    AdminRule,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::organisation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organisation.def()
    }
}

impl Related<super::admin_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AdminRule.def()
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
