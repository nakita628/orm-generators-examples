use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The mail log for seat requests and refused sign-ins: at most one row per organization,
/// kind, address and UTC day, each mailed at once.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganizationAccessNotice")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// seat_request | idp_refused
    #[sea_orm(column_type = "Text")]
    pub kind: String,
    #[sea_orm(column_type = "Text")]
    pub email: String,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub idp: Option<String>,
    /// UTC day the notice was recorded on, YYYY-MM-DD.
    #[sea_orm(column_name = "dayKey", column_type = "Text")]
    pub day_key: String,
    /// Whether the seat request it announces was also refused by the identity provider.
    #[sea_orm(column_name = "idpRefused", default_value = false)]
    pub idp_refused: bool,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    /// Set once the notice was handed to the mailer; null if a send never started.
    #[sea_orm(column_name = "mailedAt", column_type = "custom(\"timestamp(3)\")")]
    pub mailed_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
        Ok(self)
    }
}
