defmodule Documenso.AdminTag do
  @moduledoc """
  @zod.import(["import { ZAdminTagTypeSchema } from '@documenso/lib/types/admin-tags';"])
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Documenso.PrismaDateTime, autogenerate: {Documenso.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          type: String.t(),
          description: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Documenso.User.t() | nil,
          organisation: Documenso.Organisation.t() | nil,
          created_by_user: Documenso.User.t() | nil,
          admin_rule: Documenso.AdminRule.t() | nil
        }

  schema "AdminTag" do
    field(:id, :string, primary_key: true)
    # [AdminTagType] @zod.custom.use(ZAdminTagTypeSchema)
    field(:type, :string)
    field(:description, :string)
    field(:user_id, :id, source: :userId)
    field(:organisation_id, :string, source: :organisationId)
    field(:created_by_user_id, :id, source: :createdByUserId)
    field(:admin_rule_id, :string, source: :adminRuleId)
    belongs_to(:user, Documenso.User, foreign_key: :user_id, define_field: false)
    belongs_to(:organisation, Documenso.Organisation, foreign_key: :organisation_id, define_field: false, type: :string)
    belongs_to(:created_by_user, Documenso.User, foreign_key: :created_by_user_id, define_field: false)
    belongs_to(:admin_rule, Documenso.AdminRule, foreign_key: :admin_rule_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end