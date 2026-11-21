# A person asked for a full seat the organization had no room for (modules/organization/specs/
# seat-requests.feature). One row per organization and address; reopening reuses the row.
#
# email:: Lowercased address of the person asking.
# userId:: Null when only the address is known (a sign-in the identity provider refused).
# status:: OPEN | APPROVED | DECLINED
# source:: pending | lite | idp_refused: how the person was left when the request opened.
# idpRefused:: The identity provider also refused this person's sign-in.
# idpKind:: okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml
# salesNotifiedAt:: When sales were told about this request; reset when it reopens.
# seatPendingSince:: Set when an approval claims the request, cleared once its seat is given; an approval
#   still pending long after is one whose seat never landed, and is reopened.
class SeatRequest < ApplicationRecord
  self.table_name = "SeatRequest"

  attribute :idpRefused, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true, uniqueness: { scope: :email }
  validates :email, presence: true
  validates :status, presence: true
  validates :source, presence: true
  validates :openedAt, presence: true
end
