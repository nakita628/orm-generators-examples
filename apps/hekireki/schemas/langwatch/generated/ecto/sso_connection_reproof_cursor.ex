defmodule LangWatch.SsoConnectionReproofCursor do
  @moduledoc """
  Operational round-robin position for ADR-123's domain re-proof sweep.

  This is deliberately not a column on `SsoConnection`: that row is a
  whole-row event projection and its fold is its only writer, while a
  healthy re-read states no fact. Keeping the scheduling cursor beside the
  projection preserves both rules and still lets a fixed-size sweep revisit
  every connection, including one whose proof has started wavering.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          connection_id: String.t(),
          last_reproof_at: DateTime.t(),
          connection: LangWatch.SsoConnection.t() | nil
        }

  schema "SsoConnectionReproofCursor" do
    field(:last_reproof_at, LangWatch.PrismaDateTime, source: :lastReproofAt)
    field(:connection_id, :string, primary_key: true, source: :connectionId)
    belongs_to(:connection, LangWatch.SsoConnection, foreign_key: :connection_id, define_field: false, type: :string)
  end
end