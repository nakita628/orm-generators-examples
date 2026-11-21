#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Defines one Embedded Data field: the extra context stored on a response that isn't an
/// answer to a question (brand, plan, page type, a computed score, ...).
///
/// A field is either **local** — used by a single survey, which is how existing variables and
/// hidden fields are migrated — or **shared**, meaning it is part of the workspace library and can
/// be linked by many surveys. There is no stored flag for that: `surveyId` is the owner and `key`
/// is the library name, so a local field has `surveyId` set and `key` null, and a shared field the
/// other way round. Deriving it rather than storing a boolean removes a column that could drift
/// out of step with the two fields it stands for.
///
/// @property key - Library name, unique per workspace and immutable. Null for local fields
/// @property name - Human-readable display label
/// @property source - Where the value comes from (computed | ingested). `reserved` is part of the
/// shared enum but never stored as a row: those fields are a static catalog in code
/// @property dataType - Declared value type, used for coercion at ingest
/// @property defaultValue - Initial value (computed) or fallback when nothing arrives (ingested)
/// @property locked - Ingested only: ignore external writes and use the default instead
/// @property survey - The owning survey for a local field; cascades so nothing is orphaned
/// </summary>
public partial class EmbeddedData
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string? Key { get; set; }

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public EmbeddedDataSource Source { get; set; }

    public EmbeddedDataType DataType { get; set; } = EmbeddedDataType.String;

    /// <summary>
    /// [EmbeddedDataDefaultValue]
    /// </summary>
    public string? DefaultValue { get; set; }

    public bool Locked { get; set; } = false;

    public virtual Survey? Survey { get; set; }

    public string? SurveyId { get; set; }

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual ICollection<SurveyEmbeddedData> SurveyLinks { get; set; } = new List<SurveyEmbeddedData>();
}
