# Per-(organization, tool) override for which `langwatch <tool>` paths the
# CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
# `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
# row resolves to the hardcoded defaults (both true for claude/codex/gemini/
# opencode; cursor allows the gateway path only), so the policy is purely
# additive: disabling a path is an explicit admin choice. The CLI caches the
# resolved map at login and gates path selection on it.
# See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
class PlatformToolPolicy < ApplicationRecord
  self.table_name = "PlatformToolPolicy"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :allowVk, default: true
  attribute :allowOtelDirect, default: true
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true, uniqueness: { scope: :toolSlug }
  validates :toolSlug, presence: true
end
