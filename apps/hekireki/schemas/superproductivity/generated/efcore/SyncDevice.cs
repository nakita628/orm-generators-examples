#nullable enable

namespace SuperProductivity.Models;

public partial class SyncDevice
{
    public string ClientId { get; set; } = null!;

    public int UserId { get; set; }

    public string? DeviceName { get; set; }

    public string? UserAgent { get; set; }

    /// <summary>
    /// Bare semver the client last reported on download (`18.22.0`), NULL until
    /// it does. Read by the checkpoint gate (#9962); never shown to users.
    /// </summary>
    public string? AppVersion { get; set; }

    public long LastSeenAt { get; set; }

    /// <summary>
    /// Reserved: written once at device registration, never read or updated.
    /// </summary>
    public int LastAckedSeq { get; set; } = 0;

    public long CreatedAt { get; set; }

    public virtual User User { get; set; } = null!;
}
