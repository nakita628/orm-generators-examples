defmodule LangWatch.WebhookEndpoint do
  @moduledoc """
  Org-anchored outbound webhook endpoint (the webhook platform, enterprise
  only). The signing secret is AES-256-GCM encrypted at rest (shared
  encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
  once, at create or roll time. `enabledEvents` holds registry event types
  or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
  72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          destination_kind: atom(),
          url: String.t() | nil,
          sqs_queue_url: String.t() | nil,
          sqs_role_arn: String.t() | nil,
          sqs_external_id: String.t() | nil,
          sqs_access_key_id: String.t() | nil,
          sqs_secret_access_key_encrypted: String.t() | nil,
          secret_encrypted: String.t(),
          signature_scheme: String.t() | nil,
          idempotency_key: String.t() | nil,
          previous_secret_encrypted: String.t() | nil,
          previous_secret_expires_at: DateTime.t() | nil,
          enabled_events: [String.t()],
          status: atom(),
          disabled_reason: String.t() | nil,
          disabled_at: DateTime.t() | nil,
          failing_since: DateTime.t() | nil,
          max_batch_size: integer(),
          max_batch_delay_ms: integer(),
          max_in_flight: integer(),
          allow_self_signed_certificate: boolean(),
          last_success_at: DateTime.t() | nil,
          last_failure_at: DateTime.t() | nil,
          archived_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil,
          deliveries: [LangWatch.WebhookEndpointDelivery.t()]
        }

  schema "WebhookEndpoint" do
    field(:id, :string, primary_key: true)
    # Where this endpoint delivers. Everything above the last hop is shared:
    # same batching, same retry ladder, same delivery log, same signature over
    # the same bytes.
    field(:destination_kind, Ecto.Enum, values: [:http, :sqs], default: :http, source: :destinationKind)
    # The receiver URL. Present on `http` endpoints and null on the others;
    # the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
    field(:url, :string)
    # The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
    # region and the owning account are read off it rather than configured
    # separately, so they can never disagree with the queue.
    field(:sqs_queue_url, :string, source: :sqsQueueUrl)
    # The role we assume to write to the queue, and the ExternalId that proves
    # the request came from us rather than from anyone who learned the role's
    # name. The recommended credential mode.
    field(:sqs_role_arn, :string, source: :sqsRoleArn)
    field(:sqs_external_id, :string, source: :sqsExternalId)
    # A static key pair for the queue, encrypted at rest exactly like the
    # signing secret. Both halves are present or both are absent: a half-filled
    # pair is what the AWS SDK reads as "we have credentials" and stops looking
    # with, so the CHECK constraint refuses it.
    field(:sqs_access_key_id, :string, source: :sqsAccessKeyId)
    field(:sqs_secret_access_key_encrypted, :string, source: :sqsSecretAccessKeyEncrypted)
    field(:secret_encrypted, :string, source: :secretEncrypted)
    # How deliveries are formatted and signed. Null: the batch envelope signed
    # `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
    # only for endpoints migrated from governance anomaly destinations.
    field(:signature_scheme, :string, source: :signatureScheme)
    # Set only by governance's anomaly destination migration: a create that
    # repeats it in the same organization answers the endpoint it made. An
    # archived endpoint gives its key up, so a retry after cleanup creates anew.
    field(:idempotency_key, :string, source: :idempotencyKey)
    # The secret this endpoint signed with before the last roll. Deliveries
    # carry a second `v1=` computed from it until it expires, so a receiver
    # swaps secrets on its own schedule instead of dropping deliveries in the
    # gap. Null once the window closes or if the secret was never rolled.
    field(:previous_secret_encrypted, :string, source: :previousSecretEncrypted)
    # When the previous secret stops being signed with and accepted.
    field(:previous_secret_expires_at, LangWatch.PrismaDateTime, source: :previousSecretExpiresAt)
    field(:enabled_events, {:array, :string}, source: :enabledEvents)
    field(:status, Ecto.Enum, values: [:ACTIVE, :DISABLED], default: :ACTIVE)
    # Why the endpoint is disabled: "auto_failures_72h" or "manual".
    field(:disabled_reason, :string, source: :disabledReason)
    field(:disabled_at, LangWatch.PrismaDateTime, source: :disabledAt)
    # Start of the current consecutive-failure streak; cleared on any 2xx.
    # 72 hours of unbroken failures from this instant auto-disables.
    field(:failing_since, LangWatch.PrismaDateTime, source: :failingSince)
    # Delivery controls, each server-bounded (see webhookDeliveryControls
    # constants): envelopes per POST, coalescing window before a partial
    # batch ships, and concurrent POSTs to this receiver.
    field(:max_batch_size, :integer, default: 100, source: :maxBatchSize)
    field(:max_batch_delay_ms, :integer, default: 250, source: :maxBatchDelayMs)
    field(:max_in_flight, :integer, default: 4, source: :maxInFlight)
    # Self-hosted only: deliveries accept the receiver's self-signed TLS
    # certificate. Off by default; the hosted product ignores it.
    field(:allow_self_signed_certificate, :boolean, default: false, source: :allowSelfSignedCertificate)
    field(:last_success_at, LangWatch.PrismaDateTime, source: :lastSuccessAt)
    field(:last_failure_at, LangWatch.PrismaDateTime, source: :lastFailureAt)
    field(:archived_at, LangWatch.PrismaDateTime, source: :archivedAt)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    has_many(:deliveries, LangWatch.WebhookEndpointDelivery, foreign_key: :endpoint_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end