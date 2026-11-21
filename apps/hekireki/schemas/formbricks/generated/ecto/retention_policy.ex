defmodule Formbricks.RetentionPolicy do
  @moduledoc """
  One policy per organisation and kind of data: a notice `warnDays` before the action, and the action
  `periodDays` after the target's clock. The action follows from the kind of data: responses are
  deleted, surveys archived (the archive purge deletes them 30 days later, as it does every archived
  survey), members deactivated. The dates themselves are never stored; they are computed by the shared
  helper in `apps/web/modules/ee/data-retention/lib/schedule.ts`, so the UI, the emails and the sweep
  always agree.

  The allowed ranges are validated by the API. The migration repeats the ones that guard against
  acting early or re-notifying (a notice of at least 14 days and shorter than the period, an enabled
  policy always has `enabledAt`) and the structural ones (conditions on the surveys policy only) as
  CHECK constraints, which Prisma cannot express.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          entity: atom(),
          enabled: boolean(),
          enabled_at: DateTime.t() | nil,
          warn_days: integer(),
          period_days: integer(),
          conditions: atom(),
          deferred_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Formbricks.Organization.t() | nil,
          updated_by: Formbricks.User.t() | nil
        }

  schema "RetentionPolicy" do
    field(:id, :string, primary_key: true)
    field(:entity, Ecto.Enum, values: [:responses, :surveys, :members])
    field(:enabled, :boolean, default: false)
    # When the current configuration took effect: the policy was switched on or unpaused, or changed
    # in a way that could bring an action earlier (a shorter period or warning, different
    # conditions). Lengthening does not move it. The warning always runs in full from here, so no
    # change can delete data on the night it is made (ENG-3614). Also the clock for members with no
    # recorded sign-in.
    field(:enabled_at, Formbricks.PrismaDateTime, source: :enabledAt)
    # How long before the action the notice goes out. The action never runs before the notice has
    # been out for this long, so the warning always runs in full.
    field(:warn_days, :integer, default: 60, source: :warnDays)
    # The period, counted from the target's clock, after which the policy acts: responses are
    # deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
    field(:period_days, :integer, source: :periodDays)
    # Surveys policy only.
    field(:conditions, {:array, Ecto.Enum}, values: [:noResponse, :noChange, :createdBefore])
    # The last night the sweep's time ran out before it reached this policy's organisation, licensed.
    # The organisation was due and the sweep was running, so it counts as reached for the gap that
    # restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
    field(:deferred_at, Formbricks.PrismaDateTime, source: :deferredAt)
    field(:organization_id, :string, source: :organizationId)
    field(:updated_by_id, :string, source: :updatedById)
    belongs_to(:organization, Formbricks.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:updated_by, Formbricks.User, foreign_key: :updated_by_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end