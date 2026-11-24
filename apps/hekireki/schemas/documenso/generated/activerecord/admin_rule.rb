# @zod.import(["import { ZAdminRuleDefinitionSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
#
# trigger:: [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
# definition:: [AdminRuleDefinition] @zod.custom.use(ZAdminRuleDefinitionSchema)
# exemptOrganisationIds:: Escape hatch: organisations this rule is skipped for. Only meaningful on triggers that have an
#   organisation
class AdminRule < ApplicationRecord
  self.table_name = "AdminRule"

  attribute :exemptOrganisationIds, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :mode, { enforce: "ENFORCE", monitor: "MONITOR", disabled: "DISABLED" }, validate: true

  validates :name, presence: true
  validates :trigger, presence: true
  validates :definition, exclusion: { in: [ nil ] }
  validates :priority, presence: true
  validates :stopOnMatch, inclusion: { in: [ true, false ] }

  has_many :matches, class_name: "AdminRuleMatch", foreign_key: "ruleId", inverse_of: :rule, dependent: :destroy
  has_many :admin_tags, foreign_key: "adminRuleId", inverse_of: :admin_rule, dependent: :nullify
end
