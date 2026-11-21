# originalUrl:: the original full url that we want to shorten e.g. https://www.url.com
# shortUrlSlug:: slug for the short url e.g. url.com/[shortUrlSlug]
# expiresAt:: the time the short url will expire
# clicks:: the number of times the short url has redirected to the original url
class ShortURL < ApplicationRecord
  self.table_name = "ShortURL"

  attribute :clicks, default: 0
  alias_attribute :created_at, :createdAt

  validates :originalUrl, presence: true
  validates :shortUrlSlug, presence: true, uniqueness: true

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :short_urls
end
