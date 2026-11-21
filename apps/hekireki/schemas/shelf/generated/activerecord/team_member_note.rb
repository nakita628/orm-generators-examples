class TeamMemberNote < ApplicationRecord
  self.table_name = "TeamMemberNote"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { comment: "COMMENT", update: "UPDATE" }, default: :comment, prefix: true, validate: true

  validates :content, presence: true

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :team_member_notes_authored
  belongs_to :team_member, foreign_key: "teamMemberId", inverse_of: :team_member_notes
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :team_member_notes
end
