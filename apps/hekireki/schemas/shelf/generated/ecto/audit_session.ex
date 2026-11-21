defmodule Shelf.AuditSession do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          target_id: String.t() | nil,
          status: atom(),
          scope_meta: map() | nil,
          expected_asset_count: integer(),
          found_asset_count: integer(),
          missing_asset_count: integer(),
          unexpected_asset_count: integer(),
          started_at: DateTime.t() | nil,
          due_date: DateTime.t() | nil,
          completed_at: DateTime.t() | nil,
          cancelled_at: DateTime.t() | nil,
          active_scheduler_reference: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          created_by: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil,
          assignments: [Shelf.AuditAssignment.t()],
          assets: [Shelf.AuditAsset.t()],
          scans: [Shelf.AuditScan.t()],
          notes: [Shelf.AuditNote.t()],
          images: [Shelf.AuditImage.t()]
        }

  schema "AuditSession" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:target_id, :string, source: :targetId)
    field(:status, Ecto.Enum, values: [:PENDING, :ACTIVE, :COMPLETED, :CANCELLED, :ARCHIVED], default: :PENDING)
    field(:scope_meta, :map, source: :scopeMeta)
    field(:expected_asset_count, :integer, default: 0, source: :expectedAssetCount)
    field(:found_asset_count, :integer, default: 0, source: :foundAssetCount)
    field(:missing_asset_count, :integer, default: 0, source: :missingAssetCount)
    field(:unexpected_asset_count, :integer, default: 0, source: :unexpectedAssetCount)
    field(:started_at, Shelf.PrismaDateTime, source: :startedAt)
    field(:due_date, Shelf.PrismaDateTime, source: :dueDate)
    field(:completed_at, Shelf.PrismaDateTime, source: :completedAt)
    field(:cancelled_at, Shelf.PrismaDateTime, source: :cancelledAt)
    field(:active_scheduler_reference, :string, source: :activeSchedulerReference)
    field(:created_by_id, :string, source: :createdById)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:created_by, Shelf.User, foreign_key: :created_by_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:assignments, Shelf.AuditAssignment, foreign_key: :audit_session_id, references: :id)
    has_many(:assets, Shelf.AuditAsset, foreign_key: :audit_session_id, references: :id)
    has_many(:scans, Shelf.AuditScan, foreign_key: :audit_session_id, references: :id)
    has_many(:notes, Shelf.AuditNote, foreign_key: :audit_session_id, references: :id)
    has_many(:images, Shelf.AuditImage, foreign_key: :audit_session_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end