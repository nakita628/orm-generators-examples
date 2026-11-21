class Organization < ApplicationRecord
  self.table_name = "Organization"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :name, default: "Personal"
  attribute :enabledSso, default: false
  attribute :selfServiceCanSeeCustody, default: false
  attribute :selfServiceCanSeeBookings, default: false
  attribute :baseUserCanSeeCustody, default: false
  attribute :baseUserCanSeeBookings, default: false
  attribute :managersCanEditAssets, default: false
  attribute :barcodesEnabled, default: false
  attribute :usedBarcodeTrial, default: false
  attribute :auditsEnabled, default: false
  attribute :usedAuditTrial, default: false
  attribute :workspaceDisabled, default: false
  attribute :hasSequentialIdsMigrated, default: false
  attribute :showShelfBranding, default: true
  attribute :showQrCodesOnPdfs, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { personal: "PERSONAL", team: "TEAM" }, default: :personal, validate: true
  enum :currency, { aed: "AED", afn: "AFN", all: "ALL", amd: "AMD", ang: "ANG", aoa: "AOA", ars: "ARS", aud: "AUD", awg: "AWG", azn: "AZN", bam: "BAM", bbd: "BBD", bdt: "BDT", bgn: "BGN", bhd: "BHD", bif: "BIF", bmd: "BMD", bnd: "BND", bob: "BOB", brl: "BRL", bsd: "BSD", btn: "BTN", bwp: "BWP", byn: "BYN", bzd: "BZD", cad: "CAD", cdf: "CDF", chf: "CHF", clp: "CLP", cny: "CNY", cop: "COP", crc: "CRC", cup: "CUP", cve: "CVE", czk: "CZK", djf: "DJF", dkk: "DKK", dop: "DOP", dzd: "DZD", egp: "EGP", ern: "ERN", etb: "ETB", eur: "EUR", fjd: "FJD", fkp: "FKP", gbp: "GBP", gel: "GEL", ghs: "GHS", gip: "GIP", gmd: "GMD", gnf: "GNF", gtq: "GTQ", gyd: "GYD", hkd: "HKD", hnl: "HNL", htg: "HTG", huf: "HUF", idr: "IDR", ils: "ILS", inr: "INR", iqd: "IQD", irr: "IRR", isk: "ISK", jmd: "JMD", jod: "JOD", jpy: "JPY", kes: "KES", kgs: "KGS", khr: "KHR", kmf: "KMF", kpw: "KPW", krw: "KRW", kwd: "KWD", kyd: "KYD", kzt: "KZT", lak: "LAK", lbp: "LBP", lkr: "LKR", lrd: "LRD", lsl: "LSL", lyd: "LYD", mad: "MAD", mdl: "MDL", mga: "MGA", mkd: "MKD", mmk: "MMK", mnt: "MNT", mop: "MOP", mru: "MRU", mur: "MUR", mvr: "MVR", mwk: "MWK", mxn: "MXN", myr: "MYR", mzn: "MZN", nad: "NAD", ngn: "NGN", nio: "NIO", nok: "NOK", npr: "NPR", nzd: "NZD", omr: "OMR", pab: "PAB", pen: "PEN", pgk: "PGK", php: "PHP", pkr: "PKR", pln: "PLN", pyg: "PYG", qar: "QAR", ron: "RON", rsd: "RSD", rub: "RUB", rwf: "RWF", sar: "SAR", sbd: "SBD", scr: "SCR", sdg: "SDG", sek: "SEK", sgd: "SGD", shp: "SHP", sle: "SLE", sos: "SOS", srd: "SRD", ssp: "SSP", stn: "STN", svc: "SVC", syp: "SYP", szl: "SZL", thb: "THB", tjs: "TJS", tmt: "TMT", tnd: "TND", top: "TOP", try: "TRY", ttd: "TTD", twd: "TWD", tzs: "TZS", uah: "UAH", ugx: "UGX", usd: "USD", uyu: "UYU", uzs: "UZS", ves: "VES", vnd: "VND", vuv: "VUV", wst: "WST", xaf: "XAF", xcd: "XCD", xof: "XOF", xpf: "XPF", yer: "YER", zar: "ZAR", zmw: "ZMW", zwl: "ZWL" },
       default: :usd, prefix: true, validate: true
  enum :qrIdDisplayPreference, { qr_id: "QR_ID", sam_id: "SAM_ID", code128: "Code128", code39: "Code39", data_matrix: "DataMatrix", external_qr: "ExternalQR", ean13: "EAN13" },
       default: :qr_id, validate: true

  validates :imageId, uniqueness: { allow_nil: true }
  validates :customEmailFooter, length: { maximum: 500 }

  belongs_to :owner, class_name: "User", foreign_key: "userId", inverse_of: :organizations
  belongs_to :image, foreign_key: "imageId", optional: true, inverse_of: :organization
  belongs_to :sso_details, foreign_key: "ssoDetailsId", optional: true, inverse_of: :organizations
  has_one :working_hours, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_one :booking_settings, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :members, class_name: "TeamMember", foreign_key: "organizationId", inverse_of: :organization,
                     dependent: :destroy
  has_many :assets, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :asset_kits, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :asset_locations, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :scans, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :locations, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :categories, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :tags, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :qr_codes, class_name: "Qr", foreign_key: "organizationId", inverse_of: :organization, dependent: :nullify
  has_many :barcodes, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :invites, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :user_organizations, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :custom_fields, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :images, foreign_key: "ownerOrgId", inverse_of: :owner_org, dependent: :destroy
  has_many :bookings, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :kits, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :asset_models, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :asset_index_settings, class_name: "AssetIndexSettings", foreign_key: "organizationId",
                                  inverse_of: :organization, dependent: :destroy
  has_many :asset_reminders, foreign_key: "organizationId", inverse_of: :organization, dependent: :restrict_with_error
  has_many :asset_filter_presets, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :audit_sessions, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :audit_images, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :asset_imports, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :link_previews, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :pdf_templates, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :users_with_last_selected, class_name: "User", foreign_key: "lastSelectedOrganizationId",
                                      inverse_of: :last_selected_organization, dependent: :nullify
  has_many :scim_tokens, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :scim_external_ids, class_name: "UserScimExternalId", foreign_key: "organizationId",
                               inverse_of: :organization, dependent: :destroy
  has_many :role_change_logs, foreign_key: "organizationId", inverse_of: :organization, dependent: :restrict_with_error
  has_many :team_member_notes, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :activity_events, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
end
