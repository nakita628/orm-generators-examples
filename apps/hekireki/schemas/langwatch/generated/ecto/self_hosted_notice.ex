defmodule LangWatch.SelfHostedNotice do
  @moduledoc """
  A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
  report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
  tenant: every install receives every published notice that has not expired.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          severity: String.t(),
          title: String.t(),
          body: String.t(),
          url: String.t() | nil,
          fixed_in: String.t() | nil,
          status: String.t(),
          expires_at: DateTime.t() | nil,
          published_at: DateTime.t() | nil,
          published_by: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "SelfHostedNotice" do
    field(:id, :string, primary_key: true)
    # `info`, `warning` or `security`.
    field(:severity, :string)
    field(:title, :string)
    field(:body, :string)
    # An https link to read more; null where the notice names none.
    field(:url, :string)
    # The release that fixes what the notice is about; an install on it or newer hides it.
    field(:fixed_in, :string, source: :fixedIn)
    # `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent.
    field(:status, :string)
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    field(:published_at, LangWatch.PrismaDateTime, source: :publishedAt)
    # The staff member who published it, by user id.
    field(:published_by, :string, source: :publishedBy)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end