# A connection's managed users, including resources without an externalId.
# Scalar keys only; readers exclude users that no longer exist.
class ScimDirectoryUser < ApplicationRecord
  self.table_name = "ScimDirectoryUser"
  self.primary_key = %w[connectionId userId]

  validates :organizationId, presence: true
end
