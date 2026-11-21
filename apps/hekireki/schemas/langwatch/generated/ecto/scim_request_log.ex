defmodule LangWatch.ScimRequestLog do
  @moduledoc """
  Every SCIM request we could attribute, and what we answered (ADR-126).

  EVIDENCE, NOT TRUTH. A request authors nothing — an HTTP round trip that
  was refused decided nothing — so this is a table and not an event: nothing
  downstream derives from it, rebuilding the world from the log must not
  depend on how often a provider retried a GET, and rows age out of a
  retention window an event log would never have.

  NO ROW FOR A REQUEST WE CANNOT ATTRIBUTE. A token that does not verify is
  answered 401 and recorded nowhere: we do not know whose organization to
  file it under, and a table unauthenticated traffic can write is a table
  anybody on the internet can fill. What answers that case instead is
  `ScimToken.lastUsedAt` being null. The 403 IS recorded — a lapsed plan is
  a credential we recognize.

  THE PAYLOAD RULE HOLDS. A method, a resource name, a status, a reason slug
  and our own sentence. Never a token, never a header, never a provider's
  raw message — the same reason `scim_apply_failed` carries a code and not a
  message.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          connection_id: String.t() | nil,
          method: String.t(),
          resource: String.t(),
          status: integer(),
          reason: String.t() | nil,
          detail: String.t() | nil,
          occurred_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil
        }

  schema "ScimRequestLog" do
    field(:id, :string, primary_key: true)
    # Null when the refusal resolved an organization and stopped before the
    # connection — a lapsed plan is answered off the token alone.
    field(:connection_id, :string, source: :connectionId)
    field(:method, :string)
    # The SCIM resource asked for ("Users", "Users/:id", "Groups"), never the
    # raw path: a path carries ids and query strings, and this is read by
    # people rather than matched by machines.
    field(:resource, :string)
    field(:status, :integer)
    # A stable slug, only on a refusal, so a reader branches on this and never
    # on the sentence beside it.
    field(:reason, :string)
    # Our own short sentence, customer-safe.
    field(:detail, :string)
    field(:occurred_at, LangWatch.PrismaDateTime, autogenerate: true, source: :occurredAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
  end
end