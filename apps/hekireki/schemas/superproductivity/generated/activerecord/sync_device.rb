# app_version:: Bare semver the client last reported on download (`18.22.0`), NULL until
#   it does. Read by the checkpoint gate (#9962); never shown to users.
# last_acked_seq:: Reserved: written once at device registration, never read or updated.
class SyncDevice < ApplicationRecord
  self.primary_key = %w[user_id client_id]

  attribute :last_acked_seq, default: 0

  validates :last_seen_at, presence: true
  validates :created_at, presence: true

  belongs_to :user, inverse_of: :devices
end
