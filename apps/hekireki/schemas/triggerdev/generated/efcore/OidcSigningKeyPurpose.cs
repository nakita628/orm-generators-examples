#nullable enable

using NpgsqlTypes;

namespace Trigger.Models;

/// <summary>
/// Each purpose is a separate key set with its own issuer, rotation and publication.
/// </summary>
public enum OidcSigningKeyPurpose
{
    [PgName("EXTERNAL_OIDC")]
    ExternalOidc,

    [PgName("INTERNAL_RUN")]
    InternalRun,
}
