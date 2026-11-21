# Org-anchored outbound webhook endpoint (the webhook platform, enterprise
# only). The signing secret is AES-256-GCM encrypted at rest (shared
# encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
# once, at create or roll time. `enabledEvents` holds registry event types
# or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
# 72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
#
# destinationKind:: Where this endpoint delivers. Everything above the last hop is shared:
#   same batching, same retry ladder, same delivery log, same signature over
#   the same bytes.
# url:: The receiver URL. Present on `http` endpoints and null on the others;
#   the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
# sqsQueueUrl:: The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
#   region and the owning account are read off it rather than configured
#   separately, so they can never disagree with the queue.
# sqsRoleArn:: The role we assume to write to the queue, and the ExternalId that proves
#   the request came from us rather than from anyone who learned the role's
#   name. The recommended credential mode.
# sqsAccessKeyId:: A static key pair for the queue, encrypted at rest exactly like the
#   signing secret. Both halves are present or both are absent: a half-filled
#   pair is what the AWS SDK reads as "we have credentials" and stops looking
#   with, so the CHECK constraint refuses it.
# signatureScheme:: How deliveries are formatted and signed. Null: the batch envelope signed
#   `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
#   only for endpoints migrated from governance anomaly destinations.
# idempotencyKey:: Set only by governance's anomaly destination migration: a create that
#   repeats it in the same organization answers the endpoint it made. An
#   archived endpoint gives its key up, so a retry after cleanup creates anew.
# previousSecretEncrypted:: The secret this endpoint signed with before the last roll. Deliveries
#   carry a second `v1=` computed from it until it expires, so a receiver
#   swaps secrets on its own schedule instead of dropping deliveries in the
#   gap. Null once the window closes or if the secret was never rolled.
# previousSecretExpiresAt:: When the previous secret stops being signed with and accepted.
# disabledReason:: Why the endpoint is disabled: "auto_failures_72h" or "manual".
# failingSince:: Start of the current consecutive-failure streak; cleared on any 2xx.
#   72 hours of unbroken failures from this instant auto-disables.
# maxBatchSize:: Delivery controls, each server-bounded (see webhookDeliveryControls
#   constants): envelopes per POST, coalescing window before a partial
#   batch ships, and concurrent POSTs to this receiver.
# allowSelfSignedCertificate:: Self-hosted only: deliveries accept the receiver's self-signed TLS
#   certificate. Off by default; the hosted product ignores it.
class WebhookEndpoint < ApplicationRecord
  self.table_name = "WebhookEndpoint"

  attribute :enabledEvents, default: -> { [] }
  attribute :maxBatchSize, default: 100
  attribute :maxBatchDelayMs, default: 250
  attribute :maxInFlight, default: 4
  attribute :allowSelfSignedCertificate, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :destinationKind, { http: "http", sqs: "sqs" }, default: :http, validate: true
  enum :status, { active: "ACTIVE", disabled: "DISABLED" }, default: :active, validate: true

  validates :secretEncrypted, presence: true
  validates :idempotencyKey, uniqueness: { scope: :organizationId, allow_nil: true }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :webhook_endpoints
  has_many :deliveries, class_name: "WebhookEndpointDelivery", foreign_key: "endpointId", inverse_of: :endpoint,
                        dependent: :destroy
end
