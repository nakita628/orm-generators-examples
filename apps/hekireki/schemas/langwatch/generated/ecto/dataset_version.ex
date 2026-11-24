defmodule LangWatch.DatasetVersion do
  @moduledoc """
  One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          dataset_id: String.t(),
          version: integer(),
          name: String.t() | nil,
          trigger: String.t(),
          content_hash: String.t(),
          column_types: map() | list() | String.t() | number() | boolean(),
          row_count: integer(),
          size_bytes: integer(),
          chunk_count: integer(),
          restored_from_version: integer() | nil,
          author_id: String.t() | nil,
          author_kind: String.t() | nil,
          pinned_at: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "DatasetVersion" do
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    field(:dataset_id, :string, source: :datasetId)
    field(:version, :integer)
    # A name given by hand; a named version is never thinned by retention.
    field(:name, :string)
    # "autosave", "manual", "experiment_run", "restore" or "initial".
    field(:trigger, :string)
    # SHA-256 over the columns and the rows in order; equal content saves no new version.
    field(:content_hash, :string, source: :contentHash)
    field(:column_types, LangWatch.PrismaJson, source: :columnTypes)
    field(:row_count, :integer, source: :rowCount)
    field(:size_bytes, :integer, source: :sizeBytes)
    field(:chunk_count, :integer, source: :chunkCount)
    field(:restored_from_version, :integer, source: :restoredFromVersion)
    field(:author_id, :string, source: :authorId)
    # "user", "api_key" or "system": who saved the version when authorId is null.
    field(:author_kind, :string, source: :authorKind)
    # Set when an experiment pins this version; a pinned version is never thinned.
    field(:pinned_at, LangWatch.PrismaDateTime, source: :pinnedAt)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end