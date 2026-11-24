defmodule LangWatch.Session do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          session_token: String.t(),
          expires: DateTime.t(),
          ip_address: String.t() | nil,
          user_agent: String.t() | nil,
          identifier_id: String.t() | nil,
          amr: [String.t()],
          started_by_sign_up: boolean(),
          impersonating: map() | list() | String.t() | number() | boolean() | nil,
          actor_user_id: String.t() | nil,
          subject_user_id: String.t() | nil,
          impersonation_reason: String.t() | nil,
          impersonation_expires_at: DateTime.t() | nil,
          last_seen_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: LangWatch.User.t() | nil
        }

  schema "Session" do
    field(:id, :string, primary_key: true)
    field(:session_token, :string, source: :sessionToken)
    field(:expires, LangWatch.PrismaDateTime)
    field(:ip_address, :string, source: :ipAddress)
    field(:user_agent, :string, source: :userAgent)
    # Which of the user's sign-in methods minted this session (D06). NULL on
    # every session that predates the column and on any mint we could not
    # attribute - an ordinary value, not an error. Nothing reads it to decide
    # whether the session is valid, so no session was ended to add it.
    field(:identifier_id, :string, source: :identifierId)
    # What the sign-in that minted this session PROVED, as RFC 8176 method
    # references (D06). Empty on every session minted before the column, and
    # on any ordinary password sign-in. An organization that requires a second
    # factor reads this when a member reaches its data; it is never read at
    # mint and never ends a session.
    field(:amr, {:array, :string})
    # Minted while the account had never signed in: the session its sign-up
    # started, which skips the secure-account offer. False on every later one
    # and on every session that predates the column (specs/identity/passkeys.feature).
    field(:started_by_sign_up, :boolean, default: false, source: :startedBySignUp)
    # The legacy impersonation payload. NOTHING READS OR WRITES IT any more -
    # the {actor, subject} columns below replaced it, and the migration beside
    # them ended every session that still carried a value.
    #
    # It stays for ONE release, and that is the whole reason it is still here.
    # `prisma migrate deploy` runs at container start, so the first pod of a
    # new release drops a column while the previous release's pods are still
    # serving - and those pods select this column on EVERY authenticated
    # request and fail closed when the row does not come back. Dropping it in
    # the same release that stops reading it signs everybody out for the length
    # of the rollout. Expand now, contract next release.
    field(:impersonating, LangWatch.PrismaJson)
    # The impersonation claims (D06), aligned with the authz Principal
    # {actor, subject}. The actor is the operator who really made the request
    # and the subject is the person whose access they are borrowing; both are
    # recorded on every authorization decision, so the audit trail can answer
    # who really did it. These replace the legacy `impersonating` JSON blob,
    # which is dropped in the release AFTER this one.
    #
    # NULL on every ordinary session, which is all of them: an impersonation
    # is a handful of rows at a time and nothing else reads these columns.
    field(:actor_user_id, :string, source: :actorUserId)
    field(:subject_user_id, :string, source: :subjectUserId)
    # Why the operator started it. Required to start one
    # (specs/features/backoffice-user-impersonation-reason.feature) and kept
    # beside both people rather than only in the audit log, so a request under
    # impersonation carries its own justification.
    field(:impersonation_reason, :string, source: :impersonationReason)
    # When the borrowed access lapses. A session outlives its impersonation
    # window: past this the claims are ignored and the operator is themselves
    # again, with nothing ended.
    field(:impersonation_expires_at, LangWatch.PrismaDateTime, source: :impersonationExpiresAt)
    # When this session was last used, for the idle timeout an organization
    # may set (GAC-10, specs/identity/org-session-lifetime.feature).
    #
    # A COLUMN OF OUR OWN rather than `updatedAt`, because better-auth rolls a
    # live session's expiry once per `updateAge` - a day - so `updatedAt` is
    # activity to the nearest day and cannot tell an idle hour from an idle
    # minute. Written only for sessions whose owner belongs to an organization
    # that has set a window, and then at most once per quarter-window, so a
    # deployment where nobody has turned this on writes it never.
    #
    # NULL on every session minted before the column and on every session
    # under no window at all. A null READS AS `updatedAt`, which better-auth
    # keeps rolled to within a day of the last use. That fallback is
    # conservative in the direction that matters - `updatedAt` is never
    # earlier than the real last use, so it can only ever judge a session MORE
    # active than it was, never less - and it is what lets a window turned on
    # today end the sessions that have been idle since February instead of
    # waiting a day for a stamp they have never had.
    field(:last_seen_at, LangWatch.PrismaDateTime, source: :lastSeenAt)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, LangWatch.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end