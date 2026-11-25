use super::admin_rule_mode::AdminRuleMode;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZAdminRuleContextSchema, ZAdminRuleActionRunsSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AdminRuleMatch")]
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
    pub mode: AdminRuleMode,
    /// [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
    #[sea_orm(column_type = "Text")]
    pub trigger: String,
    #[sea_orm(column_name = "organisationId", column_type = "Text")]
    pub organisation_id: Option<String>,
    #[sea_orm(column_name = "envelopeId", column_type = "Text")]
    pub envelope_id: Option<String>,
    /// [AdminRuleContext] @zod.custom.use(ZAdminRuleContextSchema)
    #[sea_orm(column_name = "eventSnapshot", column_type = "JsonBinary")]
    pub event_snapshot: Json,
    /// [AdminRuleActionRuns] @zod.custom.use(ZAdminRuleActionRunsSchema)
    #[sea_orm(column_name = "actionsRun", column_type = "JsonBinary")]
    pub actions_run: Json,
    #[sea_orm(column_name = "ruleId", column_type = "Text")]
    pub rule_id: String,
    #[sea_orm(column_name = "actorUserId")]
    pub actor_user_id: Option<i32>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::admin_rule::Entity",
        from = "Column::RuleId",
        to = "super::admin_rule::Column::Id",
        on_delete = "Cascade",
        fk_name = "AdminRuleMatch_ruleId_fkey"
    )]
    Rule,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ActorUserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "AdminRuleMatch_actorUserId_fkey"
    )]
    ActorUser,
}

impl Related<super::admin_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Rule.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ActorUser.def()
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
