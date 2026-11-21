# The one-statement pre-event registration lock. `SsoConnection` remains
# event truth; this row only decides which command may append that truth.
# A slot whose projected connection is terminal may be atomically replaced.
#
# kind:: legacy | direct
class SsoConnectionRegistrationSlot < ApplicationRecord
  self.table_name = "SsoConnectionRegistrationSlot"
  self.primary_key = %w[organizationId kind]

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :connectionId, presence: true, uniqueness: true
  validates :commandId, presence: true
end
