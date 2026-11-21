#nullable enable

namespace Formbricks.Models;

/// <summary>
/// Links surveys to their supported languages.
/// Manages which languages are available for each survey.
///
/// @property language - The supported language
/// @property survey - The associated survey
/// @property default - Whether this is the default language
/// @property enabled - Whether this language is currently active
/// </summary>
public partial class SurveyLanguage
{
    public virtual Language Language { get; set; } = null!;

    public string LanguageId { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public virtual Survey Survey { get; set; } = null!;

    public bool Default { get; set; } = false;

    public bool Enabled { get; set; } = true;
}
