use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Agreement")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub content: String,
    #[sea_orm(
        column_name = "contentType",
        column_type = "Text",
        default_value = "LINK"
    )]
    pub content_type: String,
    #[sea_orm(
        column_name = "signingProvider",
        column_type = "Text",
        default_value = "LEGACY"
    )]
    pub signing_provider: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "signingExternalId", column_type = "Text")]
    pub signing_external_id: Option<String>,
    #[sea_orm(column_name = "signingEnvelopeId", column_type = "Text")]
    pub signing_envelope_id: Option<String>,
    #[sea_orm(column_name = "signingTemplateId", column_type = "Text")]
    pub signing_template_id: Option<String>,
    #[sea_orm(column_name = "requireName", default_value = true)]
    pub require_name: bool,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "deletedAt", column_type = "custom(\"timestamp(3)\")")]
    pub deleted_at: Option<DateTime>,
    #[sea_orm(column_name = "deletedBy", column_type = "Text")]
    pub deleted_by: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "Agreement_teamId_fkey"
    )]
    Team,
    #[sea_orm(has_many = "super::link::Entity")]
    Links,
    #[sea_orm(has_many = "super::agreement_response::Entity")]
    Responses,
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Links.def()
    }
}

impl Related<super::agreement_response::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Responses.def()
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
