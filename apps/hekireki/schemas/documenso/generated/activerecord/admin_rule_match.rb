# @zod.import(["import { ZAdminRuleContextSchema, ZAdminRuleActionRunsSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
#
# trigger:: [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
# eventSnapshot:: [AdminRuleContext] @zod.custom.use(ZAdminRuleContextSchema)
# actionsRun:: [AdminRuleActionRuns] @zod.custom.use(ZAdminRuleActionRunsSchema)
class AdminRuleMatch < ApplicationRecord
  self.table_name = "AdminRuleMatch"

  alias_attribute :created_at, :createdAt

  enum :mode, { enforce: "ENFORCE", monitor: "MONITOR", disabled: "DISABLED" }, validate: true

  validates :trigger, presence: true
  validates :eventSnapshot, exclusion: { in: [ nil ] }
  validates :actionsRun, exclusion: { in: [ nil ] }

  belongs_to :rule, class_name: "AdminRule", foreign_key: "ruleId", inverse_of: :matches
  belongs_to :actor_user, class_name: "User", foreign_key: "actorUserId", optional: true,
                          inverse_of: :admin_rule_matches
end
