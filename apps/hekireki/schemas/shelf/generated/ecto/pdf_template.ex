defmodule Shelf.PdfTemplate do
  @moduledoc """
  A saved layout for a generated PDF, kept per team workspace under a name:
  which columns, rows and header lines print, and in what order the rows
  come. A row is a layout, not a generated file; the person generating a PDF
  picks a template by name.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          kind: atom(),
          config: map() | list() | String.t() | number() | boolean(),
          is_default: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Shelf.Organization.t() | nil
        }

  schema "PdfTemplate" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:kind, Ecto.Enum, values: [:BOOKING], default: :BOOKING)
    # The layout's switches. Read through `parseBookingTemplateConfig`,
    # which fills any key a stored row lacks from the Checklist template's
    # values.
    field(:config, Shelf.PrismaJson, default: %{})
    # The template the PDF dialog offers first. One per workspace and kind,
    # enforced by a partial unique index in raw SQL (migration
    # 20261103120000_add_pdf_templates). Names are also unique per workspace
    # and kind ignoring case, through an expression index in the same
    # migration.
    field(:is_default, :boolean, default: false, source: :isDefault)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end