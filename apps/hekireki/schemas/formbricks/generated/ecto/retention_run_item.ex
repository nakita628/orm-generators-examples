defmodule Formbricks.RetentionRunItem do
  @moduledoc """
  What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
  never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
  `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          target_type: atom(),
          target_id: String.t(),
          target_name: String.t() | nil,
          action: atom(),
          count: integer(),
          recipient: String.t() | nil,
          skip_reason: atom() | nil,
          run: Formbricks.RetentionRun.t() | nil
        }

  schema "RetentionRunItem" do
    field(:id, :string, primary_key: true)
    field(:target_type, Ecto.Enum, values: [:survey, :user], source: :targetType)
    field(:target_id, :string, source: :targetId)
    field(:target_name, :string, source: :targetName)
    field(:action, Ecto.Enum, values: [:notified, :archived, :deactivated, :deleted, :skipped])
    # 1, except for grouped response deletions, where it is the number of responses deleted.
    field(:count, :integer, default: 1)
    # Who was emailed, for `notified` rows.
    field(:recipient, :string)
    field(:skip_reason, Ecto.Enum, values: [:exempt, :lastOwner, :otherOrganization, :noRecipient], source: :skipReason)
    field(:run_id, :string, source: :runId)
    belongs_to(:run, Formbricks.RetentionRun, foreign_key: :run_id, define_field: false, type: :string)
  end
end