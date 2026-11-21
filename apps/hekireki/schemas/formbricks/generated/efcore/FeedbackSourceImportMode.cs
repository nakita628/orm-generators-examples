#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// Which survey responses a Formbricks feedback source imports into Hub.
/// completedOnly matches the live ingestion path, which only runs on responseFinished.
/// all additionally imports partial responses — answers a respondent typed but never submitted.
/// </summary>
public enum FeedbackSourceImportMode
{
    [PgName("completedOnly")]
    CompletedOnly,

    [PgName("all")]
    All,
}
