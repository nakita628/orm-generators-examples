# Monthly counters for metered actions charged to a space, independent of
# the records they produce, so deleting a poll does not give back what it
# spent.
class SpaceUsage < ApplicationRecord
  self.table_name = "space_usage"
  self.primary_key = %w[space_id resource period]

  attribute :count, default: 0

  enum :resource, { api_poll: "api_poll", email_invite: "email_invite" }, validate: true

  belongs_to :space, inverse_of: :usage
end
