# One self-hosted install, as its daily usage report describes it (ADR-139,
# section 10). The row is created by the first report that names an instance
# id and is updated by every report after that, so the backoffice reads one
# row per install rather than folding a history on every page load.
#
# The report carries no credential, so it identifies an install and never a
# customer. `organizationId` and `issuedLicenseId` are resolved from the
# license bound to this instance, which the install proved by presenting its
# token to license sync. A field in the posted body never reaches them.
#
# instanceId:: The UUID the install minted for itself. The join key for everything else.
# version:: What the install says about itself.
# hostname:: The host it answers on, when the operator left hostname reporting on.
# installedAt:: The day the install itself says it first came up, which predates the
#   first report on an install that switched reporting on later.
# organizationId:: The customer, and the license that named them. Plain column refs rather
#   than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
#   carries no foreign keys, and the receiver resolves both from the license
#   registry rather than from the report.
# userEmailDomains:: Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
#   install that switched the optional category off.
# userDomains:: The same domains as the keys of `userEmailDomains`, flattened so an
#   operator can search for one. A JSON object's keys are the domains
#   themselves, and no query language addresses a key it does not already
#   know the name of.
# latestReport:: What the last report carried, so the list and the drawer read one row:
#   the onboarding ladder, the usage counts and the operational block, under
#   the same keys the dictionary declares.
# optionalMetricsReported:: What the last report said the operator had switched off, read from which
#   categories arrived rather than from a field the install could assert.
# reportCount:: How many reports have landed, and how many fields the last one carried
#   that this release has no name for. A count above zero says the senders
#   are ahead of the receiver.
# raisedSignals:: The lead signals already raised for this install. A signal on this list
#   is never raised again, which is what keeps a large install from
#   reporting that it is large every morning.
class SelfHostedInstance < ApplicationRecord
  self.table_name = "SelfHostedInstance"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :userDomains, default: -> { [] }
  attribute :optionalMetricsReported, default: true
  attribute :hostnameReported, default: true
  attribute :reportCount, default: 0
  attribute :lastUnknownFields, default: 0
  attribute :raisedSignals, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :instanceId, presence: true, uniqueness: true
  validates :lastSeenAt, presence: true
end
