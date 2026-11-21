#nullable enable

using System.Collections.Generic;

namespace TypeHero.Models;

public partial class Track
{
    public int Id { get; set; }

    public string Name { get; set; } = null!;

    public string Slug { get; set; } = null!;

    public string Description { get; set; } = null!;

    /// <summary>
    /// determines if a track can be seen on tracks page
    /// </summary>
    public bool Visible { get; set; } = true;

    /// <summary>
    /// use this to display coming soon UI
    /// </summary>
    public bool IsComingSoon { get; set; } = false;

    public virtual ICollection<TrackChallenge> TrackChallenges { get; set; } = new List<TrackChallenge>();

    public virtual ICollection<User> EnrolledUsers { get; set; } = new List<User>();
}
