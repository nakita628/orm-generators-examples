defmodule LangWatch.DashboardsV2 do
  @moduledoc """
  A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
  made before the split were copied here with their ids, so stars and insight
  pointers still name them. No relation, per the database rule.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          created_by_id: String.t() | nil,
          scope: atom(),
          organization_id: String.t() | nil,
          home_product: String.t(),
          excluded_origins: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "DashboardsV2" do
    field(:id, :string, primary_key: true)
    # The project that owns the board; only there can it be edited.
    field(:project_id, :string, source: :projectId)
    field(:name, :string)
    # Free text under the board's name; null until a member writes one.
    field(:description, :string)
    # The member who created it; null for boards that predate authorship.
    field(:created_by_id, :string, source: :createdById)
    # Who sees it. Only the author changes it; a board with no author stays PROJECT.
    field(:scope, Ecto.Enum, values: [:PRIVATE, :PROJECT, :ORGANIZATION], default: :PROJECT)
    # The owning project's organization, stamped when the board is first set to
    # ORGANIZATION; null before.
    field(:organization_id, :string, source: :organizationId)
    # The product whose sidebar lists it: llm-ops, me, gateway or governance (ADR-002,
    # decision 8). Every board made before homes existed is llm-ops.
    field(:home_product, :string, default: "llm-ops", source: :homeProduct)
    # Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
    # 2026-10-28); dropped in a later migration, once no serving release reads it.
    field(:excluded_origins, LangWatch.PrismaJson, source: :excludedOrigins)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end