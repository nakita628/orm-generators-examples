# The organization's GitHub connection: a GitHub App installation mapped to a
# LangWatch organization. Short-lived (1h) installation access tokens are
# minted against these rows on demand, by Langy for writes (bot-authored pull
# requests, scoped to the installation's repositories) and by pull-request
# linkage for reads. No user token and no secret are stored here, the crown
# jewel is the App private key, held only in the control-plane env.
# Spec: specs/integrations/github-connection.feature.
#
# installationId is globally unique (one GitHub installation, one row); a
# single LangWatch organization may own MANY installations (different GitHub
# accounts), so the org to installation relationship is one-to-many.
#
# installationId:: GitHub installation id (numeric; stored as string). Globally unique.
# accountLogin:: The GitHub account (organization or user) the App is installed on.
# accountType:: "Organization" or "User".
# accountId:: GitHub numeric account id (stable across renames).
# repositorySelection:: "all" or "selected" — whether the App can see every repo on the account
#   or only a chosen subset.
# repositories:: Cached repository list [{ id, fullName }] when repositorySelection =
#   "selected". Null when "all" (the full set is resolved live at mint time).
# suspendedAt:: Set when GitHub suspends the installation; cleared on unsuspend. A
#   suspended installation cannot mint tokens.
class GithubInstallation < ApplicationRecord
  self.table_name = "GithubInstallation"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :installationId, presence: true, uniqueness: true
  validates :accountLogin, presence: true
  validates :accountType, presence: true
  validates :accountId, presence: true
  validates :repositorySelection, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :github_installations
end
