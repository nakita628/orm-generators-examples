use super::email_domain_status::EmailDomainStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "EmailDomain")]
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
    #[sea_orm(default_value = "PENDING")]
    pub status: EmailDomainStatus,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub selector: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub domain: String,
    #[sea_orm(column_name = "publicKey", column_type = "Text")]
    pub public_key: String,
    #[sea_orm(column_name = "privateKey", column_type = "Text")]
    pub private_key: String,
    #[sea_orm(
        column_name = "lastVerifiedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_verified_at: Option<DateTime>,
    #[sea_orm(column_name = "organisationId", column_type = "Text")]
    pub organisation_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organisation::Entity",
        from = "Column::OrganisationId",
        to = "super::organisation::Column::Id",
        on_delete = "Cascade",
        fk_name = "EmailDomain_organisationId_fkey"
    )]
    Organisation,
    #[sea_orm(has_many = "super::organisation_email::Entity")]
    Emails,
}

impl Related<super::organisation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organisation.def()
    }
}

impl Related<super::organisation_email::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Emails.def()
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
