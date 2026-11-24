use super::admin_rule_mode::AdminRuleMode;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZAdminRuleDefinitionSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AdminRule")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    pub name: String,
    pub description: Option<String>,
    pub mode: AdminRuleMode,
    /// [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
    pub trigger: String,
    /// [AdminRuleDefinition] @zod.custom.use(ZAdminRuleDefinitionSchema)
    pub definition: Json,
    pub priority: i32,
    #[sea_orm(column_name = "stopOnMatch")]
    pub stop_on_match: bool,
    /// Escape hatch: organisations this rule is skipped for. Only meaningful on triggers that have an
    /// organisation
    #[sea_orm(column_name = "exemptOrganisationIds")]
    pub exempt_organisation_ids: Vec<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::admin_rule_match::Entity")]
    Matches,
    #[sea_orm(has_many = "super::admin_tag::Entity")]
    AdminTags,
}

impl Related<super::admin_rule_match::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Matches.def()
    }
}

impl Related<super::admin_tag::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AdminTags.def()
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
