#nullable enable

using NpgsqlTypes;

namespace Rallly.Models;

public enum SpaceUsageResource
{
    [PgName("api_poll")]
    ApiPoll,

    [PgName("email_invite")]
    EmailInvite,
}
