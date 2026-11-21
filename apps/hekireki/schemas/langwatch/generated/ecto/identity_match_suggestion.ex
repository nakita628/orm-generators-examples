defmodule LangWatch.IdentityMatchSuggestion do
  @moduledoc """
  Match candidates a background job computed and a human has yet to confirm
  (ADR-128 §12).

  Stored rather than computed when the review screen asks, and the reason is
  measured rather than stylistic: fuzzy matching is quadratic, there is no
  database route to it here, and at this document's own example size - 2,000
  discovered people against 500 accounts - plain edit distance over the
  1,000,000 pairs blocked the event loop for 2.9 seconds. Per page load,
  uncached, stalling every other request on the instance. A stored row also
  makes a pending count answerable without paying for the whole sweep.

  Rows are written only by the job, which prefilters before it scores and
  recomputes when its inputs change. Confirming one opens an `IdentityMatch`
  and deletes the suggestion.

  Suspension deliberately does NOT live here - it is on `DiscoveredPerson`,
  because a halt on automatic linking that the next recompute clears is not a
  halt.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          discovered_person_id: String.t(),
          user_id: String.t(),
          score: float(),
          computed_at: DateTime.t()
        }

  schema "IdentityMatchSuggestion" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:discovered_person_id, :string, source: :discoveredPersonId)
    # The platform user this might be. Not nullable: a suggestion with no
    # candidate is not a suggestion.
    field(:user_id, :string, source: :userId)
    # Similarity of the two name texts, after the prefilter, in [0, 1]. A
    # display hint and an ordering key - never a threshold anything auto-links
    # on, which is the whole distinction this table exists to hold.
    field(:score, :float)
    field(:computed_at, LangWatch.PrismaDateTime, source: :computedAt)
  end
end