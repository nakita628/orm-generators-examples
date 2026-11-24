#nullable enable

using NpgsqlTypes;

namespace Trigger.Models;

public enum DeploymentInitiatedByCredential
{
    [PgName("ROOT_API_KEY")]
    RootApiKey,

    [PgName("SCOPED_API_KEY")]
    ScopedApiKey,

    [PgName("PERSONAL_ACCESS_TOKEN")]
    PersonalAccessToken,

    [PgName("ORGANIZATION_ACCESS_TOKEN")]
    OrganizationAccessToken,

    [PgName("TRUSTED_PUBLISHING")]
    TrustedPublishing,
}
