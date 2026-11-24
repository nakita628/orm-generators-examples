# An AWS credential a Bedrock kill switch target signs with: a static key pair, a role to assume,
# or both. The secret access key and the external id are stored sealed.
class GovernanceAwsConnection < ApplicationRecord
  self.table_name = "GovernanceAwsConnection"

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true
  validates :name, presence: true
end
