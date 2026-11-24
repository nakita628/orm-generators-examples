defmodule LangWatch.ExperimentVersion do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          experiment_id: String.t(),
          project_id: String.t(),
          version: integer(),
          counter_version: integer(),
          auto_saved: boolean(),
          commit_message: String.t() | nil,
          author_id: String.t() | nil,
          author_label: String.t(),
          run_id: String.t() | nil,
          state: map() | list() | String.t() | number() | boolean(),
          schema_version: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "ExperimentVersion" do
    field(:id, :string, primary_key: true)
    field(:experiment_id, :string, source: :experimentId)
    field(:project_id, :string, source: :projectId)
    # The number the version is addressed and displayed by. Numbered rows hold
    # a gap-free sequence. The rolling autosave row rides on the counter
    # instead, which is what keeps it clear of that sequence.
    field(:version, :integer)
    # `Experiment.workbenchVersion` at the moment this row was last written.
    # The rolling row is rewritten in place, so this is the only column that
    # says how recent its content is, and the version list sorts and pages by
    # it. It also joins a row to the live state: the row whose `counterVersion`
    # equals the experiment's counter holds what the workbench shows now.
    field(:counter_version, :integer, source: :counterVersion)
    field(:auto_saved, :boolean, default: false, source: :autoSaved)
    field(:commit_message, :string, source: :commitMessage)
    field(:author_id, :string, source: :authorId)
    # Who wrote this version: "user", "langy" or "api".
    field(:author_label, :string, source: :authorLabel)
    # The run that wrote this version, when a run wrote it.
    #
    # A run writes its cells into the workbench state, which advances the
    # counter. Without this, the page that started that run reads the bump as
    # somebody else's write, refuses to save, and asks the reader to reload
    # over their own unsaved edits. Named here, the page adopts the version and
    # carries on.
    field(:run_id, :string, source: :runId)
    field(:state, LangWatch.PrismaJson)
    # Schema version of `state`, so a later shape change can be read forward.
    field(:schema_version, :string, default: "1", source: :schemaVersion)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end