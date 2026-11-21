defmodule Formbricks.OrganizationBilling do
  @moduledoc """
  Stores billing and Stripe synchronization data for an organization.

  @property organizationId - Reference to the organization
  @property stripeCustomerId - Stripe customer ID for cloud billing
  @property limits - Plan limits used by feature and quota checks
  @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
  @property stripe - Stripe synchronization snapshot data
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          stripe_customer_id: String.t() | nil,
          limits: map(),
          usage_cycle_anchor: DateTime.t() | nil,
          stripe: map() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Formbricks.Organization.t() | nil
        }

  schema "OrganizationBilling" do
    field(:stripe_customer_id, :string)
    # [OrganizationBillingPlanLimits]
    field(:limits, :map)
    field(:usage_cycle_anchor, Formbricks.PrismaDateTime)
    # [OrganizationStripeBilling]
    field(:stripe, :map)
    belongs_to(:organization, Formbricks.Organization, foreign_key: :organization_id, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end