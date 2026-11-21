defmodule LangWatch.ApiKey do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          lookup_id: String.t(),
          hashed_secret: String.t(),
          permission_mode: String.t(),
          created_by_user_id: String.t() | nil,
          expires_at: DateTime.t() | nil,
          last_used_at: DateTime.t() | nil,
          revoked_at: DateTime.t() | nil,
          revocation_cause: String.t() | nil,
          parent_api_key_id: String.t() | nil,
          ingest_source_type: String.t() | nil,
          ingestion_template_id: String.t() | nil,
          created_by_device_label: String.t() | nil,
          is_system_managed: boolean(),
          created_by_browser_session_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: LangWatch.User.t() | nil,
          organization: LangWatch.Organization.t() | nil,
          role_bindings: [LangWatch.RoleBinding.t()]
        }

  schema "ApiKey" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:lookup_id, :string, source: :lookupId)
    field(:hashed_secret, :string, source: :hashedSecret)
    field(:permission_mode, :string, default: "all", source: :permissionMode)
    field(:created_by_user_id, :string, source: :createdByUserId)
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    field(:last_used_at, LangWatch.PrismaDateTime, source: :lastUsedAt)
    field(:revoked_at, LangWatch.PrismaDateTime, source: :revokedAt)
    # Why the key was revoked, set with `revokedAt`: "user" when a person
    # revoked it (the API-keys page, the REST API, the devices tab), "rotation"
    # when a mint replaced it, "session" when the CLI login key it was parented
    # to was revoked, "expired" when that session ran out, "offboarded" when the
    # person's membership of the organization ended. "cap" survives on rows the
    # retired per-tool cap revoked. The CLI reads it to decide whether a dead
    # ingest key may be re-minted: a key a person revoked stays dead.
    # NULL for live keys and for keys revoked before the cause was recorded.
    field(:revocation_cause, :string, source: :revocationCause)
    # The CLI login key of the device session that minted this ingestion key.
    # Set by the CLI's personal mint, NULL for every other key. Revoking the
    # parent (logout, the devices tab, a re-login from the same device, session
    # expiry) revokes the children with it. Column ref only, no FK and no
    # index: lookups filter by the indexed `userId` first and match the parent
    # in memory over a person's few live ingest keys.
    field(:parent_api_key_id, :string, source: :parentApiKeyId)
    # Non-null marks this key as an "ingestion key": a project-scoped,
    # ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
    # installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
    # opencode / claude_cowork) is stamped as `langwatch.source` provenance on
    # every span the key authorizes. NULL for normal API / service keys.
    field(:ingest_source_type, :string, source: :ingestSourceType)
    # Set only for ingestion keys derived from an admin IngestionTemplate
    # (e.g. claude_cowork). Drives receiver OTTL application + the
    # `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
    # (which have no template) and for normal keys. Column ref only (no FK,
    # per the app-layer-integrity convention); resolved by the ingest-key service.
    field(:ingestion_template_id, :string, source: :ingestionTemplateId)
    # Human label of the CLI device session that minted this ingestion key
    # (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
    # flow's `client_info`. Denormalized so the org-wide API-keys settings page
    # can show which device a key came from without resolving another user's
    # Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
    # a CLI old enough to predate device metadata.
    field(:created_by_device_label, :string, source: :createdByDeviceLabel)
    # True for a key LangWatch minted for itself (a run, the gateway's trace
    # export, a Langy session). Hides it from listings, refuses customer edits
    # and revokes, and lets an ownerless run key act as the system. Never set
    # by a customer create; rows before the column stay false.
    field(:is_system_managed, :boolean, default: false, source: :isSystemManaged)
    # The browser session that minted a docs session key (ADR-185), so signing
    # out of that session revokes it. NULL for every other key. Unindexed:
    # sign-out filters by the indexed `userId` first.
    field(:created_by_browser_session_id, :string, source: :createdByBrowserSessionId)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:user, LangWatch.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:role_bindings, LangWatch.RoleBinding, foreign_key: :api_key_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end