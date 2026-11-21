class WorkingHours < ApplicationRecord
  self.table_name = "WorkingHours"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :enabled, default: false
  attribute :weeklySchedule,
            default: -> { JSON.parse("{\"0\":{\"isOpen\":false},\"1\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"2\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"3\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"4\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"5\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"6\":{\"isOpen\":false}}") }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, uniqueness: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :working_hours
  has_many :overrides, class_name: "WorkingHoursOverride", foreign_key: "workingHoursId", inverse_of: :working_hours,
                       dependent: :destroy
end
