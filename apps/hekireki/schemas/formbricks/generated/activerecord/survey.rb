# Represents a complete survey configuration including questions, styling, and display rules.
# Core model for the survey functionality in Formbricks.
#
# @property id - Unique identifier for the survey
# @property name - Display name of the survey
# @property type - Survey delivery method (link, web, website, app)
# @property status - Current state of the survey (draft, active, completed, etc)
# @property workspace - The workspace this survey belongs to
# @property questions - JSON array containing survey questions configuration
# @property displayOption - Rules for how often the survey can be shown
# @property triggers - Actions that can trigger this survey
# @property attributeFilters - Rules for targeting specific contacts
# @property customHeadScripts - Survey-specific custom HTML scripts (self-hosted only)
# @property customHeadScriptsMode - "add" (merge with workspace) or "replace" (override workspace)
#
# hiddenFields:: [SurveyHiddenFields]
# variables:: [SurveyVariables]
# visibility:: ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
#   the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
# visibilityVersion:: Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
#   with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
#   projection, decided from PostgreSQL facts until the projector acknowledges it.
# visibilityProjectedVersion:: The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was.
# visibilityPending:: "visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
#   trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
#   in create/update data.
# welcomeCard:: [SurveyWelcomeCard]
# questions:: [SurveyQuestions]
# blocks:: [SurveyBlocks]
# endings:: [SurveyEnding]
# inlineTriggers:: [SurveyInlineTriggers]
# surveyClosedMessage:: [SurveyClosedMessage]
# workspaceOverwrites:: [SurveyWorkspaceOverwrites]
# styling:: [SurveyStyling]
# singleUse:: [SurveySingleUse]
# autoSelectLanguage:: Open the survey in the respondent's browser language when it matches an enabled survey
#   language. Null on surveys that predate the setting, which reads as off.
# recaptcha:: [SurveyRecaptcha]
# metadata:: [SurveyLinkMetadata]
# customCss:: [CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
#   by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
# creatives:: Creatives whose Destination is this survey (ENG-3993).
# attachment_exports:: Background attachment ZIP exports of this survey (ENG-6862).
# showConditions:: [SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
#   (ENG-3873). null = no conditions, which is what every existing survey keeps.
# startLogic:: [SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
#   survey on load. null = none, which is what every existing survey keeps.
class Survey < ApplicationRecord
  self.table_name = "Survey"
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :hiddenFields, default: -> { JSON.parse("{\"enabled\": false}") }
  attribute :variables, default: -> { JSON.parse("[]") }
  attribute :ownershipInitialized, default: false
  attribute :id, default: -> { Cuid.generate }
  attribute :visibilityVersion, default: 0
  attribute :visibilityProjectedVersion, default: 0
  attribute :visibilityPending, default: false
  attribute :welcomeCard, default: -> { JSON.parse("{\"enabled\": false}") }
  attribute :questions, default: -> { JSON.parse("[]") }
  attribute :blocks, default: -> { [] }
  attribute :endings, default: -> { [] }
  attribute :keepEndingCardOpen, default: false
  attribute :delay, default: 0
  attribute :singleUse, default: -> { JSON.parse("{\"enabled\": false, \"isEncrypted\": true}") }
  attribute :isVerifyEmailEnabled, default: false
  attribute :isBackButtonHidden, default: false
  attribute :isAutoProgressingEnabled, default: false
  attribute :isCaptureIpEnabled, default: false
  attribute :is_anonymize_responses_enabled, default: false
  attribute :recaptcha, default: -> { JSON.parse("{\"enabled\": false, \"threshold\":0.1}") }
  attribute :metadata, default: -> { JSON.parse("{}") }

  enum :type, { link: "link", app: "app" }, default: :app, validate: true
  enum :visibility, { restricted: "restricted", workspace: "workspace" }, default: :workspace, validate: true
  enum :status, { draft: "draft", in_progress: "inProgress", paused: "paused", completed: "completed" }, default: :draft,
                                                                                                         validate: true
  enum :displayOption, { display_once: "displayOnce", display_multiple: "displayMultiple", display_some: "displaySome", respond_multiple: "respondMultiple" },
       default: :display_once, validate: true
  enum :customHeadScriptsMode, { add: "add", replace: "replace" }, default: :add, validate: { allow_nil: true }

  validates :id, uniqueness: { scope: :workspaceId }
  validates :name, presence: true
  validates :slug, uniqueness: { allow_nil: true }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :surveys
  belongs_to :creator, class_name: "User", foreign_key: "createdBy", optional: true, inverse_of: :surveys
  belongs_to :owner, class_name: "User", foreign_key: "ownerId", optional: true, inverse_of: :owned_surveys
  belongs_to :segment, foreign_key: "segmentId", optional: true, inverse_of: :surveys
  has_one :retention_restore, class_name: "RetentionSurveyRestore", foreign_key: "surveyId", inverse_of: :survey,
                              dependent: :destroy
  has_many :responses, foreign_key: "surveyId", inverse_of: :survey, dependent: :destroy
  has_many :quotas, class_name: "SurveyQuota", foreign_key: "surveyId", inverse_of: :survey, dependent: :destroy
  has_many :triggers, class_name: "SurveyTrigger", foreign_key: "surveyId", inverse_of: :survey, dependent: :destroy
  has_many :attribute_filters, class_name: "SurveyAttributeFilter", foreign_key: "surveyId", inverse_of: :survey,
                               dependent: :destroy
  has_many :displays, foreign_key: "surveyId", inverse_of: :survey, dependent: :destroy
  has_many :languages, class_name: "SurveyLanguage", foreign_key: "surveyId", inverse_of: :survey, dependent: :destroy
  has_many :embedded_data, class_name: "EmbeddedData", foreign_key: [ :surveyId, :workspaceId ],
                           primary_key: [ :id, :workspaceId ], inverse_of: :survey, dependent: :destroy
  has_many :embedded_data_links, class_name: "SurveyEmbeddedData", foreign_key: [ :surveyId, :workspaceId ],
                                 primary_key: [ :id, :workspaceId ], inverse_of: :survey, dependent: :destroy
  has_many :follow_ups, class_name: "SurveyFollowUp", foreign_key: "surveyId", inverse_of: :survey, dependent: :destroy
  has_many :feedback_source_mappings, class_name: "FeedbackSourceFormbricksMapping",
                                      foreign_key: [ :surveyId, :workspaceId ], primary_key: [ :id, :workspaceId ],
                                      inverse_of: :survey, dependent: :destroy
  has_many :retention_exemptions, foreign_key: "surveyId", inverse_of: :survey, dependent: :destroy
  has_many :retention_notices, foreign_key: "surveyId", inverse_of: :survey, dependent: :destroy
  has_many :creatives, foreign_key: [ :surveyId, :workspaceId ], primary_key: [ :id, :workspaceId ], inverse_of: :survey,
                       dependent: :restrict_with_error
  has_many :attachment_exports, foreign_key: [ :surveyId, :workspaceId ], primary_key: [ :id, :workspaceId ],
                                inverse_of: :survey, dependent: :destroy
end
