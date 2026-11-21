class User < ApplicationRecord
  self.table_name = "User"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :username, default: -> { Cuid.generate }
  attribute :usedFreeTrial, default: false
  attribute :onboarded, default: false
  attribute :sso, default: false
  attribute :createdWithInvite, default: false
  attribute :skipSubscriptionCheck, default: false
  attribute :hasUnpaidInvoice, default: false
  attribute :warnForNoPaymentMethod, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :dateFormat, { dd_mm_yyyy: "DD_MM_YYYY", mm_dd_yyyy: "MM_DD_YYYY", yyyy_mm_dd: "YYYY_MM_DD", mmm_dd_yyyy: "MMM_DD_YYYY", dd_mmm_yyyy: "DD_MMM_YYYY" },
       validate: { allow_nil: true }
  enum :timeFormat, { h12: "H12", h24: "H24" }, validate: { allow_nil: true }
  enum :weekStart, { monday: "MONDAY", sunday: "SUNDAY", saturday: "SATURDAY" }, validate: { allow_nil: true }
  enum :tierId, { free: "free", tier_1: "tier_1", tier_2: "tier_2", custom: "custom" }, default: :free, validate: true

  validates :email, presence: true, uniqueness: true
  validates :email, uniqueness: { scope: :username }
  validates :username, uniqueness: true
  validates :customerId, uniqueness: { allow_nil: true }

  belongs_to :last_selected_organization, class_name: "Organization", foreign_key: "lastSelectedOrganizationId",
                                          optional: true, inverse_of: :users_with_last_selected
  belongs_to :tier, foreign_key: "tierId", inverse_of: :subscribers
  has_one :contact, class_name: "UserContact", foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_one :business_intel, class_name: "UserBusinessIntel", foreign_key: "userId", inverse_of: :user,
                           dependent: :destroy
  has_one :notification_preference, foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_one :custom_tier_limit, foreign_key: "userId", inverse_of: :user, dependent: :nullify
  has_many :assets, foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_many :categories, foreign_key: "userId", inverse_of: :user, dependent: :restrict_with_error
  has_many :notes, foreign_key: "userId", inverse_of: :user, dependent: :nullify
  has_many :booking_notes, foreign_key: "userId", inverse_of: :user, dependent: :nullify
  has_many :audit_notes, foreign_key: "userId", inverse_of: :user, dependent: :nullify
  has_many :location_notes, foreign_key: "userId", inverse_of: :user, dependent: :nullify
  has_many :qr_codes, class_name: "Qr", foreign_key: "userId", inverse_of: :user, dependent: :nullify
  has_many :scans, foreign_key: "userId", inverse_of: :user, dependent: :nullify
  has_many :tags, foreign_key: "userId", inverse_of: :user, dependent: :restrict_with_error
  has_many :locations, foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_many :images, foreign_key: "userId", inverse_of: :user, dependent: :restrict_with_error
  has_many :organizations, foreign_key: "userId", inverse_of: :owner, dependent: :destroy
  has_many :custom_fields, foreign_key: "userId", inverse_of: :created_by, dependent: :restrict_with_error
  has_many :sent_invites, class_name: "Invite", foreign_key: "inviterId", inverse_of: :inviter, dependent: :destroy
  has_many :received_invites, class_name: "Invite", foreign_key: "inviteeUserId", inverse_of: :invitee_user,
                              dependent: :destroy
  has_many :team_members, foreign_key: "userId", inverse_of: :user, dependent: :nullify
  has_many :user_organizations, foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_many :bookings, foreign_key: "creatorId", inverse_of: :creator, dependent: :destroy
  has_many :custodies, class_name: "Booking", foreign_key: "custodianUserId", inverse_of: :custodian_user,
                       dependent: :destroy
  has_many :created_kits, class_name: "Kit", foreign_key: "createdById", inverse_of: :created_by,
                          dependent: :restrict_with_error
  has_many :asset_reminders, foreign_key: "createdById", inverse_of: :created_by, dependent: :restrict_with_error
  has_many :audits_created, class_name: "AuditSession", foreign_key: "createdById", inverse_of: :created_by,
                            dependent: :restrict_with_error
  has_many :asset_imports, foreign_key: "createdById", inverse_of: :created_by, dependent: :destroy
  has_many :audit_assignments, foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_many :audit_assets_scanned, class_name: "AuditAsset", foreign_key: "scannedById", inverse_of: :scanned_by,
                                  dependent: :nullify
  has_many :audit_scans, foreign_key: "scannedById", inverse_of: :scanned_by, dependent: :nullify
  has_many :audit_images_uploaded, class_name: "AuditImage", foreign_key: "uploadedById", inverse_of: :uploaded_by,
                                   dependent: :nullify
  has_many :asset_index_settings, class_name: "AssetIndexSettings", foreign_key: "userId", inverse_of: :user,
                                  dependent: :destroy
  has_many :partial_checkins, class_name: "PartialBookingCheckin", foreign_key: "checkedInById",
                              inverse_of: :checked_in_by, dependent: :restrict_with_error
  has_many :partial_checkouts, class_name: "PartialBookingCheckout", foreign_key: "checkedOutById",
                               inverse_of: :checked_out_by, dependent: :restrict_with_error
  has_many :mobile_auth_codes, foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_many :push_tokens, foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_many :created_updates, class_name: "Update", foreign_key: "createdById", inverse_of: :created_by,
                             dependent: :destroy
  has_many :update_reads, class_name: "UserUpdateRead", foreign_key: "userId", inverse_of: :user, dependent: :destroy
  has_many :asset_filter_presets, foreign_key: "ownerId", inverse_of: :owner, dependent: :destroy
  has_many :asset_models, foreign_key: "userId", inverse_of: :created_by, dependent: :destroy
  has_many :consumption_logs, foreign_key: "userId", inverse_of: :performed_by, dependent: :restrict_with_error
  has_many :scim_external_ids, class_name: "UserScimExternalId", foreign_key: "userId", inverse_of: :user,
                               dependent: :destroy
  has_many :role_changes_received, class_name: "RoleChangeLog", foreign_key: "userId", inverse_of: :user,
                                   dependent: :restrict_with_error
  has_many :role_changes_performed, class_name: "RoleChangeLog", foreign_key: "changedById", inverse_of: :changed_by,
                                    dependent: :restrict_with_error
  has_many :team_member_notes_authored, class_name: "TeamMemberNote", foreign_key: "userId", inverse_of: :user,
                                        dependent: :nullify
  has_and_belongs_to_many :roles, join_table: "_RoleToUser", foreign_key: "B", association_foreign_key: "A"
end
