#nullable enable

using NpgsqlTypes;

namespace LangWatch.Models;

/// <summary>
/// How a Slack connection posts: a bot token over the Web API (any channel it
/// is invited to, every block renders) or an incoming webhook (one channel).
/// </summary>
public enum SlackIntegrationKind
{
    [PgName("BOT")]
    Bot,

    [PgName("INCOMING_WEBHOOK")]
    IncomingWebhook,
}
