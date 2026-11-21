# Stores third-party authentication account information.
# Enables OAuth and other external authentication methods.
#
# @property id - Unique identifier for the account
# @property user - The Formbricks user who owns this account
# @property provider - The authentication provider (GitHub, Google, etc.)
# @property providerAccountId - User ID from the provider
# @property access_token - OAuth access token
# @property refresh_token - OAuth refresh token
#
# issuer:: Better Auth 1.7.0 (ENG-2343): the account's identity namespace. 1.7.0 keyed accounts on
#   (issuer, accountId) and filtered every lookup on it — findAccountByKey, findAccountOwnerByKey,
#   findCredentialAccount and updatePassword — so on 1.7.0 a null here locks the user out.
#   1.7.3 dropped the field and 1.7.7 keys accounts on (provider, providerAccountId) again
#   (ENG-3880). We keep writing it (see `account.create.before` in better-auth-hooks.ts) so a
#   rollback to 1.7.0 still finds every row; drop it once that rollback is off the table.
#
#   NULLABLE on purpose. A NOT NULL column would fail every INSERT from a 1.6 pod still serving
#   during a rolling deploy. The migration backfills every existing row; see the migration SQL for
#   the value derivation.
class Account < ApplicationRecord
  self.table_name = "Account"
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :provider, presence: true, uniqueness: { scope: :providerAccountId }
  validates :providerAccountId, presence: true
  validates :issuer, uniqueness: { scope: :providerAccountId, allow_nil: true }

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :accounts
end
