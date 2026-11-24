defmodule LangWatch.SelfHostedInstance do
  @moduledoc """
  One self-hosted install, as its daily usage report describes it (ADR-139,
  section 10). The row is created by the first report that names an instance
  id and is updated by every report after that, so the backoffice reads one
  row per install rather than folding a history on every page load.

  The report carries no credential, so it identifies an install and never a
  customer. `organizationId` and `issuedLicenseId` are resolved from the
  license bound to this instance, which the install proved by presenting its
  token to license sync. A field in the posted body never reaches them.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          instance_id: String.t(),
          first_seen_at: DateTime.t(),
          last_seen_at: DateTime.t(),
          version: String.t() | nil,
          install_method: String.t() | nil,
          chart_version: String.t() | nil,
          hostname: String.t() | nil,
          environment: String.t() | nil,
          installed_at: DateTime.t() | nil,
          report_schema_version: integer() | nil,
          organization_id: String.t() | nil,
          issued_license_id: String.t() | nil,
          user_email_domains: map() | list() | String.t() | number() | boolean() | nil,
          user_domains: [String.t()],
          latest_report: map() | list() | String.t() | number() | boolean() | nil,
          optional_metrics_reported: boolean(),
          hostname_reported: boolean(),
          report_count: integer(),
          last_unknown_fields: integer(),
          raised_signals: [String.t()],
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "SelfHostedInstance" do
    field(:id, :string, primary_key: true)
    # The UUID the install minted for itself. The join key for everything else.
    field(:instance_id, :string, source: :instanceId)
    field(:first_seen_at, LangWatch.PrismaDateTime, autogenerate: true, source: :firstSeenAt)
    field(:last_seen_at, LangWatch.PrismaDateTime, source: :lastSeenAt)
    # What the install says about itself.
    field(:version, :string)
    field(:install_method, :string, source: :installMethod)
    field(:chart_version, :string, source: :chartVersion)
    # The host it answers on, when the operator left hostname reporting on.
    field(:hostname, :string)
    field(:environment, :string)
    # The day the install itself says it first came up, which predates the
    # first report on an install that switched reporting on later.
    field(:installed_at, LangWatch.PrismaDateTime, source: :installedAt)
    field(:report_schema_version, :integer, source: :reportSchemaVersion)
    # The customer, and the license that named them. Plain column refs rather
    # than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
    # carries no foreign keys, and the receiver resolves both from the license
    # registry rather than from the report.
    field(:organization_id, :string, source: :organizationId)
    field(:issued_license_id, :string, source: :issuedLicenseId)
    # Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
    # install that switched the optional category off.
    field(:user_email_domains, LangWatch.PrismaJson, source: :userEmailDomains)
    # The same domains as the keys of `userEmailDomains`, flattened so an
    # operator can search for one. A JSON object's keys are the domains
    # themselves, and no query language addresses a key it does not already
    # know the name of.
    field(:user_domains, {:array, :string}, source: :userDomains)
    # What the last report carried, so the list and the drawer read one row:
    # the onboarding ladder, the usage counts and the operational block, under
    # the same keys the dictionary declares.
    field(:latest_report, LangWatch.PrismaJson, source: :latestReport)
    # What the last report said the operator had switched off, read from which
    # categories arrived rather than from a field the install could assert.
    field(:optional_metrics_reported, :boolean, default: true, source: :optionalMetricsReported)
    field(:hostname_reported, :boolean, default: true, source: :hostnameReported)
    # How many reports have landed, and how many fields the last one carried
    # that this release has no name for. A count above zero says the senders
    # are ahead of the receiver.
    field(:report_count, :integer, default: 0, source: :reportCount)
    field(:last_unknown_fields, :integer, default: 0, source: :lastUnknownFields)
    # The lead signals already raised for this install. A signal on this list
    # is never raised again, which is what keeps a large install from
    # reporting that it is large every morning.
    field(:raised_signals, {:array, :string}, source: :raisedSignals)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end