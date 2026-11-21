package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
 * its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
 * a token is now bound to the resource it was approved for instead of to a static allow-list.
 * `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
 */
class OauthResourceEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OauthResourceEntity>(OauthResourceTable)

    var identifier by OauthResourceTable.identifier
    var name by OauthResourceTable.name
    var accessTokenTtl by OauthResourceTable.accessTokenTtl
    var refreshTokenTtl by OauthResourceTable.refreshTokenTtl
    var signingAlgorithm by OauthResourceTable.signingAlgorithm
    var signingKeyId by OauthResourceTable.signingKeyId
    var allowedScopes by OauthResourceTable.allowedScopes
    var customClaims by OauthResourceTable.customClaims
    var dpopBoundAccessTokensRequired by OauthResourceTable.dpopBoundAccessTokensRequired
    var disabled by OauthResourceTable.disabled
    var createdAt by OauthResourceTable.createdAt
    var updatedAt by OauthResourceTable.updatedAt
    var policyVersion by OauthResourceTable.policyVersion
    var metadata by OauthResourceTable.metadata
    val clients by OauthClientResourceEntity referrersOn OauthClientResourceTable.resourceId
}
