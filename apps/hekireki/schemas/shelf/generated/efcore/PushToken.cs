#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// One Expo push token per installed companion app. A token identifies a
/// device, not a person: when someone else signs in on the same phone, the
/// row moves to them. Pushes go only to rows where `disabledAt` is null.
/// @see apps/webapp/app/modules/push/service.server.ts
/// </summary>
public partial class PushToken
{
    public string Id { get; set; } = null!;

    public string Token { get; set; } = null!;

    public DevicePlatform Platform { get; set; }

    public string? DeviceName { get; set; }

    public string? AppVersion { get; set; }

    public DateTime LastSeenAt { get; set; }

    public DateTime? DisabledAt { get; set; }

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
