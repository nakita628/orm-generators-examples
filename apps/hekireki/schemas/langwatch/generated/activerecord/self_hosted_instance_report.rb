# One usage report, exactly as it arrived. Kept so a number on an install's
# row can be traced back to the report that set it, and so a question asked
# later can be answered over the history rather than over the latest snapshot.
#
# instanceId:: The `SelfHostedInstance.instanceId` this report came from. A plain column
#   ref for the same reason the instance row's refs are.
# unknownFields:: Fields this release had no name for. Counted, never stored by name.
# payload:: The validated report, with every unknown field already dropped.
class SelfHostedInstanceReport < ApplicationRecord
  self.table_name = "SelfHostedInstanceReport"

  attribute :id, default: -> { Nanoid.generate }
  attribute :unknownFields, default: 0

  validates :instanceId, presence: true
  validates :payload, exclusion: { in: [ nil ] }
end
