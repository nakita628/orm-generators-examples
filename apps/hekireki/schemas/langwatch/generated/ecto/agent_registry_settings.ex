defmodule LangWatch.AgentRegistrySettings do
  @moduledoc """
  An organisation's agent registry settings. No row reads as the defaults.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          organization_id: String.t(),
          dormant_after_days: integer(),
          updated_at: DateTime.t()
        }

  schema "AgentRegistrySettings" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    # Days without traffic after which a listed agent reads as dormant.
    field(:dormant_after_days, :integer, default: 30, source: :dormantAfterDays)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end