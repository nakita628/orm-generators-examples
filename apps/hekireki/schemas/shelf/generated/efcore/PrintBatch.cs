#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class PrintBatch
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public bool Printed { get; set; } = false;

    public virtual ICollection<Qr> QrCodes { get; set; } = new List<Qr>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
