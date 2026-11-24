#nullable enable

using System;
using System.Collections.Generic;

namespace LiteLLM.Models;

public partial class LiteLLMAppGroupTable
{
    public string AppGroupId { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public string? Metadata { get; set; }

    public string? BudgetId { get; set; }

    public virtual LiteLLMBudgetTable? Budget { get; set; }

    public double Spend { get; set; } = 0.0;

    public string? SpendResetWindow { get; set; }

    public virtual ICollection<LiteLLMAppTable> Apps { get; set; } = new List<LiteLLMAppTable>();

    public virtual ICollection<LiteLLMAppGroupCustomerOrgTable> CustomerOrganizations { get; set; } = new List<LiteLLMAppGroupCustomerOrgTable>();

    public DateTime CreatedAt { get; set; }

    public string CreatedBy { get; set; } = null!;

    public DateTime UpdatedAt { get; set; }

    public string UpdatedBy { get; set; } = null!;
}
