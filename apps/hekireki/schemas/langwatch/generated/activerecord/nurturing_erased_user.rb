# Each account nurturing erased, as a SHA-256 of its id and never the id: a gateway key fact
# naming it, however late it lands, stores no creator or holder. Platform-wide, like the erasure.
class NurturingErasedUser < ApplicationRecord
  self.table_name = "NurturingErasedUser"
  self.primary_key = "userIdHash"
end
