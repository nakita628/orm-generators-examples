defmodule LangWatch.GithubInstallation do
  @moduledoc """
  The organization's GitHub connection: a GitHub App installation mapped to a
  LangWatch organization. Short-lived (1h) installation access tokens are
  minted against these rows on demand, by Langy for writes (bot-authored pull
  requests, scoped to the installation's repositories) and by pull-request
  linkage for reads. No user token and no secret are stored here, the crown
  jewel is the App private key, held only in the control-plane env.
  Spec: specs/integrations/github-connection.feature.

  installationId is globally unique (one GitHub installation, one row); a
  single LangWatch organization may own MANY installations (different GitHub
  accounts), so the org to installation relationship is one-to-many.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          installation_id: String.t(),
          account_login: String.t(),
          account_type: String.t(),
          account_id: String.t(),
          repository_selection: String.t(),
          repositories: map() | nil,
          suspended_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil
        }

  schema "GithubInstallation" do
    field(:id, :string, primary_key: true)
    # GitHub installation id (numeric; stored as string). Globally unique.
    field(:installation_id, :string, source: :installationId)
    # The GitHub account (organization or user) the App is installed on.
    field(:account_login, :string, source: :accountLogin)
    # "Organization" or "User".
    field(:account_type, :string, source: :accountType)
    # GitHub numeric account id (stable across renames).
    field(:account_id, :string, source: :accountId)
    # "all" or "selected" — whether the App can see every repo on the account
    # or only a chosen subset.
    field(:repository_selection, :string, source: :repositorySelection)
    # Cached repository list [{ id, fullName }] when repositorySelection =
    # "selected". Null when "all" (the full set is resolved live at mint time).
    field(:repositories, :map)
    # Set when GitHub suspends the installation; cleared on unsuspend. A
    # suspended installation cannot mint tokens.
    field(:suspended_at, LangWatch.PrismaDateTime, source: :suspendedAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end