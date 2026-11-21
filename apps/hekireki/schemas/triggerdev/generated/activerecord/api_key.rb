# description:: Optional one-line note the creator adds to say what the key is for.
# allowed_ip_ranges:: CIDR ranges this key may be used from. Empty means any IP.
class ApiKey < ApplicationRecord
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :scopes, default: -> { [] }
  attribute :allowed_ip_ranges, default: -> { [] }

  enum :issued_by, { user: "USER", system: "SYSTEM" }, default: :user, prefix: true, validate: true

  validates :name, presence: true
  validates :key_hash, presence: true, uniqueness: true
  validates :last_four, presence: true

  belongs_to :runtime_environment, inverse_of: :api_keys
  belongs_to :created_by, class_name: "User", foreign_key: "created_by_user_id", optional: true,
                          inverse_of: :created_api_keys
end
