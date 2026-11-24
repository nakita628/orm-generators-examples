#nullable enable

using System;
using System.Collections.Generic;

namespace LiteLLM.Models;

public partial class LiteLLMBudgetTable
{
    public double? RolloverMaxBudget { get; set; }

    public virtual ICollection<LiteLLMAppTable> Apps { get; set; } = new List<LiteLLMAppTable>();

    public virtual ICollection<LiteLLMAppGroupTable> AppGroups { get; set; } = new List<LiteLLMAppGroupTable>();

    public virtual ICollection<LiteLLMAppGroupCustomerOrgTable> AppGroupOrgBudgets { get; set; } = new List<LiteLLMAppGroupCustomerOrgTable>();

    public virtual ICollection<LiteLLMAppTable> AppDefaultTiers { get; set; } = new List<LiteLLMAppTable>();

    public virtual ICollection<LiteLLMAppTable> AppDefaultMembers { get; set; } = new List<LiteLLMAppTable>();

    public virtual ICollection<LiteLLMAppCustomerOrgTable> AppOrgTiers { get; set; } = new List<LiteLLMAppCustomerOrgTable>();

    public virtual ICollection<LiteLLMAppCustomerOrgTable> AppOrgBudgets { get; set; } = new List<LiteLLMAppCustomerOrgTable>();

    public virtual ICollection<LiteLLMAppCustomerMemberTable> AppMemberBudgets { get; set; } = new List<LiteLLMAppCustomerMemberTable>();

    public virtual ICollection<LiteLLMAppCustomerOrgTable> AppOrgOverrides { get; set; } = new List<LiteLLMAppCustomerOrgTable>();

    public virtual ICollection<LiteLLMAppCustomerOrgTable> AppOrgMembers { get; set; } = new List<LiteLLMAppCustomerOrgTable>();

    public virtual ICollection<LiteLLMAppCustomerMemberTable> AppMemberOverrides { get; set; } = new List<LiteLLMAppCustomerMemberTable>();

    public virtual ICollection<LiteLLMCustomerOrganizationTable> CustomerOrganizationDefaults { get; set; } = new List<LiteLLMCustomerOrganizationTable>();

    public virtual ICollection<LiteLLMCustomerOrganizationTable> CustomerOrganizations { get; set; } = new List<LiteLLMCustomerOrganizationTable>();

    public virtual ICollection<LiteLLMCustomerTeamTable> CustomerTeams { get; set; } = new List<LiteLLMCustomerTeamTable>();

    public virtual ICollection<LiteLLMCustomerMemberTable> CustomerMembers { get; set; } = new List<LiteLLMCustomerMemberTable>();

    public virtual ICollection<LiteLLMCustomerMemberTable> CustomerMemberOverrides { get; set; } = new List<LiteLLMCustomerMemberTable>();

    public virtual ICollection<LiteLLMCustomerTeamTable> CustomerTeamOverrides { get; set; } = new List<LiteLLMCustomerTeamTable>();

    public string BudgetId { get; set; } = null!;

    public string? Name { get; set; }

    public double? MaxBudget { get; set; }

    public double? SoftBudget { get; set; }

    public int? MaxParallelRequests { get; set; }

    public long? TpmLimit { get; set; }

    public long? RpmLimit { get; set; }

    public long? TpdLimit { get; set; }

    public string? ModelMaxBudget { get; set; }

    public string? ModelTpmLimit { get; set; }

    public string? ModelRpmLimit { get; set; }

    public string? BudgetDuration { get; set; }

    public string? BudgetLimits { get; set; }

    public string? CapPolicy { get; set; }

    public DateTime? BudgetResetAt { get; set; }

    public int BudgetResetCount { get; set; } = 0;

    public List<string>? AllowedModels { get; set; } = new List<string>();

    public double? TempBudgetIncrease { get; set; }

    public DateTime? TempBudgetExpiry { get; set; }

    public DateTime CreatedAt { get; set; }

    public string CreatedBy { get; set; } = null!;

    public DateTime UpdatedAt { get; set; }

    public string UpdatedBy { get; set; } = null!;

    public virtual ICollection<LiteLLMOrganizationTable> Organization { get; set; } = new List<LiteLLMOrganizationTable>();

    public virtual ICollection<LiteLLMProjectTable> Projects { get; set; } = new List<LiteLLMProjectTable>();

    public virtual ICollection<LiteLLMVerificationToken> Keys { get; set; } = new List<LiteLLMVerificationToken>();

    public virtual ICollection<LiteLLMEndUserTable> EndUsers { get; set; } = new List<LiteLLMEndUserTable>();

    public virtual ICollection<LiteLLMTagTable> Tags { get; set; } = new List<LiteLLMTagTable>();

    public virtual ICollection<LiteLLMModelAccessGroupBudgetTable> ModelAccessGroups { get; set; } = new List<LiteLLMModelAccessGroupBudgetTable>();

    public virtual ICollection<LiteLLMTeamMembership> TeamMembership { get; set; } = new List<LiteLLMTeamMembership>();

    public virtual ICollection<LiteLLMOrganizationMembership> OrganizationMembership { get; set; } = new List<LiteLLMOrganizationMembership>();
}
