defmodule LangWatch.ErasedIdentifierSuppression do
  @moduledoc """
  Hashes of erased identifiers - the do-not-reimport list.

  Without it the erasure is undone by the pipeline that produced it: the
  pullers look 30 days back, so the next day's pull re-ingests the same email
  and re-writes what we erased. Every write path carrying an identifier
  consults this list and drops a suppressed one (ADR-128 §9 step 1).

  It stores hashes and never identifiers, so it is not itself a copy of the
  data it exists to keep out. The stored `identifierHash` and the pseudonym
  written in place of the erased value are the same function of the same
  input - HMAC-SHA256(secret, identifier) - so a write path computes that digest
  once and uses it twice: as the membership test, and, on a hit, as the
  replacement value. That is why no original-to-pseudonym mapping table
  exists; one would hold the erased identifier in plaintext forever.

  Scoped to the provider as well as the organization: the same string can be a
  different person under a different provider.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          provider: String.t(),
          identifier_hash: String.t(),
          erased_at: DateTime.t()
        }

  schema "ErasedIdentifierSuppression" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:provider, :string)
    field(:identifier_hash, :string, source: :identifierHash)
    field(:erased_at, LangWatch.PrismaDateTime, source: :erasedAt)
  end
end