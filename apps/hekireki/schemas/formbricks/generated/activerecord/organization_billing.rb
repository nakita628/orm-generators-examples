# Stores billing and Stripe synchronization data for an organization.
#
# @property organizationId - Reference to the organization
# @property stripeCustomerId - Stripe customer ID for cloud billing
# @property limits - Plan limits used by feature and quota checks
# @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
# @property stripe - Stripe synchronization snapshot data
#
# limits:: [OrganizationBillingPlanLimits]
# stripe:: [OrganizationStripeBilling]
class OrganizationBilling < ApplicationRecord
  self.table_name = "OrganizationBilling"
  self.primary_key = "organization_id"

  validates :stripe_customer_id, uniqueness: { allow_nil: true }
  validates :limits, exclusion: { in: [ nil ] }

  belongs_to :organization, inverse_of: :billing
end
