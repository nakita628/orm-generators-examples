#nullable enable

using System;

namespace LangWatch.Models;

public partial class ExperimentVersion
{
    public string Id { get; set; } = null!;

    public string ExperimentId { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// The number the version is addressed and displayed by. Numbered rows hold
    /// a gap-free sequence. The rolling autosave row rides on the counter
    /// instead, which is what keeps it clear of that sequence.
    /// </summary>
    public int Version { get; set; }

    /// <summary>
    /// `Experiment.workbenchVersion` at the moment this row was last written.
    /// The rolling row is rewritten in place, so this is the only column that
    /// says how recent its content is, and the version list sorts and pages by
    /// it. It also joins a row to the live state: the row whose `counterVersion`
    /// equals the experiment's counter holds what the workbench shows now.
    /// </summary>
    public int CounterVersion { get; set; }

    public bool AutoSaved { get; set; } = false;

    public string? CommitMessage { get; set; }

    public string? AuthorId { get; set; }

    /// <summary>
    /// Who wrote this version: "user", "langy" or "api".
    /// </summary>
    public string AuthorLabel { get; set; } = null!;

    /// <summary>
    /// The run that wrote this version, when a run wrote it.
    ///
    /// A run writes its cells into the workbench state, which advances the
    /// counter. Without this, the page that started that run reads the bump as
    /// somebody else's write, refuses to save, and asks the reader to reload
    /// over their own unsaved edits. Named here, the page adopts the version and
    /// carries on.
    /// </summary>
    public string? RunId { get; set; }

    public string State { get; set; } = null!;

    /// <summary>
    /// Schema version of `state`, so a later shape change can be read forward.
    /// </summary>
    public string SchemaVersion { get; set; } = "1";

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
