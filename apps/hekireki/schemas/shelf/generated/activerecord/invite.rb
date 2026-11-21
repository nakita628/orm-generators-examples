class Invite < ApplicationRecord
  self.table_name = "Invite"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :roles, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { pending: "PENDING", accepted: "ACCEPTED", rejected: "REJECTED", invalidated: "INVALIDATED" },
       default: :pending, validate: true

  validates :inviteeEmail, presence: true
  validates :inviteCode, presence: true
  validates :inviteMessage, length: { maximum: 1000 }
  validates :expiresAt, presence: true

  belongs_to :inviter, class_name: "User", foreign_key: "inviterId", inverse_of: :sent_invites
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :invites
  belongs_to :invitee_user, class_name: "User", foreign_key: "inviteeUserId", optional: true,
                            inverse_of: :received_invites
  belongs_to :invitee_team_member, class_name: "TeamMember", foreign_key: "teamMemberId", inverse_of: :received_invites
end
