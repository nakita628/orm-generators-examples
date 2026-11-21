defmodule TypeHero.ImageUpload do
  use Ecto.Schema

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  @timestamps_opts [type: TypeHero.PrismaDateTime, autogenerate: {TypeHero.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          url: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: TypeHero.User.t() | nil
        }

  schema "ImageUpload" do
    # the url to the s3 uploaded assets
    field(:url, :string)
    # the user id fk
    field(:user_id, :binary_id, source: :userId)
    # The user that this is related to
    belongs_to(:user, TypeHero.User, foreign_key: :user_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end