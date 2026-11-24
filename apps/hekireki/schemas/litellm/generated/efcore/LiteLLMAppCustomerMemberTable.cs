#nullable enable

using System;

namespace LiteLLM.Models;

public partial class LiteLLMAppCustomerMemberTable
{
    public string AppId { get; set; } = null!;

    public string CustomerMemberId { get; set; } = null!;

    public double Spend { get; set; } = 0.0;

    public string? SpendResetWindow { get; set; }

    public string? BudgetOverrideId { get; set; }

    public string? BudgetId { get; set; }

    public virtual LiteLLMBudgetTable? Budget { get; set; }

    public DateTime? BudgetOverrideExpiresAt { get; set; }

    public virtual LiteLLMAppTable App { get; set; } = null!;

    public virtual LiteLLMCustomerMemberTable CustomerMember { get; set; } = null!;

    public virtual LiteLLMBudgetTable? BudgetOverride { get; set; }
}
