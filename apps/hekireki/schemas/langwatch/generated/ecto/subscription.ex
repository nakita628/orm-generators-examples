defmodule LangWatch.Subscription do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          plan: atom(),
          stripe_subscription_id: String.t() | nil,
          status: atom(),
          start_date: DateTime.t() | nil,
          end_date: DateTime.t() | nil,
          last_payment_failed_date: DateTime.t() | nil,
          max_members: integer() | nil,
          max_projects: integer() | nil,
          max_messages_per_month: integer() | nil,
          max_workflows: integer() | nil,
          max_retention_days: integer() | nil,
          max_members_lite: integer() | nil,
          max_members_developer: integer() | nil,
          max_teams: integer() | nil,
          max_prompts: integer() | nil,
          max_evaluators: integer() | nil,
          max_scenarios: integer() | nil,
          max_agents: integer() | nil,
          max_experiments: integer() | nil,
          max_online_evaluations: integer() | nil,
          max_datasets: integer() | nil,
          max_dashboards: integer() | nil,
          max_custom_graphs: integer() | nil,
          max_automations: integer() | nil,
          seats_kept: integer() | nil,
          seats_kept_until: DateTime.t() | nil,
          seat_line_quantity: integer() | nil,
          seat_line_quantity_at: DateTime.t() | nil,
          seat_line_write_quantities: [integer()],
          seat_line_write_since: DateTime.t() | nil,
          seat_line_write_at: DateTime.t() | nil,
          seat_line_write_attempts: integer() | nil,
          seat_line_write_flagged_at: DateTime.t() | nil,
          seat_repair_quantity: integer() | nil,
          seat_repair_lifted_quantity: integer() | nil,
          seat_repair_since: DateTime.t() | nil,
          seat_repair_attempts: integer() | nil,
          seat_repair_flagged_at: DateTime.t() | nil,
          seat_line_lease_token: String.t() | nil,
          seat_line_lease_until: DateTime.t() | nil,
          payment_recovery_invoice_id: String.t() | nil,
          payment_recovery_started_at: DateTime.t() | nil,
          payment_recovery_notice_day: integer() | nil,
          over_plan_since: DateTime.t() | nil,
          over_plan_ended_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil,
          invoices: [LangWatch.Invoice.t()]
        }

  schema "Subscription" do
    field(:id, :string, primary_key: true)
    field(:plan, Ecto.Enum, values: [:FREE, :PRO, :GROWTH, :GROWTH_SEAT_EUR_MONTHLY, :GROWTH_SEAT_EUR_ANNUAL, :GROWTH_SEAT_USD_MONTHLY, :GROWTH_SEAT_USD_ANNUAL, :ENTERPRISE, :LAUNCH, :ACCELERATE, :LAUNCH_ANNUAL, :ACCELERATE_ANNUAL])
    field(:stripe_subscription_id, :string, source: :stripeSubscriptionId)
    field(:status, Ecto.Enum, values: [:PENDING, :FAILED, :ACTIVE, :CANCELLED])
    field(:start_date, LangWatch.PrismaDateTime, source: :startDate)
    field(:end_date, LangWatch.PrismaDateTime, source: :endDate)
    field(:last_payment_failed_date, LangWatch.PrismaDateTime, source: :lastPaymentFailedDate)
    field(:max_members, :integer, source: :maxMembers)
    field(:max_projects, :integer, source: :maxProjects)
    field(:max_messages_per_month, :integer, source: :maxMessagesPerMonth)
    field(:max_workflows, :integer, source: :maxWorkflows)
    field(:max_retention_days, :integer, source: :maxRetentionDays)
    field(:max_members_lite, :integer, source: :maxMembersLite)
    # Developer seats the organization pays for; null falls to the plan's own cap.
    field(:max_members_developer, :integer, source: :maxMembersDeveloper)
    field(:max_teams, :integer, source: :maxTeams)
    field(:max_prompts, :integer, source: :maxPrompts)
    field(:max_evaluators, :integer, source: :maxEvaluators)
    field(:max_scenarios, :integer, source: :maxScenarios)
    field(:max_agents, :integer, source: :maxAgents)
    field(:max_experiments, :integer, source: :maxExperiments)
    field(:max_online_evaluations, :integer, source: :maxOnlineEvaluations)
    field(:max_datasets, :integer, source: :maxDatasets)
    field(:max_dashboards, :integer, source: :maxDashboards)
    field(:max_custom_graphs, :integer, source: :maxCustomGraphs)
    field(:max_automations, :integer, source: :maxAutomations)
    # Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil.
    field(:seats_kept, :integer, source: :seatsKept)
    field(:seats_kept_until, LangWatch.PrismaDateTime, source: :seatsKeptUntil)
    # The seat quantity the last finished seat change set the seat line to; provider events never write it.
    field(:seat_line_quantity, :integer, source: :seatLineQuantity)
    field(:seat_line_quantity_at, LangWatch.PrismaDateTime, source: :seatLineQuantityAt)
    # A Stripe seat line write recorded before it is sent and cleared once it is known to have
    # finished under the lease; one left behind is settled by the seat line repair.
    # Every seat quantity written while this record was pending.
    field(:seat_line_write_quantities, {:array, :integer}, source: :seatLineWriteQuantities)
    # The first write still pending; a later holder's write keeps it.
    field(:seat_line_write_since, LangWatch.PrismaDateTime, source: :seatLineWriteSince)
    # The latest write sent; the repair waits a grace window after it.
    field(:seat_line_write_at, LangWatch.PrismaDateTime, source: :seatLineWriteAt)
    field(:seat_line_write_attempts, :integer, source: :seatLineWriteAttempts)
    field(:seat_line_write_flagged_at, LangWatch.PrismaDateTime, source: :seatLineWriteFlaggedAt)
    # The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled.
    field(:seat_repair_quantity, :integer, source: :seatRepairQuantity)
    # The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it.
    field(:seat_repair_lifted_quantity, :integer, source: :seatRepairLiftedQuantity)
    field(:seat_repair_since, LangWatch.PrismaDateTime, source: :seatRepairSince)
    field(:seat_repair_attempts, :integer, source: :seatRepairAttempts)
    # Set once the seat line repair has failed often enough that Ops should look at it.
    field(:seat_repair_flagged_at, LangWatch.PrismaDateTime, source: :seatRepairFlaggedAt)
    # The lease a seat change or seat line repair holds while it reads and moves the seat line.
    field(:seat_line_lease_token, :string, source: :seatLineLeaseToken)
    field(:seat_line_lease_until, LangWatch.PrismaDateTime, source: :seatLineLeaseUntil)
    # An open failed payment recovery case: the unpaid renewal invoice, day 0 (its first failed
    # charge) and the last notice sent (3, 7, 14 or 21). Cleared when the case closes.
    field(:payment_recovery_invoice_id, :string, source: :paymentRecoveryInvoiceId)
    field(:payment_recovery_started_at, LangWatch.PrismaDateTime, source: :paymentRecoveryStartedAt)
    field(:payment_recovery_notice_day, :integer, source: :paymentRecoveryNoticeDay)
    # The over-plan mark: traces recorded from overPlanSince are withheld while it is open and the
    # organization is on Free with more members than Free allows; ended once that no longer holds.
    field(:over_plan_since, LangWatch.PrismaDateTime, source: :overPlanSince)
    field(:over_plan_ended_at, LangWatch.PrismaDateTime, source: :overPlanEndedAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:invoices, LangWatch.Invoice, foreign_key: :subscription_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end