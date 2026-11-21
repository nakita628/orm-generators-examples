#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a person or user who can receive and respond to surveys.
/// Contacts are workspace-specific and can have multiple attributes and responses.
///
/// @property id - Unique identifier for the contact
/// @property workspace - The workspace this contact belongs to
/// @property responses - Survey responses from this contact
/// @property attributes - Custom attributes associated with this contact
/// @property displays - Record of surveys shown to this contact
/// </summary>
public partial class Contact
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual ICollection<Response> Responses { get; set; } = new List<Response>();

    public virtual ICollection<ContactAttribute> Attributes { get; set; } = new List<ContactAttribute>();

    public virtual ICollection<Display> Displays { get; set; } = new List<Display>();
}
