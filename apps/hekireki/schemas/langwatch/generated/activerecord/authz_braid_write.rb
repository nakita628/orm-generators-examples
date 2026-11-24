class AuthzBraidWrite < ApplicationRecord
  self.table_name = "AuthzBraidWrite"
  self.primary_key = %w[organizationId commandId]

  alias_attribute :created_at, :createdAt

  validates :grantId, presence: true
  validates :readerProjectId, presence: true
  validates :ownerProjectId, presence: true
  validates :termsKey, presence: true
  validates :condition, exclusion: { in: [ nil ] }
end
