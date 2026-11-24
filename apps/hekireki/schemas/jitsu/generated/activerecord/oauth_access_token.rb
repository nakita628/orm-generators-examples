# createdAt:: @zod.custom(z.coerce.date())
# lastUsed:: @zod.custom(z.coerce.date())
# expiresAt:: @zod.custom(z.coerce.date())
class OAuthAccessToken < ApplicationRecord
  self.table_name = "OAuthAccessToken"
  self.implicit_order_column = "createdAt"

  def self.instance_method_already_implemented?(method_name)
    dangerous_attribute_method?(method_name) || super
  end
  private_class_method :instance_method_already_implemented?

  def read_attribute_for_serialization(name)
    if has_attribute?(name) && self.class.dangerous_attribute_method?(name)
      _read_attribute(name.to_s)
    else
      super
    end
  end

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :hash_value, :hash
  alias_attribute :created_at, :createdAt

  validates :hash_value, presence: true
  validates :expiresAt, presence: true

  belongs_to :refresh_token, class_name: "UserApiToken", foreign_key: "refreshTokenId", inverse_of: :access_tokens
end
