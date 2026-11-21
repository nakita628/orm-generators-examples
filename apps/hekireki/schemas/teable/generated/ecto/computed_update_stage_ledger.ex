defmodule Teable.ComputedUpdateStageLedger do
  @moduledoc """
  Durable per-stage state for budget-staged computed updates, keyed by the
  continuation chain's root task id (its scope). Rows are written once per
  record and never copied between continuation tasks:
  - kind 'excluded': processed-target exclusion ledger for partial batches;
  - kind 'frontier': seq-ordered queue of sources whose outgoing propagation
  is not finished (self-referential generations + migrated explicit seeds);
  - kind 'consumed': retired frontier sources preserved for deferred edge
  chunks, handed to the continuation as seeds at stage completion.
  Cleared when the chain's stage completes or the chain dead-letters.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          scope_id: String.t(),
          kind: String.t(),
          table_id: String.t(),
          record_id: String.t(),
          seq: integer()
        }

  schema "computed_update_stage_ledger" do
    field(:scope_id, :string, primary_key: true)
    field(:kind, :string, primary_key: true)
    field(:table_id, :string, primary_key: true)
    field(:record_id, :string, primary_key: true)
    field(:seq, :integer, default: 0)
  end
end