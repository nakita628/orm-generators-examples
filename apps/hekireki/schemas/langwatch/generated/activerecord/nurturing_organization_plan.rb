# Each organization's plan, folded from billing's subscription_changed facts: nurturing's own
# copy. The newest fact wins, and a real fact wins a tie over a catch-up; Customer.io hears a
# fact only when it changes this row.
#
# appliedEventIds:: The event ids folded in, so a redelivery folds once.
# winnerEventId:: The fact that last changed the row; only it may tell Customer.io.
class NurturingOrganizationPlan < ApplicationRecord
  self.table_name = "NurturingOrganizationPlan"
  self.primary_key = "organizationId"

  attribute :appliedEventIds, default: -> { [] }

  validates :occurredAt, presence: true
  validates :fromCatchUp, inclusion: { in: [ true, false ] }
end
