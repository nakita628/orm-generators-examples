defmodule LangWatch.NurturingProjectEvaluationDay do
  @moduledoc """
  Each project and UTC day with a completed evaluation, from evaluation's lifecycle_completed,
  written insert-ignore. Replaces the `evaluation` rows of NurturingProjectUsageView.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          project_id: String.t(),
          day: String.t()
        }

  schema "NurturingProjectEvaluationDay" do
    field(:project_id, :string, primary_key: true, source: :projectId)
    # UTC `YYYY-MM-DD`.
    field(:day, :string, primary_key: true)
  end
end