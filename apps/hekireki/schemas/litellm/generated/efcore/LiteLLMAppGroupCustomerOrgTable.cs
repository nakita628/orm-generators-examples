#nullable enable

namespace LiteLLM.Models;

public partial class LiteLLMAppGroupCustomerOrgTable
{
    public string AppGroupId { get; set; } = null!;

    public string CustomerOrgId { get; set; } = null!;

    public string? BudgetId { get; set; }

    public virtual LiteLLMBudgetTable? Budget { get; set; }

    public double Spend { get; set; } = 0.0;

    public string? SpendResetWindow { get; set; }

    public virtual LiteLLMAppGroupTable AppGroup { get; set; } = null!;

    public virtual LiteLLMCustomerOrganizationTable CustomerOrg { get; set; } = null!;
}
