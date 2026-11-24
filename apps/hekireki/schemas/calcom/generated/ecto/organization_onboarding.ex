defmodule Calcom.OrganizationOnboarding do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Calcom.PrismaDateTime, autogenerate: {Calcom.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          org_owner_email: String.t(),
          error: String.t() | nil,
          billing_period: atom(),
          price_per_seat: float(),
          seats: integer(),
          is_platform: boolean(),
          name: String.t(),
          slug: String.t(),
          logo: String.t() | nil,
          bio: String.t() | nil,
          brand_color: String.t() | nil,
          banner_url: String.t() | nil,
          is_domain_configured: boolean(),
          stripe_customer_id: String.t() | nil,
          stripe_subscription_id: String.t() | nil,
          stripe_subscription_item_id: String.t() | nil,
          invited_members: map() | list() | String.t() | number() | boolean(),
          teams: map() | list() | String.t() | number() | boolean(),
          is_complete: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          created_by: Calcom.User.t() | nil,
          organization: Calcom.Team.t() | nil
        }

  schema "OrganizationOnboarding" do
    field(:org_owner_email, :string, source: :orgOwnerEmail)
    field(:error, :string)
    field(:billing_period, Ecto.Enum, values: [:MONTHLY, :ANNUALLY], source: :billingPeriod)
    field(:price_per_seat, :float, source: :pricePerSeat)
    field(:seats, :integer)
    field(:is_platform, :boolean, default: false, source: :isPlatform)
    field(:name, :string)
    field(:slug, :string)
    field(:logo, :string)
    field(:bio, :string)
    field(:brand_color, :string, source: :brandColor)
    field(:banner_url, :string, source: :bannerUrl)
    field(:is_domain_configured, :boolean, default: false, source: :isDomainConfigured)
    field(:stripe_customer_id, :string, source: :stripeCustomerId)
    field(:stripe_subscription_id, :string, source: :stripeSubscriptionId)
    field(:stripe_subscription_item_id, :string, source: :stripeSubscriptionItemId)
    # @zod.import(["import { orgOnboardingInvitedMembersSchema } from '../../zod-utils'"]).custom.use(orgOnboardingInvitedMembersSchema)
    field(:invited_members, Calcom.PrismaJson, default: [], source: :invitedMembers)
    # @zod.import(["import { orgOnboardingTeamsSchema } from '../../zod-utils'"]).custom.use(orgOnboardingTeamsSchema)
    field(:teams, Calcom.PrismaJson, default: [])
    field(:is_complete, :boolean, default: false, source: :isComplete)
    field(:created_by_id, :id, source: :createdById)
    field(:organization_id, :id, source: :organizationId)
    belongs_to(:created_by, Calcom.User, foreign_key: :created_by_id, define_field: false, type: :id)
    belongs_to(:organization, Calcom.Team, foreign_key: :organization_id, define_field: false, type: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end