defmodule LangWatch.AuditLog do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          user_id: String.t() | nil,
          project_id: String.t() | nil,
          organization_id: String.t() | nil,
          action: String.t(),
          args: map() | nil,
          error: String.t() | nil,
          ip_address: String.t() | nil,
          user_agent: String.t() | nil,
          metadata: map() | nil,
          actor_user_id: String.t() | nil,
          target_kind: String.t() | nil,
          target_id: String.t() | nil,
          before: map() | nil,
          after: map() | nil,
          idempotency_key: String.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "AuditLog" do
    field(:id, :string, primary_key: true)
    field(:user_id, :string, source: :userId)
    field(:project_id, :string, source: :projectId)
    field(:organization_id, :string, source: :organizationId)
    field(:action, :string)
    field(:args, :map)
    field(:error, :string)
    field(:ip_address, :string, source: :ipAddress)
    field(:user_agent, :string, source: :userAgent)
    field(:metadata, :map)
    # WHO REALLY DID IT, when that is not `userId`.
    #
    # Under an impersonation the session's `user.id` is the SUBJECT — the
    # customer whose access an operator is borrowing — so a row naming only
    # `userId` filed the operator's act against the customer. An operator
    # granting a permanent way into somebody's organization appeared, in the
    # only durable record of it, to be that somebody.
    #
    # Null on an ordinary request, where the two would be the same person and
    # saying so twice adds nothing. ADR-118 / D06: every decision names both.
    field(:actor_user_id, :string, source: :actorUserId)
    field(:target_kind, :string, source: :targetKind)
    field(:target_id, :string, source: :targetId)
    field(:before, :map)
    field(:after, :map)
    # A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2).
    field(:idempotency_key, :string, source: :idempotencyKey)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end