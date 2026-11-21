#nullable enable

namespace Formbricks.Models;

/// <summary>
/// Junction table linking responses to quotas.
/// Tracks which responses counted towards specific quotas and their status.
///
/// @property response - The response that counted towards the quota
/// @property quota - The quota this response counted towards
/// @property status - Whether the response was screened in or out
/// </summary>
public partial class ResponseQuotaLink
{
    public virtual Response Response { get; set; } = null!;

    public string ResponseId { get; set; } = null!;

    public virtual SurveyQuota Quota { get; set; } = null!;

    public string QuotaId { get; set; } = null!;

    public ResponseQuotaLinkStatus Status { get; set; }
}
