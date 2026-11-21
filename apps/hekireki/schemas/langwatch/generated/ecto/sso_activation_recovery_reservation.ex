defmodule LangWatch.SsoActivationRecoveryReservation do
  @moduledoc """
  A durable claim on one live recovery path while activation is between its
  guard and projected ACTIVE fact. It has no expiry: only projection of the
  matching activation, or an explicit cancellation, may consume it.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          command_id: String.t(),
          organization_id: String.t(),
          connection_id: String.t(),
          created_at: DateTime.t()
        }

  schema "SsoActivationRecoveryReservation" do
    field(:command_id, :string, primary_key: true, source: :commandId)
    field(:organization_id, :string, source: :organizationId)
    field(:connection_id, :string, source: :connectionId)
    field(:created_at, LangWatch.PrismaDateTime, source: :createdAt)
  end
end