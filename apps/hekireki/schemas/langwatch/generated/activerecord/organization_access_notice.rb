# The mail log for seat requests and refused sign-ins: at most one row per organization,
# kind, address and UTC day, each mailed at once.
#
# kind:: seat_request | idp_refused
# dayKey:: UTC day the notice was recorded on, YYYY-MM-DD.
# idpRefused:: Whether the seat request it announces was also refused by the identity provider.
# mailedAt:: Set once the notice was handed to the mailer; null if a send never started.
class OrganizationAccessNotice < ApplicationRecord
  self.table_name = "OrganizationAccessNotice"

  attribute :idpRefused, default: false
  alias_attribute :created_at, :createdAt

  validates :organizationId, presence: true, uniqueness: { scope: [ :kind, :email, :dayKey ] }
  validates :kind, presence: true
  validates :email, presence: true
  validates :dayKey, presence: true
end
