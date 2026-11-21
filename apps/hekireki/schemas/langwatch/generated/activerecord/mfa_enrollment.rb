# The two-step verification pipeline's projection (D06 - see
# specs/identity/mfa-and-session-shape.feature). One row per PERSON, because
# one person has one setup however many organizations they belong to.
#
# A pure event-truth head like `SsoConnection`: every column is fold-written,
# rows are never deleted (DISABLED is a tombstone, so history survives being
# turned off), and a replay rebuilds it whole-row from the event log. The
# cursor columns are inline because the fold produces exactly one row per
# aggregate - the `Identifier` fan-out needed its own cursor table, this
# does not.
#
# Every column here is a lifecycle fact. There is deliberately nowhere in
# this model a secret or a backup code could be put: `consumedBackupCodeIndexes`
# holds POSITIONS, which is what lets "how many are left" be answered
# without the projection ever knowing a code.
#
# userId:: The user. Also the aggregate id and the tenant - one identity history
#   per person (ADR-029 §4).
# enrollmentId:: The current enrollment, or NULL for somebody who never started one.
# method:: totp. Never sms.
# state:: NONE | PENDING | ENABLED | EXPIRED | DISABLED
# disabledVia:: password+totp | org-admin - who ended it, and on whose authority.
# consumedBackupCodeIndexes:: Positions already spent, ascending. Never codes.
class MfaEnrollment < ApplicationRecord
  self.table_name = "MfaEnrollment"
  self.primary_key = "userId"

  attribute :backupCodeCount, default: 0
  attribute :consumedBackupCodeIndexes, default: -> { [] }
  attribute :failedCount, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :state, presence: true
  validates :occurredAt, presence: true
  validates :lastEventId, presence: true
  validates :acceptedAt, presence: true
  validates :projectionVersion, presence: true
end
