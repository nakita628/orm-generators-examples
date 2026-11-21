defmodule Shelf.UserBusinessIntel do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          how_did_you_hear_about_us: String.t() | nil,
          job_title: String.t() | nil,
          team_size: String.t() | nil,
          company_name: String.t() | nil,
          primary_use_case: String.t() | nil,
          current_solution: String.t() | nil,
          timeline: String.t() | nil,
          signup_plan: String.t() | nil,
          signup_trial: boolean() | nil,
          utm_source: String.t() | nil,
          utm_medium: String.t() | nil,
          utm_campaign: String.t() | nil,
          utm_content: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil
        }

  schema "UserBusinessIntel" do
    field(:id, :string, primary_key: true)
    field(:how_did_you_hear_about_us, :string, source: :howDidYouHearAboutUs)
    field(:job_title, :string, source: :jobTitle)
    field(:team_size, :string, source: :teamSize)
    field(:company_name, :string, source: :companyName)
    field(:primary_use_case, :string, source: :primaryUseCase)
    field(:current_solution, :string, source: :currentSolution)
    field(:timeline, :string)
    field(:signup_plan, :string, source: :signupPlan)
    field(:signup_trial, :boolean, source: :signupTrial)
    field(:utm_source, :string, source: :utmSource)
    field(:utm_medium, :string, source: :utmMedium)
    field(:utm_campaign, :string, source: :utmCampaign)
    field(:utm_content, :string, source: :utmContent)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end