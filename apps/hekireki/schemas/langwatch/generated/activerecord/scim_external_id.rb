# D08 — who the directory means, per connection.
#
# The identity provider's own identifier for a person, which is what
# survives their email changing: an address is not identity, people change
# theirs, and the same person carries different identifiers on two different
# connections. So the PAIR is the key and `externalId` alone never is.
#
# No relation to `SsoConnection` (that projection carries no foreign keys);
# the `User` relation is real, so a deleted user takes their directory
# identities with them rather than stranding a mapping that resolves to
# nobody.
class ScimExternalId < ApplicationRecord
  self.table_name = "ScimExternalId"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true
  validates :connectionId, presence: true, uniqueness: { scope: :externalId }
  validates :externalId, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :scim_external_ids
end
