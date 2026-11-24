defmodule Trigger.BatchTaskRunError do
  @moduledoc """
  Track individual run creation failures in batch processing (Run Engine v2)
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          index: integer(),
          task_identifier: String.t(),
          payload: String.t() | nil,
          options: map() | list() | String.t() | number() | boolean() | nil,
          error: String.t(),
          error_code: String.t() | nil,
          inserted_at: DateTime.t(),
          batch_task_run: Trigger.BatchTaskRun.t() | nil
        }

  schema "BatchTaskRunError" do
    field(:id, :string, primary_key: true)
    # Which item in the batch (0-based index)
    field(:index, :integer)
    # The task identifier that was being triggered
    field(:task_identifier, :string, source: :taskIdentifier)
    # The payload that failed (JSON, may be truncated)
    field(:payload, :string)
    # The options that were used
    field(:options, Trigger.PrismaJson)
    # Error message
    field(:error, :string)
    # Error code if available
    field(:error_code, :string, source: :errorCode)
    field(:batch_task_run_id, :string, source: :batchTaskRunId)
    belongs_to(:batch_task_run, Trigger.BatchTaskRun, foreign_key: :batch_task_run_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end