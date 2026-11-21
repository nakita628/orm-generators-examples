#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class KitCustody
{
    public string Id { get; set; } = null!;

    public virtual TeamMember Custodian { get; set; } = null!;

    public string CustodianId { get; set; } = null!;

    public virtual Kit Kit { get; set; } = null!;

    public string KitId { get; set; } = null!;

    public virtual ICollection<Custody> InheritedCustody { get; set; } = new List<Custody>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
