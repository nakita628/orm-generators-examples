#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One person's digest for one ISO week, accepted by our own mailer. The id is `userId:isoWeek`.
/// </summary>
public partial class DigestSend
{
    public string Id { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string IsoWeek { get; set; } = null!;

    public string TemplateId { get; set; } = null!;

    public DateTime SentAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
