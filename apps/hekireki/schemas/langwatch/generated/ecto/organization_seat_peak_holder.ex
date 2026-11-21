defmodule LangWatch.OrganizationSeatPeakHolder do
  @moduledoc """
  One member's or pending invitation's seat as the seat peak last applied it (ADR-002).
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          holder_kind: String.t(),
          holder_id: String.t(),
          seat: String.t(),
          last_changed_at: float(),
          last_change_id: String.t() | nil,
          expires_at: float() | nil,
          expired_by_fold: boolean(),
          suppressed_by: String.t() | nil,
          restored_by_fold: boolean()
        }

  schema "OrganizationSeatPeakHolder" do
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    # `member`, `invitation` or `hold` (a seat held after somebody left).
    field(:holder_kind, :string, primary_key: true, source: :holderKind)
    # The member's or the hold's userId, or the invitation's id.
    field(:holder_id, :string, primary_key: true, source: :holderId)
    # `full`, `lite`, `developer` or `none`.
    field(:seat, :string)
    field(:last_changed_at, :float, source: :lastChangedAt)
    # Absent for a seat the opening snapshot recorded.
    field(:last_change_id, :string, source: :lastChangeId)
    # An invitation's expiry; null for a member or an invitation that never expires.
    field(:expires_at, :float, source: :expiresAt)
    # The fold gave the invitation's seat back at `expiresAt`.
    field(:expired_by_fold, :boolean, default: false, source: :expiredByFold)
    # A hold whose seat a live full invitation to its address took: that invitation's id.
    field(:suppressed_by, :string, source: :suppressedBy)
    # The fold gave a suppressed hold its seat back when that invitation expired.
    field(:restored_by_fold, :boolean, default: false, source: :restoredByFold)
  end
end