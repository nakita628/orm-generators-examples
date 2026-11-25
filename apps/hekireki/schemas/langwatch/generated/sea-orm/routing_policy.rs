use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RoutingPolicy")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(
        column_name = "modelProviderIds",
        column_type = "JsonBinary",
        default_value = "[]"
    )]
    pub model_provider_ids: Json,
    #[sea_orm(
        column_name = "modelAliases",
        column_type = "JsonBinary",
        default_value = "{}"
    )]
    pub model_aliases: Json,
    #[sea_orm(column_name = "defaultModel", column_type = "Text")]
    pub default_model: Option<String>,
    #[sea_orm(
        column_name = "policyRules",
        column_type = "JsonBinary",
        default_value = "{}"
    )]
    pub policy_rules: Json,
    #[sea_orm(column_name = "isDefault", default_value = false)]
    pub is_default: bool,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: Option<String>,
    #[sea_orm(column_name = "updatedById", column_type = "Text")]
    pub updated_by_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "RoutingPolicy_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "RoutingPolicy_createdById_fkey"
    )]
    CreatedBy,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UpdatedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "RoutingPolicy_updatedById_fkey"
    )]
    UpdatedBy,
    #[sea_orm(has_many = "super::virtual_key::Entity")]
    VirtualKeys,
    #[sea_orm(has_many = "super::routing_policy_scope::Entity")]
    Scopes,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::virtual_key::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::VirtualKeys.def()
    }
}

impl Related<super::routing_policy_scope::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scopes.def()
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
