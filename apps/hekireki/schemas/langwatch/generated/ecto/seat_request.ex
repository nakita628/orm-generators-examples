defmodule LangWatch.SeatRequest do
  @moduledoc """
  A person asked for a full seat the organization had no room for (modules/organization/specs/
  seat-requests.feature). One row per organization and address; reopening reuses the row.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          email: String.t(),
          user_id: String.t() | nil,
          name: String.t() | nil,
          status: String.t(),
          source: String.t(),
          idp_refused: boolean(),
          idp_kind: String.t() | nil,
          sales_notified_at: DateTime.t() | nil,
          opened_at: DateTime.t(),
          resolved_at: DateTime.t() | nil,
          resolved_by_user_id: String.t() | nil,
          seat_pending_since: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "SeatRequest" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # Lowercased address of the person asking.
    field(:email, :string)
    # Null when only the address is known (a sign-in the identity provider refused).
    field(:user_id, :string, source: :userId)
    field(:name, :string)
    # OPEN | APPROVED | DECLINED
    field(:status, :string)
    # pending | lite | idp_refused: how the person was left when the request opened.
    field(:source, :string)
    # The identity provider also refused this person's sign-in.
    field(:idp_refused, :boolean, default: false, source: :idpRefused)
    # okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml
    field(:idp_kind, :string, source: :idpKind)
    # When sales were told about this request; reset when it reopens.
    field(:sales_notified_at, LangWatch.PrismaDateTime, source: :salesNotifiedAt)
    field(:opened_at, LangWatch.PrismaDateTime, source: :openedAt)
    field(:resolved_at, LangWatch.PrismaDateTime, source: :resolvedAt)
    field(:resolved_by_user_id, :string, source: :resolvedByUserId)
    # Set when an approval claims the request, cleared once its seat is given; an approval
    # still pending long after is one whose seat never landed, and is reopened.
    field(:seat_pending_since, LangWatch.PrismaDateTime, source: :seatPendingSince)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end