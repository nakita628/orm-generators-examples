package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Match candidates a background job computed and a human has yet to confirm
 * (ADR-128 §12).
 *
 * Stored rather than computed when the review screen asks, and the reason is
 * measured rather than stylistic: fuzzy matching is quadratic, there is no
 * database route to it here, and at this document's own example size - 2,000
 * discovered people against 500 accounts - plain edit distance over the
 * 1,000,000 pairs blocked the event loop for 2.9 seconds. Per page load,
 * uncached, stalling every other request on the instance. A stored row also
 * makes a pending count answerable without paying for the whole sweep.
 *
 * Rows are written only by the job, which prefilters before it scores and
 * recomputes when its inputs change. Confirming one opens an `IdentityMatch`
 * and deletes the suggestion.
 *
 * Suspension deliberately does NOT live here - it is on `DiscoveredPerson`,
 * because a halt on automatic linking that the next recompute clears is not a
 * halt.
 */
object IdentityMatchSuggestionTable : IdTable<String>("\"IdentityMatchSuggestion\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val discoveredPersonId = text("discoveredPersonId")

    /**
     * The platform user this might be. Not nullable: a suggestion with no
     * candidate is not a suggestion.
     */
    val userId = text("userId")

    /**
     * Similarity of the two name texts, after the prefilter, in [0, 1]. A
     * display hint and an ordering key - never a threshold anything auto-links
     * on, which is the whole distinction this table exists to hold.
     */
    val score = double("score")
    val computedAt = pgTimestamp("computedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"IdentityMatchSuggestion_pkey\"")

    init {
        index("\"IdentityMatchSuggestion_organizationId_score_idx\"", false, organizationId, score)
        uniqueIndex("IdentityMatchSuggestion_organizationId_discoveredPersonId_u_key", organizationId, discoveredPersonId, userId)
    }
}
