#nullable enable

using System;

namespace Shelf.Models;

public partial class UserContact
{
    public string Id { get; set; } = null!;

    public string? Phone { get; set; }

    public string? Street { get; set; }

    public string? City { get; set; }

    public string? StateProvince { get; set; }

    public string? ZipPostalCode { get; set; }

    public string? CountryRegion { get; set; }

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
