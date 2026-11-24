defmodule Formbricks.Workspace do
  @moduledoc """
  Main grouping mechanism for resources in Formbricks.
  Each organization can have multiple workspaces to separate different applications or products.

  @property id - Unique identifier for the workspace
  @property name - Display name of the workspace
  @property organization - Reference to parent organization
  @property surveys - Collection of surveys in this workspace
  @property styling - Workspace-wide styling configuration
  @property config - Workspace-specific configuration
  @property recontactDays - Default recontact delay for surveys
  @property placement - Default widget placement for in-app surveys
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          legacy_environment_id: String.t() | nil,
          styling: map() | list() | String.t() | number() | boolean(),
          config: map() | list() | String.t() | number() | boolean(),
          recontact_days: integer(),
          link_survey_branding: boolean(),
          in_app_survey_branding: boolean(),
          placement: atom(),
          click_outside_close: boolean(),
          overlay: atom(),
          overlay_color: String.t() | nil,
          overlay_opacity: integer() | nil,
          logo: map() | list() | String.t() | number() | boolean() | nil,
          app_setup_completed: boolean(),
          custom_head_scripts: String.t() | nil,
          custom_css: map() | list() | String.t() | number() | boolean() | nil,
          custom_css_previous: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Formbricks.Organization.t() | nil,
          languages: [Formbricks.Language.t()],
          embedded_data: [Formbricks.EmbeddedData.t()],
          workspace_teams: [Formbricks.WorkspaceTeam.t()],
          feedback_directory_workspaces: [Formbricks.FeedbackDirectoryWorkspace.t()],
          charts: [Formbricks.Chart.t()],
          dashboards: [Formbricks.Dashboard.t()],
          embedded_field_labels: [Formbricks.EmbeddedFieldLabel.t()],
          creatives: [Formbricks.Creative.t()],
          attachment_exports: [Formbricks.AttachmentExport.t()],
          surveys: [Formbricks.Survey.t()],
          contacts: [Formbricks.Contact.t()],
          action_classes: [Formbricks.ActionClass.t()],
          contact_attribute_keys: [Formbricks.ContactAttributeKey.t()],
          webhooks: [Formbricks.Webhook.t()],
          tags: [Formbricks.Tag.t()],
          segments: [Formbricks.Segment.t()],
          integrations: [Formbricks.Integration.t()],
          api_key_workspaces: [Formbricks.ApiKeyWorkspace.t()],
          feedback_sources: [Formbricks.FeedbackSource.t()],
          workflows: [Formbricks.Workflow.t()],
          workflow_runs: [Formbricks.WorkflowRun.t()],
          mcp_grants: [Formbricks.McpGrantWorkspace.t()]
        }

  schema "Workspace" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:legacy_environment_id, :string, source: :legacyEnvironmentId)
    # [Styling]
    field(:styling, Formbricks.PrismaJson, default: %{"allowStyleOverwrite" => true})
    # [WorkspaceConfig]
    field(:config, Formbricks.PrismaJson, default: %{})
    field(:recontact_days, :integer, default: 7, source: :recontactDays)
    field(:link_survey_branding, :boolean, default: true, source: :linkSurveyBranding)
    field(:in_app_survey_branding, :boolean, default: true, source: :inAppSurveyBranding)
    field(:placement, Ecto.Enum, values: [:bottomLeft, :bottomRight, :topLeft, :topRight, :center], default: :bottomRight)
    field(:click_outside_close, :boolean, default: true, source: :clickOutsideClose)
    field(:overlay, Ecto.Enum, values: [:none, :light, :dark], default: :none)
    # Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`.
    field(:overlay_color, :string, source: :overlayColor)
    field(:overlay_opacity, :integer, source: :overlayOpacity)
    # [Logo]
    field(:logo, Formbricks.PrismaJson)
    field(:app_setup_completed, :boolean, default: false, source: :appSetupCompleted)
    field(:custom_head_scripts, :string, source: :customHeadScripts)
    # [CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service.
    field(:custom_css, Formbricks.PrismaJson, source: :customCss)
    # [CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
    # submitting its source again through the current processor and authorization checks.
    field(:custom_css_previous, Formbricks.PrismaJson, source: :customCssPrevious)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Formbricks.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:languages, Formbricks.Language, foreign_key: :workspace_id, references: :id)
    has_many(:embedded_data, Formbricks.EmbeddedData, foreign_key: :workspace_id, references: :id)
    has_many(:workspace_teams, Formbricks.WorkspaceTeam, foreign_key: :workspace_id, references: :id)
    has_many(:feedback_directory_workspaces, Formbricks.FeedbackDirectoryWorkspace, foreign_key: :workspace_id, references: :id)
    has_many(:charts, Formbricks.Chart, foreign_key: :workspace_id, references: :id)
    has_many(:dashboards, Formbricks.Dashboard, foreign_key: :workspace_id, references: :id)
    has_many(:embedded_field_labels, Formbricks.EmbeddedFieldLabel, foreign_key: :workspace_id, references: :id)
    has_many(:creatives, Formbricks.Creative, foreign_key: :workspace_id, references: :id)
    has_many(:attachment_exports, Formbricks.AttachmentExport, foreign_key: :workspace_id, references: :id)
    has_many(:surveys, Formbricks.Survey, foreign_key: :workspace_id, references: :id)
    has_many(:contacts, Formbricks.Contact, foreign_key: :workspace_id, references: :id)
    has_many(:action_classes, Formbricks.ActionClass, foreign_key: :workspace_id, references: :id)
    has_many(:contact_attribute_keys, Formbricks.ContactAttributeKey, foreign_key: :workspace_id, references: :id)
    has_many(:webhooks, Formbricks.Webhook, foreign_key: :workspace_id, references: :id)
    has_many(:tags, Formbricks.Tag, foreign_key: :workspace_id, references: :id)
    has_many(:segments, Formbricks.Segment, foreign_key: :workspace_id, references: :id)
    has_many(:integrations, Formbricks.Integration, foreign_key: :workspace_id, references: :id)
    has_many(:api_key_workspaces, Formbricks.ApiKeyWorkspace, foreign_key: :workspace_id, references: :id)
    has_many(:feedback_sources, Formbricks.FeedbackSource, foreign_key: :workspace_id, references: :id)
    has_many(:workflows, Formbricks.Workflow, foreign_key: :workspace_id, references: :id)
    has_many(:workflow_runs, Formbricks.WorkflowRun, foreign_key: :workspace_id, references: :id)
    has_many(:mcp_grants, Formbricks.McpGrantWorkspace, foreign_key: :workspace_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end