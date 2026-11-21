defmodule LangWatch.NurturingTracedProject do
  @moduledoc """
  Nurturing's copy of each project that received a real trace, from trace's and project's facts.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          project_id: String.t(),
          organization_id: String.t()
        }

  schema "NurturingTracedProject" do
    field(:project_id, :string, primary_key: true, source: :projectId)
    field(:organization_id, :string, source: :organizationId)
  end
end