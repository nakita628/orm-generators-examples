use super::response_quota_link_status::ResponseQuotaLinkStatus;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Junction table linking responses to quotas.
/// Tracks which responses counted towards specific quotas and their status.
///
/// @property response - The response that counted towards the quota
/// @property quota - The quota this response counted towards
/// @property status - Whether the response was screened in or out
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ResponseQuotaLink")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "responseId", column_type = "Text")]
    pub response_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "quotaId", column_type = "Text")]
    pub quota_id: String,
    pub status: ResponseQuotaLinkStatus,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::response::Entity",
        from = "Column::ResponseId",
        to = "super::response::Column::Id",
        on_delete = "Cascade",
        fk_name = "ResponseQuotaLink_responseId_fkey"
    )]
    Response,
    #[sea_orm(
        belongs_to = "super::survey_quota::Entity",
        from = "Column::QuotaId",
        to = "super::survey_quota::Column::Id",
        on_delete = "Cascade",
        fk_name = "ResponseQuotaLink_quotaId_fkey"
    )]
    Quota,
}

impl Related<super::response::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Response.def()
    }
}

impl Related<super::survey_quota::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Quota.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
