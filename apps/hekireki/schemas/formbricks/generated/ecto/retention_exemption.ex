defmodule Formbricks.RetentionExemption do
  @moduledoc """
  Holds one survey back from one policy until `until`. Keyed by organisation + entity rather than a
  policy id, so removing and re-adding a policy doesn't orphan its exemptions.

  Active means `revokedAt IS NULL AND until > now()`, so expiry needs no job. "One active exemption
  per survey and policy" is the partial unique index `RetentionExemption_surveyId_entity_active_key`
  (`WHERE "revokedAt" IS NULL`) in
  `migration/20261007120000_add_data_retention_tables/migration.sql`; Prisma's `@@unique` has no
  `where`, so it is deliberately not declared here (see "Indexes Prisma cannot express" in the
  README). An expired row still holds that slot until it is revoked, so creating a new exemption
  must close an expired one in the same transaction.

  `surveyId` is nullable and a CHECK requires exactly one target, so a later `workflowId` is one more
  column in that CHECK. The survey's organisation can't be tied to `organizationId` in the database
  (a survey carries its workspace, not its organisation): writers derive it from the survey.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          entity: atom(),
          until: DateTime.t(),
          source: atom(),
          reason: String.t() | nil,
          revoked_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          organization: Formbricks.Organization.t() | nil,
          survey: Formbricks.Survey.t() | nil,
          created_by: Formbricks.User.t() | nil,
          revoked_by: Formbricks.User.t() | nil
        }

  schema "RetentionExemption" do
    field(:id, :string, primary_key: true)
    field(:entity, Ecto.Enum, values: [:responses, :surveys, :members])
    field(:until, Formbricks.PrismaDateTime)
    field(:source, Ecto.Enum, values: [:manual, :restore], default: :manual)
    # Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`).
    field(:reason, :string)
    field(:revoked_at, Formbricks.PrismaDateTime, source: :revokedAt)
    field(:organization_id, :string, source: :organizationId)
    field(:survey_id, :string, source: :surveyId)
    field(:created_by_id, :string, source: :createdById)
    field(:revoked_by_id, :string, source: :revokedById)
    belongs_to(:organization, Formbricks.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:survey, Formbricks.Survey, foreign_key: :survey_id, define_field: false, type: :string)
    belongs_to(:created_by, Formbricks.User, foreign_key: :created_by_id, define_field: false, type: :string)
    belongs_to(:revoked_by, Formbricks.User, foreign_key: :revoked_by_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end