defmodule Postiz.Clipping do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Postiz.PrismaDateTime, autogenerate: {Postiz.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          url: String.t(),
          status: String.t(),
          error: String.t() | nil,
          title: String.t() | nil,
          thumbnail: String.t() | nil,
          duration: integer() | nil,
          max_clips: integer(),
          fit: String.t(),
          integrations: String.t(),
          credits_id: String.t() | nil,
          deleted_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Postiz.Organization.t() | nil,
          clips: [Postiz.ClippingClip.t()]
        }

  schema "Clipping" do
    field(:url, :string)
    field(:status, :string, default: "analysing")
    field(:error, :string)
    field(:title, :string)
    field(:thumbnail, :string)
    field(:duration, :integer)
    field(:max_clips, :integer, default: 5, source: :maxClips)
    field(:fit, :string, default: "blur")
    field(:integrations, :string)
    field(:credits_id, :string, source: :creditsId)
    field(:deleted_at, Postiz.PrismaDateTime, source: :deletedAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Postiz.Organization, foreign_key: :organization_id, define_field: false)
    has_many(:clips, Postiz.ClippingClip, foreign_key: :clipping_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end