defmodule Shelf.AuditAssignment do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          role: atom() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          audit_session: Shelf.AuditSession.t() | nil,
          user: Shelf.User.t() | nil
        }

  schema "AuditAssignment" do
    field(:id, :string, primary_key: true)
    field(:role, Ecto.Enum, values: [:LEAD, :PARTICIPANT])
    field(:audit_session_id, :string, source: :auditSessionId)
    field(:user_id, :string, source: :userId)
    belongs_to(:audit_session, Shelf.AuditSession, foreign_key: :audit_session_id, define_field: false, type: :string)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end