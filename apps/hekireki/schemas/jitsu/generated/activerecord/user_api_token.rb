# createdAt:: @zod.custom(z.coerce.date())
# updatedAt:: @zod.custom(z.coerce.date())
# lastUsed:: @zod.custom(z.coerce.date())
# expiresAt:: @zod.custom(z.coerce.date())
class UserApiToken < ApplicationRecord
  self.table_name = "UserApiToken"
  self.inheritance_column = nil
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
  alias_attribute :updated_at, :updatedAt

  validates :hint, presence: true
  validates :hash_value, presence: true

  belongs_to :user, class_name: "UserProfile", foreign_key: "userId", inverse_of: :tokens
  belongs_to :oauth_client, class_name: "OAuthClient", foreign_key: "oauthClientId", optional: true, inverse_of: :tokens
  has_many :access_tokens, class_name: "OAuthAccessToken", foreign_key: "refreshTokenId", inverse_of: :refresh_token,
                           dependent: :restrict_with_error
end
