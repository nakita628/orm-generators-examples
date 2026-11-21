class UserUpdateRead < ApplicationRecord
  self.table_name = "UserUpdateRead"

  attribute :id, default: -> { Cuid.generate }

  validates :userId, uniqueness: { scope: :updateId }

  belongs_to :user, foreign_key: "userId", inverse_of: :update_reads
  belongs_to :user_update_read_update, class_name: "Update", foreign_key: "updateId", inverse_of: :user_reads
end
