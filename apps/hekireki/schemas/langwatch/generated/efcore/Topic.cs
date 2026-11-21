#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// PROJECTION of the topic-clustering event stream (ADR-051): topics are
/// `topics_recorded` events; this table is written ONLY by
/// TopicModelFoldProjection (cursor: TopicModelProjection) and is
/// rebuildable by replay. Kept named `Topic` because it is the domain
/// entity every surface reads — the `*Projection` suffix is reserved for
/// cursor/read-model tables.
/// </summary>
public partial class Topic
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? ParentId { get; set; }

    public virtual Topic? Parent { get; set; }

    public virtual ICollection<Topic> Subtopics { get; set; } = new List<Topic>();

    public string EmbeddingsModel { get; set; } = null!;

    public string Centroid { get; set; } = null!;

    public double P95Distance { get; set; }

    public bool AutomaticallyGenerated { get; set; } = true;

    /// <summary>
    /// The `topics_recorded` event that recorded this topic (its id in the
    /// event log). Null means the row predates event ownership and has not
    /// been seeded onto the stream yet; the projection sets it on every row
    /// it writes.
    /// </summary>
    public string? LastEventId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
