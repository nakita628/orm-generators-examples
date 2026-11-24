defmodule TypeHero.ImageUpload do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: TypeHero.PrismaDateTime, autogenerate: {TypeHero.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          url: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: TypeHero.User.t() | nil
        }

  schema "ImageUpload" do
    # the url to the s3 uploaded assets
    field(:url, :string)
    # the user id fk
    field(:user_id, :string, source: :userId)
    # The user that this is related to
    belongs_to(:user, TypeHero.User, foreign_key: :user_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end