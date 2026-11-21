defmodule LangWatch.SelfHostedInstanceDiagnostic do
  @moduledoc """
  One upgrade diagnostics bundle a self-hosted install sent over Connect, on
  LangWatch Cloud. The bytes live in object storage under
  `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
  points at them and feeds Cloud admin and the Slack signal. Plain column
  refs, no relation, as the usage report rows.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          instance_id: String.t(),
          step_id: String.t(),
          step_description: String.t() | nil,
          release: String.t() | nil,
          attempt: integer(),
          trigger: String.t(),
          byte_length: integer(),
          object_key: String.t(),
          received_at: DateTime.t(),
          expires_at: DateTime.t(),
          slack_announced_at: DateTime.t() | nil
        }

  schema "SelfHostedInstanceDiagnostic" do
    field(:id, :string, primary_key: true)
    # The `SelfHostedInstance.instanceId` the licence caller presented.
    field(:instance_id, :string, source: :instanceId)
    field(:step_id, :string, source: :stepId)
    field(:step_description, :string, source: :stepDescription)
    field(:release, :string)
    field(:attempt, :integer)
    # `operator` or `automatic`.
    field(:trigger, :string)
    field(:byte_length, :integer, source: :byteLength)
    field(:object_key, :string, source: :objectKey)
    field(:received_at, LangWatch.PrismaDateTime, autogenerate: true, source: :receivedAt)
    # When the lifecycle rule removes the bundle; the row is listed until then.
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    # Set on the report that posted this step and release's one Slack message.
    field(:slack_announced_at, LangWatch.PrismaDateTime, source: :slackAnnouncedAt)
  end
end