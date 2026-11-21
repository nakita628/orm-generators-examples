class BookingNote < ApplicationRecord
  self.table_name = "BookingNote"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { comment: "COMMENT", update: "UPDATE" }, default: :comment, prefix: true, validate: true

  validates :content, presence: true

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :booking_notes
  belongs_to :booking, foreign_key: "bookingId", inverse_of: :notes
end
