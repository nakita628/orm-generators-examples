defmodule Plunk.Contact do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Plunk.PrismaDateTime, autogenerate: {Plunk.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          email: String.t(),
          data: map() | list() | String.t() | number() | boolean() | nil,
          subscribed: boolean(),
          snoozed_until: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: Plunk.Project.t() | nil,
          emails: [Plunk.Email.t()],
          workflow_executions: [Plunk.WorkflowExecution.t()],
          events: [Plunk.Event.t()],
          segment_memberships: [Plunk.SegmentMembership.t()],
          throttle_claims: [Plunk.ThrottleClaim.t()]
        }

  schema "contacts" do
    field(:email, :string)
    field(:data, Plunk.PrismaJson)
    field(:subscribed, :boolean, default: true)
    # When set, `subscribed = false` is temporary rather than permanent: the snooze
    # sweep resubscribes the contact once this passes. Suppression itself is still
    # carried entirely by `subscribed`, so every send path filters snoozed contacts
    # out without knowing this column exists.
    #
    # Cleared by every other write to `subscribed` — including the bounce and
    # complaint paths, so a suppressed address can never be woken back up.
    field(:snoozed_until, Plunk.PrismaDateTime, source: :snoozedUntil)
    field(:project_id, :string, source: :projectId)
    belongs_to(:project, Plunk.Project, foreign_key: :project_id, define_field: false)
    has_many(:emails, Plunk.Email, foreign_key: :contact_id)
    has_many(:workflow_executions, Plunk.WorkflowExecution, foreign_key: :contact_id)
    has_many(:events, Plunk.Event, foreign_key: :contact_id)
    has_many(:segment_memberships, Plunk.SegmentMembership, foreign_key: :contact_id)
    has_many(:throttle_claims, Plunk.ThrottleClaim, foreign_key: :contact_id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end