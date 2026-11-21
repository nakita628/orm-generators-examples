#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One row per erased account, holding only its creation time: no id, name or address. It keeps
/// the install's first-created account knowable after an erasure deletes that row.
/// Spec: modules/user/specs/user.feature
/// </summary>
public partial class UserErasure
{
    public string Id { get; set; } = null!;

    public DateTime AccountCreatedAt { get; set; }

    public DateTime ErasedAt { get; set; }
}
