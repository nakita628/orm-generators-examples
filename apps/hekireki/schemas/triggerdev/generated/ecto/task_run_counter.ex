defmodule Trigger.TaskRunCounter do
  @moduledoc """
  deprecated, we hadn't included the project id in the unique constraint
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          task_identifier: String.t(),
          last_number: integer()
        }

  schema "TaskRunCounter" do
    field(:task_identifier, :string, primary_key: true, source: :taskIdentifier)
    field(:last_number, :integer, default: 0, source: :lastNumber)
  end
end