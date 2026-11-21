#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Represents an attribute value associated with a contact.
/// Used to store custom properties and characteristics of contacts.
///
/// @property id - Unique identifier for the attribute
/// @property attributeKey - Reference to the attribute definition
/// @property contact - The contact this attribute belongs to
/// @property value - The string value of the attribute (used for string type + backwards compatibility)
/// @property valueNumber - Native numeric storage for number type attributes
/// @property valueDate - Native date storage for date type attributes
/// </summary>
public partial class ContactAttribute
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ContactAttributeKey AttributeKey { get; set; } = null!;

    public string AttributeKeyId { get; set; } = null!;

    public virtual Contact Contact { get; set; } = null!;

    public string ContactId { get; set; } = null!;

    public string Value { get; set; } = null!;

    public double? ValueNumber { get; set; }

    public DateTime? ValueDate { get; set; }
}
