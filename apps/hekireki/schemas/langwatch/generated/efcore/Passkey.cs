#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
/// public key and the signature counter are the plugin's to read and write.
/// A passkey also becomes an `Identifier` row with provider `passkey`, which
/// is what makes it a sign-in method like any other - this table is the
/// credential, that one is the identity.
///
/// `aaguid` names the authenticator MODEL, never a device or a person, and is
/// only present at registration. Privacy-preserving platforms report all
/// zeroes, so a label derived from it is best-effort by construction.
/// </summary>
public partial class Passkey
{
    public string Id { get; set; } = null!;

    public string? Name { get; set; }

    public string PublicKey { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string CredentialID { get; set; } = null!;

    public int Counter { get; set; } = 0;

    public string DeviceType { get; set; } = null!;

    public bool BackedUp { get; set; } = false;

    public string? Transports { get; set; }

    public string? Aaguid { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
