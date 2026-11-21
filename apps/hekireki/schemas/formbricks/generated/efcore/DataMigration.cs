#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Tracks database schema migrations.
/// Used to manage and track the state of data structure changes.
///
/// @property id - Unique identifier for the migration
/// @property name - Name of the migration
/// @property status - Current state of the migration (pending, applied, failed)
/// @property startedAt - When the migration began
/// @property finishedAt - When the migration completed
/// </summary>
public partial class DataMigration
{
    public string Id { get; set; } = null!;

    public DateTime StartedAt { get; set; }

    public DateTime? FinishedAt { get; set; }

    public string Name { get; set; } = null!;

    public DataMigrationStatus Status { get; set; }
}
