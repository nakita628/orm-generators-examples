defmodule Formbricks.ActionClass do
  @moduledoc """
  Represents a user action that can trigger surveys.
  Used to define points in the user journey where surveys can be shown.

  @property id - Unique identifier for the action
  @property name - Display name of the action
  @property type - Whether this is a code or no-code action
  @property key - Unique identifier used in code implementation
  @property noCodeConfig - Configuration for no-code setup
  @property workspace - The workspace this action belongs to
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          type: atom(),
          key: String.t() | nil,
          no_code_config: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          workspace: Formbricks.Workspace.t() | nil,
          survey_triggers: [Formbricks.SurveyTrigger.t()],
          creative_triggers: [Formbricks.CreativeTrigger.t()]
        }

  schema "ActionClass" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:type, Ecto.Enum, values: [:code, :noCode])
    field(:key, :string)
    # [ActionClassNoCodeConfig]
    field(:no_code_config, Formbricks.PrismaJson, source: :noCodeConfig)
    field(:workspace_id, :string, source: :workspaceId)
    belongs_to(:workspace, Formbricks.Workspace, foreign_key: :workspace_id, define_field: false, type: :string)
    has_many(:survey_triggers, Formbricks.SurveyTrigger, foreign_key: :action_class_id, references: :id)
    has_many(:creative_triggers, Formbricks.CreativeTrigger, foreign_key: :action_class_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end