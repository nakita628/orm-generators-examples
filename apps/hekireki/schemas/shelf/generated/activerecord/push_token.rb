# One Expo push token per installed companion app. A token identifies a
# device, not a person: when someone else signs in on the same phone, the
# row moves to them. Pushes go only to rows where `disabledAt` is null.
# @see apps/webapp/app/modules/push/service.server.ts
class PushToken < ApplicationRecord
  self.table_name = "PushToken"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :platform, { ios: "IOS", android: "ANDROID" }, validate: true

  validates :token, presence: true, uniqueness: true
  validates :lastSeenAt, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :push_tokens
end
