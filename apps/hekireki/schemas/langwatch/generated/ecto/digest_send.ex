defmodule LangWatch.DigestSend do
  @moduledoc """
  One person's digest for one ISO week, accepted by our own mailer. The id is `userId:isoWeek`.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          user_id: String.t(),
          organization_id: String.t(),
          iso_week: String.t(),
          template_id: String.t(),
          sent_at: DateTime.t(),
          inserted_at: DateTime.t()
        }

  schema "DigestSend" do
    field(:id, :string, primary_key: true)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    field(:iso_week, :string, source: :isoWeek)
    field(:template_id, :string, source: :templateId)
    field(:sent_at, LangWatch.PrismaDateTime, source: :sentAt)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end