#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Stores a user's response to a survey, including their answers and metadata.
/// Each response is linked to a specific survey and optionally to a contact.
///
/// @property id - Unique identifier for the response
/// @property finished - Whether the survey was completed
/// @property survey - The associated survey
/// @property contact - The optional contact who provided the response
/// @property data - JSON object containing the actual response data
/// @property variables - Custom variables used in the response
/// @property ttc - Time to completion metrics
/// @property meta - Additional metadata about the response
/// </summary>
public partial class Response
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public bool Finished { get; set; } = false;

    public virtual Survey Survey { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public virtual Contact? Contact { get; set; }

    public string? ContactId { get; set; }

    public string? EndingId { get; set; }

    /// <summary>
    /// [ResponseData]
    /// </summary>
    public string Data { get; set; } = "{}";

    /// <summary>
    /// [ResponseVariables]
    /// </summary>
    public string Variables { get; set; } = "{}";

    /// <summary>
    /// [ResponseTtc]
    /// </summary>
    public string Ttc { get; set; } = "{}";

    /// <summary>
    /// [ResponseMeta]
    /// </summary>
    public string Meta { get; set; } = "{}";

    /// <summary>
    /// Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
    /// or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
    /// from the client. Null means no ingest boundary has written it, which is not the same claim as
    /// an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
    /// and surfacing that in the UI (v2) would have nothing to read.
    /// [ResponseIngestFlags]
    /// </summary>
    public string? IngestFlags { get; set; }

    /// <summary>
    /// What the Cal.com elements reported when the respondent booked, `{ [Question ID]: booking }`. `data`
    /// keeps the string "booked" for those questions, so this column is purely additive: it is omitted from
    /// every read without a `select` (`PRISMA_GLOBAL_OMIT`), and only the webhook body (as
    /// `calBookingsByQuestionId`) and `/api/v3/responses` read it. Server-derived from the incoming `data`,
    /// never accepted as a field of its own.
    /// [ResponseCalBookings]
    /// </summary>
    public string? CalBookings { get; set; }

    /// <summary>
    /// A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
    /// export by default and from Feedback Records, and nothing else: it is never deleted and every
    /// count, list, webhook and integration still includes it (ENG-6780).
    /// </summary>
    public ResponseOutcome? Outcome { get; set; }

    /// <summary>
    /// What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
    /// from them, and per element its options, matrix rows and columns, by stable ids. Written by the
    /// survey while they answer; null means unknown (older responses and bundles), read as before.
    /// Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
    /// webhook or integration.
    /// [ResponseNotShown]
    /// </summary>
    public string? NotShown { get; set; }

    public virtual ICollection<TagsOnResponses> Tags { get; set; } = new List<TagsOnResponses>();

    public virtual ICollection<ResponseQuotaLink> QuotaLinks { get; set; } = new List<ResponseQuotaLink>();

    /// <summary>
    /// [ResponseContactAttributes]
    /// </summary>
    public string? ContactAttributes { get; set; }

    public string? SingleUseId { get; set; }

    public string? Language { get; set; }

    public string? DisplayId { get; set; }

    public virtual Display? Display { get; set; }

    public virtual ICollection<WorkflowRun> WorkflowRuns { get; set; } = new List<WorkflowRun>();
}
