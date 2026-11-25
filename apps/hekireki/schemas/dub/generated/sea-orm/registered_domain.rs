use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RegisteredDomain")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub slug: String,
    #[sea_orm(column_name = "projectId", column_type = "String(StringLen::N(191))")]
    pub project_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "domainId", column_type = "String(StringLen::N(191))")]
    pub domain_id: Option<String>,
    #[sea_orm(
        column_name = "autoRenewalDisabledAt",
        column_type = "custom(\"datetime(3)\")"
    )]
    pub auto_renewal_disabled_at: Option<DateTime>,
    #[sea_orm(column_name = "renewalFee", default_value = 1200)]
    pub renewal_fee: i32,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"datetime(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"datetime(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "RegisteredDomain_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::domain::Entity",
        from = "Column::DomainId",
        to = "super::domain::Column::Id",
        on_delete = "SetNull",
        fk_name = "RegisteredDomain_domainId_fkey"
    )]
    Domain,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::domain::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Domain.def()
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
