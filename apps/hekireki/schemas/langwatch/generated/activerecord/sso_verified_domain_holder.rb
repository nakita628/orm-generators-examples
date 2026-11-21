# A connection currently entitled to use an organization-owned verified
# domain. Usually there is one holder. During an Auth0 migration there are
# exactly two: the grandfathered predecessor and its direct replacement.
#
# The projection repository enforces that relationship before inserting a
# second holder. The flattened migration also installs partial unique live
# legacy/direct slots on `SsoConnection`, closing concurrent third-writer
# races that Prisma cannot describe declaratively.
class SsoVerifiedDomainHolder < ApplicationRecord
  self.table_name = "SsoVerifiedDomainHolder"
  self.primary_key = %w[domain connectionId]

  belongs_to :ownership, class_name: "SsoVerifiedDomain", foreign_key: [ :domain, :organizationId ],
                         primary_key: [ :domain, :organizationId ], inverse_of: :holders
end
