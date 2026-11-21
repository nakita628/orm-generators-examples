defmodule LangWatch.IdentitySignInSwitch do
  @moduledoc """
  Whether a user signs in on the identifier path: identity's own switch, set when the
  identifier backfill proves the user and pinned off by an operator's rollback. No row is the
  legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          user_id: String.t(),
          on_identifiers: boolean(),
          pinned: boolean(),
          switched_at: DateTime.t(),
          last_applied_event_id: String.t() | nil,
          updated_at: DateTime.t()
        }

  schema "IdentitySignInSwitch" do
    field(:user_id, :string, primary_key: true, source: :userId)
    field(:on_identifiers, :boolean, source: :onIdentifiers)
    field(:pinned, :boolean, default: false)
    field(:switched_at, LangWatch.PrismaDateTime, source: :switchedAt)
    field(:last_applied_event_id, :string, source: :lastAppliedEventId)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end