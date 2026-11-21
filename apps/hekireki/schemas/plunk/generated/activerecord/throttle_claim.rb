# One transactional notification held for one recipient (or for the whole project) by a
# send with `delivery.throttle`. While it is held, the same notification is not sent again.
#
# The claim is taken by inserting (or taking over an expired) row, so the unique constraint
# decides the race between concurrent sends, like `IdempotencyKey`. It is held from the
# moment the send is accepted, re-anchored to the real send time once SES has the message,
# and released when the email never gets there (cancelled, or failed before sending).
class ThrottleClaim < ApplicationRecord
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { SecureRandom.uuid }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :key, presence: true, uniqueness: { scope: [ :projectId, :recipientKey ] }
  validates :recipientKey, presence: true

  belongs_to :project, foreign_key: "projectId", inverse_of: :throttle_claims
  belongs_to :contact, foreign_key: "contactId", optional: true, inverse_of: :throttle_claims
  belongs_to :workflow, foreign_key: "workflowId", optional: true, inverse_of: :throttle_claims
end
