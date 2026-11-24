class ShareOverview < ApplicationRecord
  self.table_name = "shares"
  self.primary_key = "id"

  attribute :public, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :id, presence: true, uniqueness: true
  validates :projectId, uniqueness: true

  belongs_to :project, foreign_key: "projectId", inverse_of: :share
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :share_overview
end
