defmodule LangWatch.OrganizationSeatLease do
  @moduledoc """
  Held by the one seat-taking write running on an organization (seat-requests.feature). A
  row past expiresAt is free to take; a holder deletes its own row when its write is done.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          organization_id: String.t(),
          holder: String.t(),
          expires_at: DateTime.t(),
          inserted_at: DateTime.t()
        }

  schema "OrganizationSeatLease" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    # Random per hold, so a release deletes only its own lease.
    field(:holder, :string)
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end