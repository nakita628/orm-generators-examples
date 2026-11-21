#nullable enable

namespace Calcom.Models;

public partial class BookingReference
{
    public int Id { get; set; }

    /// <summary>
    /// @zod.string.min(1)
    /// </summary>
    public string Type { get; set; } = null!;

    /// <summary>
    /// @zod.string.min(1)
    /// </summary>
    public string Uid { get; set; } = null!;

    public string? MeetingId { get; set; }

    public string? ThirdPartyRecurringEventId { get; set; }

    public string? MeetingPassword { get; set; }

    public string? MeetingUrl { get; set; }

    public virtual Booking? Booking { get; set; }

    public int? BookingId { get; set; }

    public string? ExternalCalendarId { get; set; }

    public bool? Deleted { get; set; }

    public virtual Credential? Credential { get; set; }

    public int? CredentialId { get; set; }

    public virtual DelegationCredential? DelegationCredential { get; set; }

    public string? DelegationCredentialId { get; set; }
}
