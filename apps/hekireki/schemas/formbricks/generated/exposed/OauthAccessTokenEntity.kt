package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Better Auth opaque access token records. The token column stores Better Auth's hashed lookup value.
 * JWT access tokens are verified through JWKS and bounded by the configured short lifetime.
 */
class OauthAccessTokenEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OauthAccessTokenEntity>(OauthAccessTokenTable)

    var token by OauthAccessTokenTable.token
    var clientId by OauthAccessTokenTable.clientId
    var client by OauthClientEntity referencedOn OauthAccessTokenTable.clientId
    var sessionId by OauthAccessTokenTable.sessionId
    var session by SessionEntity optionalReferencedOn OauthAccessTokenTable.sessionId
    var userId by OauthAccessTokenTable.userId
    var user by UserEntity optionalReferencedOn OauthAccessTokenTable.userId
    var referenceId by OauthAccessTokenTable.referenceId
    var refreshId by OauthAccessTokenTable.refreshId
    var refresh by OauthRefreshTokenEntity optionalReferencedOn OauthAccessTokenTable.refreshId
    var expiresAt by OauthAccessTokenTable.expiresAt
    var createdAt by OauthAccessTokenTable.createdAt
    var scopes by OauthAccessTokenTable.scopes

    /**
     * Better Auth 1.7 (ENG-2343). `resources` is the RFC 8707 resource set the grant approved — the
     * binding this whole upgrade exists to get (GHSA-p2fr-6hmx-4528).
     */
    var authorizationCodeId by OauthAccessTokenTable.authorizationCodeId
    var resources by OauthAccessTokenTable.resources
    var requestedUserInfoClaims by OauthAccessTokenTable.requestedUserInfoClaims
    var revoked by OauthAccessTokenTable.revoked
    var confirmation by OauthAccessTokenTable.confirmation
}
