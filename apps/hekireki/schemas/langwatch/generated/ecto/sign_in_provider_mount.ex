defmodule LangWatch.SignInProviderMount do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          provider: String.t() | nil,
          mounted: boolean(),
          last_applied_event_id: String.t(),
          updated_at: DateTime.t()
        }

  schema "SignInProviderMount" do
    # Fixed, so this table holds one row: licensing's copy of sso's sign-in
    # provider fact (ADR-179).
    field(:id, :string, primary_key: true, default: "self")
    # The configured provider's name; null when the deployment signs in by email.
    field(:provider, :string)
    field(:mounted, :boolean)
    # The last sso event applied: the row's idempotency guard and its version.
    field(:last_applied_event_id, :string, source: :lastAppliedEventId)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end