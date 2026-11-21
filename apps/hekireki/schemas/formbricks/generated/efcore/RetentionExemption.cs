#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Holds one survey back from one policy until `until`. Keyed by organisation + entity rather than a
/// policy id, so removing and re-adding a policy doesn't orphan its exemptions.
///
/// Active means `revokedAt IS NULL AND until &gt; now()`, so expiry needs no job. "One active exemption
/// per survey and policy" is the partial unique index `RetentionExemption_surveyId_entity_active_key`
/// (`WHERE "revokedAt" IS NULL`) in
/// `migration/20261007120000_add_data_retention_tables/migration.sql`; Prisma's `@@unique` has no
/// `where`, so it is deliberately not declared here (see "Indexes Prisma cannot express" in the
/// README). An expired row still holds that slot until it is revoked, so creating a new exemption
/// must close an expired one in the same transaction.
///
/// `surveyId` is nullable and a CHECK requires exactly one target, so a later `workflowId` is one more
/// column in that CHECK. The survey's organisation can't be tied to `organizationId` in the database
/// (a survey carries its workspace, not its organisation): writers derive it from the survey.
/// </summary>
public partial class RetentionExemption
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public RetentionEntity Entity { get; set; }

    public virtual Survey? Survey { get; set; }

    public string? SurveyId { get; set; }

    public DateTime Until { get; set; }

    public RetentionExemptionSource Source { get; set; } = RetentionExemptionSource.Manual;

    /// <summary>
    /// Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`).
    /// </summary>
    public string? Reason { get; set; }

    public virtual User? CreatedBy { get; set; }

    public string? CreatedById { get; set; }

    public DateTime? RevokedAt { get; set; }

    public virtual User? RevokedBy { get; set; }

    public string? RevokedById { get; set; }
}
