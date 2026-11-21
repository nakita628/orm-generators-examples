#nullable enable

namespace SuperProductivity.Models;

/// <summary>
/// One sealed capture waiting for a device. No relation to its grant: items
/// outlive a grant that is revoked without discarding them.
/// </summary>
public partial class CaptureItem
{
    public string Id { get; set; } = null!;

    public int UserId { get; set; }

    public string GrantId { get; set; } = null!;

    public string Ciphertext { get; set; } = null!;

    /// <summary>
    /// Task id the claiming device must use; replaced when another device takes over.
    /// </summary>
    public string TaskId { get; set; } = null!;

    public long CreatedAt { get; set; }

    public string? LeaseDeviceId { get; set; }

    public long? LeaseUntil { get; set; }

    public virtual User User { get; set; } = null!;
}
