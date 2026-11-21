package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Per-recipient transient sign-time state held across the credential-scope
 * OAuth round-trip. The `recipientId @unique` constraint caps each recipient
 * to at most one in-flight session — re-clicking Sign UPSERTs the row.
 * Cleaned up transitively when the Recipient is deleted.
 * @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
 */
class CscSessionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CscSessionEntity>(CscSessionTable)

    /**
     * Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
     * flows through the Recipient cascade.
     */
    var envelopeId by CscSessionTable.envelopeId

    /**
     * Pinned at prep time; stable across overlay re-renders so the captured
     * signedAttrs digest matches the bytes the TSP signs.
     */
    var signingTime by CscSessionTable.signingTime

    /**
     * Contract between prep and sign: an array of
     * `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
     * `documentDataId` pins the immutable bytes that produced the captured
     * digest so sign-time can re-derive the same hash against the same content.
     * [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
     */
    var itemsJson by CscSessionTable.itemsJson

    /**
     * Populated by the credential-scope OAuth callback. Both nullable at row
     * creation; set together when the SAD is exchanged.
     */
    var encryptedSad by CscSessionTable.encryptedSad
    var sadExpiresAt by CscSessionTable.sadExpiresAt
    var createdAt by CscSessionTable.createdAt
    var recipientId by CscSessionTable.recipientId
    var recipient by RecipientEntity referencedOn CscSessionTable.recipientId
}
