defmodule Formbricks.FeedbackSourceFieldMapping do
  @moduledoc """
  Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
  Maps source fields to Hub FeedbackRecord fields.

  @property id - Unique identifier for the mapping
  @property feedbackSource - The parent feedback source
  @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
  @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
  @property staticValue - If set, use this value instead of reading from sourceFieldId
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          feedback_source_id: String.t(),
          workspace_id: String.t(),
          source_field_id: String.t(),
          target_field_id: String.t(),
          static_value: String.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "FeedbackSourceFieldMapping" do
    field(:id, :string, primary_key: true)
    field(:feedback_source_id, :string)
    field(:workspace_id, :string, source: :workspaceId)
    field(:source_field_id, :string)
    field(:target_field_id, :string)
    field(:static_value, :string)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end