defmodule LangWatch.AiToolEntryDepartment do
  @moduledoc """
  Per-department scope binding for AiToolEntry. Empty set = org-wide.
  Non-empty = entry visible only to members of those departments (an
  OrganizationUser whose `departmentId` is in the set). Mirrors
  AiToolEntryTeam's shape. `departmentId` is a plain column ref
  (relationMode="prisma"), matching how Department is referenced
  elsewhere; the service validates the id belongs to the org.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          department_id: String.t(),
          entry: LangWatch.AiToolEntry.t() | nil
        }

  schema "AiToolEntryDepartment" do
    field(:id, :string, primary_key: true)
    field(:department_id, :string, source: :departmentId)
    field(:entry_id, :string, source: :entryId)
    belongs_to(:entry, LangWatch.AiToolEntry, foreign_key: :entry_id, define_field: false, type: :string)
  end
end