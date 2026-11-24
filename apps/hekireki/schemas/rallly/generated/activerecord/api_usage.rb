# Superseded by SpaceUsage (resource api_poll). Unused; kept so old
# deployments keep working during the switch, then dropped.
class ApiUsage < ApplicationRecord
  self.table_name = "api_usage"
  self.primary_key = %w[space_id resource period]

  attribute :count, default: 0

  enum :resource, { poll: "poll" }, validate: true

  belongs_to :space, inverse_of: :api_usage
end
