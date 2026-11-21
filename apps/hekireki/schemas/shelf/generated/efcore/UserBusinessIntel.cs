#nullable enable

using System;

namespace Shelf.Models;

public partial class UserBusinessIntel
{
    public string Id { get; set; } = null!;

    public string? HowDidYouHearAboutUs { get; set; }

    public string? JobTitle { get; set; }

    public string? TeamSize { get; set; }

    public string? CompanyName { get; set; }

    public string? PrimaryUseCase { get; set; }

    public string? CurrentSolution { get; set; }

    public string? Timeline { get; set; }

    public string? SignupPlan { get; set; }

    public bool? SignupTrial { get; set; }

    public string? UtmSource { get; set; }

    public string? UtmMedium { get; set; }

    public string? UtmCampaign { get; set; }

    public string? UtmContent { get; set; }

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
