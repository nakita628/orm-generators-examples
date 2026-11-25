use super::gateway_guardrail_assignment_scope::GatewayGuardrailAssignmentScope;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "gateway_guardrail_assignments")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub guardrail_id: String,
    #[sea_orm(column_type = "Text")]
    pub organization_id: String,
    pub scope: GatewayGuardrailAssignmentScope,
    #[sea_orm(column_type = "Text")]
    pub api_key_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::gateway_guardrail::Entity",
        from = "Column::GuardrailId",
        to = "super::gateway_guardrail::Column::Id",
        on_delete = "Cascade",
        fk_name = "gateway_guardrail_assignments_guardrail_id_fkey"
    )]
    Guardrail,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "gateway_guardrail_assignments_organization_id_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::api_key::Entity",
        from = "Column::ApiKeyId",
        to = "super::api_key::Column::Id",
        on_delete = "Cascade",
        fk_name = "gateway_guardrail_assignments_api_key_id_fkey"
    )]
    ApiKey,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "gateway_guardrail_assignments_user_id_fkey"
    )]
    User,
}

impl Related<super::gateway_guardrail::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Guardrail.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::api_key::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ApiKey.def()
    }
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
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
