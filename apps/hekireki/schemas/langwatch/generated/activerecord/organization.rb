# kind:: `customer` for every organization a person founds or an operator provisions; `system` for
#   the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
#   Spec: modules/ops/specs/checkup-canary-project.feature.
# datasetAttachmentMaxMb:: The largest file a dataset image or file cell accepts for this
#   organization, in mebibytes. Null means the default. Set by an operator
#   from the back office only; every other dataset size limit scales with it.
# privateClickHouseLabel:: The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
#   organization module the first time it is seen. Once set, a missing route refuses the
#   organization instead of serving it from the shared ClickHouse. Null: never private.
#   Spec: specs/private-dataplane/clickhouse-routing.feature.
# privateS3Label:: The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature.
# privateDataplaneSeenAt:: When either label was first recorded.
# maxSessionDurationDays:: Phase 8 - admin-configurable max lifetime for CLI device-flow
#   sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
#   gh-style flows). Companies that need finite session lifetime
#   (eg. SOC 2 / ISO 27001 customers with mandated session
#   expiration) can set a positive value. The /api/auth/cli/refresh
#   endpoint rejects refresh attempts where now - issued_at exceeds
#   this value, regardless of the underlying refresh-token TTL.
#   Spec: specs/ai-governance/sessions/session-ttl.feature.
# mfaRequired:: Every member of this organization can prove a second factor (D06). A
#   MEMBERSHIP CONDITION, not a session policy: it is asked when a member
#   reaches this organization's data, it holds the ones who cannot prove
#   one at an enrollment gate for THIS organization alone, and it ends no
#   session, ever. Turning it on signs nobody out.
#   Spec: specs/identity/mfa-and-session-shape.feature.
# docsPlaygroundSessionsEnabled:: Whether members may connect the documentation playground to their own access with a
#   short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
#   no key is minted and live ones stop authenticating. The pasted-key path is untouched.
#   Spec: modules/api-key/specs/docs-session-keys.feature.
# docsPlaygroundSessionsDisabledAt:: When the docs playground was last turned off, written in the same update as the switch. A
#   docs session key minted at or before it never authenticates again, nor any key minted under
#   one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
# lockoutAfterFailedAttempts:: How many consecutive failed sign-ins lock an account (GAC-09).
#   0 = never lock, which is the default and what every organization that
#   has not asked for this keeps. Counts EVERY credential together - a
#   password, a one-time code, a backup code, a passkey - because five per
#   factor hands an attacker one budget per factor.
#   Spec: specs/identity/org-account-lockout.feature.
# lockoutMinutes:: How long a lock lasts, in minutes. Ignored while the threshold is 0.
# sessionIdleTimeoutMinutes:: How long a BROWSER session may sit idle before it ends, in minutes
#   (GAC-10). 0 = no idle timeout. Activity rolls the window, which is the
#   "refreshed every hour" half of the control.
#
#   A day for organizations created from now on; organizations that predate
#   the column keep 0 and are unchanged, because the migration adds the
#   column at zero and moves the default afterwards.
#
#   NOT `maxSessionDurationDays`, which bounds CLI and device sessions and
#   is enforced when the CLI refreshes. Two surfaces, two lifetimes, two
#   decisions an administrator makes separately.
#   Spec: specs/identity/org-session-lifetime.feature.
# sessionMaxLifetimeMinutes:: The ceiling on a browser session from the moment it was minted, in
#   minutes, which no amount of activity extends. 0 = no ceiling, the
#   default. Offered beside the idle timeout because a session kept alive
#   forever by a polling page is not really bounded.
# domainJoin:: How colleagues on a matching domain get in (D12, ADR-117):
#   `off` (invisible to them entirely) | `request` (they ask, an admin
#   approves) | `auto` (a verified colleague walks straight in).
#
#   `request` is the default for a self-serve organization; `auto` is never
#   a default and never inferred - an administrator turns it on, names the
#   domain in `joinDomains` while doing it, and the licence gate holds that
#   value while letting `request` through.
# joinerRole:: The organisation role handed to a person admitted without an invitation:
#   a self-service domain join (`domainJoin`) or an SSO-admitted login.
#   `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
#   organisation settings page only; the join flow never asks, and an
#   invitation always names its own role.
# joinDomains:: The domains an administrator NAMED when turning automatic joining on.
#   Empty means automatic joining admits nobody whatever `domainJoin` says -
#   which is what keeps a setting left on from admitting a domain nobody
#   ever agreed to.
# traceSharingEnabled:: Org-level trace-sharing kill switch. Off = trace share links cannot be
#   created in ANY project of this org, and existing links stop resolving —
#   the org toggle is the global switch, the per-project
#   `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
#   `presenceEnabled`). Effective sharing = org AND project. ADR-039.
# supportContact:: Org-configurable contact surfaced to users in CLI "contact your
#   admin" messages and the in-app budget-exceeded banner. Free text:
#   accepts an email, a URL (internal ticketing system), or any other
#   instruction. When null, callers fall back to the first ADMIN
#   member's email (`resolveSupportContact` helper).
# primaryIntent:: Signup intent, editable later as the org "Primary use" setting.
#   AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
#   NULL means intent unset (every pre-fork org and any org created
#   outside onboarding) — the home resolver keeps legacy behavior.
#   ADR-038.
# selfHostedCustomer:: A customer that runs LangWatch self-hosted. The hosted usage, budgets and
#   invoices of its connected installs attach to this organization (ADR-139).
# connectServicesDisabled:: LangWatch-hosted services this organization has switched OFF, named the
#   way its license names them ("instant_evals"). A service the license is
#   entitled to is on unless it is listed here, so a customer who bought
#   hosted judging gets it without finding a settings page first, and an
#   administrator who switched one off keeps it off across license changes
#   (ADR-139, section 2). A license that names no service reaches nothing,
#   whatever this column holds.
# instantEvalsEnabledAt:: When a member of this organization switched Instant Evals on from the
#   search bar, which is the organization's consent to have judged text sent
#   to the judge's provider. Null until then: a self-serve organization is
#   never opted in on its behalf, and an enterprise one is switched on by an
#   operator through the release flag after asking, never here.
# instantEvalsEnabledByUserId:: The member who switched them on, so the audit trail names a person and
#   not just a moment.
# connectLastSyncAt:: When the last license sync succeeded.
# connectLastSyncError:: The code of the failure the last sync ended on, cleared on the next
#   success. Settings, Connect renders the copy for it.
class Organization < ApplicationRecord
  self.table_name = "Organization"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :kind, default: "customer"
  attribute :maxSessionDurationDays, default: 0
  attribute :mfaRequired, default: false
  attribute :docsPlaygroundSessionsEnabled, default: true
  attribute :lockoutAfterFailedAttempts, default: 0
  attribute :lockoutMinutes, default: 30
  attribute :sessionIdleTimeoutMinutes, default: 1440
  attribute :sessionMaxLifetimeMinutes, default: 0
  attribute :signedDPA, default: false
  attribute :useCustomElasticsearch, default: false
  attribute :useCustomS3, default: false
  attribute :domainJoin, default: "request"
  attribute :joinDomains, default: -> { [] }
  attribute :presenceEnabled, default: true
  attribute :traceSharingEnabled, default: true
  attribute :selfHostedCustomer, default: false
  attribute :connectServicesDisabled, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :joinerRole, { admin: "ADMIN", member: "MEMBER", external: "EXTERNAL", developer: "DEVELOPER" }, default: :member,
                                                                                                        validate: true
  enum :primaryIntent, { agent_governance: "AGENT_GOVERNANCE", llm_ops: "LLM_OPS" }, validate: { allow_nil: true }
  enum :currency, { usd: "USD", eur: "EUR" }, default: :eur, validate: true
  enum :pricingModel, { tiered: "TIERED", seat_event: "SEAT_EVENT" }, default: :tiered, validate: true

  validates :name, presence: true
  validates :slug, presence: true, uniqueness: true
  validates :ssoDomain, uniqueness: { allow_nil: true }
  validates :stripeCustomerId, uniqueness: { allow_nil: true }

  has_one :connected_billing, class_name: "ConnectedBillingAccount", foreign_key: "organizationId",
                              inverse_of: :organization, dependent: :restrict_with_error
  has_many :members, class_name: "OrganizationUser", foreign_key: "organizationId", inverse_of: :organization,
                     dependent: :restrict_with_error
  has_many :teams, foreign_key: "organizationId", inverse_of: :organization, dependent: :restrict_with_error
  has_many :webhook_endpoints, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :organization_invite, foreign_key: "organizationId", inverse_of: :organization,
                                 dependent: :restrict_with_error
  has_many :custom_roles, foreign_key: "organizationId", inverse_of: :organization, dependent: :restrict_with_error
  has_many :notifications, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :subscriptions, foreign_key: "organizationId", inverse_of: :organization, dependent: :restrict_with_error
  has_many :scim_tokens, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :scim_request_log, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :prompt_tags, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :groups, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :role_bindings, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :api_keys, foreign_key: "organizationId", inverse_of: :organization, dependent: :restrict_with_error
  has_many :issued_licenses, foreign_key: "organizationId", inverse_of: :organization, dependent: :restrict_with_error
  has_many :gateway_change_events, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :routing_policies, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :ingestion_sources, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :anomaly_rules, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :anomaly_alerts, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :ai_tool_entries, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :ingestion_templates, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :github_installations, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :github_pull_requests, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :github_branch_pr_checks, class_name: "GithubBranchPullRequestCheck", foreign_key: "organizationId",
                                     inverse_of: :organization, dependent: :destroy
end
