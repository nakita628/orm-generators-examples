defmodule Langfuse.VerifiedDomain do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Langfuse.PrismaDateTime, autogenerate: {Langfuse.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          domain: String.t(),
          verification_token: String.t(),
          verified_at: DateTime.t() | nil,
          created_by_user_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Langfuse.Organization.t() | nil
        }

  schema "verified_domains" do
    field(:domain, :string)
    field(:verification_token, :string, autogenerate: {Ecto.UUID, :generate, []})
    field(:verified_at, Langfuse.PrismaDateTime)
    field(:created_by_user_id, :string)
    belongs_to(:organization, Langfuse.Organization, foreign_key: :organization_id)
    timestamps(inserted_at_source: :created_at)
  end
end