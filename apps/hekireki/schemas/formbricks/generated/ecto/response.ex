defmodule Formbricks.Response do
  @moduledoc """
  Stores a user's response to a survey, including their answers and metadata.
  Each response is linked to a specific survey and optionally to a contact.

  @property id - Unique identifier for the response
  @property finished - Whether the survey was completed
  @property survey - The associated survey
  @property contact - The optional contact who provided the response
  @property data - JSON object containing the actual response data
  @property variables - Custom variables used in the response
  @property ttc - Time to completion metrics
  @property meta - Additional metadata about the response
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          finished: boolean(),
          ending_id: String.t() | nil,
          data: map() | list() | String.t() | number() | boolean(),
          variables: map() | list() | String.t() | number() | boolean(),
          ttc: map() | list() | String.t() | number() | boolean(),
          meta: map() | list() | String.t() | number() | boolean(),
          ingest_flags: map() | list() | String.t() | number() | boolean() | nil,
          outcome: atom() | nil,
          not_shown: map() | list() | String.t() | number() | boolean() | nil,
          contact_attributes: map() | list() | String.t() | number() | boolean() | nil,
          single_use_id: String.t() | nil,
          language: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          survey: Formbricks.Survey.t() | nil,
          contact: Formbricks.Contact.t() | nil,
          display: Formbricks.Display.t() | nil,
          tags: [Formbricks.TagsOnResponses.t()],
          quota_links: [Formbricks.ResponseQuotaLink.t()],
          workflow_runs: [Formbricks.WorkflowRun.t()]
        }

  schema "Response" do
    field(:id, :string, primary_key: true)
    field(:finished, :boolean, default: false)
    field(:ending_id, :string, source: :endingId)
    # [ResponseData]
    field(:data, Formbricks.PrismaJson, default: %{})
    # [ResponseVariables]
    field(:variables, Formbricks.PrismaJson, default: %{})
    # [ResponseTtc]
    field(:ttc, Formbricks.PrismaJson, default: %{})
    # [ResponseMeta]
    field(:meta, Formbricks.PrismaJson, default: %{})
    # Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
    # or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
    # from the client. Null means no ingest boundary has written it, which is not the same claim as
    # an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
    # and surfacing that in the UI (v2) would have nothing to read.
    # [ResponseIngestFlags]
    field(:ingest_flags, Formbricks.PrismaJson)
    # A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
    # export by default and from Feedback Records, and nothing else: it is never deleted and every
    # count, list, webhook and integration still includes it (ENG-6780).
    field(:outcome, Ecto.Enum, values: [:lowQuality])
    # What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
    # from them, and per element its options, matrix rows and columns, by stable ids. Written by the
    # survey while they answer; null means unknown (older responses and bundles), read as before.
    # Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
    # webhook or integration.
    # [ResponseNotShown]
    field(:not_shown, Formbricks.PrismaJson)
    # [ResponseContactAttributes]
    field(:contact_attributes, Formbricks.PrismaJson, source: :contactAttributes)
    field(:single_use_id, :string, source: :singleUseId)
    field(:language, :string)
    field(:survey_id, :string, source: :surveyId)
    field(:contact_id, :string, source: :contactId)
    field(:display_id, :string, source: :displayId)
    belongs_to(:survey, Formbricks.Survey, foreign_key: :survey_id, define_field: false, type: :string)
    belongs_to(:contact, Formbricks.Contact, foreign_key: :contact_id, define_field: false, type: :string)
    belongs_to(:display, Formbricks.Display, foreign_key: :display_id, define_field: false, type: :string)
    has_many(:tags, Formbricks.TagsOnResponses, foreign_key: :response_id, references: :id)
    has_many(:quota_links, Formbricks.ResponseQuotaLink, foreign_key: :response_id, references: :id)
    has_many(:workflow_runs, Formbricks.WorkflowRun, foreign_key: :response_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end