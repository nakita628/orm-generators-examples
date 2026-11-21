# Every ClickHouse `TenantId` this organization has ever written governance
# rows under, appended on first use and never pruned.
#
# The live `resolveGovProjectId` lookup cannot be this translation: it filters
# `archivedAt: null` and so returns null forever once someone archives the
# governance project, while the write path re-reads the same project by slug
# with no such filter and keeps landing rows under the old tenant. The
# organization would map to zero tenants on read and one on write - a
# permanent split-brain in which the cost screen reports "no governance data"
# and an erasure job erases nothing and reports success.
#
# Reads resolve against the whole history; erasure walks all of it, because
# personal data does not stop existing in a tenant that stopped being current.
#
# tenantId:: A hidden governance project id this organization has written rows under.
class GovernanceTenantHistory < ApplicationRecord
  self.table_name = "GovernanceTenantHistory"

  attribute :id, default: -> { Nanoid.generate }

  validates :organizationId, presence: true, uniqueness: { scope: :tenantId }
  validates :tenantId, presence: true
  validates :firstUsedAt, presence: true
  validates :lastUsedAt, presence: true
end
