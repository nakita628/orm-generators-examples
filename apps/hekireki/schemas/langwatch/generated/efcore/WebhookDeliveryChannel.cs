#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// Which webhook pipeline wrote a delivery row, and therefore which columns
/// carry its tenancy: `platform` rows anchor on organizationId + endpointId,
/// `automations` rows on projectId + triggerId.
/// </summary>
public enum WebhookDeliveryChannel
{
    [PgName("platform")]
    Platform,

    [PgName("automations")]
    Automations,
}
