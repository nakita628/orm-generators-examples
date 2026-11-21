defmodule LangWatch.AgentRegistryBaseline do
  @moduledoc """
  An organisation's registry baseline: its first reconcile, written as already pushed so the
  agents it already had are not announced. Every reconcile after it pushes its changes.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          baselined_at: DateTime.t()
        }

  schema "AgentRegistryBaseline" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    field(:baselined_at, LangWatch.PrismaDateTime, source: :baselinedAt)
  end
end