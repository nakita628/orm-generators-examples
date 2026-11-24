#nullable enable

using System;

namespace LiteLLM.Models;

public partial class LiteLLMAppCustomerOrgTable
{
    public string AppId { get; set; } = null!;

    public string CustomerOrgId { get; set; } = null!;

    public double Spend { get; set; } = 0.0;

    public string? SpendResetWindow { get; set; }

    public string? TierBudgetId { get; set; }

    public string? BudgetId { get; set; }

    public virtual LiteLLMBudgetTable? Budget { get; set; }

    public string? BudgetOverrideId { get; set; }

    public DateTime? BudgetOverrideExpiresAt { get; set; }

    public string? CustomerMemberBudgetId { get; set; }

    public virtual LiteLLMAppTable App { get; set; } = null!;

    public virtual LiteLLMCustomerOrganizationTable CustomerOrg { get; set; } = null!;

    public virtual LiteLLMBudgetTable? TierBudget { get; set; }

    public virtual LiteLLMBudgetTable? BudgetOverride { get; set; }

    public virtual LiteLLMBudgetTable? CustomerMemberBudget { get; set; }
}
