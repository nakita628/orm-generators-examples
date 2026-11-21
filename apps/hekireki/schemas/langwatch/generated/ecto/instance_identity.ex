defmodule LangWatch.InstanceIdentity do
  @moduledoc """
  The identity this install presents to LangWatch (ADR-139, section 2).

  Exactly one row, minted on first use and never rewritten. It survives
  restarts, backups and hostname changes, and an install restored into
  another database mints another one, which is what the registry refuses as
  the wrong instance.

  It replaced an identity built from the organization name and id, which
  leaked the customer's name in cleartext, named one organization rather than
  one install, and could not be joined to a license row without splitting a
  string.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          instance_id: String.t(),
          last_report_at: DateTime.t() | nil,
          last_report_error: String.t() | nil,
          optional_metrics_opt_out: boolean(),
          hostname_opt_out: boolean(),
          upgrade_diagnostics_opt_out: boolean(),
          inserted_at: DateTime.t()
        }

  schema "InstanceIdentity" do
    # Fixed, so this table holds one row. A second insert loses the race on the
    # primary key and re-reads the winner's id.
    field(:id, :string, primary_key: true, default: "self")
    # A random UUID. Carries nothing about the customer.
    field(:instance_id, :string, source: :instanceId)
    # When the usage report last reached LangWatch, and why the last attempt
    # did not. Both are what the checkup page reads to say whether reporting
    # works, which nothing recorded before: the sender logged success as soon
    # as `fetch` resolved and never looked at the response.
    field(:last_report_at, LangWatch.PrismaDateTime, source: :lastReportAt)
    field(:last_report_error, :string, source: :lastReportError)
    # The optional category of the usage report, switched off from the checkup
    # page. With it off the report still says which release runs and how many
    # projects the install carries, and says nothing about what anyone does
    # with it (ADR-139, section 10).
    field(:optional_metrics_opt_out, :boolean, default: false, source: :optionalMetricsOptOut)
    # Hostname, which names the customer's own network and so has a switch of
    # its own rather than travelling with its category.
    field(:hostname_opt_out, :boolean, default: false, source: :hostnameOptOut)
    # Sending a failed upgrade step's diagnostics to LangWatch on its own,
    # switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
    field(:upgrade_diagnostics_opt_out, :boolean, default: false, source: :upgradeDiagnosticsOptOut)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end