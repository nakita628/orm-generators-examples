#nullable enable

using System;

namespace LangWatch.Models;

public partial class ScenarioVersion
{
    public string Id { get; set; } = null!;

    public string ScenarioId { get; set; } = null!;

    /// <summary>
    /// Deleting the case takes its history with it. Under relationMode="prisma"
    /// the cascade is the client's, so a raw SQL delete still has to clear these
    /// rows itself.
    /// </summary>
    public virtual Scenario Scenario { get; set; } = null!;

    /// <summary>
    /// The project of the parent case, carried here so every read filters by
    /// tenant without a join. The service is what keeps the pair honest: no
    /// composite foreign key can, under relationMode="prisma".
    /// </summary>
    public string ProjectId { get; set; } = null!;

    public int Version { get; set; }

    public string? AuthorId { get; set; }

    /// <summary>
    /// Who wrote this version: "user", "api", "cli" or "langy".
    /// </summary>
    public string AuthorLabel { get; set; } = null!;

    public string? ChangeDescription { get; set; }

    /// <summary>
    /// Envelope { schemaVersion, fields, changedFields } holding the editable
    /// field set as saved and the names of the fields this save changed.
    /// </summary>
    public string Snapshot { get; set; } = null!;

    /// <summary>
    /// Schema version of `snapshot`, so a later shape change can be read forward.
    /// </summary>
    public int SchemaVersion { get; set; } = 1;

    public DateTime CreatedAt { get; set; }
}
