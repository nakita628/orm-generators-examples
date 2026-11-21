class Update < ApplicationRecord
  self.table_name = "Update"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :targetRoles, default: -> { [] }
  attribute :clickCount, default: 0
  attribute :viewCount, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { draft: "DRAFT", published: "PUBLISHED" }, default: :draft, validate: true

  validates :title, presence: true
  validates :content, presence: true
  validates :publishDate, presence: true

  belongs_to :created_by, class_name: "User", foreign_key: "createdById", inverse_of: :created_updates
  has_many :user_reads, class_name: "UserUpdateRead", foreign_key: "updateId", inverse_of: :user_update_read_update,
                        dependent: :destroy
end
