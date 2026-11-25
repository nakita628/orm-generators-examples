defmodule LangWatch.NurturingUncopiedProject do
  @moduledoc """
  retired: unread since v3.28.0. The trace-day folds count a trace with no placement copy, so no
  retry clock reads or writes this table; an operator step may drop it once old writers are gone.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          project_id: String.t(),
          first_seen_at: DateTime.t()
        }

  schema "NurturingUncopiedProject" do
    field(:project_id, :string, primary_key: true, source: :projectId)
    field(:first_seen_at, LangWatch.PrismaDateTime, autogenerate: true, source: :firstSeenAt)
  end
end