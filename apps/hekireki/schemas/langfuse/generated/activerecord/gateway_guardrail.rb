class GatewayGuardrail < ApplicationRecord
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :config, default: -> { JSON.parse("{}") }

  enum :type, { regex: "REGEX", bedrock: "BEDROCK" }, default: :regex, validate: true
  enum :action, { block: "BLOCK", log: "LOG", redact: "REDACT" }, validate: true

  validates :name, presence: true, uniqueness: { scope: :organization_id }

  belongs_to :organization, inverse_of: :gateway_guardrails
  belongs_to :created_by_user, class_name: "User", foreign_key: "created_by", optional: true,
                               inverse_of: :gateway_guardrails_created
  has_many :assignments, class_name: "GatewayGuardrailAssignment", foreign_key: "guardrail_id", inverse_of: :guardrail,
                         dependent: :destroy
end
