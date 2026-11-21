defmodule Formbricks.Survey do
  @moduledoc """
  Represents a complete survey configuration including questions, styling, and display rules.
  Core model for the survey functionality in Formbricks.

  @property id - Unique identifier for the survey
  @property name - Display name of the survey
  @property type - Survey delivery method (link, web, website, app)
  @property status - Current state of the survey (draft, active, completed, etc)
  @property workspace - The workspace this survey belongs to
  @property questions - JSON array containing survey questions configuration
  @property displayOption - Rules for how often the survey can be shown
  @property triggers - Actions that can trigger this survey
  @property attributeFilters - Rules for targeting specific contacts
  @property customHeadScripts - Survey-specific custom HTML scripts (self-hosted only)
  @property customHeadScriptsMode - "add" (merge with workspace) or "replace" (override workspace)
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          hidden_fields: map(),
          variables: map(),
          ownership_initialized: boolean(),
          id: String.t(),
          name: String.t(),
          redirect_url: String.t() | nil,
          type: atom(),
          visibility: atom(),
          visibility_version: integer(),
          visibility_projected_version: integer(),
          visibility_pending: boolean(),
          visibility_changed_at: DateTime.t() | nil,
          visibility_changed_by_id: String.t() | nil,
          status: atom(),
          welcome_card: map(),
          questions: map(),
          blocks: [map()],
          endings: [map()],
          display_option: atom(),
          recontact_days: integer() | nil,
          display_limit: integer() | nil,
          inline_triggers: map() | nil,
          auto_close: integer() | nil,
          keep_ending_card_open: boolean(),
          auto_complete: integer() | nil,
          delay: integer(),
          publish_on: DateTime.t() | nil,
          close_on: DateTime.t() | nil,
          archived_at: DateTime.t() | nil,
          survey_closed_message: map() | nil,
          workspace_overwrites: map() | nil,
          styling: map() | nil,
          single_use: map() | nil,
          is_verify_email_enabled: boolean(),
          is_back_button_hidden: boolean(),
          is_auto_progressing_enabled: boolean(),
          is_capture_ip_enabled: boolean(),
          is_anonymize_responses_enabled: boolean(),
          pin: String.t() | nil,
          display_percentage: Decimal.t() | nil,
          show_language_switch: boolean() | nil,
          auto_select_language: boolean() | nil,
          recaptcha: map() | nil,
          metadata: map(),
          slug: String.t() | nil,
          custom_head_scripts: String.t() | nil,
          custom_head_scripts_mode: atom() | nil,
          custom_css: map() | nil,
          show_conditions: map() | nil,
          start_logic: map() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          workspace: Formbricks.Workspace.t() | nil,
          creator: Formbricks.User.t() | nil,
          owner: Formbricks.User.t() | nil,
          segment: Formbricks.Segment.t() | nil,
          retention_restore: Formbricks.RetentionSurveyRestore.t() | nil,
          responses: [Formbricks.Response.t()],
          quotas: [Formbricks.SurveyQuota.t()],
          triggers: [Formbricks.SurveyTrigger.t()],
          attribute_filters: [Formbricks.SurveyAttributeFilter.t()],
          displays: [Formbricks.Display.t()],
          languages: [Formbricks.SurveyLanguage.t()],
          follow_ups: [Formbricks.SurveyFollowUp.t()],
          retention_exemptions: [Formbricks.RetentionExemption.t()],
          retention_notices: [Formbricks.RetentionNotice.t()]
        }

  schema "Survey" do
    # [SurveyHiddenFields]
    field(:hidden_fields, :map, default: %{"enabled" => false}, source: :hiddenFields)
    # [SurveyVariables]
    field(:variables, :map)
    field(:ownership_initialized, :boolean, default: false, source: :ownershipInitialized)
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:redirect_url, :string, source: :redirectUrl)
    field(:type, Ecto.Enum, values: [:link, :app], default: :app)
    # ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
    # the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
    field(:visibility, Ecto.Enum, values: [:restricted, :workspace], default: :workspace)
    # Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
    # with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
    # projection, decided from PostgreSQL facts until the projector acknowledges it.
    field(:visibility_version, :integer, default: 0, source: :visibilityVersion)
    # The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was.
    field(:visibility_projected_version, :integer, default: 0, source: :visibilityProjectedVersion)
    # "visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
    # trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
    # in create/update data.
    field(:visibility_pending, :boolean, default: false, source: :visibilityPending)
    field(:visibility_changed_at, Formbricks.PrismaDateTime, source: :visibilityChangedAt)
    field(:visibility_changed_by_id, :string, source: :visibilityChangedById)
    field(:status, Ecto.Enum, values: [:draft, :inProgress, :paused, :completed], default: :draft)
    # [SurveyWelcomeCard]
    field(:welcome_card, :map, default: %{"enabled" => false}, source: :welcomeCard)
    # [SurveyQuestions]
    field(:questions, :map)
    # [SurveyBlocks]
    field(:blocks, {:array, :map})
    # [SurveyEnding]
    field(:endings, {:array, :map})
    field(:display_option, Ecto.Enum, values: [:displayOnce, :displayMultiple, :displaySome, :respondMultiple], default: :displayOnce, source: :displayOption)
    field(:recontact_days, :integer, source: :recontactDays)
    field(:display_limit, :integer, source: :displayLimit)
    # [SurveyInlineTriggers]
    field(:inline_triggers, :map, source: :inlineTriggers)
    field(:auto_close, :integer, source: :autoClose)
    field(:keep_ending_card_open, :boolean, default: false, source: :keepEndingCardOpen)
    field(:auto_complete, :integer, source: :autoComplete)
    field(:delay, :integer, default: 0)
    field(:publish_on, Formbricks.PrismaDateTime, source: :publishOn)
    field(:close_on, Formbricks.PrismaDateTime, source: :closeOn)
    field(:archived_at, Formbricks.PrismaDateTime, source: :archivedAt)
    # [SurveyClosedMessage]
    field(:survey_closed_message, :map, source: :surveyClosedMessage)
    # [SurveyWorkspaceOverwrites]
    field(:workspace_overwrites, :map, source: :workspaceOverwrites)
    # [SurveyStyling]
    field(:styling, :map)
    # [SurveySingleUse]
    field(:single_use, :map, default: %{"enabled" => false, "isEncrypted" => true}, source: :singleUse)
    field(:is_verify_email_enabled, :boolean, default: false, source: :isVerifyEmailEnabled)
    field(:is_back_button_hidden, :boolean, default: false, source: :isBackButtonHidden)
    field(:is_auto_progressing_enabled, :boolean, default: false, source: :isAutoProgressingEnabled)
    field(:is_capture_ip_enabled, :boolean, default: false, source: :isCaptureIpEnabled)
    field(:is_anonymize_responses_enabled, :boolean, default: false)
    field(:pin, :string)
    field(:display_percentage, :decimal, source: :displayPercentage)
    field(:show_language_switch, :boolean, source: :showLanguageSwitch)
    # Open the survey in the respondent's browser language when it matches an enabled survey
    # language. Null on surveys that predate the setting, which reads as off.
    field(:auto_select_language, :boolean, source: :autoSelectLanguage)
    # [SurveyRecaptcha]
    field(:recaptcha, :map, default: %{"enabled" => false, "threshold" => 0.1})
    # [SurveyLinkMetadata]
    field(:metadata, :map, default: %{})
    field(:slug, :string)
    field(:custom_head_scripts, :string, source: :customHeadScripts)
    field(:custom_head_scripts_mode, Ecto.Enum, values: [:add, :replace], default: :add, source: :customHeadScriptsMode)
    # [CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
    # by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
    field(:custom_css, :map, source: :customCss)
    # [SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
    # (ENG-3873). null = no conditions, which is what every existing survey keeps.
    field(:show_conditions, :map, source: :showConditions)
    # [SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
    # survey on load. null = none, which is what every existing survey keeps.
    field(:start_logic, :map, source: :startLogic)
    field(:workspace_id, :string, source: :workspaceId)
    field(:created_by, :string, source: :createdBy)
    field(:owner_id, :string, source: :ownerId)
    field(:segment_id, :string, source: :segmentId)
    belongs_to(:workspace, Formbricks.Workspace, foreign_key: :workspace_id, define_field: false, type: :string)
    belongs_to(:creator, Formbricks.User, foreign_key: :created_by, define_field: false, type: :string)
    belongs_to(:owner, Formbricks.User, foreign_key: :owner_id, define_field: false, type: :string)
    belongs_to(:segment, Formbricks.Segment, foreign_key: :segment_id, define_field: false, type: :string)
    has_one(:retention_restore, Formbricks.RetentionSurveyRestore, foreign_key: :survey_id, references: :id)
    has_many(:responses, Formbricks.Response, foreign_key: :survey_id, references: :id)
    has_many(:quotas, Formbricks.SurveyQuota, foreign_key: :survey_id, references: :id)
    has_many(:triggers, Formbricks.SurveyTrigger, foreign_key: :survey_id, references: :id)
    has_many(:attribute_filters, Formbricks.SurveyAttributeFilter, foreign_key: :survey_id, references: :id)
    has_many(:displays, Formbricks.Display, foreign_key: :survey_id, references: :id)
    has_many(:languages, Formbricks.SurveyLanguage, foreign_key: :survey_id, references: :id)
    has_many(:follow_ups, Formbricks.SurveyFollowUp, foreign_key: :survey_id, references: :id)
    has_many(:retention_exemptions, Formbricks.RetentionExemption, foreign_key: :survey_id, references: :id)
    has_many(:retention_notices, Formbricks.RetentionNotice, foreign_key: :survey_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end