#nullable enable

using System;

namespace LiteLLM.Models;

public partial class LiteLLMBackgroundResponseCreator
{
    public string ResponseId { get; set; } = null!;

    public string UpstreamUrl { get; set; } = null!;

    public string CreateContext { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime? PollLeaseExpiresAt { get; set; }
}
