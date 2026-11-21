#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// Creatives (ENG-3993): page elements the web SDK renders next to surveys. Every type lives in
/// `Creative`; `content` and `styling` are JSON validated per type by the API v3 schemas, so a new type is
/// a new payload variant, not new columns.
/// </summary>
public enum CreativeType
{
    [PgName("feedbackButton")]
    FeedbackButton,

    [PgName("banner")]
    Banner,
}
