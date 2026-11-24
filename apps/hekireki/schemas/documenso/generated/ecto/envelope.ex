defmodule Documenso.Envelope do
  @moduledoc """
  @zod.import(["import { ZDocumentAuthOptionsSchema } from '@documenso/lib/types/document-auth';", "import { ZDocumentFormValuesSchema } from '@documenso/lib/types/document-form-values';", "import { ZSignatureLevelSchema } from '@documenso/lib/types/signature-level';"])
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Documenso.PrismaDateTime, autogenerate: {Documenso.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          secondary_id: String.t(),
          external_id: String.t() | nil,
          type: atom(),
          completed_at: DateTime.t() | nil,
          deleted_at: DateTime.t() | nil,
          title: String.t(),
          status: atom(),
          source: atom(),
          qr_token: String.t() | nil,
          signature_level: String.t(),
          internal_version: integer(),
          use_legacy_field_insertion: boolean(),
          auth_options: map() | list() | String.t() | number() | boolean() | nil,
          form_values: map() | list() | String.t() | number() | boolean() | nil,
          visibility: atom(),
          template_type: atom(),
          public_title: String.t(),
          public_description: String.t(),
          template_id: integer() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Documenso.User.t() | nil,
          team: Documenso.Team.t() | nil,
          folder: Documenso.Folder.t() | nil,
          document_meta: Documenso.DocumentMeta.t() | nil,
          direct_link: Documenso.TemplateDirectLink.t() | nil,
          envelope_items: [Documenso.EnvelopeItem.t()],
          recipients: [Documenso.Recipient.t()],
          fields: [Documenso.Field.t()],
          contents: [Documenso.EnvelopeContent.t()],
          share_links: [Documenso.DocumentShareLink.t()],
          audit_logs: [Documenso.DocumentAuditLog.t()],
          envelope_attachments: [Documenso.EnvelopeAttachment.t()]
        }

  schema "Envelope" do
    field(:id, :string, primary_key: true)
    field(:secondary_id, :string, source: :secondaryId)
    # @zod.string.describe("A custom external ID you can use to identify the document.")
    field(:external_id, :string, source: :externalId)
    field(:type, Ecto.Enum, values: [:DOCUMENT, :TEMPLATE])
    field(:completed_at, Documenso.PrismaDateTime, source: :completedAt)
    field(:deleted_at, Documenso.PrismaDateTime, source: :deletedAt)
    field(:title, :string)
    field(:status, Ecto.Enum, values: [:DRAFT, :PENDING, :COMPLETED, :REJECTED, :CANCELLED], default: :DRAFT)
    field(:source, Ecto.Enum, values: [:DOCUMENT, :TEMPLATE, :TEMPLATE_DIRECT_LINK])
    # @zod.string.describe("The token for viewing the document using the QR code on the certificate.")
    field(:qr_token, :string, source: :qrToken)
    # [SignatureLevel] @zod.custom.use(ZSignatureLevelSchema)
    field(:signature_level, :string, source: :signatureLevel)
    field(:internal_version, :integer, source: :internalVersion)
    field(:use_legacy_field_insertion, :boolean, default: false, source: :useLegacyFieldInsertion)
    # [DocumentAuthOptions] @zod.custom.use(ZDocumentAuthOptionsSchema)
    field(:auth_options, Documenso.PrismaJson, source: :authOptions)
    # [DocumentFormValues] @zod.custom.use(ZDocumentFormValuesSchema)
    field(:form_values, Documenso.PrismaJson, source: :formValues)
    field(:visibility, Ecto.Enum, values: [:EVERYONE, :MANAGER_AND_ABOVE, :ADMIN], default: :EVERYONE)
    field(:template_type, Ecto.Enum, values: [:PUBLIC, :PRIVATE, :ORGANISATION], default: :PRIVATE, source: :templateType)
    field(:public_title, :string, default: "", source: :publicTitle)
    field(:public_description, :string, default: "", source: :publicDescription)
    field(:template_id, :integer, source: :templateId)
    # @zod.number.describe("The ID of the user that created this document.")
    field(:user_id, :id, source: :userId)
    field(:team_id, :id, source: :teamId)
    field(:folder_id, :string, source: :folderId)
    field(:document_meta_id, :string, source: :documentMetaId)
    belongs_to(:user, Documenso.User, foreign_key: :user_id, define_field: false)
    belongs_to(:team, Documenso.Team, foreign_key: :team_id, define_field: false)
    belongs_to(:folder, Documenso.Folder, foreign_key: :folder_id, define_field: false, type: :string)
    belongs_to(:document_meta, Documenso.DocumentMeta, foreign_key: :document_meta_id, define_field: false, type: :string)
    has_one(:direct_link, Documenso.TemplateDirectLink, foreign_key: :envelope_id, references: :id)
    has_many(:envelope_items, Documenso.EnvelopeItem, foreign_key: :envelope_id, references: :id)
    has_many(:recipients, Documenso.Recipient, foreign_key: :envelope_id, references: :id)
    has_many(:fields, Documenso.Field, foreign_key: :envelope_id, references: :id)
    has_many(:contents, Documenso.EnvelopeContent, foreign_key: :envelope_id, references: :id)
    has_many(:share_links, Documenso.DocumentShareLink, foreign_key: :envelope_id, references: :id)
    has_many(:audit_logs, Documenso.DocumentAuditLog, foreign_key: :envelope_id, references: :id)
    has_many(:envelope_attachments, Documenso.EnvelopeAttachment, foreign_key: :envelope_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end