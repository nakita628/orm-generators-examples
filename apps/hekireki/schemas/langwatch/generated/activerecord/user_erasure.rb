# One row per erased account, holding only its creation time: no id, name or address. It keeps
# the install's first-created account knowable after an erasure deletes that row.
# Spec: modules/user/specs/user.feature
class UserErasure < ApplicationRecord
  self.table_name = "UserErasure"

  attribute :id, default: -> { Nanoid.generate }

  validates :accountCreatedAt, presence: true
end
