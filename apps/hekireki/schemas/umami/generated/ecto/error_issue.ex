defmodule Umami.ErrorIssue do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Umami.PrismaDateTime, autogenerate: {Umami.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          website_id: Ecto.UUID.t(),
          fingerprint: String.t(),
          grouping_version: integer(),
          title: String.t(),
          status: String.t(),
          resolved_at: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "error_issue" do
    field(:id, Ecto.UUID, primary_key: true, source: :issue_id)
    field(:website_id, Ecto.UUID)
    field(:fingerprint, :string)
    field(:grouping_version, :integer, default: 1)
    field(:title, :string)
    field(:status, :string, default: "unresolved")
    field(:resolved_at, Umami.PrismaDateTime)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end