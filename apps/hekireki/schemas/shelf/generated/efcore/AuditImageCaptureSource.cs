#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// How an audit photo reached Shelf, as reported by the client. The phone app
/// reports CAMERA only for its own camera, which cannot hand back an old
/// picture, so it is the strongest evidence short of a modified client; a
/// library photo's date is only what its file claims; the web app cannot tell
/// a camera shot from a file; PHONE is a phone app version that does not say
/// which.
/// </summary>
public enum AuditImageCaptureSource
{
    [PgName("CAMERA")]
    Camera,

    [PgName("LIBRARY")]
    Library,

    [PgName("WEB_UPLOAD")]
    WebUpload,

    [PgName("PHONE")]
    Phone,
}
