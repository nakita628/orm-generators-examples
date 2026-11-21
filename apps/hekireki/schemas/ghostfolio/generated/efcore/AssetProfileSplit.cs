#nullable enable

using System;

namespace Ghostfolio.Models;

/// <summary>
/// A stock split of an asset profile. The numerator is the number of shares
/// held after the split, the denominator the number of shares held before, for
/// example 4 and 1 for a 4:1 split or 1 and 10 for a 1:10 reverse split.
///
/// The ratio is stored as two integers instead of the resulting factor, so that
/// it stays exact, for example 1/3 for a 1:3 reverse split.
///
/// Activities before the date are adjusted by the ratio: the quantity is
/// multiplied by it and the unit price divided by it, so that the total value
/// of an activity stays the same.
///
/// Market data is already split-adjusted by the data providers and must not be
/// adjusted again.
/// </summary>
public partial class AssetProfileSplit
{
    public DateTime CreatedAt { get; set; }

    public DateTime Date { get; set; }

    public int Denominator { get; set; }

    public string Id { get; set; } = null!;

    public int Numerator { get; set; }

    public virtual SymbolProfile SymbolProfile { get; set; } = null!;

    public string SymbolProfileId { get; set; } = null!;

    public DateTime UpdatedAt { get; set; }
}
