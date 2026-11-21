#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// The last hop a webhook endpoint delivers over. Lowercase because these
/// values are the wire spelling: the same strings appear in the REST DTO,
/// both SDKs and the CLI.
/// </summary>
public enum WebhookDestinationKind
{
    [PgName("http")]
    Http,

    [PgName("sqs")]
    Sqs,
}
