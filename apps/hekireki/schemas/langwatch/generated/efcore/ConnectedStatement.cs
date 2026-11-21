#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The monthly usage statement one connected customer was sent (ADR-139,
/// section 7). One row per account per month keeps the statement to one send,
/// whatever runs the tick.
/// </summary>
public partial class ConnectedStatement
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The `ConnectedBillingAccount` row. A plain column ref.
    /// </summary>
    public string AccountId { get; set; } = null!;

    /// <summary>
    /// First day of the month the statement covers, in UTC.
    /// </summary>
    public DateTime Month { get; set; }

    public DateTime SentAt { get; set; }
}
