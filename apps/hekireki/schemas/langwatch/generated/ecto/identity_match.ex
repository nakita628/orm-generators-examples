defmodule LangWatch.IdentityMatch do
  @moduledoc """
  The dated match table: "`user_abc123` on Anthropic is platform user #42,
  valid from March 3", closed when it stops being true.

  Dating is what makes a re-issued email address survivable: the leaver's link
  closes at offboarding, the new hire gets a new link, and January's spend
  stays with January's person. A correction closes the old row and opens a new
  one; nothing is ever silently rewritten.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          discovered_person_id: String.t(),
          user_id: String.t() | nil,
          evidence_kind: String.t(),
          valid_from: DateTime.t(),
          valid_to: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "IdentityMatch" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # The provider-side identity this link is about.
    field(:discovered_person_id, :string, source: :discoveredPersonId)
    # The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
    # the row and its dates remain, so the history of the link is not rewritten
    # by the erasure either.
    field(:user_id, :string, source: :userId)
    # directory_id | verified_email | human_confirmed. Exact verified-email
    # equality is the primary evidence; a directory id corroborates but never
    # stands alone (ADR-128 §12).
    field(:evidence_kind, :string, source: :evidenceKind)
    field(:valid_from, LangWatch.PrismaDateTime, source: :validFrom)
    # Null means the link is open. Offboarding and corrections close it.
    field(:valid_to, LangWatch.PrismaDateTime, source: :validTo)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end