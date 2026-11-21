package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** Better Auth OAuth 2.1 client registrations for MCP and future OAuth clients. */
class OauthClientEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OauthClientEntity>(OauthClientTable)

    var clientId by OauthClientTable.clientId
    var clientSecret by OauthClientTable.clientSecret
    var disabled by OauthClientTable.disabled
    var skipConsent by OauthClientTable.skipConsent
    var enableEndSession by OauthClientTable.enableEndSession
    var subjectType by OauthClientTable.subjectType
    var scopes by OauthClientTable.scopes
    var userId by OauthClientTable.userId
    var user by UserEntity optionalReferencedOn OauthClientTable.userId
    var createdAt by OauthClientTable.createdAt
    var updatedAt by OauthClientTable.updatedAt
    var name by OauthClientTable.name
    var uri by OauthClientTable.uri
    var icon by OauthClientTable.icon
    var contacts by OauthClientTable.contacts
    var tos by OauthClientTable.tos
    var policy by OauthClientTable.policy
    var softwareId by OauthClientTable.softwareId
    var softwareVersion by OauthClientTable.softwareVersion
    var softwareStatement by OauthClientTable.softwareStatement
    var redirectUris by OauthClientTable.redirectUris
    var postLogoutRedirectUris by OauthClientTable.postLogoutRedirectUris
    var tokenEndpointAuthMethod by OauthClientTable.tokenEndpointAuthMethod
    var grantTypes by OauthClientTable.grantTypes
    var responseTypes by OauthClientTable.responseTypes

    /**
     * Unwritten by Better Auth 1.7 (it dropped both from its schema), kept deliberately: they are
     * nullable, so keeping them makes the 1.7 migration purely additive and a rollback to 1.6 needs no
     * down-migration. Drop them in a separate cleanup after a bake period.
     */
    var public by OauthClientTable.public
    var type by OauthClientTable.type
    var requirePKCE by OauthClientTable.requirePKCE
    var referenceId by OauthClientTable.referenceId
    var metadata by OauthClientTable.metadata

    /** Better Auth 1.7 (ENG-2343). */
    var clientDiscoveryId by OauthClientTable.clientDiscoveryId
    var clientCredentialsScopes by OauthClientTable.clientCredentialsScopes
    var backchannelLogoutUri by OauthClientTable.backchannelLogoutUri
    var backchannelLogoutSessionRequired by OauthClientTable.backchannelLogoutSessionRequired
    var applicationType by OauthClientTable.applicationType
    var jwks by OauthClientTable.jwks
    var jwksUri by OauthClientTable.jwksUri
    var dpopBoundAccessTokens by OauthClientTable.dpopBoundAccessTokens
    val accessTokens by OauthAccessTokenEntity referrersOn OauthAccessTokenTable.clientId
    val refreshTokens by OauthRefreshTokenEntity referrersOn OauthRefreshTokenTable.clientId
    val consents by OauthConsentEntity referrersOn OauthConsentTable.clientId
    val resources by OauthClientResourceEntity referrersOn OauthClientResourceTable.clientId
}
