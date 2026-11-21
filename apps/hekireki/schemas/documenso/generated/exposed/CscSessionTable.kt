package documenso.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Per-recipient transient sign-time state held across the credential-scope
 * OAuth round-trip. The `recipientId @unique` constraint caps each recipient
 * to at most one in-flight session — re-clicking Sign UPSERTs the row.
 * Cleaned up transitively when the Recipient is deleted.
 * @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
 */
object CscSessionTable : IdTable<String>("\"CscSession\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()

    /**
     * Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
     * flows through the Recipient cascade.
     */
    val envelopeId = text("envelopeId")

    /**
     * Pinned at prep time; stable across overlay re-renders so the captured
     * signedAttrs digest matches the bytes the TSP signs.
     */
    val signingTime = pgTimestamp("signingTime", 3)

    /**
     * Contract between prep and sign: an array of
     * `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
     * `documentDataId` pins the immutable bytes that produced the captured
     * digest so sign-time can re-derive the same hash against the same content.
     * [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
     */
    val itemsJson = jsonb<String>("itemsJson", { it }, { it })

    /**
     * Populated by the credential-scope OAuth callback. Both nullable at row
     * creation; set together when the SAD is exchanged.
     */
    val encryptedSad = binary("encryptedSad").nullable()
    val sadExpiresAt = pgTimestamp("sadExpiresAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val recipientId =
        reference(
            "recipientId",
            RecipientTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CscSession_recipientId_fkey\"",
        ).uniqueIndex("\"CscSession_recipientId_key\"")
    override val primaryKey = PrimaryKey(id, name = "\"CscSession_pkey\"")
}
