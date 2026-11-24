# Better Auth OAuth 2.1 client registrations for MCP and future OAuth clients.
#
# public:: Unwritten by Better Auth 1.7 (it dropped both from its schema), kept deliberately: they are
#   nullable, so keeping them makes the 1.7 migration purely additive and a rollback to 1.6 needs no
#   down-migration. Drop them in a separate cleanup after a bake period.
# clientDiscoveryId:: Better Auth 1.7 (ENG-2343).
class OauthClient < ApplicationRecord
  self.table_name = "oauthClient"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :disabled, default: false
  attribute :scopes, default: -> { [] }
  attribute :contacts, default: -> { [] }
  attribute :redirectUris, default: -> { [] }
  attribute :postLogoutRedirectUris, default: -> { [] }
  attribute :grantTypes, default: -> { [] }
  attribute :responseTypes, default: -> { [] }
  attribute :clientCredentialsScopes, default: -> { [] }
  attribute :dpopBoundAccessTokens, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :clientId, presence: true, uniqueness: true

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :oauth_clients
  has_many :access_tokens, class_name: "OauthAccessToken", foreign_key: "clientId", primary_key: "clientId",
                           inverse_of: :client, dependent: :destroy
  has_many :refresh_tokens, class_name: "OauthRefreshToken", foreign_key: "clientId", primary_key: "clientId",
                            inverse_of: :client, dependent: :destroy
  has_many :consents, class_name: "OauthConsent", foreign_key: "clientId", primary_key: "clientId", inverse_of: :client,
                      dependent: :destroy
  has_many :resources, class_name: "OauthClientResource", foreign_key: "clientId", primary_key: "clientId",
                       inverse_of: :client, dependent: :destroy
  has_many :mcp_grants, foreign_key: "clientId", primary_key: "clientId", inverse_of: :client, dependent: :destroy
end
