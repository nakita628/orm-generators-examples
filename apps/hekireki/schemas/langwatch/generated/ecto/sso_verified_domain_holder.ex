defmodule LangWatch.SsoVerifiedDomainHolder do
  @moduledoc """
  A connection currently entitled to use an organization-owned verified
  domain. Usually there is one holder. During an Auth0 migration there are
  exactly two: the grandfathered predecessor and its direct replacement.

  The projection repository enforces that relationship before inserting a
  second holder. The flattened migration also installs partial unique live
  legacy/direct slots on `SsoConnection`, closing concurrent third-writer
  races that Prisma cannot describe declaratively.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          domain: String.t(),
          connection_id: String.t(),
          organization_id: String.t()
        }

  schema "SsoVerifiedDomainHolder" do
    field(:domain, :string, primary_key: true)
    field(:connection_id, :string, primary_key: true, source: :connectionId)
    field(:organization_id, :string, source: :organizationId)
  end
end