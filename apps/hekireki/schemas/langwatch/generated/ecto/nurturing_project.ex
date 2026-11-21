defmodule LangWatch.NurturingProject do
  @moduledoc """
  Nurturing's read model: which organization a project belongs to, from project's created event.
  The per-email usage places a project and lists an organization's projects from these rows only.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          project_id: String.t(),
          organization_id: String.t()
        }

  schema "NurturingProject" do
    field(:project_id, :string, primary_key: true, source: :projectId)
    field(:organization_id, :string, source: :organizationId)
  end
end