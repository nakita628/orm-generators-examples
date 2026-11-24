defmodule Shelf.User do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          email: String.t(),
          username: String.t(),
          first_name: String.t() | nil,
          last_name: String.t() | nil,
          display_name: String.t() | nil,
          profile_picture: String.t() | nil,
          used_free_trial: boolean(),
          onboarded: boolean(),
          customer_id: String.t() | nil,
          sso: boolean(),
          created_with_invite: boolean(),
          skip_subscription_check: boolean(),
          has_unpaid_invoice: boolean(),
          warn_for_no_payment_method: boolean(),
          deleted_at: DateTime.t() | nil,
          last_mobile_active_at: DateTime.t() | nil,
          date_format: atom() | nil,
          time_format: atom() | nil,
          week_start: atom() | nil,
          time_zone: String.t() | nil,
          referral_source: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          last_selected_organization: Shelf.Organization.t() | nil,
          tier: Shelf.Tier.t() | nil,
          contact: Shelf.UserContact.t() | nil,
          business_intel: Shelf.UserBusinessIntel.t() | nil,
          notification_preference: Shelf.NotificationPreference.t() | nil,
          custom_tier_limit: Shelf.CustomTierLimit.t() | nil,
          assets: [Shelf.Asset.t()],
          categories: [Shelf.Category.t()],
          notes: [Shelf.Note.t()],
          booking_notes: [Shelf.BookingNote.t()],
          audit_notes: [Shelf.AuditNote.t()],
          location_notes: [Shelf.LocationNote.t()],
          qr_codes: [Shelf.Qr.t()],
          scans: [Shelf.Scan.t()],
          tags: [Shelf.Tag.t()],
          locations: [Shelf.Location.t()],
          images: [Shelf.Image.t()],
          organizations: [Shelf.Organization.t()],
          custom_fields: [Shelf.CustomField.t()],
          sent_invites: [Shelf.Invite.t()],
          received_invites: [Shelf.Invite.t()],
          team_members: [Shelf.TeamMember.t()],
          user_organizations: [Shelf.UserOrganization.t()],
          bookings: [Shelf.Booking.t()],
          custodies: [Shelf.Booking.t()],
          approved_bookings: [Shelf.Booking.t()],
          created_kits: [Shelf.Kit.t()],
          asset_reminders: [Shelf.AssetReminder.t()],
          audits_created: [Shelf.AuditSession.t()],
          asset_imports: [Shelf.AssetImport.t()],
          audit_assignments: [Shelf.AuditAssignment.t()],
          audit_assets_scanned: [Shelf.AuditAsset.t()],
          audit_scans: [Shelf.AuditScan.t()],
          audit_images_uploaded: [Shelf.AuditImage.t()],
          asset_index_settings: [Shelf.AssetIndexSettings.t()],
          partial_checkins: [Shelf.PartialBookingCheckin.t()],
          partial_checkouts: [Shelf.PartialBookingCheckout.t()],
          mobile_auth_codes: [Shelf.MobileAuthCode.t()],
          push_tokens: [Shelf.PushToken.t()],
          created_updates: [Shelf.Update.t()],
          update_reads: [Shelf.UserUpdateRead.t()],
          asset_filter_presets: [Shelf.AssetFilterPreset.t()],
          asset_models: [Shelf.AssetModel.t()],
          consumption_logs: [Shelf.ConsumptionLog.t()],
          scim_external_ids: [Shelf.UserScimExternalId.t()],
          role_changes_received: [Shelf.RoleChangeLog.t()],
          role_changes_performed: [Shelf.RoleChangeLog.t()],
          team_member_notes_authored: [Shelf.TeamMemberNote.t()],
          roles: [Shelf.Role.t()]
        }

  schema "User" do
    field(:id, :string, primary_key: true)
    field(:email, :string)
    field(:username, :string)
    field(:first_name, :string, source: :firstName)
    field(:last_name, :string, source: :lastName)
    field(:display_name, :string, source: :displayName)
    field(:profile_picture, :string, source: :profilePicture)
    field(:used_free_trial, :boolean, default: false, source: :usedFreeTrial)
    field(:onboarded, :boolean, default: false)
    field(:customer_id, :string, source: :customerId)
    field(:sso, :boolean, default: false)
    field(:created_with_invite, :boolean, default: false, source: :createdWithInvite)
    field(:skip_subscription_check, :boolean, default: false, source: :skipSubscriptionCheck)
    field(:has_unpaid_invoice, :boolean, default: false, source: :hasUnpaidInvoice)
    field(:warn_for_no_payment_method, :boolean, default: false, source: :warnForNoPaymentMethod)
    field(:deleted_at, Shelf.PrismaDateTime, source: :deletedAt)
    field(:last_mobile_active_at, Shelf.PrismaDateTime, source: :lastMobileActiveAt)
    field(:date_format, Ecto.Enum, values: [:DD_MM_YYYY, :MM_DD_YYYY, :YYYY_MM_DD, :MMM_DD_YYYY, :DD_MMM_YYYY], source: :dateFormat)
    field(:time_format, Ecto.Enum, values: [:H12, :H24], source: :timeFormat)
    field(:week_start, Ecto.Enum, values: [:MONDAY, :SUNDAY, :SATURDAY], source: :weekStart)
    field(:time_zone, :string, source: :timeZone)
    field(:referral_source, :string, source: :referralSource)
    field(:last_selected_organization_id, :string, source: :lastSelectedOrganizationId)
    field(:tier_id, :string, default: "free", source: :tierId)
    belongs_to(:last_selected_organization, Shelf.Organization, foreign_key: :last_selected_organization_id, define_field: false, type: :string)
    belongs_to(:tier, Shelf.Tier, foreign_key: :tier_id, define_field: false, type: :string)
    has_one(:contact, Shelf.UserContact, foreign_key: :user_id, references: :id)
    has_one(:business_intel, Shelf.UserBusinessIntel, foreign_key: :user_id, references: :id)
    has_one(:notification_preference, Shelf.NotificationPreference, foreign_key: :user_id, references: :id)
    has_one(:custom_tier_limit, Shelf.CustomTierLimit, foreign_key: :user_id, references: :id)
    has_many(:assets, Shelf.Asset, foreign_key: :user_id, references: :id)
    has_many(:categories, Shelf.Category, foreign_key: :user_id, references: :id)
    has_many(:notes, Shelf.Note, foreign_key: :user_id, references: :id)
    has_many(:booking_notes, Shelf.BookingNote, foreign_key: :user_id, references: :id)
    has_many(:audit_notes, Shelf.AuditNote, foreign_key: :user_id, references: :id)
    has_many(:location_notes, Shelf.LocationNote, foreign_key: :user_id, references: :id)
    has_many(:qr_codes, Shelf.Qr, foreign_key: :user_id, references: :id)
    has_many(:scans, Shelf.Scan, foreign_key: :user_id, references: :id)
    has_many(:tags, Shelf.Tag, foreign_key: :user_id, references: :id)
    has_many(:locations, Shelf.Location, foreign_key: :user_id, references: :id)
    has_many(:images, Shelf.Image, foreign_key: :user_id, references: :id)
    has_many(:organizations, Shelf.Organization, foreign_key: :user_id, references: :id)
    has_many(:custom_fields, Shelf.CustomField, foreign_key: :user_id, references: :id)
    has_many(:sent_invites, Shelf.Invite, foreign_key: :inviter_id, references: :id)
    has_many(:received_invites, Shelf.Invite, foreign_key: :invitee_user_id, references: :id)
    has_many(:team_members, Shelf.TeamMember, foreign_key: :user_id, references: :id)
    has_many(:user_organizations, Shelf.UserOrganization, foreign_key: :user_id, references: :id)
    has_many(:bookings, Shelf.Booking, foreign_key: :creator_id, references: :id)
    has_many(:custodies, Shelf.Booking, foreign_key: :custodian_user_id, references: :id)
    has_many(:approved_bookings, Shelf.Booking, foreign_key: :approved_by_id, references: :id)
    has_many(:created_kits, Shelf.Kit, foreign_key: :created_by_id, references: :id)
    has_many(:asset_reminders, Shelf.AssetReminder, foreign_key: :created_by_id, references: :id)
    has_many(:audits_created, Shelf.AuditSession, foreign_key: :created_by_id, references: :id)
    has_many(:asset_imports, Shelf.AssetImport, foreign_key: :created_by_id, references: :id)
    has_many(:audit_assignments, Shelf.AuditAssignment, foreign_key: :user_id, references: :id)
    has_many(:audit_assets_scanned, Shelf.AuditAsset, foreign_key: :scanned_by_id, references: :id)
    has_many(:audit_scans, Shelf.AuditScan, foreign_key: :scanned_by_id, references: :id)
    has_many(:audit_images_uploaded, Shelf.AuditImage, foreign_key: :uploaded_by_id, references: :id)
    has_many(:asset_index_settings, Shelf.AssetIndexSettings, foreign_key: :user_id, references: :id)
    has_many(:partial_checkins, Shelf.PartialBookingCheckin, foreign_key: :checked_in_by_id, references: :id)
    has_many(:partial_checkouts, Shelf.PartialBookingCheckout, foreign_key: :checked_out_by_id, references: :id)
    has_many(:mobile_auth_codes, Shelf.MobileAuthCode, foreign_key: :user_id, references: :id)
    has_many(:push_tokens, Shelf.PushToken, foreign_key: :user_id, references: :id)
    has_many(:created_updates, Shelf.Update, foreign_key: :created_by_id, references: :id)
    has_many(:update_reads, Shelf.UserUpdateRead, foreign_key: :user_id, references: :id)
    has_many(:asset_filter_presets, Shelf.AssetFilterPreset, foreign_key: :owner_id, references: :id)
    has_many(:asset_models, Shelf.AssetModel, foreign_key: :user_id, references: :id)
    has_many(:consumption_logs, Shelf.ConsumptionLog, foreign_key: :user_id, references: :id)
    has_many(:scim_external_ids, Shelf.UserScimExternalId, foreign_key: :user_id, references: :id)
    has_many(:role_changes_received, Shelf.RoleChangeLog, foreign_key: :user_id, references: :id)
    has_many(:role_changes_performed, Shelf.RoleChangeLog, foreign_key: :changed_by_id, references: :id)
    has_many(:team_member_notes_authored, Shelf.TeamMemberNote, foreign_key: :user_id, references: :id)
    many_to_many(:roles, Shelf.Role, join_through: "_RoleToUser", join_keys: [B: :id, A: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end