defmodule Shelf.AssetImport do
  @moduledoc """
  One CSV content import. Holds the parsed rows from preview until the import
  completes, so the job imports exactly the file the user previewed.
  A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
  active import per workspace. Prisma cannot express it; it lives in migration
  20261015120100_asset_import_one_active_per_workspace.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          status: atom(),
          phase: atom(),
          file_name: String.t() | nil,
          total_rows: integer(),
          processed_rows: integer(),
          rows: map() | nil,
          resolved: map() | nil,
          can_use_barcodes: boolean(),
          error_message: String.t() | nil,
          warnings: map() | nil,
          warning_count: integer(),
          locked_until: DateTime.t() | nil,
          confirmed_at: DateTime.t() | nil,
          started_at: DateTime.t() | nil,
          completed_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          created_by: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "AssetImport" do
    field(:id, :string, primary_key: true)
    field(:status, Ecto.Enum, values: [:DRAFT, :QUEUED, :RUNNING, :COMPLETED, :FAILED], default: :DRAFT)
    field(:phase, Ecto.Enum, values: [:PREPARE, :ROWS, :FINALIZE], default: :PREPARE)
    field(:file_name, :string, source: :fileName)
    field(:total_rows, :integer, source: :totalRows)
    # Rows committed so far. Advanced in the same transaction as each chunk.
    field(:processed_rows, :integer, default: 0, source: :processedRows)
    # Parsed rows (`StoredImportRow[]`). Cleared when the import completes.
    field(:rows, :map)
    # Name to id maps written by the PREPARE step (`ResolvedImportEntities`).
    field(:resolved, :map)
    field(:can_use_barcodes, :boolean, default: false, source: :canUseBarcodes)
    # User-facing reason the import failed. Null unless status is FAILED.
    field(:error_message, :string, source: :errorMessage)
    # Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500.
    field(:warnings, :map)
    # Exact number of warnings, including any beyond the stored cap.
    field(:warning_count, :integer, default: 0, source: :warningCount)
    # A step holds the import until this time. Null when no step is running.
    field(:locked_until, Shelf.PrismaDateTime, source: :lockedUntil)
    field(:confirmed_at, Shelf.PrismaDateTime, source: :confirmedAt)
    field(:started_at, Shelf.PrismaDateTime, source: :startedAt)
    field(:completed_at, Shelf.PrismaDateTime, source: :completedAt)
    field(:created_by_id, :string, source: :createdById)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:created_by, Shelf.User, foreign_key: :created_by_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end