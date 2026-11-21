# Tracks database schema migrations.
# Used to manage and track the state of data structure changes.
#
# @property id - Unique identifier for the migration
# @property name - Name of the migration
# @property status - Current state of the migration (pending, applied, failed)
# @property startedAt - When the migration began
# @property finishedAt - When the migration completed
class DataMigration < ApplicationRecord
  self.table_name = "DataMigration"

  attribute :id, default: -> { Cuid.generate }

  enum :status, { pending: "pending", applied: "applied", failed: "failed" }, validate: true

  validates :name, presence: true, uniqueness: true
end
