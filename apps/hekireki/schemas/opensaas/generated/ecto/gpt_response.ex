defmodule OpenSaas.GptResponse do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: OpenSaas.PrismaDateTime, autogenerate: {OpenSaas.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          content: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: OpenSaas.User.t() | nil
        }

  schema "GptResponse" do
    field(:content, :string)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, OpenSaas.User, foreign_key: :user_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end