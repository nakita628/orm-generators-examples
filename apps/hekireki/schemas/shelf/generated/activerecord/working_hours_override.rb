class WorkingHoursOverride < ApplicationRecord
  self.table_name = "WorkingHoursOverride"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :isOpen, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :date, presence: true, uniqueness: { scope: :workingHoursId }

  belongs_to :working_hours, foreign_key: "workingHoursId", inverse_of: :overrides
end
