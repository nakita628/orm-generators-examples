defmodule LangWatch.VerificationToken do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          identifier: String.t(),
          token: String.t(),
          expires: DateTime.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "VerificationToken" do
    field(:id, :string, primary_key: true)
    field(:identifier, :string)
    # Not unique alone: better-auth stores plain values here (a user id, a counter). See WEB-19201.
    field(:token, :string)
    field(:expires, LangWatch.PrismaDateTime)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end