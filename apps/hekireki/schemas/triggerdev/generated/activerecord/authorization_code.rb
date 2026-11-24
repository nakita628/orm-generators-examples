# Used to generate PersonalAccessTokens, they're one-time use
#
# requesterIp:: Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires.
# codeChallenge:: S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it.
# callbackPort:: Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it.
# approvalCodeHash:: sha256 of the approval code, which only the approving browser ever sees.
# requestedByUserId:: The user who started the OAuth authorization. Only they can approve it.
class AuthorizationCode < ApplicationRecord
  self.table_name = "AuthorizationCode"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :scope, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :code, presence: true, uniqueness: true

  belongs_to :personal_access_token, foreign_key: "personalAccessTokenId", optional: true,
                                     inverse_of: :authorization_codes
  belongs_to :oauth_client, class_name: "OAuthClient", foreign_key: "oauthClientId", optional: true,
                            inverse_of: :authorization_codes
end
