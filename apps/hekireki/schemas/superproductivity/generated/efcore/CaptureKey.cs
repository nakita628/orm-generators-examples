#nullable enable

namespace SuperProductivity.Models;

/// <summary>
/// Public half of the user's capture key. The row outlives switching the
/// feature off: `disabledAt` stops a stale device re-registering that key.
/// </summary>
public partial class CaptureKey
{
    public int UserId { get; set; }

    /// <summary>
    /// Base64 SPKI DER of an ECDH P-256 public key.
    /// </summary>
    public string PublicKey { get; set; } = null!;

    public string Fingerprint { get; set; } = null!;

    public long CreatedAt { get; set; }

    public long? DisabledAt { get; set; }

    public virtual User User { get; set; } = null!;
}
