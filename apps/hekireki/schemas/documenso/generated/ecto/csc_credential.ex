defmodule Documenso.CscCredential do
  @moduledoc """
  Per-recipient cached credential metadata for the configured Cloud Signature
  Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
  symmetric-encrypted service-scope access token. One row per recipient on a
  CSC envelope; absent for SES recipients.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Documenso.PrismaDateTime, autogenerate: {Documenso.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          provider_id: String.t(),
          credential_id: String.t(),
          cert_cache: binary() | nil,
          signature_algorithm: String.t(),
          key_type: String.t(),
          digest_algorithm: String.t(),
          key_len_bits: integer() | nil,
          sign_algo_params: String.t() | nil,
          service_token_ciphertext: binary() | nil,
          service_token_expires_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          recipient: Documenso.Recipient.t() | nil
        }

  schema "CscCredential" do
    field(:id, :string, primary_key: true)
    field(:provider_id, :string, source: :providerId)
    field(:credential_id, :string, source: :credentialId)
    # TSP-validated certificate chain (length-prefixed). Nullable until the
    # service-scope OAuth callback successfully retrieves and validates it.
    field(:cert_cache, :binary, source: :certCache)
    # Algorithm metadata derived from the CSC `credentials/info` response.
    # Held as flat scalars so each is independently queryable.
    field(:signature_algorithm, :string, source: :signatureAlgorithm)
    field(:key_type, :string, source: :keyType)
    field(:digest_algorithm, :string, source: :digestAlgorithm)
    field(:key_len_bits, :integer, source: :keyLenBits)
    field(:sign_algo_params, :string, source: :signAlgoParams)
    # Service-scope access token, symmetric-encrypted at the application layer
    # before persistence. Nullable until the service-scope OAuth callback runs.
    field(:service_token_ciphertext, :binary, source: :serviceTokenCiphertext)
    field(:service_token_expires_at, Documenso.PrismaDateTime, source: :serviceTokenExpiresAt)
    field(:recipient_id, :id, source: :recipientId)
    belongs_to(:recipient, Documenso.Recipient, foreign_key: :recipient_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end