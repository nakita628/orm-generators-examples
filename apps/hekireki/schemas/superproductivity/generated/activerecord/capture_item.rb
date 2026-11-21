# One sealed capture waiting for a device. No relation to its grant: items
# outlive a grant that is revoked without discarding them.
#
# task_id:: Task id the claiming device must use; replaced when another device takes over.
class CaptureItem < ApplicationRecord
  validates :grant_id, presence: true
  validates :ciphertext, presence: true
  validates :task_id, presence: true
  validates :created_at, presence: true

  belongs_to :user, inverse_of: :capture_items
end
