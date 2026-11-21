# Better Auth JWT signing keys for OAuth/OIDC access and ID tokens.
#
# alg:: Better Auth 1.7 (ENG-2343): the JWS algorithm this key signs with, so a keyring can hold more
#   than one alg and `getLatestKeyByAlg` can pick the right key per resource.
#
#   NOT optional upstream in practice, despite being declared `required: false`: `createJwk` always
#   puts `alg` in the row it writes (defaulting to EdDSA), so without this column every key mint is
#   rejected by Prisma as an unknown argument — which breaks JWT signing, and with it the whole MCP
#   OAuth flow, on any deployment that has yet to mint a key. Nullable because rows written before
#   1.7 have no value; upstream reads treat `alg: null` as the configured default alg, so existing
#   keys keep working untouched.
# crv:: Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
#   the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
#   same always-written path as `alg` above.
class Jwks < ApplicationRecord
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  validates :publicKey, presence: true
  validates :privateKey, presence: true
end
