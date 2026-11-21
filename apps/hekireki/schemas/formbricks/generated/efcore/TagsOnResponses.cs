#nullable enable

namespace Formbricks.Models;

/// <summary>
/// Junction table linking tags to responses.
/// Enables many-to-many relationship between tags and responses.
///
/// @property response - The tagged response
/// @property tag - The tag applied to the response
/// </summary>
public partial class TagsOnResponses
{
    public string ResponseId { get; set; } = null!;

    public virtual Response Response { get; set; } = null!;

    public string TagId { get; set; } = null!;

    public virtual Tag Tag { get; set; } = null!;
}
