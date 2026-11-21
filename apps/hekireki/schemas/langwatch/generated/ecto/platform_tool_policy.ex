defmodule LangWatch.PlatformToolPolicy do
  @moduledoc """
  Per-(organization, tool) override for which `langwatch <tool>` paths the
  CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
  `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
  row resolves to the hardcoded defaults (both true for claude/codex/gemini/
  opencode; cursor allows the gateway path only), so the policy is purely
  additive: disabling a path is an explicit admin choice. The CLI caches the
  resolved map at login and gates path selection on it.
  See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          tool_slug: String.t(),
          allow_vk: boolean(),
          allow_otel_direct: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "PlatformToolPolicy" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:tool_slug, :string, source: :toolSlug)
    field(:allow_vk, :boolean, default: true, source: :allowVk)
    field(:allow_otel_direct, :boolean, default: true, source: :allowOtelDirect)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end