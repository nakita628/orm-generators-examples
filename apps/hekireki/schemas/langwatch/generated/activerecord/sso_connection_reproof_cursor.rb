# Operational round-robin position for ADR-123's domain re-proof sweep.
#
# This is deliberately not a column on `SsoConnection`: that row is a
# whole-row event projection and its fold is its only writer, while a
# healthy re-read states no fact. Keeping the scheduling cursor beside the
# projection preserves both rules and still lets a fixed-size sweep revisit
# every connection, including one whose proof has started wavering.
class SsoConnectionReproofCursor < ApplicationRecord
  self.table_name = "SsoConnectionReproofCursor"
  self.primary_key = "connectionId"

  validates :lastReproofAt, presence: true

  belongs_to :connection, class_name: "SsoConnection", foreign_key: "connectionId", inverse_of: :reproof_cursor
end
