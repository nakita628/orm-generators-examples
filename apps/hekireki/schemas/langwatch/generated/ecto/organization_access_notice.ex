defmodule LangWatch.OrganizationAccessNotice do
  @moduledoc """
  The mail log for seat requests and refused sign-ins: at most one row per organization,
  kind, address and UTC day, each mailed at once.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          kind: String.t(),
          email: String.t(),
          name: String.t() | nil,
          idp: String.t() | nil,
          day_key: String.t(),
          idp_refused: boolean(),
          mailed_at: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "OrganizationAccessNotice" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # seat_request | idp_refused
    field(:kind, :string)
    field(:email, :string)
    field(:name, :string)
    field(:idp, :string)
    # UTC day the notice was recorded on, YYYY-MM-DD.
    field(:day_key, :string, source: :dayKey)
    # Whether the seat request it announces was also refused by the identity provider.
    field(:idp_refused, :boolean, default: false, source: :idpRefused)
    # Set once the notice was handed to the mailer; null if a send never started.
    field(:mailed_at, LangWatch.PrismaDateTime, source: :mailedAt)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end