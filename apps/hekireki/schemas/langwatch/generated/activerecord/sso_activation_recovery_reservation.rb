# A durable claim on one live recovery path while activation is between its
# guard and projected ACTIVE fact. It has no expiry: only projection of the
# matching activation, or an explicit cancellation, may consume it.
class SsoActivationRecoveryReservation < ApplicationRecord
  self.table_name = "SsoActivationRecoveryReservation"
  self.primary_key = "commandId"

  alias_attribute :created_at, :createdAt

  validates :organizationId, presence: true, uniqueness: { scope: :connectionId }
  validates :connectionId, presence: true
end
