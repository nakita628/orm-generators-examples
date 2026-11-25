# Each gateway virtual key, folded from gateway's vk_lifecycle facts: nurturing's own copy. The
# creator comes from the created fact alone; the holder and purpose are each the latest non-null
# value by the fact's time (stamped beside them); the state times keep their latest.
class NurturingVirtualKey < ApplicationRecord
  self.table_name = "NurturingVirtualKey"
  self.primary_key = "virtualKeyId"

  validates :organizationId, presence: true
  validates :name, presence: true
  validates :occurredAt, presence: true
end
