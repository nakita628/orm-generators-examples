# The registry of licenses LangWatch has issued (ADR-139). It lives on
# LangWatch Cloud; a self-hosted install has the table and no rows.
#
# The row never holds the token an install presents. `tokenHash` is the
# SHA-256 of that token, so read access to this table is not credential
# access. Status is derived: revoked, superseded, past `expiresAt`, else active.
#
# licenseId:: The `licenseId` inside the signed license (`lic-<uuid>`).
# tokenHash:: SHA-256 of the `lwl_` token. Looked up on every credential resolution.
# organizationId:: The customer organization on LangWatch Cloud. Null for a license that is
#   recorded but not linked to a customer yet, such as a self-serve purchase:
#   it resolves to nothing until an operator links it.
# organizationName:: As signed into the license.
# maxMembersLite:: The lite seats sold: null when not stated or when unlimited. Never a stand-in number.
# liteSeatsUnlimited:: True when the license sells unlimited lite seats; maxMembersLite is then null.
# maxMembersDeveloper:: The Developer seats sold: null when not stated, which sells none.
# issuedById:: The operator who issued or registered it. Null for the purchase flow and the script.
# supersededAt:: Set when the install first presents the license that replaced this one.
# replacesId:: The id of the IssuedLicense this one replaces, when it was reissued. A
#   plain column ref (relationMode="prisma" cannot declare a self-relation on
#   Postgres); the registry service validates it. Unique, so a license is
#   replaced by one license at most.
# pendingDeliveryLicense:: A reissued license waiting to be delivered over sync, encrypted. Erased
#   when the install first presents it. The only license text the registry holds.
# services:: Hosted services this license may call, such as `instant_evals`.
# seatRateCents:: The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats.
# commitUsdCents:: The prepaid usage commit for the term, in USD cents.
# overageEnabled:: Whether usage may continue past the commit, up to `overageMaxUsdCents`.
# instanceId:: The install this license is bound to. Set on first use, cleared by an operator.
# lastSyncAt:: The last license sync this install reported: when it arrived, the
#   LangWatch version that sent it, and the seats in use at that moment.
#   Null until the install syncs for the first time.
# virtualKeyId:: The managed gateway key (purpose CONNECT) that hosted-service traffic of
#   this license runs under. One per license, so spend rows name the install.
#   Created when the license first resolves. A plain column, not a relation.
# seatsRaisedFrom:: The seats of the license this one replaced, set only when a seat change
#   raised a linked license. Billing reads it through licensing to invoice
#   the added seats, so the fact commits with the license or not at all.
class IssuedLicense < ApplicationRecord
  self.table_name = "IssuedLicense"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :liteSeatsUnlimited, default: false
  attribute :services, default: -> { [] }
  attribute :commitUsdCents, default: 0
  attribute :overageEnabled, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :source, { backoffice: "BACKOFFICE", purchase: "PURCHASE", script: "SCRIPT", legacy_import: "LEGACY_IMPORT" },
       validate: true
  enum :seatCurrency, { usd: "USD", eur: "EUR" }, validate: { allow_nil: true }

  validates :licenseId, presence: true, uniqueness: true
  validates :tokenHash, presence: true, uniqueness: true
  validates :organizationName, presence: true
  validates :email, presence: true
  validates :planType, presence: true
  validates :maxMembers, presence: true
  validates :issuedAt, presence: true
  validates :expiresAt, presence: true
  validates :replacesId, uniqueness: { allow_nil: true }
  validates :virtualKeyId, uniqueness: { allow_nil: true }

  belongs_to :organization, foreign_key: "organizationId", optional: true, inverse_of: :issued_licenses
end
