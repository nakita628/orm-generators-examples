defmodule LangWatch.NurturingUncopiedProject do
  @moduledoc """
  A project nurturing saw traces of with no placement copy, and when it first did: every worker
  retries its trace batches only within one window from then, across restarts. Deleted when the
  copy lands.
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