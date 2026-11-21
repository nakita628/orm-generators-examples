defmodule LangWatch.FeatureFlagExperimentSetting do
  @moduledoc """
  One experiment setting per (flag, subject). A USER row is that person's
  own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
  for that exact scope. No row means inherit.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          flag_key: String.t(),
          subject_type: atom(),
          subject_id: String.t(),
          enabled: boolean(),
          changed_by_user_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "FeatureFlagExperimentSetting" do
    field(:flag_key, :string, primary_key: true, source: :flagKey)
    field(:subject_type, Ecto.Enum, values: [:USER, :ORGANIZATION, :PROJECT], source: :subjectType)
    field(:subject_id, :string, primary_key: true, source: :subjectId)
    field(:enabled, :boolean)
    field(:changed_by_user_id, :string, source: :changedByUserId)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end