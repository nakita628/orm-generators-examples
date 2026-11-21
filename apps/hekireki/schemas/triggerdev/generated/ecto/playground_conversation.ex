defmodule Trigger.PlaygroundConversation do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          chat_id: String.t(),
          title: String.t(),
          agent_slug: String.t(),
          client_data: map() | nil,
          messages: map() | nil,
          last_event_id: String.t() | nil,
          user_id: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          run: Trigger.TaskRun.t() | nil,
          project: Trigger.Project.t() | nil,
          runtime_environment: Trigger.RuntimeEnvironment.t() | nil
        }

  schema "PlaygroundConversation" do
    field(:id, :string, primary_key: true)
    # The chat session ID used by the transport
    field(:chat_id, :string, source: :chatId)
    # User-editable conversation title (auto-generated from first message)
    field(:title, :string, default: "New conversation")
    # Which agent this conversation is with
    field(:agent_slug, :string, source: :agentSlug)
    # The client data JSON used for this conversation
    field(:client_data, :map, source: :clientData)
    # Accumulated UIMessages from completed turns (for resume without stream replay)
    field(:messages, :map)
    # Last SSE event ID — resume from this position to avoid replaying old turns
    field(:last_event_id, :string, source: :lastEventId)
    # The user who started this conversation
    field(:user_id, :string, source: :userId)
    # The current active run backing this conversation (null if no run yet)
    field(:run_id, :string, source: :runId)
    field(:project_id, :string, source: :projectId)
    field(:runtime_environment_id, :string, source: :runtimeEnvironmentId)
    belongs_to(:run, Trigger.TaskRun, foreign_key: :run_id, define_field: false, type: :string)
    belongs_to(:project, Trigger.Project, foreign_key: :project_id, define_field: false, type: :string)
    belongs_to(:runtime_environment, Trigger.RuntimeEnvironment, foreign_key: :runtime_environment_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end