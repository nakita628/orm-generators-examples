#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// Who may read a survey. `workspace` is everyone with workspace access (the behaviour before ENG-3282);
/// `restricted` is the owner and the organization's owners and managers only.
/// </summary>
public enum SurveyVisibility
{
    [PgName("restricted")]
    Restricted,

    [PgName("workspace")]
    Workspace,
}
