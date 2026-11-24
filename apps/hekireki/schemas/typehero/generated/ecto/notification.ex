defmodule TypeHero.Notification do
  use Ecto.Schema

  @primary_key {:id, :id, autogenerate: true}
  @timestamps_opts [type: TypeHero.PrismaDateTime, autogenerate: {TypeHero.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: integer(),
          is_read: boolean(),
          type: atom(),
          content: String.t() | nil,
          blurb: String.t() | nil,
          url: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          to_user: TypeHero.User.t() | nil,
          from_user: TypeHero.User.t() | nil,
          comment: TypeHero.Comment.t() | nil,
          challenge: TypeHero.Challenge.t() | nil,
          shared_solution: TypeHero.SharedSolution.t() | nil
        }

  schema "Notification" do
    field(:is_read, :boolean, default: false, source: :isRead)
    field(:type, Ecto.Enum, values: [:LIKE, :REPLY, :MENTION, :MISC, :SYSTEM])
    # could be comment text or solution title to avoid having to join to grab this data
    field(:content, :string)
    # can be used to replace the blurb (eg: mention you in a comment)
    field(:blurb, :string)
    # use this instead of building the url on the fly (avoids a lot of joins and perf hits to build a string)
    field(:url, :string)
    field(:to_user_id, :string, source: :toUserId)
    field(:from_user_id, :string, source: :fromUserId)
    field(:comment_id, :id, source: :commentId)
    field(:challenge_id, :id, source: :challengeId)
    field(:shared_solution_id, :id, source: :sharedSolutionId)
    belongs_to(:to_user, TypeHero.User, foreign_key: :to_user_id, define_field: false, type: :string)
    belongs_to(:from_user, TypeHero.User, foreign_key: :from_user_id, define_field: false, type: :string)
    belongs_to(:comment, TypeHero.Comment, foreign_key: :comment_id, define_field: false)
    belongs_to(:challenge, TypeHero.Challenge, foreign_key: :challenge_id, define_field: false)
    belongs_to(:shared_solution, TypeHero.SharedSolution, foreign_key: :shared_solution_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end