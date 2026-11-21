defmodule Shelf.AuditImage do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          image_url: String.t(),
          thumbnail_url: String.t() | nil,
          description: String.t() | nil,
          captured_at: DateTime.t() | nil,
          capture_source: atom() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          audit_session: Shelf.AuditSession.t() | nil,
          audit_asset: Shelf.AuditAsset.t() | nil,
          uploaded_by: Shelf.User.t() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "AuditImage" do
    field(:id, :string, primary_key: true)
    field(:image_url, :string, source: :imageUrl)
    field(:thumbnail_url, :string, source: :thumbnailUrl)
    field(:description, :string)
    # When the photo was taken, from its EXIF data (or the phone app's copy of
    # it). Null when the photo carried no usable date. EXIF can be edited, so
    # this is the best signal available, not proof.
    field(:captured_at, Shelf.PrismaDateTime, source: :capturedAt)
    # How the photo reached Shelf. Null only for photos uploaded before this
    # field existed; a phone upload with no or an unknown source is PHONE.
    field(:capture_source, Ecto.Enum, values: [:CAMERA, :LIBRARY, :WEB_UPLOAD, :PHONE], source: :captureSource)
    field(:audit_session_id, :string, source: :auditSessionId)
    field(:audit_asset_id, :string, source: :auditAssetId)
    field(:uploaded_by_id, :string, source: :uploadedById)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:audit_session, Shelf.AuditSession, foreign_key: :audit_session_id, define_field: false, type: :string)
    belongs_to(:audit_asset, Shelf.AuditAsset, foreign_key: :audit_asset_id, define_field: false, type: :string)
    belongs_to(:uploaded_by, Shelf.User, foreign_key: :uploaded_by_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end