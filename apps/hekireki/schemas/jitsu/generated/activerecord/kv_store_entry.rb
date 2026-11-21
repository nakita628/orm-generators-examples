# expires_at:: @zod.custom(z.coerce.date())
class KvStoreEntry < ApplicationRecord
  self.table_name = "kv"
  self.primary_key = "key"

  validates :value, exclusion: { in: [ nil ] }
end
