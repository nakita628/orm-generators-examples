defmodule Formbricks.EmbeddedData do
  @moduledoc """
  Defines one Embedded Data field: the extra context stored on a response that isn't an
  answer to a question (brand, plan, page type, a computed score, ...).

  A field is either **local** — used by a single survey, which is how existing variables and
  hidden fields are migrated — or **shared**, meaning it is part of the workspace library and can
  be linked by many surveys. There is no stored flag for that: `surveyId` is the owner and `key`
  is the library name, so a local field has `surveyId` set and `key` null, and a shared field the
  other way round. Deriving it rather than storing a boolean removes a column that could drift
  out of step with the two fields it stands for.

  @property key - Library name, unique per workspace and immutable. Null for local fields
  @property name - Human-readable display label
  @property source - Where the value comes from (computed | ingested). `reserved` is part of the
  shared enum but never stored as a row: those fields are a static catalog in code
  @property dataType - Declared value type, used for coercion at ingest
  @property defaultValue - Initial value (computed) or fallback when nothing arrives (ingested)
  @property locked - Ingested only: ignore external writes and use the default instead
  @property survey - The owning survey for a local field; cascades so nothing is orphaned
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          key: String.t() | nil,
          name: String.t(),
          description: String.t() | nil,
          source: atom(),
          data_type: atom(),
          default_value: map() | nil,
          locked: boolean(),
          survey_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          workspace: Formbricks.Workspace.t() | nil
        }

  schema "EmbeddedData" do
    field(:id, :string, primary_key: true)
    field(:key, :string)
    field(:name, :string)
    field(:description, :string)
    field(:source, Ecto.Enum, values: [:computed, :ingested, :reserved])
    field(:data_type, Ecto.Enum, values: [:string, :number, :boolean, :date], default: :string, source: :dataType)
    # [EmbeddedDataDefaultValue]
    field(:default_value, :map, source: :defaultValue)
    field(:locked, :boolean, default: false)
    field(:survey_id, :string, source: :surveyId)
    field(:workspace_id, :string, source: :workspaceId)
    belongs_to(:workspace, Formbricks.Workspace, foreign_key: :workspace_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end