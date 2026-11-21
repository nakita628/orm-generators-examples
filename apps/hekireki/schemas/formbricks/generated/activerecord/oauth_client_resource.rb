# Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
# `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
# endpoint, which is why the migration backfills one row per existing client.
# Both sides reference natural keys, matching the other oauth child tables.
class OauthClientResource < ApplicationRecord
  self.table_name = "oauthClientResource"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  validates :clientId, uniqueness: { scope: :resourceId }

  belongs_to :client, class_name: "OauthClient", foreign_key: "clientId", primary_key: "clientId",
                      inverse_of: :resources
  belongs_to :resource, class_name: "OauthResource", foreign_key: "resourceId", primary_key: "identifier",
                        inverse_of: :clients
end
