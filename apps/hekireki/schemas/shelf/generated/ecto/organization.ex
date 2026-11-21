defmodule Shelf.Organization do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          type: atom(),
          currency: atom(),
          enabled_sso: boolean(),
          self_service_can_see_custody: boolean(),
          self_service_can_see_bookings: boolean(),
          base_user_can_see_custody: boolean(),
          base_user_can_see_bookings: boolean(),
          managers_can_edit_assets: boolean(),
          barcodes_enabled: boolean(),
          barcodes_enabled_at: DateTime.t() | nil,
          used_barcode_trial: boolean(),
          audits_enabled: boolean(),
          audits_enabled_at: DateTime.t() | nil,
          used_audit_trial: boolean(),
          workspace_disabled: boolean(),
          has_sequential_ids_migrated: boolean(),
          qr_id_display_preference: atom(),
          show_shelf_branding: boolean(),
          show_qr_codes_on_pdfs: boolean(),
          custom_email_footer: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          owner: Shelf.User.t() | nil,
          image: Shelf.Image.t() | nil,
          sso_details: Shelf.SsoDetails.t() | nil,
          working_hours: Shelf.WorkingHours.t() | nil,
          booking_settings: Shelf.BookingSettings.t() | nil,
          members: [Shelf.TeamMember.t()],
          assets: [Shelf.Asset.t()],
          asset_kits: [Shelf.AssetKit.t()],
          asset_locations: [Shelf.AssetLocation.t()],
          scans: [Shelf.Scan.t()],
          locations: [Shelf.Location.t()],
          categories: [Shelf.Category.t()],
          tags: [Shelf.Tag.t()],
          qr_codes: [Shelf.Qr.t()],
          barcodes: [Shelf.Barcode.t()],
          invites: [Shelf.Invite.t()],
          user_organizations: [Shelf.UserOrganization.t()],
          custom_fields: [Shelf.CustomField.t()],
          images: [Shelf.Image.t()],
          bookings: [Shelf.Booking.t()],
          kits: [Shelf.Kit.t()],
          asset_models: [Shelf.AssetModel.t()],
          asset_index_settings: [Shelf.AssetIndexSettings.t()],
          asset_reminders: [Shelf.AssetReminder.t()],
          asset_filter_presets: [Shelf.AssetFilterPreset.t()],
          audit_sessions: [Shelf.AuditSession.t()],
          audit_images: [Shelf.AuditImage.t()],
          asset_imports: [Shelf.AssetImport.t()],
          link_previews: [Shelf.LinkPreview.t()],
          pdf_templates: [Shelf.PdfTemplate.t()],
          users_with_last_selected: [Shelf.User.t()],
          scim_tokens: [Shelf.ScimToken.t()],
          scim_external_ids: [Shelf.UserScimExternalId.t()],
          role_change_logs: [Shelf.RoleChangeLog.t()],
          team_member_notes: [Shelf.TeamMemberNote.t()],
          activity_events: [Shelf.ActivityEvent.t()]
        }

  schema "Organization" do
    field(:id, :string, primary_key: true)
    field(:name, :string, default: "Personal")
    field(:type, Ecto.Enum, values: [:PERSONAL, :TEAM], default: :PERSONAL)
    field(:currency, Ecto.Enum, values: [:AED, :AFN, :ALL, :AMD, :ANG, :AOA, :ARS, :AUD, :AWG, :AZN, :BAM, :BBD, :BDT, :BGN, :BHD, :BIF, :BMD, :BND, :BOB, :BRL, :BSD, :BTN, :BWP, :BYN, :BZD, :CAD, :CDF, :CHF, :CLP, :CNY, :COP, :CRC, :CUP, :CVE, :CZK, :DJF, :DKK, :DOP, :DZD, :EGP, :ERN, :ETB, :EUR, :FJD, :FKP, :GBP, :GEL, :GHS, :GIP, :GMD, :GNF, :GTQ, :GYD, :HKD, :HNL, :HTG, :HUF, :IDR, :ILS, :INR, :IQD, :IRR, :ISK, :JMD, :JOD, :JPY, :KES, :KGS, :KHR, :KMF, :KPW, :KRW, :KWD, :KYD, :KZT, :LAK, :LBP, :LKR, :LRD, :LSL, :LYD, :MAD, :MDL, :MGA, :MKD, :MMK, :MNT, :MOP, :MRU, :MUR, :MVR, :MWK, :MXN, :MYR, :MZN, :NAD, :NGN, :NIO, :NOK, :NPR, :NZD, :OMR, :PAB, :PEN, :PGK, :PHP, :PKR, :PLN, :PYG, :QAR, :RON, :RSD, :RUB, :RWF, :SAR, :SBD, :SCR, :SDG, :SEK, :SGD, :SHP, :SLE, :SOS, :SRD, :SSP, :STN, :SVC, :SYP, :SZL, :THB, :TJS, :TMT, :TND, :TOP, :TRY, :TTD, :TWD, :TZS, :UAH, :UGX, :USD, :UYU, :UZS, :VES, :VND, :VUV, :WST, :XAF, :XCD, :XOF, :XPF, :YER, :ZAR, :ZMW, :ZWL], default: :USD)
    field(:enabled_sso, :boolean, default: false, source: :enabledSso)
    field(:self_service_can_see_custody, :boolean, default: false, source: :selfServiceCanSeeCustody)
    field(:self_service_can_see_bookings, :boolean, default: false, source: :selfServiceCanSeeBookings)
    field(:base_user_can_see_custody, :boolean, default: false, source: :baseUserCanSeeCustody)
    field(:base_user_can_see_bookings, :boolean, default: false, source: :baseUserCanSeeBookings)
    field(:managers_can_edit_assets, :boolean, default: false, source: :managersCanEditAssets)
    field(:barcodes_enabled, :boolean, default: false, source: :barcodesEnabled)
    field(:barcodes_enabled_at, Shelf.PrismaDateTime, source: :barcodesEnabledAt)
    field(:used_barcode_trial, :boolean, default: false, source: :usedBarcodeTrial)
    field(:audits_enabled, :boolean, default: false, source: :auditsEnabled)
    field(:audits_enabled_at, Shelf.PrismaDateTime, source: :auditsEnabledAt)
    field(:used_audit_trial, :boolean, default: false, source: :usedAuditTrial)
    field(:workspace_disabled, :boolean, default: false, source: :workspaceDisabled)
    field(:has_sequential_ids_migrated, :boolean, default: false, source: :hasSequentialIdsMigrated)
    field(:qr_id_display_preference, Ecto.Enum, values: [:QR_ID, :SAM_ID, :Code128, :Code39, :DataMatrix, :ExternalQR, :EAN13], default: :QR_ID, source: :qrIdDisplayPreference)
    field(:show_shelf_branding, :boolean, default: true, source: :showShelfBranding)
    field(:show_qr_codes_on_pdfs, :boolean, default: true, source: :showQrCodesOnPdfs)
    field(:custom_email_footer, :string, source: :customEmailFooter)
    field(:user_id, :string, source: :userId)
    field(:image_id, :string, source: :imageId)
    field(:sso_details_id, :string, source: :ssoDetailsId)
    belongs_to(:owner, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:image, Shelf.Image, foreign_key: :image_id, define_field: false, type: :string)
    belongs_to(:sso_details, Shelf.SsoDetails, foreign_key: :sso_details_id, define_field: false, type: :string)
    has_one(:working_hours, Shelf.WorkingHours, foreign_key: :organization_id, references: :id)
    has_one(:booking_settings, Shelf.BookingSettings, foreign_key: :organization_id, references: :id)
    has_many(:members, Shelf.TeamMember, foreign_key: :organization_id, references: :id)
    has_many(:assets, Shelf.Asset, foreign_key: :organization_id, references: :id)
    has_many(:asset_kits, Shelf.AssetKit, foreign_key: :organization_id, references: :id)
    has_many(:asset_locations, Shelf.AssetLocation, foreign_key: :organization_id, references: :id)
    has_many(:scans, Shelf.Scan, foreign_key: :organization_id, references: :id)
    has_many(:locations, Shelf.Location, foreign_key: :organization_id, references: :id)
    has_many(:categories, Shelf.Category, foreign_key: :organization_id, references: :id)
    has_many(:tags, Shelf.Tag, foreign_key: :organization_id, references: :id)
    has_many(:qr_codes, Shelf.Qr, foreign_key: :organization_id, references: :id)
    has_many(:barcodes, Shelf.Barcode, foreign_key: :organization_id, references: :id)
    has_many(:invites, Shelf.Invite, foreign_key: :organization_id, references: :id)
    has_many(:user_organizations, Shelf.UserOrganization, foreign_key: :organization_id, references: :id)
    has_many(:custom_fields, Shelf.CustomField, foreign_key: :organization_id, references: :id)
    has_many(:images, Shelf.Image, foreign_key: :owner_org_id, references: :id)
    has_many(:bookings, Shelf.Booking, foreign_key: :organization_id, references: :id)
    has_many(:kits, Shelf.Kit, foreign_key: :organization_id, references: :id)
    has_many(:asset_models, Shelf.AssetModel, foreign_key: :organization_id, references: :id)
    has_many(:asset_index_settings, Shelf.AssetIndexSettings, foreign_key: :organization_id, references: :id)
    has_many(:asset_reminders, Shelf.AssetReminder, foreign_key: :organization_id, references: :id)
    has_many(:asset_filter_presets, Shelf.AssetFilterPreset, foreign_key: :organization_id, references: :id)
    has_many(:audit_sessions, Shelf.AuditSession, foreign_key: :organization_id, references: :id)
    has_many(:audit_images, Shelf.AuditImage, foreign_key: :organization_id, references: :id)
    has_many(:asset_imports, Shelf.AssetImport, foreign_key: :organization_id, references: :id)
    has_many(:link_previews, Shelf.LinkPreview, foreign_key: :organization_id, references: :id)
    has_many(:pdf_templates, Shelf.PdfTemplate, foreign_key: :organization_id, references: :id)
    has_many(:users_with_last_selected, Shelf.User, foreign_key: :last_selected_organization_id, references: :id)
    has_many(:scim_tokens, Shelf.ScimToken, foreign_key: :organization_id, references: :id)
    has_many(:scim_external_ids, Shelf.UserScimExternalId, foreign_key: :organization_id, references: :id)
    has_many(:role_change_logs, Shelf.RoleChangeLog, foreign_key: :organization_id, references: :id)
    has_many(:team_member_notes, Shelf.TeamMemberNote, foreign_key: :organization_id, references: :id)
    has_many(:activity_events, Shelf.ActivityEvent, foreign_key: :organization_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end