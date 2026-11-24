defmodule TypeHero.Challenge do
  use Ecto.Schema

  @primary_key {:id, :id, autogenerate: true}
  @timestamps_opts [type: TypeHero.PrismaDateTime, autogenerate: {TypeHero.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: integer(),
          difficulty: atom(),
          name: String.t(),
          slug: String.t(),
          description: String.t(),
          short_description: String.t(),
          code: String.t(),
          tests: String.t(),
          status: atom(),
          tsconfig: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: TypeHero.User.t() | nil,
          vote: [TypeHero.Vote.t()],
          submission: [TypeHero.Submission.t()],
          comment: [TypeHero.Comment.t()],
          bookmark: [TypeHero.Bookmark.t()],
          shared_solution: [TypeHero.SharedSolution.t()],
          report: [TypeHero.Report.t()],
          track_challenge: [TypeHero.TrackChallenge.t()],
          notification: [TypeHero.Notification.t()],
          tags: [TypeHero.Tag.t()]
        }

  schema "Challenge" do
    field(:difficulty, Ecto.Enum, values: [:BEGINNER, :EASY, :MEDIUM, :HARD, :EXTREME, :EVENT])
    field(:name, :string)
    field(:slug, :string)
    # this is the description to describe the challenge on description panel
    field(:description, :string)
    # this is a short blue thats shown on the explore card
    field(:short_description, :string, source: :shortDescription)
    # Starter code for the challenge
    field(:code, :string)
    # Test cases that need to be passed in order have the submission be accepted.
    field(:tests, :string)
    field(:status, Ecto.Enum, values: [:DECLINED, :BANNED, :PENDING, :ACTIVE], default: :PENDING)
    field(:tsconfig, TypeHero.PrismaJson)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, TypeHero.User, foreign_key: :user_id, define_field: false, type: :string)
    has_many(:vote, TypeHero.Vote, foreign_key: :challenge_id)
    has_many(:submission, TypeHero.Submission, foreign_key: :challenge_id)
    has_many(:comment, TypeHero.Comment, foreign_key: :root_challenge_id)
    has_many(:bookmark, TypeHero.Bookmark, foreign_key: :challenge_id)
    has_many(:shared_solution, TypeHero.SharedSolution, foreign_key: :challenge_id)
    has_many(:report, TypeHero.Report, foreign_key: :challenge_id)
    has_many(:track_challenge, TypeHero.TrackChallenge, foreign_key: :challenge_id)
    has_many(:notification, TypeHero.Notification, foreign_key: :challenge_id)
    many_to_many(:tags, TypeHero.Tag, join_through: "_ChallengeToTag", join_keys: [A: :id, B: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end