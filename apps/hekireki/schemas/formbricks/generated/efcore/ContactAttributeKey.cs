#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Defines the possible attributes that can be assigned to contacts.
/// Acts as a schema for contact attributes within a workspace.
///
/// @property id - Unique identifier for the attribute key
/// @property isUnique - Whether the attribute must have unique values across contacts
/// @property key - The attribute identifier used in the system
/// @property name - Display name for the attribute
/// @property type - Whether this is a default or custom attribute
/// @property dataType - The data type of the attribute (string, number, date)
/// @property workspace - The workspace this attribute belongs to
/// </summary>
public partial class ContactAttributeKey
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public bool IsUnique { get; set; } = false;

    public string Key { get; set; } = null!;

    public string? Name { get; set; }

    public string? Description { get; set; }

    public ContactAttributeType Type { get; set; } = ContactAttributeType.Custom;

    public ContactAttributeDataType DataType { get; set; } = ContactAttributeDataType.String;

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual ICollection<ContactAttribute> Attributes { get; set; } = new List<ContactAttribute>();

    public virtual ICollection<SurveyAttributeFilter> AttributeFilters { get; set; } = new List<SurveyAttributeFilter>();
}
