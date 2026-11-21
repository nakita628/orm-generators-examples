defmodule LangWatch.NurturingPersonalProject do
  @moduledoc """
  Nurturing's copy of each personal project's owner, from organization's and project's facts.
  The first owner recorded stays: a personal project's owner never changes.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          project_id: String.t(),
          organization_id: String.t(),
          owner_user_id: String.t()
        }

  schema "NurturingPersonalProject" do
    field(:project_id, :string, primary_key: true, source: :projectId)
    field(:organization_id, :string, source: :organizationId)
    field(:owner_user_id, :string, source: :ownerUserId)
  end
end