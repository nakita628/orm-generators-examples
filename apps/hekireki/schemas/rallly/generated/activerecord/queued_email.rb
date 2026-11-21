# An email waiting to be sent by the queue, one row per recipient. The
# email is built when it is sent, from the row its subject points at.
#
# kind:: Which handler builds the email. A string rather than an enum so a new
#   kind needs no migration; the claim only takes kinds the build knows.
# subject_id:: The row the email is about, interpreted by the kind. Not a foreign key:
#   kinds point at different tables, and a subject deleted before the send
#   makes its handler skip the email.
# batch_id:: Groups the emails one action queued, so the action can send its own
#   first batch at once and its progress can be read back.
# user_id:: The account whose action queued the email. Banning it skips whatever
#   is still pending.
# send_after:: Not claimed before this time. Unset sends at the next run.
class QueuedEmail < ApplicationRecord
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :attempts, default: 0

  enum :status, { pending: "pending", sent: "sent", skipped: "skipped", failed: "failed" }, default: :pending,
                                                                                            validate: true

  validates :kind, presence: true
  validates :subject_id, presence: true
  validates :batch_id, presence: true

  belongs_to :user, inverse_of: :queued_emails
end
