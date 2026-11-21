# matchAll:: Whether every filter must match, or any one of them ("Show mail matching
#   all/any of these conditions"). A split with no filters is the plain inbox.
# bundleInInbox:: Shown as one collapsed row in the inbox when the account bundles its inbox.
class MailSplit < ApplicationRecord
  self.table_name = "MailSplit"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :matchAll, default: true
  attribute :bundleInInbox, default: false
  attribute :hideFromAll, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true, uniqueness: { scope: :emailAccountId }
  validates :order, presence: true

  belongs_to :email_account, foreign_key: "emailAccountId", inverse_of: :mail_splits
  has_many :filters, class_name: "MailSplitFilter", foreign_key: "mailSplitId", inverse_of: :mail_split,
                     dependent: :destroy
end
