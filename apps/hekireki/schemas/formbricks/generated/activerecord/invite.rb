# Represents pending invitations to join an organization.
# Used to manage the process of adding new users to an organization.
# Once accepted, invites are converted into memberships.
#
# @property id - Unique identifier for the invite
# @property email - Email address of the invited user
# @property name - Optional display name for the invited user
# @property organization - The organization sending the invite
# @property creator - The user who created the invite
# @property acceptor - The user who accepted the invite (if accepted)
# @property expiresAt - When the invite becomes invalid
# @property role - Intended role for the invited user
# @property teamIds - Teams the user will be added to upon acceptance
class Invite < ApplicationRecord
  self.table_name = "Invite"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { SecureRandom.uuid }
  attribute :teamIds, default: -> { [] }
  alias_attribute :created_at, :createdAt

  enum :role, { owner: "owner", manager: "manager", member: "member", billing: "billing", dashboard_viewer: "dashboard_viewer" },
       default: :member, validate: true

  validates :email, presence: true
  validates :expiresAt, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :invites
  belongs_to :creator, class_name: "User", foreign_key: "creatorId", inverse_of: :invites_created
  belongs_to :acceptor, class_name: "User", foreign_key: "acceptorId", optional: true, inverse_of: :invites_accepted
end
