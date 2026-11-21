# ADR-201: a project's key for signing the tags on its calls. Owned by the API key module.
class SigningKey < ApplicationRecord
  self.table_name = "SigningKey"

  alias_attribute :created_at, :createdAt

  validates :organizationId, presence: true
  validates :projectId, presence: true
  validates :name, presence: true
  validates :publicId, presence: true, uniqueness: true
  validates :secretCiphertext, presence: true
  validates :createdByUserId, presence: true
end
