#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class CustomField
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? HelpText { get; set; }

    public bool Required { get; set; } = false;

    public bool Active { get; set; } = true;

    public CustomFieldType Type { get; set; } = CustomFieldType.Text;

    public List<string>? Options { get; set; }

    /// <summary>
    /// How a TEXT field's values are read. LINK requires a web link in every
    /// value; any other type stays PLAIN (enforced by a CHECK constraint).
    /// </summary>
    public CustomFieldFormat Format { get; set; } = CustomFieldFormat.Plain;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual User CreatedBy { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual ICollection<Category> Categories { get; set; } = new List<Category>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? DeletedAt { get; set; }

    public virtual ICollection<AssetCustomFieldValue> AssetCustomFieldsValues { get; set; } = new List<AssetCustomFieldValue>();
}
