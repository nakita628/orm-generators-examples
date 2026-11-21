#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The identity this install presents to LangWatch (ADR-139, section 2).
///
/// Exactly one row, minted on first use and never rewritten. It survives
/// restarts, backups and hostname changes, and an install restored into
/// another database mints another one, which is what the registry refuses as
/// the wrong instance.
///
/// It replaced an identity built from the organization name and id, which
/// leaked the customer's name in cleartext, named one organization rather than
/// one install, and could not be joined to a license row without splitting a
/// string.
/// </summary>
public partial class InstanceIdentity
{
    /// <summary>
    /// Fixed, so this table holds one row. A second insert loses the race on the
    /// primary key and re-reads the winner's id.
    /// </summary>
    public string Id { get; set; } = "self";

    /// <summary>
    /// A random UUID. Carries nothing about the customer.
    /// </summary>
    public string InstanceId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// When the usage report last reached LangWatch, and why the last attempt
    /// did not. Both are what the checkup page reads to say whether reporting
    /// works, which nothing recorded before: the sender logged success as soon
    /// as `fetch` resolved and never looked at the response.
    /// </summary>
    public DateTime? LastReportAt { get; set; }

    public string? LastReportError { get; set; }

    /// <summary>
    /// The optional category of the usage report, switched off from the checkup
    /// page. With it off the report still says which release runs and how many
    /// projects the install carries, and says nothing about what anyone does
    /// with it (ADR-139, section 10).
    /// </summary>
    public bool OptionalMetricsOptOut { get; set; } = false;

    /// <summary>
    /// Hostname, which names the customer's own network and so has a switch of
    /// its own rather than travelling with its category.
    /// </summary>
    public bool HostnameOptOut { get; set; } = false;

    /// <summary>
    /// Sending a failed upgrade step's diagnostics to LangWatch on its own,
    /// switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
    /// </summary>
    public bool UpgradeDiagnosticsOptOut { get; set; } = false;
}
