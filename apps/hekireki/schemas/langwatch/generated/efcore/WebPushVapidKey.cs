#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// This installation's VAPID key pair, generated on first use so Web Push works
/// without configuration. Fixed id, so the table holds one row: a second
/// insert loses the race on the primary key and reads the winner's pair.
/// </summary>
public partial class WebPushVapidKey
{
    public string Id { get; set; } = "self";

    /// <summary>
    /// The uncompressed P-256 public key, base64url. The browser subscribes with it.
    /// </summary>
    public string PublicKey { get; set; } = null!;

    /// <summary>
    /// The private key, base64url, encrypted with the process encryption key.
    /// </summary>
    public string PrivateKeyEncrypted { get; set; } = null!;

    public DateTime CreatedAt { get; set; }
}
