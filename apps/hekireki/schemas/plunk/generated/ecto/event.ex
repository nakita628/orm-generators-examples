defmodule Plunk.Event do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Plunk.PrismaDateTime, autogenerate: {Plunk.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          data: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          project: Plunk.Project.t() | nil,
          contact: Plunk.Contact.t() | nil,
          email: Plunk.Email.t() | nil
        }

  schema "events" do
    field(:name, :string)
    field(:data, Plunk.PrismaJson)
    field(:project_id, :string, source: :projectId)
    field(:contact_id, :string, source: :contactId)
    field(:email_id, :string, source: :emailId)
    belongs_to(:project, Plunk.Project, foreign_key: :project_id, define_field: false)
    belongs_to(:contact, Plunk.Contact, foreign_key: :contact_id, define_field: false)
    belongs_to(:email, Plunk.Email, foreign_key: :email_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end