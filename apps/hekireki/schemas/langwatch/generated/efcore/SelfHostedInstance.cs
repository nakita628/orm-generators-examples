#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// One self-hosted install, as its daily usage report describes it (ADR-139,
/// section 10). The row is created by the first report that names an instance
/// id and is updated by every report after that, so the backoffice reads one
/// row per install rather than folding a history on every page load.
///
/// The report carries no credential, so it identifies an install and never a
/// customer. `organizationId` and `issuedLicenseId` are resolved from the
/// license bound to this instance, which the install proved by presenting its
/// token to license sync. A field in the posted body never reaches them.
/// </summary>
public partial class SelfHostedInstance
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The UUID the install minted for itself. The join key for everything else.
    /// </summary>
    public string InstanceId { get; set; } = null!;

    public DateTime FirstSeenAt { get; set; }

    public DateTime LastSeenAt { get; set; }

    /// <summary>
    /// What the install says about itself.
    /// </summary>
    public string? Version { get; set; }

    public string? InstallMethod { get; set; }

    public string? ChartVersion { get; set; }

    /// <summary>
    /// The host it answers on, when the operator left hostname reporting on.
    /// </summary>
    public string? Hostname { get; set; }

    public string? Environment { get; set; }

    /// <summary>
    /// The day the install itself says it first came up, which predates the
    /// first report on an install that switched reporting on later.
    /// </summary>
    public DateTime? InstalledAt { get; set; }

    public int? ReportSchemaVersion { get; set; }

    /// <summary>
    /// The customer, and the license that named them. Plain column refs rather
    /// than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
    /// carries no foreign keys, and the receiver resolves both from the license
    /// registry rather than from the report.
    /// </summary>
    public string? OrganizationId { get; set; }

    public string? IssuedLicenseId { get; set; }

    /// <summary>
    /// Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
    /// install that switched the optional category off.
    /// </summary>
    public string? UserEmailDomains { get; set; }

    /// <summary>
    /// The same domains as the keys of `userEmailDomains`, flattened so an
    /// operator can search for one. A JSON object's keys are the domains
    /// themselves, and no query language addresses a key it does not already
    /// know the name of.
    /// </summary>
    public List<string>? UserDomains { get; set; } = new List<string>();

    /// <summary>
    /// What the last report carried, so the list and the drawer read one row:
    /// the onboarding ladder, the usage counts and the operational block, under
    /// the same keys the dictionary declares.
    /// </summary>
    public string? LatestReport { get; set; }

    /// <summary>
    /// What the last report said the operator had switched off, read from which
    /// categories arrived rather than from a field the install could assert.
    /// </summary>
    public bool OptionalMetricsReported { get; set; } = true;

    public bool HostnameReported { get; set; } = true;

    /// <summary>
    /// How many reports have landed, and how many fields the last one carried
    /// that this release has no name for. A count above zero says the senders
    /// are ahead of the receiver.
    /// </summary>
    public int ReportCount { get; set; } = 0;

    public int LastUnknownFields { get; set; } = 0;

    /// <summary>
    /// The lead signals already raised for this install. A signal on this list
    /// is never raised again, which is what keeps a large install from
    /// reporting that it is large every morning.
    /// </summary>
    public List<string>? RaisedSignals { get; set; } = new List<string>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
