#nullable enable

using System;
using System.Collections.Generic;

namespace LiteLLM.Models;

public partial class LiteLLMAppTable
{
    public string AppId { get; set; } = null!;

    public string? AppGroupId { get; set; }

    public virtual LiteLLMAppGroupTable? AppGroup { get; set; }

    public string? ResourceType { get; set; }

    public string? ResourceId { get; set; }

    public string Alias { get; set; } = null!;

    public double Spend { get; set; } = 0.0;

    public string? SpendResetWindow { get; set; }

    public string? BudgetId { get; set; }

    public virtual LiteLLMBudgetTable? Budget { get; set; }

    public string? DefaultTierBudgetId { get; set; }

    public virtual LiteLLMBudgetTable? DefaultTierBudget { get; set; }

    public string? DefaultCustomerMemberBudgetId { get; set; }

    public virtual LiteLLMBudgetTable? DefaultCustomerMemberBudget { get; set; }

    public bool AutoCreateCustomers { get; set; } = true;

    public double? CustomerMargin { get; set; }

    public virtual ICollection<LiteLLMAppCustomerOrgTable> CustomerOrganizations { get; set; } = new List<LiteLLMAppCustomerOrgTable>();

    public virtual ICollection<LiteLLMAppCustomerMemberTable> CustomerMembers { get; set; } = new List<LiteLLMAppCustomerMemberTable>();

    public DateTime CreatedAt { get; set; }

    public string CreatedBy { get; set; } = null!;

    public DateTime UpdatedAt { get; set; }

    public string UpdatedBy { get; set; } = null!;
}
