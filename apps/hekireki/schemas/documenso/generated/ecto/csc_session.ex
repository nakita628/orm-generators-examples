defmodule Documenso.CscSession do
  @moduledoc """
  Per-recipient transient sign-time state held across the credential-scope
  OAuth round-trip. The `recipientId @unique` constraint caps each recipient
  to at most one in-flight session — re-clicking Sign UPSERTs the row.
  Cleaned up transitively when the Recipient is deleted.
  @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Documenso.PrismaDateTime, autogenerate: {Documenso.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          envelope_id: String.t(),
          signing_time: DateTime.t(),
          items_json: map() | list() | String.t() | number() | boolean(),
          encrypted_sad: binary() | nil,
          sad_expires_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          recipient: Documenso.Recipient.t() | nil
        }

  schema "CscSession" do
    field(:id, :string, primary_key: true)
    # Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
    # flows through the Recipient cascade.
    field(:envelope_id, :string, source: :envelopeId)
    # Pinned at prep time; stable across overlay re-renders so the captured
    # signedAttrs digest matches the bytes the TSP signs.
    field(:signing_time, Documenso.PrismaDateTime, source: :signingTime)
    # Contract between prep and sign: an array of
    # `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
    # `documentDataId` pins the immutable bytes that produced the captured
    # digest so sign-time can re-derive the same hash against the same content.
    # [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
    field(:items_json, Documenso.PrismaJson, source: :itemsJson)
    # Populated by the credential-scope OAuth callback. Both nullable at row
    # creation; set together when the SAD is exchanged.
    field(:encrypted_sad, :binary, source: :encryptedSad)
    field(:sad_expires_at, Documenso.PrismaDateTime, source: :sadExpiresAt)
    field(:recipient_id, :id, source: :recipientId)
    belongs_to(:recipient, Documenso.Recipient, foreign_key: :recipient_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end