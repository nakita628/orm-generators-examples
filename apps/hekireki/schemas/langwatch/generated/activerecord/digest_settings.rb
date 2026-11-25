# The digest's one settings row. `armed` is the weekly send's master switch; it starts off.
class DigestSettings < ApplicationRecord
  self.table_name = "DigestSettings"

  attribute :id, default: "singleton"
  attribute :armed, default: false
  alias_attribute :updated_at, :updatedAt
end
