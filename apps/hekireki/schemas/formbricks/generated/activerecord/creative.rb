# A feedback button or a banner. Opens at most one Destination: a survey (`surveyId`) or a link
# (`linkUrl`), never both.
#
# @property type - Fixed at creation.
# @property content - Per-type JSON (`FeedbackButtonContent` / `BannerContent` in the v3 spec), texts keyed by locale code.
# @property styling - Per-type JSON (`FeedbackButtonStyling` / `BannerStyling` in the v3 spec).
# @property survey - The Destination survey, in the same workspace by construction of the composite key.
# @property showConditions - Show when (`ZShowConditions`, shared with surveys); null always shows.
# @property autoSelectLanguage - The browser's language picks the texts' language when the site sets none; null (never chosen, as on Survey) is off.
class Creative < ApplicationRecord
  self.table_name = "Creative"
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  enum :type, { feedback_button: "feedbackButton", banner: "banner" }, validate: true
  enum :status, { draft: "draft", in_progress: "inProgress", paused: "paused", completed: "completed" }, default: :draft,
                                                                                                         validate: true
  enum :showOn, { page_view: "pageView", action: "action" }, default: :page_view, validate: true

  validates :name, presence: true
  validates :content, exclusion: { in: [ nil ] }
  validates :styling, exclusion: { in: [ nil ] }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :creatives
  belongs_to :creator, class_name: "User", foreign_key: "createdBy", optional: true, inverse_of: :creatives
  belongs_to :survey, foreign_key: [ :surveyId, :workspaceId ], primary_key: [ :id, :workspaceId ], optional: true,
                      inverse_of: :creatives
  has_many :triggers, class_name: "CreativeTrigger", foreign_key: "creativeId", inverse_of: :creative,
                      dependent: :destroy
  has_many :languages, class_name: "CreativeLanguage", foreign_key: "creativeId", inverse_of: :creative,
                       dependent: :destroy
  has_many :daily_stats, class_name: "CreativeDailyStat", foreign_key: "creativeId", inverse_of: :creative,
                         dependent: :destroy
  has_many :daily_page_stats, class_name: "CreativeDailyPageStat", foreign_key: "creativeId", inverse_of: :creative,
                              dependent: :destroy
end
