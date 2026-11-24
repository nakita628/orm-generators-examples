defmodule LangWatch.LangyPullRequest do
  @moduledoc """
  One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
  checks look it up by repository and number to wake the conversation; the daily pull request
  cap counts a person's rows by day. Rebuilt by a projection replay.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          repository_full_name: String.t(),
          pr_number: integer(),
          conversation_id: String.t(),
          user_id: String.t(),
          opened_at: float(),
          source_event_id: String.t()
        }

  schema "LangyPullRequest" do
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    # `owner/repo`, lowercased: GitHub matches repository names without case.
    field(:repository_full_name, :string, source: :repositoryFullName)
    field(:pr_number, :integer, source: :prNumber)
    field(:conversation_id, :string, source: :conversationId)
    # The conversation's owner, who opened it through Langy.
    field(:user_id, :string, source: :userId)
    # When the answer that opened it was recorded (epoch ms).
    field(:opened_at, :float, source: :openedAt)
    # The agent_responded event it came from.
    field(:source_event_id, :string, source: :sourceEventId)
  end
end