use super::category_filter_type::CategoryFilterType;
use super::logical_operator::LogicalOperator;
use super::system_type::SystemType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Rule")]
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
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    #[sea_orm(default_value = true)]
    pub automate: bool,
    #[sea_orm(column_name = "runOnThreads", default_value = false)]
    pub run_on_threads: bool,
    #[sea_orm(column_name = "emailAccountId", column_type = "Text")]
    pub email_account_id: String,
    #[sea_orm(column_name = "organizationRuleId", column_type = "Text")]
    pub organization_rule_id: Option<String>,
    #[sea_orm(column_name = "organizationRuleMemberEnabled")]
    pub organization_rule_member_enabled: Option<bool>,
    #[sea_orm(column_name = "conditionalOperator", default_value = "AND")]
    pub conditional_operator: LogicalOperator,
    #[sea_orm(column_type = "Text")]
    pub instructions: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "groupId", column_type = "Text")]
    pub group_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub from: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub to: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub subject: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub body: Option<String>,
    #[sea_orm(column_name = "categoryFilterType")]
    pub category_filter_type: Option<CategoryFilterType>,
    #[sea_orm(column_name = "systemType")]
    pub system_type: Option<SystemType>,
    #[sea_orm(column_name = "promptText", column_type = "Text")]
    pub prompt_text: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::email_account::Entity",
        from = "Column::EmailAccountId",
        to = "super::email_account::Column::Id",
        on_delete = "Cascade",
        fk_name = "Rule_emailAccountId_fkey"
    )]
    EmailAccount,
    #[sea_orm(
        belongs_to = "super::organization_rule::Entity",
        from = "Column::OrganizationRuleId",
        to = "super::organization_rule::Column::Id",
        on_delete = "Cascade",
        fk_name = "Rule_organizationRuleId_fkey"
    )]
    OrganizationRule,
    #[sea_orm(
        belongs_to = "super::group::Entity",
        from = "Column::GroupId",
        to = "super::group::Column::Id",
        on_delete = "Cascade",
        fk_name = "Rule_groupId_fkey"
    )]
    Group,
    #[sea_orm(has_many = "super::action::Entity")]
    Actions,
    #[sea_orm(has_many = "super::attachment_source::Entity")]
    AttachmentSources,
    #[sea_orm(has_many = "super::executed_rule::Entity")]
    ExecutedRules,
    #[sea_orm(has_many = "super::classification_feedback::Entity")]
    ClassificationFeedback,
    #[sea_orm(has_many = "super::rule_history::Entity")]
    History,
}

impl Related<super::email_account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmailAccount.def()
    }
}

impl Related<super::organization_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrganizationRule.def()
    }
}

impl Related<super::group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Group.def()
    }
}

impl Related<super::action::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Actions.def()
    }
}

impl Related<super::attachment_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttachmentSources.def()
    }
}

impl Related<super::executed_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ExecutedRules.def()
    }
}

impl Related<super::classification_feedback::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ClassificationFeedback.def()
    }
}

impl Related<super::rule_history::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::History.def()
    }
}

impl Related<super::category::Entity> for Entity {
    fn to() -> RelationDef {
        super::category_to_rule::Relation::Category.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::category_to_rule::Relation::Rule.def().rev())
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
