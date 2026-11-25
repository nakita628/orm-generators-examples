class GatewayGuardrailAssignment < ApplicationRecord
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  enum :scope, { organization: "ORGANIZATION", api_key: "API_KEY", user: "USER" }, validate: true

  belongs_to :guardrail, class_name: "GatewayGuardrail", inverse_of: :assignments
  belongs_to :organization, inverse_of: :gateway_guardrail_assignments
  belongs_to :api_key, optional: true, inverse_of: :gateway_guardrail_assignments
  belongs_to :user, optional: true, inverse_of: :gateway_guardrail_assignments
end
