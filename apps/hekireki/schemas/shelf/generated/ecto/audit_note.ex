defmodule Shelf.AuditNote do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          content: String.t(),
          type: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil,
          audit_session: Shelf.AuditSession.t() | nil,
          audit_asset: Shelf.AuditAsset.t() | nil
        }

  schema "AuditNote" do
    field(:id, :string, primary_key: true)
    field(:content, :string)
    field(:type, Ecto.Enum, values: [:COMMENT, :UPDATE], default: :COMMENT)
    field(:user_id, :string, source: :userId)
    field(:audit_session_id, :string, source: :auditSessionId)
    field(:audit_asset_id, :string, source: :auditAssetId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:audit_session, Shelf.AuditSession, foreign_key: :audit_session_id, define_field: false, type: :string)
    belongs_to(:audit_asset, Shelf.AuditAsset, foreign_key: :audit_asset_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end