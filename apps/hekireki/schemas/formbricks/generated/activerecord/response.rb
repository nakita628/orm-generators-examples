# Stores a user's response to a survey, including their answers and metadata.
# Each response is linked to a specific survey and optionally to a contact.
#
# @property id - Unique identifier for the response
# @property finished - Whether the survey was completed
# @property survey - The associated survey
# @property contact - The optional contact who provided the response
# @property data - JSON object containing the actual response data
# @property variables - Custom variables used in the response
# @property ttc - Time to completion metrics
# @property meta - Additional metadata about the response
#
# data:: [ResponseData]
# variables:: [ResponseVariables]
# ttc:: [ResponseTtc]
# meta:: [ResponseMeta]
# ingest_flags:: Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
#   or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
#   from the client. Null means no ingest boundary has written it, which is not the same claim as
#   an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
#   and surfacing that in the UI (v2) would have nothing to read.
#   [ResponseIngestFlags]
# cal_bookings_by_question_id:: What the Cal.com elements reported when the respondent booked, `{ [Question ID]: booking }`. `data`
#   keeps the string "booked" for those questions, so this column is purely additive: it is omitted from
#   every read without a `select` (`PRISMA_GLOBAL_OMIT`), and only the webhook body (as
#   `calBookingsByQuestionId`) and `/api/v3/responses` read it. Server-derived from the incoming `data`,
#   never accepted as a field of its own.
#   [ResponseCalBookings]
# outcome:: A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
#   export by default and from Feedback Records, and nothing else: it is never deleted and every
#   count, list, webhook and integration still includes it (ENG-6780).
# not_shown:: What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
#   from them, and per element its options, matrix rows and columns, by stable ids. Written by the
#   survey while they answer; null means unknown (older responses and bundles), read as before.
#   Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
#   webhook or integration.
#   [ResponseNotShown]
# contactAttributes:: [ResponseContactAttributes]
class Response < ApplicationRecord
  self.table_name = "Response"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :finished, default: false
  attribute :data, default: -> { JSON.parse("{}") }
  attribute :variables, default: -> { JSON.parse("{}") }
  attribute :ttc, default: -> { JSON.parse("{}") }
  attribute :meta, default: -> { JSON.parse("{}") }

  enum :outcome, { low_quality: "lowQuality" }, validate: { allow_nil: true }

  validates :singleUseId, uniqueness: { scope: :surveyId, allow_nil: true }
  validates :displayId, uniqueness: { allow_nil: true }

  belongs_to :survey, foreign_key: "surveyId", inverse_of: :responses
  belongs_to :contact, foreign_key: "contactId", optional: true, inverse_of: :responses
  belongs_to :display, foreign_key: "displayId", optional: true, inverse_of: :response
  has_many :tags, class_name: "TagsOnResponses", foreign_key: "responseId", inverse_of: :response, dependent: :destroy
  has_many :quota_links, class_name: "ResponseQuotaLink", foreign_key: "responseId", inverse_of: :response,
                         dependent: :destroy
  has_many :workflow_runs, foreign_key: "responseId", inverse_of: :response, dependent: :destroy
end
