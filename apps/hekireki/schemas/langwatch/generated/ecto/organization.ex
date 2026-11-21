defmodule LangWatch.Organization do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          kind: String.t(),
          phone_number: String.t() | nil,
          slug: String.t(),
          usage_spending_max_limit: integer() | nil,
          dataset_attachment_max_mb: integer() | nil,
          private_click_house_label: String.t() | nil,
          private_s3_label: String.t() | nil,
          private_dataplane_seen_at: DateTime.t() | nil,
          max_session_duration_days: integer(),
          mfa_required: boolean(),
          docs_playground_sessions_enabled: boolean(),
          docs_playground_sessions_disabled_at: DateTime.t() | nil,
          lockout_after_failed_attempts: integer(),
          lockout_minutes: integer(),
          session_idle_timeout_minutes: integer(),
          session_max_lifetime_minutes: integer(),
          signup_data: map() | nil,
          signed_dpa: boolean(),
          elasticsearch_node_url: String.t() | nil,
          elasticsearch_api_key: String.t() | nil,
          use_custom_elasticsearch: boolean(),
          s3_endpoint: String.t() | nil,
          s3_access_key_id: String.t() | nil,
          s3_secret_access_key: String.t() | nil,
          s3_bucket: String.t() | nil,
          use_custom_s3: boolean(),
          sent_plan_limit_alert: DateTime.t() | nil,
          sso_domain: String.t() | nil,
          sso_provider: String.t() | nil,
          domain_join: String.t(),
          joiner_role: atom(),
          join_domains: [String.t()],
          presence_enabled: boolean(),
          trace_sharing_enabled: boolean(),
          support_contact: String.t() | nil,
          primary_intent: atom() | nil,
          promo_code: String.t() | nil,
          stripe_customer_id: String.t() | nil,
          currency: atom(),
          pricing_model: atom(),
          license: String.t() | nil,
          license_expires_at: DateTime.t() | nil,
          license_last_validated_at: DateTime.t() | nil,
          self_hosted_customer: boolean(),
          connect_services_disabled: [String.t()],
          instant_evals_enabled_at: DateTime.t() | nil,
          instant_evals_enabled_by_user_id: String.t() | nil,
          connect_last_sync_at: DateTime.t() | nil,
          connect_last_sync_error: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          connected_billing: LangWatch.ConnectedBillingAccount.t() | nil,
          members: [LangWatch.OrganizationUser.t()],
          teams: [LangWatch.Team.t()],
          webhook_endpoints: [LangWatch.WebhookEndpoint.t()],
          organization_invite: [LangWatch.OrganizationInvite.t()],
          custom_roles: [LangWatch.CustomRole.t()],
          notifications: [LangWatch.Notification.t()],
          subscriptions: [LangWatch.Subscription.t()],
          scim_tokens: [LangWatch.ScimToken.t()],
          scim_request_log: [LangWatch.ScimRequestLog.t()],
          prompt_tags: [LangWatch.PromptTag.t()],
          groups: [LangWatch.Group.t()],
          role_bindings: [LangWatch.RoleBinding.t()],
          api_keys: [LangWatch.ApiKey.t()],
          issued_licenses: [LangWatch.IssuedLicense.t()],
          gateway_change_events: [LangWatch.GatewayChangeEvent.t()],
          routing_policies: [LangWatch.RoutingPolicy.t()],
          ingestion_sources: [LangWatch.IngestionSource.t()],
          anomaly_rules: [LangWatch.AnomalyRule.t()],
          anomaly_alerts: [LangWatch.AnomalyAlert.t()],
          ai_tool_entries: [LangWatch.AiToolEntry.t()],
          ingestion_templates: [LangWatch.IngestionTemplate.t()],
          github_installations: [LangWatch.GithubInstallation.t()],
          github_pull_requests: [LangWatch.GithubPullRequest.t()],
          github_branch_pr_checks: [LangWatch.GithubBranchPullRequestCheck.t()]
        }

  schema "Organization" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    # `customer` for every organization a person founds or an operator provisions; `system` for
    # the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
    # Spec: modules/ops/specs/checkup-canary-project.feature.
    field(:kind, :string, default: "customer")
    field(:phone_number, :string, source: :phoneNumber)
    field(:slug, :string)
    field(:usage_spending_max_limit, :integer, source: :usageSpendingMaxLimit)
    # The largest file a dataset image or file cell accepts for this
    # organization, in mebibytes. Null means the default. Set by an operator
    # from the back office only; every other dataset size limit scales with it.
    field(:dataset_attachment_max_mb, :integer, source: :datasetAttachmentMaxMb)
    # The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
    # organization module the first time it is seen. Once set, a missing route refuses the
    # organization instead of serving it from the shared ClickHouse. Null: never private.
    # Spec: specs/private-dataplane/clickhouse-routing.feature.
    field(:private_click_house_label, :string, source: :privateClickHouseLabel)
    # The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature.
    field(:private_s3_label, :string, source: :privateS3Label)
    # When either label was first recorded.
    field(:private_dataplane_seen_at, LangWatch.PrismaDateTime, source: :privateDataplaneSeenAt)
    # Phase 8 - admin-configurable max lifetime for CLI device-flow
    # sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
    # gh-style flows). Companies that need finite session lifetime
    # (eg. SOC 2 / ISO 27001 customers with mandated session
    # expiration) can set a positive value. The /api/auth/cli/refresh
    # endpoint rejects refresh attempts where now - issued_at exceeds
    # this value, regardless of the underlying refresh-token TTL.
    # Spec: specs/ai-governance/sessions/session-ttl.feature.
    field(:max_session_duration_days, :integer, default: 0, source: :maxSessionDurationDays)
    # Every member of this organization can prove a second factor (D06). A
    # MEMBERSHIP CONDITION, not a session policy: it is asked when a member
    # reaches this organization's data, it holds the ones who cannot prove
    # one at an enrollment gate for THIS organization alone, and it ends no
    # session, ever. Turning it on signs nobody out.
    # Spec: specs/identity/mfa-and-session-shape.feature.
    field(:mfa_required, :boolean, default: false, source: :mfaRequired)
    # Whether members may connect the documentation playground to their own access with a
    # short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
    # no key is minted and live ones stop authenticating. The pasted-key path is untouched.
    # Spec: modules/api-key/specs/docs-session-keys.feature.
    field(:docs_playground_sessions_enabled, :boolean, default: true, source: :docsPlaygroundSessionsEnabled)
    # When the docs playground was last turned off, written in the same update as the switch. A
    # docs session key minted at or before it never authenticates again, nor any key minted under
    # one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
    field(:docs_playground_sessions_disabled_at, LangWatch.PrismaDateTime, source: :docsPlaygroundSessionsDisabledAt)
    # How many consecutive failed sign-ins lock an account (GAC-09).
    # 0 = never lock, which is the default and what every organization that
    # has not asked for this keeps. Counts EVERY credential together - a
    # password, a one-time code, a backup code, a passkey - because five per
    # factor hands an attacker one budget per factor.
    # Spec: specs/identity/org-account-lockout.feature.
    field(:lockout_after_failed_attempts, :integer, default: 0, source: :lockoutAfterFailedAttempts)
    # How long a lock lasts, in minutes. Ignored while the threshold is 0.
    field(:lockout_minutes, :integer, default: 30, source: :lockoutMinutes)
    # How long a BROWSER session may sit idle before it ends, in minutes
    # (GAC-10). 0 = no idle timeout. Activity rolls the window, which is the
    # "refreshed every hour" half of the control.
    #
    # A day for organizations created from now on; organizations that predate
    # the column keep 0 and are unchanged, because the migration adds the
    # column at zero and moves the default afterwards.
    #
    # NOT `maxSessionDurationDays`, which bounds CLI and device sessions and
    # is enforced when the CLI refreshes. Two surfaces, two lifetimes, two
    # decisions an administrator makes separately.
    # Spec: specs/identity/org-session-lifetime.feature.
    field(:session_idle_timeout_minutes, :integer, default: 1440, source: :sessionIdleTimeoutMinutes)
    # The ceiling on a browser session from the moment it was minted, in
    # minutes, which no amount of activity extends. 0 = no ceiling, the
    # default. Offered beside the idle timeout because a session kept alive
    # forever by a polling page is not really bounded.
    field(:session_max_lifetime_minutes, :integer, default: 0, source: :sessionMaxLifetimeMinutes)
    field(:signup_data, :map, source: :signupData)
    field(:signed_dpa, :boolean, default: false, source: :signedDPA)
    field(:elasticsearch_node_url, :string, source: :elasticsearchNodeUrl)
    field(:elasticsearch_api_key, :string, source: :elasticsearchApiKey)
    field(:use_custom_elasticsearch, :boolean, default: false, source: :useCustomElasticsearch)
    field(:s3_endpoint, :string, source: :s3Endpoint)
    field(:s3_access_key_id, :string, source: :s3AccessKeyId)
    field(:s3_secret_access_key, :string, source: :s3SecretAccessKey)
    field(:s3_bucket, :string, source: :s3Bucket)
    field(:use_custom_s3, :boolean, default: false, source: :useCustomS3)
    field(:sent_plan_limit_alert, LangWatch.PrismaDateTime, source: :sentPlanLimitAlert)
    field(:sso_domain, :string, source: :ssoDomain)
    field(:sso_provider, :string, source: :ssoProvider)
    # How colleagues on a matching domain get in (D12, ADR-117):
    # `off` (invisible to them entirely) | `request` (they ask, an admin
    # approves) | `auto` (a verified colleague walks straight in).
    #
    # `request` is the default for a self-serve organization; `auto` is never
    # a default and never inferred - an administrator turns it on, names the
    # domain in `joinDomains` while doing it, and the licence gate holds that
    # value while letting `request` through.
    field(:domain_join, :string, default: "request", source: :domainJoin)
    # The organisation role handed to a person admitted without an invitation:
    # a self-service domain join (`domainJoin`) or an SSO-admitted login.
    # `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
    # organisation settings page only; the join flow never asks, and an
    # invitation always names its own role.
    field(:joiner_role, Ecto.Enum, values: [:ADMIN, :MEMBER, :EXTERNAL, :DEVELOPER], default: :MEMBER, source: :joinerRole)
    # The domains an administrator NAMED when turning automatic joining on.
    # Empty means automatic joining admits nobody whatever `domainJoin` says -
    # which is what keeps a setting left on from admitting a domain nobody
    # ever agreed to.
    field(:join_domains, {:array, :string}, source: :joinDomains)
    field(:presence_enabled, :boolean, default: true, source: :presenceEnabled)
    # Org-level trace-sharing kill switch. Off = trace share links cannot be
    # created in ANY project of this org, and existing links stop resolving —
    # the org toggle is the global switch, the per-project
    # `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
    # `presenceEnabled`). Effective sharing = org AND project. ADR-039.
    field(:trace_sharing_enabled, :boolean, default: true, source: :traceSharingEnabled)
    # Org-configurable contact surfaced to users in CLI "contact your
    # admin" messages and the in-app budget-exceeded banner. Free text:
    # accepts an email, a URL (internal ticketing system), or any other
    # instruction. When null, callers fall back to the first ADMIN
    # member's email (`resolveSupportContact` helper).
    field(:support_contact, :string, source: :supportContact)
    # Signup intent, editable later as the org "Primary use" setting.
    # AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
    # NULL means intent unset (every pre-fork org and any org created
    # outside onboarding) — the home resolver keeps legacy behavior.
    # ADR-038.
    field(:primary_intent, Ecto.Enum, values: [:AGENT_GOVERNANCE, :LLM_OPS], source: :primaryIntent)
    field(:promo_code, :string, source: :promoCode)
    field(:stripe_customer_id, :string, source: :stripeCustomerId)
    field(:currency, Ecto.Enum, values: [:USD, :EUR], default: :EUR)
    field(:pricing_model, Ecto.Enum, values: [:TIERED, :SEAT_EVENT], default: :TIERED, source: :pricingModel)
    field(:license, :string)
    field(:license_expires_at, LangWatch.PrismaDateTime, source: :licenseExpiresAt)
    field(:license_last_validated_at, LangWatch.PrismaDateTime, source: :licenseLastValidatedAt)
    # A customer that runs LangWatch self-hosted. The hosted usage, budgets and
    # invoices of its connected installs attach to this organization (ADR-139).
    field(:self_hosted_customer, :boolean, default: false, source: :selfHostedCustomer)
    # LangWatch-hosted services this organization has switched OFF, named the
    # way its license names them ("instant_evals"). A service the license is
    # entitled to is on unless it is listed here, so a customer who bought
    # hosted judging gets it without finding a settings page first, and an
    # administrator who switched one off keeps it off across license changes
    # (ADR-139, section 2). A license that names no service reaches nothing,
    # whatever this column holds.
    field(:connect_services_disabled, {:array, :string}, source: :connectServicesDisabled)
    # When a member of this organization switched Instant Evals on from the
    # search bar, which is the organization's consent to have judged text sent
    # to the judge's provider. Null until then: a self-serve organization is
    # never opted in on its behalf, and an enterprise one is switched on by an
    # operator through the release flag after asking, never here.
    field(:instant_evals_enabled_at, LangWatch.PrismaDateTime, source: :instantEvalsEnabledAt)
    # The member who switched them on, so the audit trail names a person and
    # not just a moment.
    field(:instant_evals_enabled_by_user_id, :string, source: :instantEvalsEnabledByUserId)
    # When the last license sync succeeded.
    field(:connect_last_sync_at, LangWatch.PrismaDateTime, source: :connectLastSyncAt)
    # The code of the failure the last sync ended on, cleared on the next
    # success. Settings, Connect renders the copy for it.
    field(:connect_last_sync_error, :string, source: :connectLastSyncError)
    has_one(:connected_billing, LangWatch.ConnectedBillingAccount, foreign_key: :organization_id, references: :id)
    has_many(:members, LangWatch.OrganizationUser, foreign_key: :organization_id, references: :id)
    has_many(:teams, LangWatch.Team, foreign_key: :organization_id, references: :id)
    has_many(:webhook_endpoints, LangWatch.WebhookEndpoint, foreign_key: :organization_id, references: :id)
    has_many(:organization_invite, LangWatch.OrganizationInvite, foreign_key: :organization_id, references: :id)
    has_many(:custom_roles, LangWatch.CustomRole, foreign_key: :organization_id, references: :id)
    has_many(:notifications, LangWatch.Notification, foreign_key: :organization_id, references: :id)
    has_many(:subscriptions, LangWatch.Subscription, foreign_key: :organization_id, references: :id)
    has_many(:scim_tokens, LangWatch.ScimToken, foreign_key: :organization_id, references: :id)
    has_many(:scim_request_log, LangWatch.ScimRequestLog, foreign_key: :organization_id, references: :id)
    has_many(:prompt_tags, LangWatch.PromptTag, foreign_key: :organization_id, references: :id)
    has_many(:groups, LangWatch.Group, foreign_key: :organization_id, references: :id)
    has_many(:role_bindings, LangWatch.RoleBinding, foreign_key: :organization_id, references: :id)
    has_many(:api_keys, LangWatch.ApiKey, foreign_key: :organization_id, references: :id)
    has_many(:issued_licenses, LangWatch.IssuedLicense, foreign_key: :organization_id, references: :id)
    has_many(:gateway_change_events, LangWatch.GatewayChangeEvent, foreign_key: :organization_id, references: :id)
    has_many(:routing_policies, LangWatch.RoutingPolicy, foreign_key: :organization_id, references: :id)
    has_many(:ingestion_sources, LangWatch.IngestionSource, foreign_key: :organization_id, references: :id)
    has_many(:anomaly_rules, LangWatch.AnomalyRule, foreign_key: :organization_id, references: :id)
    has_many(:anomaly_alerts, LangWatch.AnomalyAlert, foreign_key: :organization_id, references: :id)
    has_many(:ai_tool_entries, LangWatch.AiToolEntry, foreign_key: :organization_id, references: :id)
    has_many(:ingestion_templates, LangWatch.IngestionTemplate, foreign_key: :organization_id, references: :id)
    has_many(:github_installations, LangWatch.GithubInstallation, foreign_key: :organization_id, references: :id)
    has_many(:github_pull_requests, LangWatch.GithubPullRequest, foreign_key: :organization_id, references: :id)
    has_many(:github_branch_pr_checks, LangWatch.GithubBranchPullRequestCheck, foreign_key: :organization_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end