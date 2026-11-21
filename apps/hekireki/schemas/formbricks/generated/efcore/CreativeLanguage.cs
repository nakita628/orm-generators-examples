#nullable enable

namespace Formbricks.Models;

/// <summary>
/// The languages a creative's texts are written in. Mirrors `SurveyLanguage`; exactly one is the default.
/// </summary>
public partial class CreativeLanguage
{
    public virtual Language Language { get; set; } = null!;

    public string LanguageId { get; set; } = null!;

    public virtual Creative Creative { get; set; } = null!;

    public string CreativeId { get; set; } = null!;

    public bool Default { get; set; } = false;

    public bool Enabled { get; set; } = true;
}
