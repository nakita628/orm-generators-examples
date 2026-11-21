defmodule Formbricks.Jwks do
  @moduledoc """
  Better Auth JWT signing keys for OAuth/OIDC access and ID tokens.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          public_key: String.t(),
          private_key: String.t(),
          created_at: DateTime.t(),
          expires_at: DateTime.t() | nil,
          alg: String.t() | nil,
          crv: String.t() | nil
        }

  schema "jwks" do
    field(:id, :string, primary_key: true)
    field(:public_key, :string, source: :publicKey)
    field(:private_key, :string, source: :privateKey)
    field(:created_at, Formbricks.PrismaDateTime, source: :createdAt)
    field(:expires_at, Formbricks.PrismaDateTime, source: :expiresAt)
    # Better Auth 1.7 (ENG-2343): the JWS algorithm this key signs with, so a keyring can hold more
    # than one alg and `getLatestKeyByAlg` can pick the right key per resource.
    #
    # NOT optional upstream in practice, despite being declared `required: false`: `createJwk` always
    # puts `alg` in the row it writes (defaulting to EdDSA), so without this column every key mint is
    # rejected by Prisma as an unknown argument — which breaks JWT signing, and with it the whole MCP
    # OAuth flow, on any deployment that has yet to mint a key. Nullable because rows written before
    # 1.7 have no value; upstream reads treat `alg: null` as the configured default alg, so existing
    # keys keep working untouched.
    field(:alg, :string)
    # Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
    # the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
    # same always-written path as `alg` above.
    field(:crv, :string)
  end
end