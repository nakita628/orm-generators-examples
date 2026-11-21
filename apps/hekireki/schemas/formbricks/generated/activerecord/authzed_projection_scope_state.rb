# ENG-3282: per-scope readiness of a SpiceDB projection. A row with `readyAt` set means the scope was
# backfilled and audited clean, so the evaluator may enforce it. Written only by the backfill CLI.
class AuthzedProjectionScopeState < ApplicationRecord
  self.table_name = "AuthzedProjectionScopeState"
  self.primary_key = "scope"

  alias_attribute :updated_at, :updatedAt
end
