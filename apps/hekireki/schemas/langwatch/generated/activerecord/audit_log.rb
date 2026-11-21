# actorUserId:: WHO REALLY DID IT, when that is not `userId`.
#
#   Under an impersonation the session's `user.id` is the SUBJECT — the
#   customer whose access an operator is borrowing — so a row naming only
#   `userId` filed the operator's act against the customer. An operator
#   granting a permanent way into somebody's organization appeared, in the
#   only durable record of it, to be that somebody.
#
#   Null on an ordinary request, where the two would be the same person and
#   saying so twice adds nothing. ADR-118 / D06: every decision names both.
# idempotencyKey:: A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2).
class AuditLog < ApplicationRecord
  self.table_name = "AuditLog"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  validates :action, presence: true
  validates :idempotencyKey, uniqueness: { allow_nil: true }
end
