package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** Better Auth OAuth user consent grants for delegated clients. */
class OauthConsentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OauthConsentEntity>(OauthConsentTable)

    var clientId by OauthConsentTable.clientId
    var client by OauthClientEntity referencedOn OauthConsentTable.clientId
    var userId by OauthConsentTable.userId
    var user by UserEntity optionalReferencedOn OauthConsentTable.userId
    var referenceId by OauthConsentTable.referenceId
    var scopes by OauthConsentTable.scopes
    var createdAt by OauthConsentTable.createdAt
    var updatedAt by OauthConsentTable.updatedAt

    /**
     * Better Auth 1.7 (ENG-2343). /authorize compares the requested resources against these; without
     * them every pre-existing connection is forced back through the consent screen.
     */
    var resources by OauthConsentTable.resources
    var requestedUserInfoClaims by OauthConsentTable.requestedUserInfoClaims
}
