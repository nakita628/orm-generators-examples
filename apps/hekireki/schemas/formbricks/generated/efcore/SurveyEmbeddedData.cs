#nullable enable

namespace Formbricks.Models;

/// <summary>
/// Links a survey to an Embedded Data field and records the key that field's value is stored
/// under inside that survey.
///
/// `workspaceId` is part of both foreign keys on purpose: `surveyId` and `embeddedDataId` alone are
/// independent, so nothing would stop a survey in one workspace linking a field defined in another.
/// Referencing `(id, workspaceId)` on both sides makes a cross-workspace pair unrepresentable, and
/// puts the workspace scope every query needs on the row instead of behind a join through Survey.
///
/// @property storageKey - The key the value lives under: a cuid for computed fields, the field
/// name for ingested ones. Migrated fields keep their original key, which is what lets existing
/// recall tokens, logic conditions and stored responses keep resolving untouched.
///
/// @property order - Position within this survey's field list, dense from 0. Lives on the link
/// rather than on `EmbeddedData` because a shared definition sits at a different position in each
/// survey that links it. Deliberately has no default: every writer knows the position it is writing,
/// so a missing one should be a type error rather than a row that silently sorts to the front.
/// </summary>
public partial class SurveyEmbeddedData
{
    public string Id { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public string EmbeddedDataId { get; set; } = null!;

    public string StorageKey { get; set; } = null!;

    public int Order { get; set; }

    public virtual Survey Survey { get; set; } = null!;

    public virtual EmbeddedData EmbeddedData { get; set; } = null!;
}
