#nullable enable

using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.ChangeTracking;
using Microsoft.EntityFrameworkCore.Metadata.Conventions;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using Microsoft.EntityFrameworkCore.ValueGeneration;
using Npgsql;
using Npgsql.EntityFrameworkCore.PostgreSQL.Infrastructure;
using Visus.Cuid;

namespace Shelf.Models;

public partial class ShelfContext : DbContext
{
    public ShelfContext(DbContextOptions<ShelfContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Image> Images { get; set; }

    public virtual DbSet<User> Users { get; set; }

    public virtual DbSet<UserContact> UserContacts { get; set; }

    public virtual DbSet<UserBusinessIntel> UserBusinessIntels { get; set; }

    public virtual DbSet<Asset> Assets { get; set; }

    public virtual DbSet<AssetFilterPreset> AssetFilterPresets { get; set; }

    public virtual DbSet<AssetIndexSettings> AssetIndexSettings { get; set; }

    public virtual DbSet<PdfTemplate> PdfTemplates { get; set; }

    public virtual DbSet<Category> Categories { get; set; }

    public virtual DbSet<AssetModel> AssetModels { get; set; }

    public virtual DbSet<BookingModelRequest> BookingModelRequests { get; set; }

    public virtual DbSet<Tag> Tags { get; set; }

    public virtual DbSet<Note> Notes { get; set; }

    public virtual DbSet<BookingNote> BookingNotes { get; set; }

    public virtual DbSet<LocationNote> LocationNotes { get; set; }

    public virtual DbSet<Qr> Qrs { get; set; }

    public virtual DbSet<Barcode> Barcodes { get; set; }

    public virtual DbSet<PrintBatch> PrintBatches { get; set; }

    public virtual DbSet<ReportFound> ReportFounds { get; set; }

    public virtual DbSet<Scan> Scans { get; set; }

    public virtual DbSet<Location> Locations { get; set; }

    public virtual DbSet<Role> Roles { get; set; }

    public virtual DbSet<TeamMember> TeamMembers { get; set; }

    public virtual DbSet<Custody> Custodies { get; set; }

    public virtual DbSet<Organization> Organizations { get; set; }

    public virtual DbSet<UserOrganization> UserOrganizations { get; set; }

    public virtual DbSet<SsoDetails> SsoDetails { get; set; }

    public virtual DbSet<ScimToken> ScimTokens { get; set; }

    public virtual DbSet<UserScimExternalId> UserScimExternalIds { get; set; }

    public virtual DbSet<Tier> Tiers { get; set; }

    public virtual DbSet<TierLimit> TierLimits { get; set; }

    public virtual DbSet<CustomTierLimit> CustomTierLimits { get; set; }

    public virtual DbSet<CustomField> CustomFields { get; set; }

    public virtual DbSet<AssetCustomFieldValue> AssetCustomFieldValues { get; set; }

    public virtual DbSet<AssetModelCustomFieldDefault> AssetModelCustomFieldDefaults { get; set; }

    public virtual DbSet<Invite> Invites { get; set; }

    public virtual DbSet<Announcement> Announcements { get; set; }

    public virtual DbSet<ConsumptionLog> ConsumptionLogs { get; set; }

    public virtual DbSet<BookingAsset> BookingAssets { get; set; }

    public virtual DbSet<Booking> Bookings { get; set; }

    public virtual DbSet<BookingSettings> BookingSettings { get; set; }

    public virtual DbSet<PartialBookingCheckin> PartialBookingCheckins { get; set; }

    public virtual DbSet<PartialBookingCheckout> PartialBookingCheckouts { get; set; }

    public virtual DbSet<MobileAuthCode> MobileAuthCodes { get; set; }

    public virtual DbSet<PushToken> PushTokens { get; set; }

    public virtual DbSet<NotificationPreference> NotificationPreferences { get; set; }

    public virtual DbSet<Kit> Kits { get; set; }

    public virtual DbSet<KitCustody> KitCustodies { get; set; }

    public virtual DbSet<AssetKit> AssetKits { get; set; }

    public virtual DbSet<AssetLocation> AssetLocations { get; set; }

    public virtual DbSet<AssetReminder> AssetReminders { get; set; }

    public virtual DbSet<WorkingHours> WorkingHours { get; set; }

    public virtual DbSet<WorkingHoursOverride> WorkingHoursOverrides { get; set; }

    public virtual DbSet<Update> Updates { get; set; }

    public virtual DbSet<UserUpdateRead> UserUpdateReads { get; set; }

    public virtual DbSet<AuditNote> AuditNotes { get; set; }

    public virtual DbSet<TeamMemberNote> TeamMemberNotes { get; set; }

    public virtual DbSet<AuditSession> AuditSessions { get; set; }

    public virtual DbSet<AssetImport> AssetImports { get; set; }

    public virtual DbSet<LinkPreview> LinkPreviews { get; set; }

    public virtual DbSet<AuditAssignment> AuditAssignments { get; set; }

    public virtual DbSet<AuditAsset> AuditAssets { get; set; }

    public virtual DbSet<AuditScan> AuditScans { get; set; }

    public virtual DbSet<AuditImage> AuditImages { get; set; }

    public virtual DbSet<RoleChangeLog> RoleChangeLogs { get; set; }

    public virtual DbSet<ActivityEvent> ActivityEvents { get; set; }

    public static void MapEnums(NpgsqlDbContextOptionsBuilder npgsql)
    {
        npgsql.MapEnum<AssetStatus>("AssetStatus");
        npgsql.MapEnum<AssetType>("AssetType");
        npgsql.MapEnum<ConsumptionType>("ConsumptionType");
        npgsql.MapEnum<ConsumptionCategory>("ConsumptionCategory");
        npgsql.MapEnum<PdfTemplateKind>("PdfTemplateKind");
        npgsql.MapEnum<AssetIndexMode>("AssetIndexMode");
        npgsql.MapEnum<TagUseFor>("TagUseFor");
        npgsql.MapEnum<NoteType>("NoteType");
        npgsql.MapEnum<BarcodeType>("BarcodeType");
        npgsql.MapEnum<ErrorCorrection>("ErrorCorrection");
        npgsql.MapEnum<ScanCodeType>("ScanCodeType");
        npgsql.MapEnum<ScanSource>("ScanSource");
        npgsql.MapEnum<Roles>("Roles");
        npgsql.MapEnum<OrganizationType>("OrganizationType");
        npgsql.MapEnum<QrIdDisplayPreference>("QrIdDisplayPreference");
        npgsql.MapEnum<DateFormatPreference>("DateFormatPreference");
        npgsql.MapEnum<TimeFormatPreference>("TimeFormatPreference");
        npgsql.MapEnum<WeekStartPreference>("WeekStartPreference");
        npgsql.MapEnum<OrganizationRoles>("OrganizationRoles");
        npgsql.MapEnum<TierId>("TierId");
        npgsql.MapEnum<CustomFieldType>("CustomFieldType");
        npgsql.MapEnum<CustomFieldFormat>("CustomFieldFormat");
        npgsql.MapEnum<Currency>("Currency");
        npgsql.MapEnum<InviteStatuses>("InviteStatuses");
        npgsql.MapEnum<BookingStatus>("BookingStatus");
        npgsql.MapEnum<DevicePlatform>("DevicePlatform");
        npgsql.MapEnum<KitStatus>("KitStatus");
        npgsql.MapEnum<ReminderRecurrenceUnit>("ReminderRecurrenceUnit");
        npgsql.MapEnum<UpdateStatus>("UpdateStatus");
        npgsql.MapEnum<AuditStatus>("AuditStatus");
        npgsql.MapEnum<AuditAssetStatus>("AuditAssetStatus");
        npgsql.MapEnum<AuditAssignmentRole>("AuditAssignmentRole");
        npgsql.MapEnum<AssetImportStatus>("AssetImportStatus");
        npgsql.MapEnum<AssetImportPhase>("AssetImportPhase");
        npgsql.MapEnum<LinkPreviewOutcome>("LinkPreviewOutcome");
        npgsql.MapEnum<AuditImageCaptureSource>("AuditImageCaptureSource");
        npgsql.MapEnum<RoleChangeSource>("RoleChangeSource");
        npgsql.MapEnum<ActivityEntity>("ActivityEntity");
        npgsql.MapEnum<ActivityAction>("ActivityAction");
    }

    public static void MapEnums(NpgsqlDataSourceBuilder dataSource)
    {
        dataSource.MapEnum<AssetStatus>("AssetStatus");
        dataSource.MapEnum<AssetType>("AssetType");
        dataSource.MapEnum<ConsumptionType>("ConsumptionType");
        dataSource.MapEnum<ConsumptionCategory>("ConsumptionCategory");
        dataSource.MapEnum<PdfTemplateKind>("PdfTemplateKind");
        dataSource.MapEnum<AssetIndexMode>("AssetIndexMode");
        dataSource.MapEnum<TagUseFor>("TagUseFor");
        dataSource.MapEnum<NoteType>("NoteType");
        dataSource.MapEnum<BarcodeType>("BarcodeType");
        dataSource.MapEnum<ErrorCorrection>("ErrorCorrection");
        dataSource.MapEnum<ScanCodeType>("ScanCodeType");
        dataSource.MapEnum<ScanSource>("ScanSource");
        dataSource.MapEnum<Roles>("Roles");
        dataSource.MapEnum<OrganizationType>("OrganizationType");
        dataSource.MapEnum<QrIdDisplayPreference>("QrIdDisplayPreference");
        dataSource.MapEnum<DateFormatPreference>("DateFormatPreference");
        dataSource.MapEnum<TimeFormatPreference>("TimeFormatPreference");
        dataSource.MapEnum<WeekStartPreference>("WeekStartPreference");
        dataSource.MapEnum<OrganizationRoles>("OrganizationRoles");
        dataSource.MapEnum<TierId>("TierId");
        dataSource.MapEnum<CustomFieldType>("CustomFieldType");
        dataSource.MapEnum<CustomFieldFormat>("CustomFieldFormat");
        dataSource.MapEnum<Currency>("Currency");
        dataSource.MapEnum<InviteStatuses>("InviteStatuses");
        dataSource.MapEnum<BookingStatus>("BookingStatus");
        dataSource.MapEnum<DevicePlatform>("DevicePlatform");
        dataSource.MapEnum<KitStatus>("KitStatus");
        dataSource.MapEnum<ReminderRecurrenceUnit>("ReminderRecurrenceUnit");
        dataSource.MapEnum<UpdateStatus>("UpdateStatus");
        dataSource.MapEnum<AuditStatus>("AuditStatus");
        dataSource.MapEnum<AuditAssetStatus>("AuditAssetStatus");
        dataSource.MapEnum<AuditAssignmentRole>("AuditAssignmentRole");
        dataSource.MapEnum<AssetImportStatus>("AssetImportStatus");
        dataSource.MapEnum<AssetImportPhase>("AssetImportPhase");
        dataSource.MapEnum<LinkPreviewOutcome>("LinkPreviewOutcome");
        dataSource.MapEnum<AuditImageCaptureSource>("AuditImageCaptureSource");
        dataSource.MapEnum<RoleChangeSource>("RoleChangeSource");
        dataSource.MapEnum<ActivityEntity>("ActivityEntity");
        dataSource.MapEnum<ActivityAction>("ActivityAction");
    }

    protected override void ConfigureConventions(ModelConfigurationBuilder configurationBuilder)
    {
        configurationBuilder.Conventions.Remove<ForeignKeyIndexConvention>();
        ConfigureConventionsPartial(configurationBuilder);
    }

    partial void ConfigureConventionsPartial(ModelConfigurationBuilder configurationBuilder);

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.HasPostgresEnum<AssetStatus>(name: "AssetStatus");
        modelBuilder.HasPostgresEnum<AssetType>(name: "AssetType");
        modelBuilder.HasPostgresEnum<ConsumptionType>(name: "ConsumptionType");
        modelBuilder.HasPostgresEnum<ConsumptionCategory>(name: "ConsumptionCategory");
        modelBuilder.HasPostgresEnum<PdfTemplateKind>(name: "PdfTemplateKind");
        modelBuilder.HasPostgresEnum<AssetIndexMode>(name: "AssetIndexMode");
        modelBuilder.HasPostgresEnum<TagUseFor>(name: "TagUseFor");
        modelBuilder.HasPostgresEnum<NoteType>(name: "NoteType");
        modelBuilder.HasPostgresEnum<BarcodeType>(name: "BarcodeType");
        modelBuilder.HasPostgresEnum<ErrorCorrection>(name: "ErrorCorrection");
        modelBuilder.HasPostgresEnum<ScanCodeType>(name: "ScanCodeType");
        modelBuilder.HasPostgresEnum<ScanSource>(name: "ScanSource");
        modelBuilder.HasPostgresEnum<Roles>(name: "Roles");
        modelBuilder.HasPostgresEnum<OrganizationType>(name: "OrganizationType");
        modelBuilder.HasPostgresEnum<QrIdDisplayPreference>(name: "QrIdDisplayPreference");
        modelBuilder.HasPostgresEnum<DateFormatPreference>(name: "DateFormatPreference");
        modelBuilder.HasPostgresEnum<TimeFormatPreference>(name: "TimeFormatPreference");
        modelBuilder.HasPostgresEnum<WeekStartPreference>(name: "WeekStartPreference");
        modelBuilder.HasPostgresEnum<OrganizationRoles>(name: "OrganizationRoles");
        modelBuilder.HasPostgresEnum<TierId>(name: "TierId");
        modelBuilder.HasPostgresEnum<CustomFieldType>(name: "CustomFieldType");
        modelBuilder.HasPostgresEnum<CustomFieldFormat>(name: "CustomFieldFormat");
        modelBuilder.HasPostgresEnum<Currency>(name: "Currency");
        modelBuilder.HasPostgresEnum<InviteStatuses>(name: "InviteStatuses");
        modelBuilder.HasPostgresEnum<BookingStatus>(name: "BookingStatus");
        modelBuilder.HasPostgresEnum<DevicePlatform>(name: "DevicePlatform");
        modelBuilder.HasPostgresEnum<KitStatus>(name: "KitStatus");
        modelBuilder.HasPostgresEnum<ReminderRecurrenceUnit>(name: "ReminderRecurrenceUnit");
        modelBuilder.HasPostgresEnum<UpdateStatus>(name: "UpdateStatus");
        modelBuilder.HasPostgresEnum<AuditStatus>(name: "AuditStatus");
        modelBuilder.HasPostgresEnum<AuditAssetStatus>(name: "AuditAssetStatus");
        modelBuilder.HasPostgresEnum<AuditAssignmentRole>(name: "AuditAssignmentRole");
        modelBuilder.HasPostgresEnum<AssetImportStatus>(name: "AssetImportStatus");
        modelBuilder.HasPostgresEnum<AssetImportPhase>(name: "AssetImportPhase");
        modelBuilder.HasPostgresEnum<LinkPreviewOutcome>(name: "LinkPreviewOutcome");
        modelBuilder.HasPostgresEnum<AuditImageCaptureSource>(name: "AuditImageCaptureSource");
        modelBuilder.HasPostgresEnum<RoleChangeSource>(name: "RoleChangeSource");
        modelBuilder.HasPostgresEnum<ActivityEntity>(name: "ActivityEntity");
        modelBuilder.HasPostgresEnum<ActivityAction>(name: "ActivityAction");

        modelBuilder.Entity<Image>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Image_pkey");

            entity.ToTable("Image");

            entity.HasIndex(e => e.OwnerOrgId, "Image_ownerOrgId_idx");

            entity.HasIndex(e => e.UserId, "Image_userId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.ContentType).HasColumnName("contentType");
            entity.Property(e => e.AltText).HasColumnName("altText");
            entity.Property(e => e.Blob).HasColumnName("blob");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.OwnerOrgId).HasColumnName("ownerOrgId");
            entity.Property(e => e.UserId).HasColumnName("userId");

            entity.HasOne(d => d.OwnerOrg).WithMany(p => p.Images)
                .HasForeignKey(d => d.OwnerOrgId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Image_ownerOrgId_fkey");

            entity.HasOne(d => d.User).WithMany(p => p.Images)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("Image_userId_fkey");
        });

        modelBuilder.Entity<User>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("User_pkey");

            entity.ToTable("User");

            entity.HasIndex(e => e.Email, "User_email_key").IsUnique();

            entity.HasIndex(e => e.Username, "User_username_key").IsUnique();

            entity.HasIndex(e => e.CustomerId, "User_customerId_key").IsUnique();

            entity.HasIndex(e => new { e.FirstName, e.LastName }, "User_firstName_lastName_idx");

            entity.HasIndex(e => e.TierId, "User_tierId_idx");

            entity.HasIndex(e => e.LastSelectedOrganizationId, "User_lastSelectedOrganizationId_idx");

            entity.HasIndex(e => e.LastMobileActiveAt, "User_lastMobileActiveAt_idx");

            entity.HasIndex(e => new { e.Email, e.Username }, "User_email_username_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Email).HasColumnName("email");
            entity.Property(e => e.Username)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("username");
            entity.Property(e => e.FirstName).HasColumnName("firstName");
            entity.Property(e => e.LastName).HasColumnName("lastName");
            entity.Property(e => e.DisplayName).HasColumnName("displayName");
            entity.Property(e => e.ProfilePicture).HasColumnName("profilePicture");
            entity.Property(e => e.UsedFreeTrial)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("usedFreeTrial");
            entity.Property(e => e.Onboarded)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("onboarded");
            entity.Property(e => e.CustomerId).HasColumnName("customerId");
            entity.Property(e => e.Sso)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("sso");
            entity.Property(e => e.CreatedWithInvite)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("createdWithInvite");
            entity.Property(e => e.SkipSubscriptionCheck)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("skipSubscriptionCheck");
            entity.Property(e => e.HasUnpaidInvoice)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("hasUnpaidInvoice");
            entity.Property(e => e.WarnForNoPaymentMethod)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("warnForNoPaymentMethod");
            entity.Property(e => e.LastSelectedOrganizationId).HasColumnName("lastSelectedOrganizationId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.DeletedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("deletedAt");
            entity.Property(e => e.LastMobileActiveAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("lastMobileActiveAt");
            entity.Property(e => e.DateFormat).HasColumnName("dateFormat");
            entity.Property(e => e.TimeFormat).HasColumnName("timeFormat");
            entity.Property(e => e.WeekStart).HasColumnName("weekStart");
            entity.Property(e => e.TimeZone).HasColumnName("timeZone");
            entity.Property(e => e.TierId)
                .ValueGeneratedNever()
                .HasDefaultValue(TierId.Free)
                .HasColumnName("tierId");
            entity.Property(e => e.ReferralSource).HasColumnName("referralSource");

            entity.HasOne(d => d.LastSelectedOrganization).WithMany(p => p.UsersWithLastSelected)
                .HasForeignKey(d => d.LastSelectedOrganizationId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("User_lastSelectedOrganizationId_fkey");

            entity.HasOne(d => d.Tier).WithMany(p => p.Subscribers)
                .HasForeignKey(d => d.TierId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("User_tierId_fkey");
        });

        modelBuilder.Entity<UserContact>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("UserContact_pkey");

            entity.ToTable("UserContact");

            entity.HasIndex(e => e.UserId, "UserContact_userId_key").IsUnique();

            entity.HasIndex(e => e.UserId, "UserContact_userId_idx");

            entity.HasIndex(e => e.Phone, "UserContact_phone_idx");

            entity.HasIndex(e => new { e.City, e.StateProvince }, "UserContact_city_stateProvince_idx");

            entity.HasIndex(e => e.CountryRegion, "UserContact_countryRegion_idx");

            entity.HasIndex(e => e.ZipPostalCode, "UserContact_zipPostalCode_idx");

            entity.HasIndex(e => new { e.City, e.CountryRegion }, "UserContact_city_countryRegion_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Phone).HasColumnName("phone");
            entity.Property(e => e.Street).HasColumnName("street");
            entity.Property(e => e.City).HasColumnName("city");
            entity.Property(e => e.StateProvince).HasColumnName("stateProvince");
            entity.Property(e => e.ZipPostalCode).HasColumnName("zipPostalCode");
            entity.Property(e => e.CountryRegion).HasColumnName("countryRegion");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithOne(p => p.Contact)
                .HasForeignKey<UserContact>(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("UserContact_userId_fkey");
        });

        modelBuilder.Entity<UserBusinessIntel>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("UserBusinessIntel_pkey");

            entity.ToTable("UserBusinessIntel");

            entity.HasIndex(e => e.UserId, "UserBusinessIntel_userId_key").IsUnique();

            entity.HasIndex(e => e.UserId, "UserBusinessIntel_userId_idx");

            entity.HasIndex(e => e.CompanyName, "UserBusinessIntel_companyName_idx");

            entity.HasIndex(e => e.JobTitle, "UserBusinessIntel_jobTitle_idx");

            entity.HasIndex(e => e.TeamSize, "UserBusinessIntel_teamSize_idx");

            entity.HasIndex(e => e.UtmSource, "UserBusinessIntel_utmSource_idx");

            entity.HasIndex(e => e.UtmCampaign, "UserBusinessIntel_utmCampaign_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.HowDidYouHearAboutUs).HasColumnName("howDidYouHearAboutUs");
            entity.Property(e => e.JobTitle).HasColumnName("jobTitle");
            entity.Property(e => e.TeamSize).HasColumnName("teamSize");
            entity.Property(e => e.CompanyName).HasColumnName("companyName");
            entity.Property(e => e.PrimaryUseCase).HasColumnName("primaryUseCase");
            entity.Property(e => e.CurrentSolution).HasColumnName("currentSolution");
            entity.Property(e => e.Timeline).HasColumnName("timeline");
            entity.Property(e => e.SignupPlan).HasColumnName("signupPlan");
            entity.Property(e => e.SignupTrial).HasColumnName("signupTrial");
            entity.Property(e => e.UtmSource).HasColumnName("utmSource");
            entity.Property(e => e.UtmMedium).HasColumnName("utmMedium");
            entity.Property(e => e.UtmCampaign).HasColumnName("utmCampaign");
            entity.Property(e => e.UtmContent).HasColumnName("utmContent");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithOne(p => p.BusinessIntel)
                .HasForeignKey<UserBusinessIntel>(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("UserBusinessIntel_userId_fkey");
        });

        modelBuilder.Entity<Asset>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Asset_pkey");

            entity.ToTable("Asset");

            entity.HasIndex(e => e.PreferredBarcodeId, "Asset_preferredBarcodeId_key").IsUnique();

            entity.HasIndex(e => new { e.Title, e.Description }, "Asset_title_description_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops", "gin_trgm_ops");

            entity.HasIndex(e => e.SequentialId, "Asset_sequentialId_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.HasIndex(e => new { e.OrganizationId, e.Title, e.Status, e.AvailableToBook }, "Asset_organizationId_compound_idx");

            entity.HasIndex(e => new { e.Status, e.OrganizationId }, "Asset_status_organizationId_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.ArchivedAt }, "Asset_organizationId_archivedAt_idx");

            entity.HasIndex(e => new { e.CreatedAt, e.OrganizationId }, "Asset_createdAt_organizationId_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.CreatedAt, e.Id }, "Asset_organizationId_createdAt_id_idx").IsDescending(false, true, false);

            entity.HasIndex(e => new { e.Valuation, e.OrganizationId }, "Asset_valuation_organizationId_idx");

            entity.HasIndex(e => new { e.CategoryId, e.OrganizationId }, "Asset_categoryId_organizationId_idx");

            entity.HasIndex(e => new { e.AssetModelId, e.OrganizationId }, "Asset_assetModelId_organizationId_idx");

            entity.HasIndex(e => e.SequentialId, "Asset_sequentialId_idx");

            entity.HasIndex(e => e.UserId, "Asset_userId_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.SequentialId }, "Asset_organizationId_sequentialId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Title).HasColumnName("title");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.MainImage).HasColumnName("mainImage");
            entity.Property(e => e.ThumbnailImage).HasColumnName("thumbnailImage");
            entity.Property(e => e.MainImageExpiration)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("mainImageExpiration");
            entity.Property(e => e.Status)
                .ValueGeneratedNever()
                .HasDefaultValue(AssetStatus.Available)
                .HasColumnName("status");
            entity.Property(e => e.Valuation).HasColumnName("value");
            entity.Property(e => e.AvailableToBook)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("availableToBook");
            entity.Property(e => e.SequentialId).HasColumnName("sequentialId");
            entity.Property(e => e.ArchivedAt)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("archivedAt");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(AssetType.Individual)
                .HasColumnName("type");
            entity.Property(e => e.AssetModelId).HasColumnName("assetModelId");
            entity.Property(e => e.Quantity).HasColumnName("quantity");
            entity.Property(e => e.MinQuantity).HasColumnName("minQuantity");
            entity.Property(e => e.ConsumptionType).HasColumnName("consumptionType");
            entity.Property(e => e.UnitOfMeasure).HasColumnName("unitOfMeasure");
            entity.Property(e => e.LowStockNotifiedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("lowStockNotifiedAt");
            entity.Property(e => e.StockLedgerStartedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("stockLedgerStartedAt");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CategoryId).HasColumnName("categoryId");
            entity.Property(e => e.PreferredBarcodeId).HasColumnName("preferredBarcodeId");

            entity.HasOne(d => d.AssetModel).WithMany(p => p.Assets)
                .HasForeignKey(d => d.AssetModelId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Asset_assetModelId_fkey");

            entity.HasOne(d => d.User).WithMany(p => p.Assets)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Asset_userId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.Assets)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Asset_organizationId_fkey");

            entity.HasOne(d => d.Category).WithMany(p => p.Assets)
                .HasForeignKey(d => d.CategoryId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Asset_categoryId_fkey");

            entity.HasOne(d => d.PreferredBarcode).WithOne(p => p.PreferredByAsset)
                .HasForeignKey<Asset>(d => d.PreferredBarcodeId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Asset_preferredBarcodeId_fkey");

            entity.HasMany(d => d.Tags).WithMany(p => p.Assets)
                .UsingEntity<Dictionary<string, object>>(
                    "_AssetToTag",
                    r => r.HasOne<Tag>().WithMany()
                        .HasForeignKey("B")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_AssetToTag_B_fkey"),
                    l => l.HasOne<Asset>().WithMany()
                        .HasForeignKey("A")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_AssetToTag_A_fkey"),
                    j =>
                    {
                        j.HasKey("A", "B").HasName("_AssetToTag_AB_pkey");
                        j.ToTable("_AssetToTag");
                        j.HasIndex(new[] { "B" }, "_AssetToTag_B_index");
                    });
        });

        modelBuilder.Entity<AssetFilterPreset>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetFilterPreset_pkey");

            entity.ToTable("AssetFilterPreset");

            entity.HasIndex(e => new { e.OrganizationId, e.OwnerId }, "asset_filter_presets_owner_lookup_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.OwnerId, e.Name }, "asset_filter_presets_owner_name_unique").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.OwnerId).HasColumnName("ownerId");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Query).HasColumnName("query");
            entity.Property(e => e.Starred)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("starred");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Organization).WithMany(p => p.AssetFilterPresets)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetFilterPreset_organizationId_fkey");

            entity.HasOne(d => d.Owner).WithMany(p => p.AssetFilterPresets)
                .HasForeignKey(d => d.OwnerId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetFilterPreset_ownerId_fkey");
        });

        modelBuilder.Entity<AssetIndexSettings>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetIndexSettings_pkey");

            entity.ToTable("AssetIndexSettings");

            entity.HasIndex(e => e.OrganizationId, "AssetIndexSettings_organizationId_idx");

            entity.HasIndex(e => new { e.UserId, e.OrganizationId }, "AssetIndexSettings_userId_organizationId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.Mode)
                .ValueGeneratedNever()
                .HasDefaultValue(AssetIndexMode.Simple)
                .HasColumnName("mode");
            entity.Property(e => e.Columns)
                .ValueGeneratedNever()
                .HasDefaultValue("[ {\"name\": \"id\", \"visible\": true, \"position\": 0}, {\"name\": \"status\", \"visible\": true, \"position\": 1}, {\"name\": \"description\", \"visible\": true, \"position\": 2}, {\"name\": \"valuation\", \"visible\": true, \"position\": 3}, {\"name\": \"createdAt\", \"visible\": true, \"position\": 4}, {\"name\": \"category\", \"visible\": true, \"position\": 5}, {\"name\": \"tags\", \"visible\": true, \"position\": 6}, {\"name\": \"location\", \"visible\": true, \"position\": 7}, {\"name\": \"kit\", \"visible\": true, \"position\": 8}, {\"name\": \"custody\", \"visible\": true, \"position\": 9} ]")
                .HasColumnType("jsonb")
                .HasColumnName("columns");
            entity.Property(e => e.FreezeColumn)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("freezeColumn");
            entity.Property(e => e.ShowAssetImage)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("showAssetImage");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithMany(p => p.AssetIndexSettings)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetIndexSettings_userId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.AssetIndexSettings)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetIndexSettings_organizationId_fkey");
        });

        modelBuilder.Entity<PdfTemplate>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("PdfTemplate_pkey");

            entity.ToTable("PdfTemplate");

            entity.HasIndex(e => new { e.OrganizationId, e.Kind, e.Name }, "PdfTemplate_organizationId_kind_name_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Kind)
                .ValueGeneratedNever()
                .HasDefaultValue(PdfTemplateKind.Booking)
                .HasColumnName("kind");
            entity.Property(e => e.Config)
                .ValueGeneratedNever()
                .HasDefaultValue("{}")
                .HasColumnType("jsonb")
                .HasColumnName("config");
            entity.Property(e => e.IsDefault)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("isDefault");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Organization).WithMany(p => p.PdfTemplates)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("PdfTemplate_organizationId_fkey");
        });

        modelBuilder.Entity<Category>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Category_pkey");

            entity.ToTable("Category");

            entity.HasIndex(e => e.OrganizationId, "Category_organizationId_idx");

            entity.HasIndex(e => e.UserId, "Category_userId_idx");

            entity.HasIndex(e => e.Name, "Category_name_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.Color).HasColumnName("color");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");

            entity.HasOne(d => d.User).WithMany(p => p.Categories)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("Category_userId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.Categories)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Category_organizationId_fkey");

            entity.HasMany(d => d.CustomFields).WithMany(p => p.Categories)
                .UsingEntity<Dictionary<string, object>>(
                    "_CategoryToCustomField",
                    r => r.HasOne<CustomField>().WithMany()
                        .HasForeignKey("B")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_CategoryToCustomField_B_fkey"),
                    l => l.HasOne<Category>().WithMany()
                        .HasForeignKey("A")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_CategoryToCustomField_A_fkey"),
                    j =>
                    {
                        j.HasKey("A", "B").HasName("_CategoryToCustomField_AB_pkey");
                        j.ToTable("_CategoryToCustomField");
                        j.HasIndex(new[] { "B" }, "_CategoryToCustomField_B_index");
                    });
        });

        modelBuilder.Entity<AssetModel>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetModel_pkey");

            entity.ToTable("AssetModel");

            entity.HasIndex(e => new { e.OrganizationId, e.Name }, "AssetModel_organizationId_name_idx");

            entity.HasIndex(e => e.UserId, "AssetModel_userId_idx");

            entity.HasIndex(e => e.Name, "AssetModel_name_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.Image).HasColumnName("image");
            entity.Property(e => e.ThumbnailImage).HasColumnName("thumbnailImage");
            entity.Property(e => e.DefaultCategoryId).HasColumnName("defaultCategoryId");
            entity.Property(e => e.DefaultValuation).HasColumnName("defaultValuation");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.DefaultCategory).WithMany(p => p.AssetModelDefaults)
                .HasForeignKey(d => d.DefaultCategoryId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AssetModel_defaultCategoryId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.AssetModels)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetModel_organizationId_fkey");

            entity.HasOne(d => d.CreatedBy).WithMany(p => p.AssetModels)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetModel_userId_fkey");
        });

        modelBuilder.Entity<BookingModelRequest>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("BookingModelRequest_pkey");

            entity.ToTable("BookingModelRequest");

            entity.HasIndex(e => e.AssetModelId, "BookingModelRequest_assetModelId_idx");

            entity.HasIndex(e => new { e.BookingId, e.AssetModelId }, "BookingModelRequest_bookingId_assetModelId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.BookingId).HasColumnName("bookingId");
            entity.Property(e => e.AssetModelId).HasColumnName("assetModelId");
            entity.Property(e => e.Quantity).HasColumnName("quantity");
            entity.Property(e => e.FulfilledQuantity)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("fulfilledQuantity");
            entity.Property(e => e.FulfilledAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("fulfilledAt");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Booking).WithMany(p => p.ModelRequests)
                .HasForeignKey(d => d.BookingId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("BookingModelRequest_bookingId_fkey");

            entity.HasOne(d => d.AssetModel).WithMany(p => p.ModelRequests)
                .HasForeignKey(d => d.AssetModelId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("BookingModelRequest_assetModelId_fkey");
        });

        modelBuilder.Entity<Tag>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Tag_pkey");

            entity.ToTable("Tag");

            entity.HasIndex(e => e.OrganizationId, "Tag_organizationId_idx");

            entity.HasIndex(e => e.UserId, "Tag_userId_idx");

            entity.HasIndex(e => e.Name, "Tag_name_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.Color).HasColumnName("color");
            entity.Property(e => e.UseFor)
                .ValueGeneratedNever()
                .HasDefaultValue(new List<TagUseFor> { TagUseFor.Asset })
                .HasColumnName("useFor");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithMany(p => p.Tags)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("Tag_userId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.Tags)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Tag_organizationId_fkey");
        });

        modelBuilder.Entity<Note>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Note_pkey");

            entity.ToTable("Note");

            entity.HasIndex(e => e.AssetId, "Note_assetId_idx");

            entity.HasIndex(e => e.UserId, "Note_userId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(NoteType.Comment)
                .HasColumnName("type");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.AssetId).HasColumnName("assetId");

            entity.HasOne(d => d.User).WithMany(p => p.Notes)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Note_userId_fkey");

            entity.HasOne(d => d.Asset).WithMany(p => p.Notes)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Note_assetId_fkey");
        });

        modelBuilder.Entity<BookingNote>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("BookingNote_pkey");

            entity.ToTable("BookingNote");

            entity.HasIndex(e => e.BookingId, "BookingNote_bookingId_idx");

            entity.HasIndex(e => e.UserId, "BookingNote_userId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(NoteType.Comment)
                .HasColumnName("type");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.BookingId).HasColumnName("bookingId");

            entity.HasOne(d => d.User).WithMany(p => p.BookingNotes)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("BookingNote_userId_fkey");

            entity.HasOne(d => d.Booking).WithMany(p => p.Notes)
                .HasForeignKey(d => d.BookingId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("BookingNote_bookingId_fkey");
        });

        modelBuilder.Entity<LocationNote>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("LocationNote_pkey");

            entity.ToTable("LocationNote");

            entity.HasIndex(e => e.LocationId, "LocationNote_locationId_idx");

            entity.HasIndex(e => e.UserId, "LocationNote_userId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(NoteType.Comment)
                .HasColumnName("type");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.LocationId).HasColumnName("locationId");

            entity.HasOne(d => d.User).WithMany(p => p.LocationNotes)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("LocationNote_userId_fkey");

            entity.HasOne(d => d.Location).WithMany(p => p.Notes)
                .HasForeignKey(d => d.LocationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("LocationNote_locationId_fkey");
        });

        modelBuilder.Entity<Qr>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Qr_pkey");

            entity.ToTable("Qr");

            entity.HasIndex(e => e.AssetId, "Qr_assetId_idx");

            entity.HasIndex(e => e.KitId, "Qr_kitId_idx");

            entity.HasIndex(e => e.UserId, "Qr_userId_idx");

            entity.HasIndex(e => e.OrganizationId, "Qr_organizationId_idx");

            entity.HasIndex(e => e.BatchId, "Qr_batchId_idx");

            entity.HasIndex(e => e.Id, "Qr_id_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Version)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("version");
            entity.Property(e => e.ErrorCorrection)
                .ValueGeneratedNever()
                .HasDefaultValue(ErrorCorrection.L)
                .HasColumnName("errorCorrection");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.KitId).HasColumnName("kitId");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.BatchId).HasColumnName("batchId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Asset).WithMany(p => p.QrCodes)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Qr_assetId_fkey");

            entity.HasOne(d => d.Kit).WithMany(p => p.QrCodes)
                .HasForeignKey(d => d.KitId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Qr_kitId_fkey");

            entity.HasOne(d => d.User).WithMany(p => p.QrCodes)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Qr_userId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.QrCodes)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Qr_organizationId_fkey");

            entity.HasOne(d => d.Batch).WithMany(p => p.QrCodes)
                .HasForeignKey(d => d.BatchId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Qr_batchId_fkey");
        });

        modelBuilder.Entity<Barcode>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Barcode_pkey");

            entity.ToTable("Barcode");

            entity.HasIndex(e => new { e.OrganizationId, e.Value }, "Barcode_organizationId_value_idx");

            entity.HasIndex(e => e.AssetId, "Barcode_assetId_idx");

            entity.HasIndex(e => e.KitId, "Barcode_kitId_idx");

            entity.HasIndex(e => e.OrganizationId, "Barcode_organizationId_idx");

            entity.HasIndex(e => e.Value, "Barcode_value_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.HasIndex(e => new { e.OrganizationId, e.Value }, "Barcode_organizationId_value_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Value).HasColumnName("value");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(BarcodeType.Code128)
                .HasColumnName("type");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.KitId).HasColumnName("kitId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Asset).WithMany(p => p.Barcodes)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Barcode_assetId_fkey");

            entity.HasOne(d => d.Kit).WithMany(p => p.Barcodes)
                .HasForeignKey(d => d.KitId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Barcode_kitId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.Barcodes)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Barcode_organizationId_fkey");
        });

        modelBuilder.Entity<PrintBatch>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("PrintBatch_pkey");

            entity.ToTable("PrintBatch");

            entity.HasIndex(e => e.Name, "PrintBatch_name_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Printed)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("printed");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
        });

        modelBuilder.Entity<ReportFound>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("ReportFound_pkey");

            entity.ToTable("ReportFound");

            entity.HasIndex(e => e.AssetId, "ReportFound_assetId_idx");

            entity.HasIndex(e => e.KitId, "ReportFound_kitId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Email).HasColumnName("email");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.KitId).HasColumnName("kitId");

            entity.HasOne(d => d.Asset).WithMany(p => p.Reports)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ReportFound_assetId_fkey");

            entity.HasOne(d => d.Kit).WithMany(p => p.Reports)
                .HasForeignKey(d => d.KitId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ReportFound_kitId_fkey");
        });

        modelBuilder.Entity<Scan>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Scan_pkey");

            entity.ToTable("Scan");

            entity.HasIndex(e => e.QrId, "Scan_qrId_idx");

            entity.HasIndex(e => e.UserId, "Scan_userId_idx");

            entity.HasIndex(e => new { e.RawQrId, e.CreatedAt }, "Scan_rawQrId_createdAt_idx");

            entity.HasIndex(e => new { e.AssetId, e.CreatedAt }, "Scan_assetId_createdAt_idx");

            entity.HasIndex(e => new { e.KitId, e.CreatedAt }, "Scan_kitId_createdAt_idx");

            entity.HasIndex(e => e.BarcodeId, "Scan_barcodeId_idx");

            entity.HasIndex(e => e.OrganizationId, "Scan_organizationId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Latitude).HasColumnName("latitude");
            entity.Property(e => e.Longitude).HasColumnName("longitude");
            entity.Property(e => e.UserAgent).HasColumnName("userAgent");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.QrId).HasColumnName("qrId");
            entity.Property(e => e.RawQrId).HasColumnName("rawQrId");
            entity.Property(e => e.CodeType)
                .ValueGeneratedNever()
                .HasDefaultValue(ScanCodeType.Qr)
                .HasColumnName("codeType");
            entity.Property(e => e.Code).HasColumnName("code");
            entity.Property(e => e.BarcodeId).HasColumnName("barcodeId");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.KitId).HasColumnName("kitId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.Source).HasColumnName("source");
            entity.Property(e => e.ManuallyGenerated)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("manuallyGenerated");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithMany(p => p.Scans)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Scan_userId_fkey");

            entity.HasOne(d => d.Qr).WithMany(p => p.Scans)
                .HasForeignKey(d => d.QrId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Scan_qrId_fkey");

            entity.HasOne(d => d.Barcode).WithMany(p => p.Scans)
                .HasForeignKey(d => d.BarcodeId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Scan_barcodeId_fkey");

            entity.HasOne(d => d.Asset).WithMany(p => p.Scans)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Scan_assetId_fkey");

            entity.HasOne(d => d.Kit).WithMany(p => p.Scans)
                .HasForeignKey(d => d.KitId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Scan_kitId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.Scans)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Scan_organizationId_fkey");
        });

        modelBuilder.Entity<Location>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Location_pkey");

            entity.ToTable("Location");

            entity.HasIndex(e => e.ImageId, "Location_imageId_key").IsUnique();

            entity.HasIndex(e => e.OrganizationId, "Location_organizationId_idx");

            entity.HasIndex(e => e.UserId, "Location_userId_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.ParentId }, "Location_organizationId_parentId_idx");

            entity.HasIndex(e => e.Name, "Location_name_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.HasIndex(e => e.Description, "Location_description_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.HasIndex(e => e.Address, "Location_address_trgm_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.Address).HasColumnName("address");
            entity.Property(e => e.Latitude).HasColumnName("latitude");
            entity.Property(e => e.Longitude).HasColumnName("longitude");
            entity.Property(e => e.ImageUrl).HasColumnName("imageUrl");
            entity.Property(e => e.ThumbnailUrl).HasColumnName("thumbnailUrl");
            entity.Property(e => e.ImageId).HasColumnName("imageId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.ParentId).HasColumnName("parentId");

            entity.HasOne(d => d.Image).WithOne(p => p.Location)
                .HasForeignKey<Location>(d => d.ImageId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Location_imageId_fkey");

            entity.HasOne(d => d.User).WithMany(p => p.Locations)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Location_userId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.Locations)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Location_organizationId_fkey");

            entity.HasOne(d => d.Parent).WithMany(p => p.Children)
                .HasForeignKey(d => d.ParentId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Location_parentId_fkey");
        });

        modelBuilder.Entity<Role>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Role_pkey");

            entity.ToTable("Role");

            entity.HasIndex(e => e.Name, "Role_name_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name)
                .ValueGeneratedNever()
                .HasDefaultValue(global::Shelf.Models.Roles.User)
                .HasColumnName("name");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasMany(d => d.Users).WithMany(p => p.Roles)
                .UsingEntity<Dictionary<string, object>>(
                    "_RoleToUser",
                    r => r.HasOne<User>().WithMany()
                        .HasForeignKey("B")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_RoleToUser_B_fkey"),
                    l => l.HasOne<Role>().WithMany()
                        .HasForeignKey("A")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_RoleToUser_A_fkey"),
                    j =>
                    {
                        j.HasKey("A", "B").HasName("_RoleToUser_AB_pkey");
                        j.ToTable("_RoleToUser");
                        j.HasIndex(new[] { "B" }, "_RoleToUser_B_index");
                    });
        });

        modelBuilder.Entity<TeamMember>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("TeamMember_pkey");

            entity.ToTable("TeamMember");

            entity.HasIndex(e => e.Name, "TeamMember_name_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops");

            entity.HasIndex(e => e.OrganizationId, "TeamMember_organizationId_idx");

            entity.HasIndex(e => e.UserId, "TeamMember_userId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.DeletedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("deletedAt");

            entity.HasOne(d => d.Organization).WithMany(p => p.Members)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("TeamMember_organizationId_fkey");

            entity.HasOne(d => d.User).WithMany(p => p.TeamMembers)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("TeamMember_userId_fkey");
        });

        modelBuilder.Entity<Custody>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Custody_pkey");

            entity.ToTable("Custody");

            entity.HasIndex(e => new { e.AssetId, e.TeamMemberId }, "Custody_assetId_teamMemberId_idx");

            entity.HasIndex(e => e.TeamMemberId, "Custody_teamMemberId_idx");

            entity.HasIndex(e => e.KitCustodyId, "Custody_kitCustodyId_idx");

            entity.HasIndex(e => e.LocationId, "Custody_locationId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.TeamMemberId).HasColumnName("teamMemberId");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.KitCustodyId).HasColumnName("kitCustodyId");
            entity.Property(e => e.Quantity)
                .ValueGeneratedNever()
                .HasDefaultValue(1)
                .HasColumnName("quantity");
            entity.Property(e => e.LocationId).HasColumnName("locationId");
            entity.Property(e => e.SourceUnknown)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("sourceUnknown");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Custodian).WithMany(p => p.Custodies)
                .HasForeignKey(d => d.TeamMemberId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("Custody_teamMemberId_fkey");

            entity.HasOne(d => d.Asset).WithMany(p => p.Custody)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Custody_assetId_fkey");

            entity.HasOne(d => d.KitCustody).WithMany(p => p.InheritedCustody)
                .HasForeignKey(d => d.KitCustodyId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Custody_kitCustodyId_fkey");

            entity.HasOne(d => d.Location).WithMany(p => p.Custodies)
                .HasForeignKey(d => d.LocationId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Custody_locationId_fkey");
        });

        modelBuilder.Entity<Organization>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Organization_pkey");

            entity.ToTable("Organization");

            entity.HasIndex(e => e.ImageId, "Organization_imageId_key").IsUnique();

            entity.HasIndex(e => e.UserId, "Organization_userId_idx");

            entity.HasIndex(e => e.SsoDetailsId, "Organization_ssoDetailsId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name)
                .ValueGeneratedNever()
                .HasDefaultValue("Personal")
                .HasColumnName("name");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(OrganizationType.Personal)
                .HasColumnName("type");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.Currency)
                .ValueGeneratedNever()
                .HasDefaultValue(Currency.Usd)
                .HasColumnName("currency");
            entity.Property(e => e.ImageId).HasColumnName("imageId");
            entity.Property(e => e.EnabledSso)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("enabledSso");
            entity.Property(e => e.SsoDetailsId).HasColumnName("ssoDetailsId");
            entity.Property(e => e.SelfServiceCanSeeCustody)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("selfServiceCanSeeCustody");
            entity.Property(e => e.SelfServiceCanSeeBookings)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("selfServiceCanSeeBookings");
            entity.Property(e => e.BaseUserCanSeeCustody)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("baseUserCanSeeCustody");
            entity.Property(e => e.BaseUserCanSeeBookings)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("baseUserCanSeeBookings");
            entity.Property(e => e.ManagersCanEditAssets)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("managersCanEditAssets");
            entity.Property(e => e.BarcodesEnabled)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("barcodesEnabled");
            entity.Property(e => e.BarcodesEnabledAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("barcodesEnabledAt");
            entity.Property(e => e.UsedBarcodeTrial)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("usedBarcodeTrial");
            entity.Property(e => e.AuditsEnabled)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("auditsEnabled");
            entity.Property(e => e.AuditsEnabledAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("auditsEnabledAt");
            entity.Property(e => e.UsedAuditTrial)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("usedAuditTrial");
            entity.Property(e => e.WorkspaceDisabled)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("workspaceDisabled");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.HasSequentialIdsMigrated)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("hasSequentialIdsMigrated");
            entity.Property(e => e.QrIdDisplayPreference)
                .ValueGeneratedNever()
                .HasDefaultValue(QrIdDisplayPreference.QrId)
                .HasColumnName("qrIdDisplayPreference");
            entity.Property(e => e.ShowShelfBranding)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("showShelfBranding");
            entity.Property(e => e.ShowQrCodesOnPdfs)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("showQrCodesOnPdfs");
            entity.Property(e => e.CustomEmailFooter)
                .HasMaxLength(500)
                .HasColumnName("customEmailFooter");

            entity.HasOne(d => d.Owner).WithMany(p => p.Organizations)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Organization_userId_fkey");

            entity.HasOne(d => d.Image).WithOne(p => p.Organization)
                .HasForeignKey<Organization>(d => d.ImageId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Organization_imageId_fkey");

            entity.HasOne(d => d.SsoDetails).WithMany(p => p.Organizations)
                .HasForeignKey(d => d.SsoDetailsId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Organization_ssoDetailsId_fkey");
        });

        modelBuilder.Entity<UserOrganization>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("UserOrganization_pkey");

            entity.ToTable("UserOrganization");

            entity.HasIndex(e => e.CalendarTokenId, "UserOrganization_calendarTokenId_key").IsUnique();

            entity.HasIndex(e => e.OrganizationId, "UserOrganization_organizationId_idx");

            entity.HasIndex(e => new { e.UserId, e.OrganizationId }, "UserOrganization_userId_organizationId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.Roles).HasColumnName("roles");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.CalendarTokenId).HasColumnName("calendarTokenId");

            entity.HasOne(d => d.User).WithMany(p => p.UserOrganizations)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("UserOrganization_userId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.UserOrganizations)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("UserOrganization_organizationId_fkey");
        });

        modelBuilder.Entity<SsoDetails>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("SsoDetails_pkey");

            entity.ToTable("SsoDetails");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Domain).HasColumnName("domain");
            entity.Property(e => e.BaseUserGroupId).HasColumnName("baseUserGroupId");
            entity.Property(e => e.SelfServiceGroupId).HasColumnName("selfServiceGroupId");
            entity.Property(e => e.AdminGroupId).HasColumnName("adminGroupId");
            entity.Property(e => e.ManagerGroupId).HasColumnName("managerGroupId");
            entity.Property(e => e.RequireSsoLogin)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("requireSsoLogin");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
        });

        modelBuilder.Entity<ScimToken>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("ScimToken_pkey");

            entity.ToTable("ScimToken");

            entity.HasIndex(e => e.TokenHash, "ScimToken_tokenHash_key").IsUnique();

            entity.HasIndex(e => e.OrganizationId, "ScimToken_organizationId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.TokenHash).HasColumnName("tokenHash");
            entity.Property(e => e.Label).HasColumnName("label");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedById).HasColumnName("createdById");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.LastUsedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("lastUsedAt");

            entity.HasOne(d => d.Organization).WithMany(p => p.ScimTokens)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ScimToken_organizationId_fkey");
        });

        modelBuilder.Entity<UserScimExternalId>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("UserScimExternalId_pkey");

            entity.ToTable("UserScimExternalId");

            entity.HasIndex(e => e.OrganizationId, "UserScimExternalId_organizationId_idx");

            entity.HasIndex(e => new { e.UserId, e.OrganizationId }, "UserScimExternalId_userId_organizationId_key").IsUnique();

            entity.HasIndex(e => new { e.OrganizationId, e.ScimExternalId }, "UserScimExternalId_organizationId_scimExternalId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.ScimExternalId).HasColumnName("scimExternalId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithMany(p => p.ScimExternalIds)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("UserScimExternalId_userId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.ScimExternalIds)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("UserScimExternalId_organizationId_fkey");
        });

        modelBuilder.Entity<Tier>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Tier_pkey");

            entity.ToTable("Tier");

            entity.HasIndex(e => e.TierLimitId, "Tier_tierLimitId_key").IsUnique();

            entity.Property(e => e.Id).HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.TierLimitId).HasColumnName("tierLimitId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.TierLimit).WithOne(p => p.Tier)
                .HasForeignKey<Tier>(d => d.TierLimitId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Tier_tierLimitId_fkey");
        });

        modelBuilder.Entity<TierLimit>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("TierLimit_pkey");

            entity.ToTable("TierLimit");

            entity.Property(e => e.Id).HasColumnName("id");
            entity.Property(e => e.CanImportAssets)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("canImportAssets");
            entity.Property(e => e.CanExportAssets)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("canExportAssets");
            entity.Property(e => e.CanImportNRM)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("canImportNRM");
            entity.Property(e => e.CanHideShelfBranding)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("canHideShelfBranding");
            entity.Property(e => e.CanUseReports)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("canUseReports");
            entity.Property(e => e.CanUseRecurringReminders)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("canUseRecurringReminders");
            entity.Property(e => e.CanUsePushNotifications)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("canUsePushNotifications");
            entity.Property(e => e.MaxCustomFields)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("maxCustomFields");
            entity.Property(e => e.MaxOrganizations)
                .ValueGeneratedNever()
                .HasDefaultValue(1)
                .HasColumnName("maxOrganizations");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
        });

        modelBuilder.Entity<CustomTierLimit>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("CustomTierLimit_pkey");

            entity.ToTable("CustomTierLimit");

            entity.HasIndex(e => e.UserId, "CustomTierLimit_userId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.CanImportAssets)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("canImportAssets");
            entity.Property(e => e.CanExportAssets)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("canExportAssets");
            entity.Property(e => e.CanImportNRM)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("canImportNRM");
            entity.Property(e => e.CanHideShelfBranding)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("canHideShelfBranding");
            entity.Property(e => e.CanUseReports)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("canUseReports");
            entity.Property(e => e.CanUseRecurringReminders)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("canUseRecurringReminders");
            entity.Property(e => e.CanUsePushNotifications)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("canUsePushNotifications");
            entity.Property(e => e.MaxCustomFields)
                .ValueGeneratedNever()
                .HasDefaultValue(1000)
                .HasColumnName("maxCustomFields");
            entity.Property(e => e.MaxOrganizations)
                .ValueGeneratedNever()
                .HasDefaultValue(1)
                .HasColumnName("maxOrganizations");
            entity.Property(e => e.IsEnterprise)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("isEnterprise");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithOne(p => p.CustomTierLimit)
                .HasForeignKey<CustomTierLimit>(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("CustomTierLimit_userId_fkey");
        });

        modelBuilder.Entity<CustomField>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("CustomField_pkey");

            entity.ToTable("CustomField");

            entity.HasIndex(e => e.OrganizationId, "CustomField_organizationId_idx");

            entity.HasIndex(e => e.UserId, "CustomField_userId_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.DeletedAt }, "CustomField_organizationId_deletedAt_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.HelpText).HasColumnName("helpText");
            entity.Property(e => e.Required)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("required");
            entity.Property(e => e.Active)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("active");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(CustomFieldType.Text)
                .HasColumnName("type");
            entity.Property(e => e.Options).HasColumnName("options");
            entity.Property(e => e.Format)
                .ValueGeneratedNever()
                .HasDefaultValue(CustomFieldFormat.Plain)
                .HasColumnName("format");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.DeletedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("deletedAt");

            entity.HasOne(d => d.Organization).WithMany(p => p.CustomFields)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("CustomField_organizationId_fkey");

            entity.HasOne(d => d.CreatedBy).WithMany(p => p.CustomFields)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("CustomField_userId_fkey");
        });

        modelBuilder.Entity<AssetCustomFieldValue>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetCustomFieldValue_pkey");

            entity.ToTable("AssetCustomFieldValue");

            entity.HasIndex(e => new { e.AssetId, e.CustomFieldId }, "AssetCustomFieldValue_lookup_idx");

            entity.HasIndex(e => e.CustomFieldId, "AssetCustomFieldValue_customFieldId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Value)
                .HasColumnType("jsonb")
                .HasColumnName("value");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.CustomFieldId).HasColumnName("customFieldId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Asset).WithMany(p => p.CustomFields)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetCustomFieldValue_assetId_fkey");

            entity.HasOne(d => d.CustomField).WithMany(p => p.AssetCustomFieldsValues)
                .HasForeignKey(d => d.CustomFieldId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetCustomFieldValue_customFieldId_fkey");
        });

        modelBuilder.Entity<AssetModelCustomFieldDefault>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetModelCustomFieldDefault_pkey");

            entity.ToTable("AssetModelCustomFieldDefault");

            entity.HasIndex(e => e.CustomFieldId, "AssetModelCustomFieldDefault_customFieldId_idx");

            entity.HasIndex(e => new { e.AssetModelId, e.CustomFieldId }, "AssetModelCustomFieldDefault_assetModelId_customFieldId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Value)
                .HasColumnType("jsonb")
                .HasColumnName("value");
            entity.Property(e => e.AssetModelId).HasColumnName("assetModelId");
            entity.Property(e => e.CustomFieldId).HasColumnName("customFieldId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.AssetModel).WithMany(p => p.CustomFieldDefaults)
                .HasForeignKey(d => d.AssetModelId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetModelCustomFieldDefault_assetModelId_fkey");
        });

        modelBuilder.Entity<Invite>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Invite_pkey");

            entity.ToTable("Invite");

            entity.HasIndex(e => e.InviteeUserId, "Invite_inviteeUserId_idx");

            entity.HasIndex(e => e.InviterId, "Invite_inviterId_idx");

            entity.HasIndex(e => e.OrganizationId, "Invite_organizationId_idx");

            entity.HasIndex(e => e.TeamMemberId, "Invite_teamMemberId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.InviterId).HasColumnName("inviterId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.InviteeUserId).HasColumnName("inviteeUserId");
            entity.Property(e => e.TeamMemberId).HasColumnName("teamMemberId");
            entity.Property(e => e.InviteeEmail).HasColumnName("inviteeEmail");
            entity.Property(e => e.Status)
                .ValueGeneratedNever()
                .HasDefaultValue(InviteStatuses.Pending)
                .HasColumnName("status");
            entity.Property(e => e.InviteCode).HasColumnName("inviteCode");
            entity.Property(e => e.Roles).HasColumnName("roles");
            entity.Property(e => e.InviteMessage)
                .HasMaxLength(1000)
                .HasColumnName("inviteMessage");
            entity.Property(e => e.ExpiresAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("expiresAt");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Inviter).WithMany(p => p.SentInvites)
                .HasForeignKey(d => d.InviterId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Invite_inviterId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.Invites)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Invite_organizationId_fkey");

            entity.HasOne(d => d.InviteeUser).WithMany(p => p.ReceivedInvites)
                .HasForeignKey(d => d.InviteeUserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Invite_inviteeUserId_fkey");

            entity.HasOne(d => d.InviteeTeamMember).WithMany(p => p.ReceivedInvites)
                .HasForeignKey(d => d.TeamMemberId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("Invite_teamMemberId_fkey");
        });

        modelBuilder.Entity<Announcement>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Announcement_pkey");

            entity.ToTable("Announcement");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.Link).HasColumnName("link");
            entity.Property(e => e.LinkText).HasColumnName("linkText");
            entity.Property(e => e.Published)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("published");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
        });

        modelBuilder.Entity<ConsumptionLog>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("ConsumptionLog_pkey");

            entity.ToTable("ConsumptionLog");

            entity.HasIndex(e => new { e.AssetId, e.CreatedAt }, "ConsumptionLog_assetId_createdAt_idx");

            entity.HasIndex(e => e.UserId, "ConsumptionLog_userId_idx");

            entity.HasIndex(e => e.BookingId, "ConsumptionLog_bookingId_idx");

            entity.HasIndex(e => e.CustodianId, "ConsumptionLog_custodianId_idx");

            entity.HasIndex(e => e.BookingAssetId, "ConsumptionLog_bookingAssetId_idx");

            entity.HasIndex(e => e.LocationId, "ConsumptionLog_locationId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.Category).HasColumnName("category");
            entity.Property(e => e.Quantity).HasColumnName("quantity");
            entity.Property(e => e.Note).HasColumnName("note");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.BookingId).HasColumnName("bookingId");
            entity.Property(e => e.CustodianId).HasColumnName("custodianId");
            entity.Property(e => e.BookingAssetId).HasColumnName("bookingAssetId");
            entity.Property(e => e.LocationId).HasColumnName("locationId");
            entity.Property(e => e.StockChange).HasColumnName("stockChange");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");

            entity.HasOne(d => d.Asset).WithMany(p => p.ConsumptionLogs)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ConsumptionLog_assetId_fkey");

            entity.HasOne(d => d.PerformedBy).WithMany(p => p.ConsumptionLogs)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("ConsumptionLog_userId_fkey");

            entity.HasOne(d => d.Booking).WithMany(p => p.ConsumptionLogs)
                .HasForeignKey(d => d.BookingId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("ConsumptionLog_bookingId_fkey");

            entity.HasOne(d => d.Custodian).WithMany(p => p.ConsumptionLogs)
                .HasForeignKey(d => d.CustodianId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("ConsumptionLog_custodianId_fkey");

            entity.HasOne(d => d.Location).WithMany(p => p.ConsumptionLogs)
                .HasForeignKey(d => d.LocationId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("ConsumptionLog_locationId_fkey");
        });

        modelBuilder.Entity<BookingAsset>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("BookingAsset_pkey");

            entity.ToTable("BookingAsset");

            entity.HasIndex(e => e.AssetId, "BookingAsset_assetId_idx");

            entity.HasIndex(e => e.AssetKitId, "BookingAsset_assetKitId_idx");

            entity.HasIndex(e => e.SourceKitId, "BookingAsset_sourceKitId_idx");

            entity.HasIndex(e => e.SourceLocationId, "BookingAsset_sourceLocationId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.BookingId).HasColumnName("bookingId");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.Quantity)
                .ValueGeneratedNever()
                .HasDefaultValue(1)
                .HasColumnName("quantity");
            entity.Property(e => e.BookingModelRequestId).HasColumnName("bookingModelRequestId");
            entity.Property(e => e.AssetKitId).HasColumnName("assetKitId");
            entity.Property(e => e.SourceKitId).HasColumnName("sourceKitId");
            entity.Property(e => e.CheckedOutAt)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("checkedOutAt");
            entity.Property(e => e.CheckedOutById).HasColumnName("checkedOutById");
            entity.Property(e => e.CheckedInAt)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("checkedInAt");
            entity.Property(e => e.CheckedInById).HasColumnName("checkedInById");
            entity.Property(e => e.CheckedOutQuantity)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("checkedOutQuantity");
            entity.Property(e => e.SourceLocationId).HasColumnName("sourceLocationId");

            entity.HasOne(d => d.Booking).WithMany(p => p.BookingAssets)
                .HasForeignKey(d => d.BookingId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("BookingAsset_bookingId_fkey");

            entity.HasOne(d => d.Asset).WithMany(p => p.BookingAssets)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("BookingAsset_assetId_fkey");
        });

        modelBuilder.Entity<Booking>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Booking_pkey");

            entity.ToTable("Booking");

            entity.HasIndex(e => e.CreatorId, "Booking_creatorId_idx");

            entity.HasIndex(e => e.CustodianTeamMemberId, "Booking_custodianTeamMemberId_idx");

            entity.HasIndex(e => e.CustodianUserId, "Booking_custodianUserId_idx");

            entity.HasIndex(e => e.OrganizationId, "Booking_organizationId_idx");

            entity.HasIndex(e => e.Status, "Booking_status_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.From, e.Id }, "Booking_organizationId_from_id_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Status)
                .ValueGeneratedNever()
                .HasDefaultValue(BookingStatus.Draft)
                .HasColumnName("status");
            entity.Property(e => e.Description)
                .ValueGeneratedNever()
                .HasDefaultValue("")
                .HasColumnName("description");
            entity.Property(e => e.ActiveSchedulerReference).HasColumnName("activeSchedulerReference");
            entity.Property(e => e.CreatorId).HasColumnName("creatorId");
            entity.Property(e => e.CustodianUserId).HasColumnName("custodianUserId");
            entity.Property(e => e.CustodianTeamMemberId).HasColumnName("custodianTeamMemberId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasPrecision(3)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcConverter>()
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("updatedAt");
            entity.Property(e => e.From)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("from");
            entity.Property(e => e.To)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("to");
            entity.Property(e => e.OriginalFrom)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("originalFrom");
            entity.Property(e => e.OriginalTo)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("originalTo");
            entity.Property(e => e.AutoArchivedAt)
                .HasPrecision(3)
                .HasConversion<UtcConverter>()
                .HasColumnName("autoArchivedAt");
            entity.Property(e => e.ArchivedWithoutCheckin)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("archivedWithoutCheckin");
            entity.Property(e => e.CancellationReason).HasColumnName("cancellationReason");

            entity.HasOne(d => d.Creator).WithMany(p => p.Bookings)
                .HasForeignKey(d => d.CreatorId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Booking_creatorId_fkey");

            entity.HasOne(d => d.CustodianUser).WithMany(p => p.Custodies)
                .HasForeignKey(d => d.CustodianUserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Booking_custodianUserId_fkey");

            entity.HasOne(d => d.CustodianTeamMember).WithMany(p => p.Bookings)
                .HasForeignKey(d => d.CustodianTeamMemberId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Booking_custodianTeamMemberId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.Bookings)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Booking_organizationId_fkey");

            entity.HasMany(d => d.Tags).WithMany(p => p.Bookings)
                .UsingEntity<Dictionary<string, object>>(
                    "_BookingToTag",
                    r => r.HasOne<Tag>().WithMany()
                        .HasForeignKey("B")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_BookingToTag_B_fkey"),
                    l => l.HasOne<Booking>().WithMany()
                        .HasForeignKey("A")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_BookingToTag_A_fkey"),
                    j =>
                    {
                        j.HasKey("A", "B").HasName("_BookingToTag_AB_pkey");
                        j.ToTable("_BookingToTag");
                        j.HasIndex(new[] { "B" }, "_BookingToTag_B_index");
                    });

            entity.HasMany(d => d.NotificationRecipients).WithMany(p => p.BookingNotifications)
                .UsingEntity<Dictionary<string, object>>(
                    "_BookingNotificationRecipients",
                    r => r.HasOne<TeamMember>().WithMany()
                        .HasForeignKey("B")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_BookingNotificationRecipients_B_fkey"),
                    l => l.HasOne<Booking>().WithMany()
                        .HasForeignKey("A")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_BookingNotificationRecipients_A_fkey"),
                    j =>
                    {
                        j.HasKey("A", "B").HasName("_BookingNotificationRecipients_AB_pkey");
                        j.ToTable("_BookingNotificationRecipients");
                        j.HasIndex(new[] { "B" }, "_BookingNotificationRecipients_B_index");
                    });
        });

        modelBuilder.Entity<BookingSettings>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("BookingSettings_pkey");

            entity.ToTable("BookingSettings");

            entity.HasIndex(e => e.OrganizationId, "BookingSettings_organizationId_key").IsUnique();

            entity.HasIndex(e => e.OrganizationId, "BookingSettings_organizationId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.BufferStartTime)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("bufferStartTime");
            entity.Property(e => e.TagsRequired)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("tagsRequired");
            entity.Property(e => e.MaxBookingLength).HasColumnName("maxBookingLength");
            entity.Property(e => e.MaxBookingLengthSkipClosedDays)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("maxBookingLengthSkipClosedDays");
            entity.Property(e => e.AutoArchiveBookings)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("autoArchiveBookings");
            entity.Property(e => e.AutoArchiveDays)
                .ValueGeneratedNever()
                .HasDefaultValue(2)
                .HasColumnName("autoArchiveDays");
            entity.Property(e => e.AutoArchiveExpiredReservations)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("autoArchiveExpiredReservations");
            entity.Property(e => e.RequireExplicitCheckinForAdmin)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("requireExplicitCheckinForAdmin");
            entity.Property(e => e.RequireExplicitCheckinForSelfService)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("requireExplicitCheckinForSelfService");
            entity.Property(e => e.RequireExplicitCheckoutForAdmin)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("requireExplicitCheckoutForAdmin");
            entity.Property(e => e.RequireExplicitCheckoutForSelfService)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("requireExplicitCheckoutForSelfService");
            entity.Property(e => e.CountKitsAsSingleUnit)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("countKitsAsSingleUnit");
            entity.Property(e => e.NotifyBookingCreator)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("notifyBookingCreator");
            entity.Property(e => e.NotifyAdminsOnNewBooking)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("notifyAdminsOnNewBooking");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Organization).WithOne(p => p.BookingSettings)
                .HasForeignKey<BookingSettings>(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("BookingSettings_organizationId_fkey");

            entity.HasMany(d => d.AlwaysNotifyTeamMembers).WithMany(p => p.AlwaysNotifyBookingSettings)
                .UsingEntity<Dictionary<string, object>>(
                    "_BookingSettingsAlwaysNotify",
                    r => r.HasOne<TeamMember>().WithMany()
                        .HasForeignKey("B")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_BookingSettingsAlwaysNotify_B_fkey"),
                    l => l.HasOne<BookingSettings>().WithMany()
                        .HasForeignKey("A")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_BookingSettingsAlwaysNotify_A_fkey"),
                    j =>
                    {
                        j.HasKey("A", "B").HasName("_BookingSettingsAlwaysNotify_AB_pkey");
                        j.ToTable("_BookingSettingsAlwaysNotify");
                        j.HasIndex(new[] { "B" }, "_BookingSettingsAlwaysNotify_B_index");
                    });
        });

        modelBuilder.Entity<PartialBookingCheckin>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("PartialBookingCheckin_pkey");

            entity.ToTable("PartialBookingCheckin");

            entity.HasIndex(e => e.BookingId, "PartialBookingCheckin_bookingId_idx");

            entity.HasIndex(e => e.CheckedInById, "PartialBookingCheckin_checkedInById_idx");

            entity.HasIndex(e => e.CheckinTimestamp, "PartialBookingCheckin_checkinTimestamp_idx");

            entity.HasIndex(e => new { e.BookingId, e.CheckinTimestamp }, "PartialBookingCheckin_bookingId_checkinTimestamp_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.AssetIds).HasColumnName("assetIds");
            entity.Property(e => e.CheckinCount).HasColumnName("checkinCount");
            entity.Property(e => e.CheckinTimestamp)
                .HasValueGenerator<UtcNowGenerator>()
                .HasPrecision(3)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcConverter>()
                .HasColumnName("checkinTimestamp");
            entity.Property(e => e.BookingId).HasColumnName("bookingId");
            entity.Property(e => e.CheckedInById).HasColumnName("checkedInById");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Booking).WithMany(p => p.PartialCheckins)
                .HasForeignKey(d => d.BookingId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("PartialBookingCheckin_bookingId_fkey");

            entity.HasOne(d => d.CheckedInBy).WithMany(p => p.PartialCheckins)
                .HasForeignKey(d => d.CheckedInById)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("PartialBookingCheckin_checkedInById_fkey");
        });

        modelBuilder.Entity<PartialBookingCheckout>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("PartialBookingCheckout_pkey");

            entity.ToTable("PartialBookingCheckout");

            entity.HasIndex(e => e.BookingId, "PartialBookingCheckout_bookingId_idx");

            entity.HasIndex(e => e.CheckedOutById, "PartialBookingCheckout_checkedOutById_idx");

            entity.HasIndex(e => e.CheckoutTimestamp, "PartialBookingCheckout_checkoutTimestamp_idx");

            entity.HasIndex(e => new { e.BookingId, e.CheckoutTimestamp }, "PartialBookingCheckout_bookingId_checkoutTimestamp_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.AssetIds).HasColumnName("assetIds");
            entity.Property(e => e.Quantities)
                .ValueGeneratedNever()
                .HasDefaultValue(new List<int>())
                .HasColumnName("quantities");
            entity.Property(e => e.BookingAssetIds)
                .ValueGeneratedNever()
                .HasDefaultValue(new List<string>())
                .HasColumnName("bookingAssetIds");
            entity.Property(e => e.CheckoutCount).HasColumnName("checkoutCount");
            entity.Property(e => e.CheckoutTimestamp)
                .HasValueGenerator<UtcNowGenerator>()
                .HasPrecision(3)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcConverter>()
                .HasColumnName("checkoutTimestamp");
            entity.Property(e => e.BookingId).HasColumnName("bookingId");
            entity.Property(e => e.CheckedOutById).HasColumnName("checkedOutById");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Booking).WithMany(p => p.PartialCheckouts)
                .HasForeignKey(d => d.BookingId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("PartialBookingCheckout_bookingId_fkey");

            entity.HasOne(d => d.CheckedOutBy).WithMany(p => p.PartialCheckouts)
                .HasForeignKey(d => d.CheckedOutById)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("PartialBookingCheckout_checkedOutById_fkey");
        });

        modelBuilder.Entity<MobileAuthCode>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("MobileAuthCode_pkey");

            entity.ToTable("MobileAuthCode");

            entity.HasIndex(e => e.CodeHash, "MobileAuthCode_codeHash_key").IsUnique();

            entity.HasIndex(e => e.UserId, "MobileAuthCode_userId_idx");

            entity.HasIndex(e => e.ExpiresAt, "MobileAuthCode_expiresAt_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.CodeHash).HasColumnName("codeHash");
            entity.Property(e => e.CodeChallenge).HasColumnName("codeChallenge");
            entity.Property(e => e.SessionCiphertext).HasColumnName("sessionCiphertext");
            entity.Property(e => e.ConsumedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("consumedAt");
            entity.Property(e => e.ExpiresAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("expiresAt");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");

            entity.HasOne(d => d.User).WithMany(p => p.MobileAuthCodes)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("MobileAuthCode_userId_fkey");
        });

        modelBuilder.Entity<PushToken>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("PushToken_pkey");

            entity.ToTable("PushToken");

            entity.HasIndex(e => e.Token, "PushToken_token_key").IsUnique();

            entity.HasIndex(e => e.UserId, "PushToken_userId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Token).HasColumnName("token");
            entity.Property(e => e.Platform).HasColumnName("platform");
            entity.Property(e => e.DeviceName).HasColumnName("deviceName");
            entity.Property(e => e.AppVersion).HasColumnName("appVersion");
            entity.Property(e => e.LastSeenAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("lastSeenAt");
            entity.Property(e => e.DisabledAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("disabledAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithMany(p => p.PushTokens)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("PushToken_userId_fkey");
        });

        modelBuilder.Entity<NotificationPreference>(entity =>
        {
            entity.HasKey(e => e.UserId).HasName("NotificationPreference_pkey");

            entity.ToTable("NotificationPreference");

            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.PushEnabled)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("pushEnabled");
            entity.Property(e => e.Bookings)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("bookings");
            entity.Property(e => e.Audits)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("audits");
            entity.Property(e => e.StockAndReminders)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("stockAndReminders");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.User).WithOne(p => p.NotificationPreference)
                .HasForeignKey<NotificationPreference>(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("NotificationPreference_userId_fkey");
        });

        modelBuilder.Entity<Kit>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Kit_pkey");

            entity.ToTable("Kit");

            entity.HasIndex(e => e.CreatedById, "Kit_createdById_idx");

            entity.HasIndex(e => e.OrganizationId, "Kit_organizationId_idx");

            entity.HasIndex(e => new { e.CategoryId, e.OrganizationId }, "Kit_categoryId_organizationId_idx");

            entity.HasIndex(e => new { e.CategoryId, e.OrganizationId, e.CreatedAt }, "Kit_categoryId_organizationId_createdAt_idx");

            entity.HasIndex(e => new { e.CategoryId, e.OrganizationId, e.Name }, "Kit_categoryId_organizationId_name_idx");

            entity.HasIndex(e => new { e.CategoryId, e.OrganizationId, e.Status }, "Kit_categoryId_organizationId_status_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.Status)
                .ValueGeneratedNever()
                .HasDefaultValue(KitStatus.Available)
                .HasColumnName("status");
            entity.Property(e => e.Image).HasColumnName("image");
            entity.Property(e => e.ImageExpiration)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("imageExpiration");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedById).HasColumnName("createdById");
            entity.Property(e => e.CategoryId).HasColumnName("categoryId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.LocationId).HasColumnName("locationId");

            entity.HasOne(d => d.Organization).WithMany(p => p.Kits)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Kit_organizationId_fkey");

            entity.HasOne(d => d.CreatedBy).WithMany(p => p.CreatedKits)
                .HasForeignKey(d => d.CreatedById)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("Kit_createdById_fkey");

            entity.HasOne(d => d.Category).WithMany(p => p.Kits)
                .HasForeignKey(d => d.CategoryId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Kit_categoryId_fkey");

            entity.HasOne(d => d.Location).WithMany(p => p.Kits)
                .HasForeignKey(d => d.LocationId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("Kit_locationId_fkey");
        });

        modelBuilder.Entity<KitCustody>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("KitCustody_pkey");

            entity.ToTable("KitCustody");

            entity.HasIndex(e => e.KitId, "KitCustody_kitId_key").IsUnique();

            entity.HasIndex(e => e.CustodianId, "KitCustody_custodianId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.CustodianId).HasColumnName("custodianId");
            entity.Property(e => e.KitId).HasColumnName("kitId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Custodian).WithMany(p => p.KitCustodies)
                .HasForeignKey(d => d.CustodianId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("KitCustody_custodianId_fkey");

            entity.HasOne(d => d.Kit).WithOne(p => p.Custody)
                .HasForeignKey<KitCustody>(d => d.KitId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("KitCustody_kitId_fkey");
        });

        modelBuilder.Entity<AssetKit>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetKit_pkey");

            entity.ToTable("AssetKit");

            entity.HasIndex(e => e.KitId, "AssetKit_kitId_idx");

            entity.HasIndex(e => e.OrganizationId, "AssetKit_organizationId_idx");

            entity.HasIndex(e => new { e.AssetId, e.KitId }, "AssetKit_assetId_kitId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.KitId).HasColumnName("kitId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.Quantity)
                .ValueGeneratedNever()
                .HasDefaultValue(1)
                .HasColumnName("quantity");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Asset).WithMany(p => p.AssetKits)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetKit_assetId_fkey");

            entity.HasOne(d => d.Kit).WithMany(p => p.AssetKits)
                .HasForeignKey(d => d.KitId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetKit_kitId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.AssetKits)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetKit_organizationId_fkey");
        });

        modelBuilder.Entity<AssetLocation>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetLocation_pkey");

            entity.ToTable("AssetLocation");

            entity.HasIndex(e => e.LocationId, "AssetLocation_locationId_idx");

            entity.HasIndex(e => e.OrganizationId, "AssetLocation_organizationId_idx");

            entity.HasIndex(e => e.AssetId, "AssetLocation_assetId_idx");

            entity.HasIndex(e => e.AssetKitId, "AssetLocation_assetKitId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.LocationId).HasColumnName("locationId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.Quantity)
                .ValueGeneratedNever()
                .HasDefaultValue(1)
                .HasColumnName("quantity");
            entity.Property(e => e.AssetKitId).HasColumnName("assetKitId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Asset).WithMany(p => p.AssetLocations)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetLocation_assetId_fkey");

            entity.HasOne(d => d.Location).WithMany(p => p.AssetLocations)
                .HasForeignKey(d => d.LocationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetLocation_locationId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.AssetLocations)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetLocation_organizationId_fkey");

            entity.HasOne(d => d.AssetKit).WithMany(p => p.AssetLocations)
                .HasForeignKey(d => d.AssetKitId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetLocation_assetKitId_fkey");
        });

        modelBuilder.Entity<AssetReminder>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetReminder_pkey");

            entity.ToTable("AssetReminder");

            entity.HasIndex(e => new { e.AssetId, e.AlertDateTime }, "AssetReminder_assetId_alertDateTime_idx");

            entity.HasIndex(e => new { e.Name, e.Message }, "AssetReminder_name_message_idx")
                .HasMethod("gin")
                .HasOperators("gin_trgm_ops", "gin_trgm_ops");

            entity.HasIndex(e => new { e.OrganizationId, e.AlertDateTime, e.AssetId }, "AssetReminder_organizationId_alertDateTime_assetId_idx");

            entity.HasIndex(e => new { e.AlertDateTime, e.ActiveSchedulerReference }, "AssetReminder_alertDateTime_activeSchedulerReference_idx");

            entity.HasIndex(e => e.CreatedById, "AssetReminder_createdById_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Message).HasColumnName("message");
            entity.Property(e => e.AlertDateTime)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("alertDateTime");
            entity.Property(e => e.ActiveSchedulerReference).HasColumnName("activeSchedulerReference");
            entity.Property(e => e.RecurrenceUnit).HasColumnName("recurrenceUnit");
            entity.Property(e => e.RecurrenceInterval).HasColumnName("recurrenceInterval");
            entity.Property(e => e.RecurrenceTimezone).HasColumnName("recurrenceTimezone");
            entity.Property(e => e.RecurrenceEndsAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("recurrenceEndsAt");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.CreatedById).HasColumnName("createdById");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Organization).WithMany(p => p.AssetReminders)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("AssetReminder_organizationId_fkey");

            entity.HasOne(d => d.Asset).WithMany(p => p.Reminders)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetReminder_assetId_fkey");

            entity.HasOne(d => d.CreatedBy).WithMany(p => p.AssetReminders)
                .HasForeignKey(d => d.CreatedById)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("AssetReminder_createdById_fkey");

            entity.HasMany(d => d.TeamMembers).WithMany(p => p.AssetReminders)
                .UsingEntity<Dictionary<string, object>>(
                    "_AssetReminderToTeamMember",
                    r => r.HasOne<TeamMember>().WithMany()
                        .HasForeignKey("B")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_AssetReminderToTeamMember_B_fkey"),
                    l => l.HasOne<AssetReminder>().WithMany()
                        .HasForeignKey("A")
                        .OnDelete(DeleteBehavior.Cascade)
                        .HasConstraintName("_AssetReminderToTeamMember_A_fkey"),
                    j =>
                    {
                        j.HasKey("A", "B").HasName("_AssetReminderToTeamMember_AB_pkey");
                        j.ToTable("_AssetReminderToTeamMember");
                        j.HasIndex(new[] { "B" }, "_AssetReminderToTeamMember_B_index");
                    });
        });

        modelBuilder.Entity<WorkingHours>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("WorkingHours_pkey");

            entity.ToTable("WorkingHours");

            entity.HasIndex(e => e.OrganizationId, "WorkingHours_organizationId_key").IsUnique();

            entity.HasIndex(e => e.OrganizationId, "WorkingHours_organizationId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Enabled)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("enabled");
            entity.Property(e => e.WeeklySchedule)
                .ValueGeneratedNever()
                .HasDefaultValue("{\"0\":{\"isOpen\":false},\"1\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"2\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"3\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"4\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"5\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"6\":{\"isOpen\":false}}")
                .HasColumnType("jsonb")
                .HasColumnName("weeklySchedule");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.Organization).WithOne(p => p.WorkingHours)
                .HasForeignKey<WorkingHours>(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("WorkingHours_organizationId_fkey");
        });

        modelBuilder.Entity<WorkingHoursOverride>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("WorkingHoursOverride_pkey");

            entity.ToTable("WorkingHoursOverride");

            entity.HasIndex(e => new { e.WorkingHoursId, e.Date }, "WorkingHoursOverride_workingHoursId_date_idx");

            entity.HasIndex(e => new { e.Date, e.IsOpen }, "WorkingHoursOverride_date_isOpen_idx");

            entity.HasIndex(e => new { e.WorkingHoursId, e.Date }, "WorkingHoursOverride_workingHoursId_date_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Date).HasColumnName("date");
            entity.Property(e => e.IsOpen)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("isOpen");
            entity.Property(e => e.OpenTime).HasColumnName("openTime");
            entity.Property(e => e.CloseTime).HasColumnName("closeTime");
            entity.Property(e => e.Reason).HasColumnName("reason");
            entity.Property(e => e.WorkingHoursId).HasColumnName("workingHoursId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.WorkingHours).WithMany(p => p.Overrides)
                .HasForeignKey(d => d.WorkingHoursId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("WorkingHoursOverride_workingHoursId_fkey");
        });

        modelBuilder.Entity<Update>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("Update_pkey");

            entity.ToTable("Update");

            entity.HasIndex(e => new { e.Status, e.PublishDate }, "Update_status_publishDate_idx");

            entity.HasIndex(e => e.PublishDate, "Update_publishDate_idx");

            entity.HasIndex(e => e.CreatedById, "Update_createdById_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Title).HasColumnName("title");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.Url).HasColumnName("url");
            entity.Property(e => e.ImageUrl).HasColumnName("imageUrl");
            entity.Property(e => e.PublishDate)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("publishDate");
            entity.Property(e => e.Status)
                .ValueGeneratedNever()
                .HasDefaultValue(UpdateStatus.Draft)
                .HasColumnName("status");
            entity.Property(e => e.TargetRoles).HasColumnName("targetRoles");
            entity.Property(e => e.ClickCount)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("clickCount");
            entity.Property(e => e.ViewCount)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("viewCount");
            entity.Property(e => e.CreatedById).HasColumnName("createdById");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.CreatedBy).WithMany(p => p.CreatedUpdates)
                .HasForeignKey(d => d.CreatedById)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("Update_createdById_fkey");
        });

        modelBuilder.Entity<UserUpdateRead>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("UserUpdateRead_pkey");

            entity.ToTable("UserUpdateRead");

            entity.HasIndex(e => e.UserId, "UserUpdateRead_userId_idx");

            entity.HasIndex(e => e.UpdateId, "UserUpdateRead_updateId_idx");

            entity.HasIndex(e => e.ReadAt, "UserUpdateRead_readAt_idx");

            entity.HasIndex(e => new { e.UserId, e.UpdateId }, "UserUpdateRead_userId_updateId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.UpdateId).HasColumnName("updateId");
            entity.Property(e => e.ReadAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("readAt");

            entity.HasOne(d => d.User).WithMany(p => p.UpdateReads)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("UserUpdateRead_userId_fkey");

            entity.HasOne(d => d.Update).WithMany(p => p.UserReads)
                .HasForeignKey(d => d.UpdateId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("UserUpdateRead_updateId_fkey");
        });

        modelBuilder.Entity<AuditNote>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AuditNote_pkey");

            entity.ToTable("AuditNote");

            entity.HasIndex(e => e.AuditSessionId, "AuditNote_auditSessionId_idx");

            entity.HasIndex(e => e.UserId, "AuditNote_userId_idx");

            entity.HasIndex(e => e.AuditAssetId, "AuditNote_auditAssetId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(NoteType.Comment)
                .HasColumnName("type");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.AuditSessionId).HasColumnName("auditSessionId");
            entity.Property(e => e.AuditAssetId).HasColumnName("auditAssetId");

            entity.HasOne(d => d.User).WithMany(p => p.AuditNotes)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AuditNote_userId_fkey");

            entity.HasOne(d => d.AuditSession).WithMany(p => p.Notes)
                .HasForeignKey(d => d.AuditSessionId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditNote_auditSessionId_fkey");

            entity.HasOne(d => d.AuditAsset).WithMany(p => p.Notes)
                .HasForeignKey(d => d.AuditAssetId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AuditNote_auditAssetId_fkey");
        });

        modelBuilder.Entity<TeamMemberNote>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("TeamMemberNote_pkey");

            entity.ToTable("TeamMemberNote");

            entity.HasIndex(e => e.TeamMemberId, "TeamMemberNote_teamMemberId_idx");

            entity.HasIndex(e => e.UserId, "TeamMemberNote_userId_idx");

            entity.HasIndex(e => e.OrganizationId, "TeamMemberNote_organizationId_idx");

            entity.HasIndex(e => new { e.TeamMemberId, e.OrganizationId }, "TeamMemberNote_teamMemberId_organizationId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.Type)
                .ValueGeneratedNever()
                .HasDefaultValue(NoteType.Comment)
                .HasColumnName("type");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.TeamMemberId).HasColumnName("teamMemberId");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");

            entity.HasOne(d => d.User).WithMany(p => p.TeamMemberNotesAuthored)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("TeamMemberNote_userId_fkey");

            entity.HasOne(d => d.TeamMember).WithMany(p => p.TeamMemberNotes)
                .HasForeignKey(d => d.TeamMemberId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("TeamMemberNote_teamMemberId_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.TeamMemberNotes)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("TeamMemberNote_organizationId_fkey");
        });

        modelBuilder.Entity<AuditSession>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AuditSession_pkey");

            entity.ToTable("AuditSession");

            entity.HasIndex(e => new { e.OrganizationId, e.Status }, "AuditSession_organizationId_status_idx");

            entity.HasIndex(e => e.CreatedById, "AuditSession_createdById_idx");

            entity.HasIndex(e => new { e.Status, e.CreatedAt }, "AuditSession_status_createdAt_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.TargetId).HasColumnName("targetId");
            entity.Property(e => e.Status)
                .ValueGeneratedNever()
                .HasDefaultValue(AuditStatus.Pending)
                .HasColumnName("status");
            entity.Property(e => e.ScopeMeta)
                .HasColumnType("jsonb")
                .HasColumnName("scopeMeta");
            entity.Property(e => e.ExpectedAssetCount)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("expectedAssetCount");
            entity.Property(e => e.FoundAssetCount)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("foundAssetCount");
            entity.Property(e => e.MissingAssetCount)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("missingAssetCount");
            entity.Property(e => e.UnexpectedAssetCount)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("unexpectedAssetCount");
            entity.Property(e => e.StartedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("startedAt");
            entity.Property(e => e.DueDate)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("dueDate");
            entity.Property(e => e.CompletedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("completedAt");
            entity.Property(e => e.CancelledAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("cancelledAt");
            entity.Property(e => e.ActiveSchedulerReference).HasColumnName("activeSchedulerReference");
            entity.Property(e => e.CreatedById).HasColumnName("createdById");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.CreatedBy).WithMany(p => p.AuditsCreated)
                .HasForeignKey(d => d.CreatedById)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("AuditSession_createdById_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.AuditSessions)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditSession_organizationId_fkey");
        });

        modelBuilder.Entity<AssetImport>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AssetImport_pkey");

            entity.ToTable("AssetImport");

            entity.HasIndex(e => new { e.OrganizationId, e.Status }, "AssetImport_organizationId_status_idx");

            entity.HasIndex(e => new { e.CreatedById, e.OrganizationId, e.Status }, "AssetImport_createdById_organizationId_status_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Status)
                .ValueGeneratedNever()
                .HasDefaultValue(AssetImportStatus.Draft)
                .HasColumnName("status");
            entity.Property(e => e.Phase)
                .ValueGeneratedNever()
                .HasDefaultValue(AssetImportPhase.Prepare)
                .HasColumnName("phase");
            entity.Property(e => e.FileName).HasColumnName("fileName");
            entity.Property(e => e.TotalRows).HasColumnName("totalRows");
            entity.Property(e => e.ProcessedRows)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("processedRows");
            entity.Property(e => e.Rows)
                .HasColumnType("jsonb")
                .HasColumnName("rows");
            entity.Property(e => e.Resolved)
                .HasColumnType("jsonb")
                .HasColumnName("resolved");
            entity.Property(e => e.CanUseBarcodes)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("canUseBarcodes");
            entity.Property(e => e.ErrorMessage).HasColumnName("errorMessage");
            entity.Property(e => e.Warnings)
                .HasColumnType("jsonb")
                .HasColumnName("warnings");
            entity.Property(e => e.WarningCount)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("warningCount");
            entity.Property(e => e.LockedUntil)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("lockedUntil");
            entity.Property(e => e.ConfirmedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("confirmedAt");
            entity.Property(e => e.StartedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("startedAt");
            entity.Property(e => e.CompletedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("completedAt");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.CreatedById).HasColumnName("createdById");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");

            entity.HasOne(d => d.CreatedBy).WithMany(p => p.AssetImports)
                .HasForeignKey(d => d.CreatedById)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetImport_createdById_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.AssetImports)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AssetImport_organizationId_fkey");
        });

        modelBuilder.Entity<LinkPreview>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("LinkPreview_pkey");

            entity.ToTable("LinkPreview");

            entity.HasIndex(e => new { e.OrganizationId, e.Url }, "LinkPreview_organizationId_url_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.Url).HasColumnName("url");
            entity.Property(e => e.PageTitle).HasColumnName("pageTitle");
            entity.Property(e => e.Outcome).HasColumnName("outcome");
            entity.Property(e => e.Attempts)
                .ValueGeneratedNever()
                .HasDefaultValue(0)
                .HasColumnName("attempts");
            entity.Property(e => e.FetchedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("fetchedAt");
            entity.Property(e => e.NextFetchAt)
                .ValueGeneratedNever()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("nextFetchAt");
            entity.Property(e => e.CustomTitle).HasColumnName("customTitle");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");

            entity.HasOne(d => d.Organization).WithMany(p => p.LinkPreviews)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("LinkPreview_organizationId_fkey");
        });

        modelBuilder.Entity<AuditAssignment>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AuditAssignment_pkey");

            entity.ToTable("AuditAssignment");

            entity.HasIndex(e => e.UserId, "AuditAssignment_userId_idx");

            entity.HasIndex(e => new { e.AuditSessionId, e.UserId }, "AuditAssignment_auditSessionId_userId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.AuditSessionId).HasColumnName("auditSessionId");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.Role).HasColumnName("role");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.AuditSession).WithMany(p => p.Assignments)
                .HasForeignKey(d => d.AuditSessionId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditAssignment_auditSessionId_fkey");

            entity.HasOne(d => d.User).WithMany(p => p.AuditAssignments)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditAssignment_userId_fkey");
        });

        modelBuilder.Entity<AuditAsset>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AuditAsset_pkey");

            entity.ToTable("AuditAsset");

            entity.HasIndex(e => e.Status, "AuditAsset_status_idx");

            entity.HasIndex(e => e.ScannedById, "AuditAsset_scannedById_idx");

            entity.HasIndex(e => e.AssetId, "AuditAsset_assetId_idx");

            entity.HasIndex(e => new { e.AuditSessionId, e.AssetId }, "AuditAsset_auditSessionId_assetId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.AuditSessionId).HasColumnName("auditSessionId");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.Expected)
                .ValueGeneratedNever()
                .HasDefaultValue(true)
                .HasColumnName("expected");
            entity.Property(e => e.Status)
                .ValueGeneratedNever()
                .HasDefaultValue(AuditAssetStatus.Pending)
                .HasColumnName("status");
            entity.Property(e => e.ScannedById).HasColumnName("scannedById");
            entity.Property(e => e.ScannedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("scannedAt");
            entity.Property(e => e.Metadata)
                .HasColumnType("jsonb")
                .HasColumnName("metadata");
            entity.Property(e => e.ExpectedQuantity).HasColumnName("expectedQuantity");
            entity.Property(e => e.FoundQuantity).HasColumnName("foundQuantity");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.AuditSession).WithMany(p => p.Assets)
                .HasForeignKey(d => d.AuditSessionId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditAsset_auditSessionId_fkey");

            entity.HasOne(d => d.Asset).WithMany(p => p.AuditEntries)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditAsset_assetId_fkey");

            entity.HasOne(d => d.ScannedBy).WithMany(p => p.AuditAssetsScanned)
                .HasForeignKey(d => d.ScannedById)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AuditAsset_scannedById_fkey");
        });

        modelBuilder.Entity<AuditScan>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AuditScan_pkey");

            entity.ToTable("AuditScan");

            entity.HasIndex(e => new { e.AuditSessionId, e.ScannedAt }, "AuditScan_auditSessionId_scannedAt_idx");

            entity.HasIndex(e => e.AuditAssetId, "AuditScan_auditAssetId_idx");

            entity.HasIndex(e => e.AssetId, "AuditScan_assetId_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.AuditSessionId).HasColumnName("auditSessionId");
            entity.Property(e => e.AuditAssetId).HasColumnName("auditAssetId");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.ScannedById).HasColumnName("scannedById");
            entity.Property(e => e.Code).HasColumnName("code");
            entity.Property(e => e.AssetTitle).HasColumnName("assetTitle");
            entity.Property(e => e.WasExpected).HasColumnName("wasExpected");
            entity.Property(e => e.Metadata)
                .HasColumnType("jsonb")
                .HasColumnName("metadata");
            entity.Property(e => e.ScannedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("scannedAt");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");

            entity.HasOne(d => d.AuditSession).WithMany(p => p.Scans)
                .HasForeignKey(d => d.AuditSessionId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditScan_auditSessionId_fkey");

            entity.HasOne(d => d.AuditAsset).WithMany(p => p.Scans)
                .HasForeignKey(d => d.AuditAssetId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AuditScan_auditAssetId_fkey");

            entity.HasOne(d => d.Asset).WithMany(p => p.AuditScans)
                .HasForeignKey(d => d.AssetId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AuditScan_assetId_fkey");

            entity.HasOne(d => d.ScannedBy).WithMany(p => p.AuditScans)
                .HasForeignKey(d => d.ScannedById)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AuditScan_scannedById_fkey");
        });

        modelBuilder.Entity<AuditImage>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("AuditImage_pkey");

            entity.ToTable("AuditImage");

            entity.HasIndex(e => e.AuditSessionId, "AuditImage_auditSessionId_idx");

            entity.HasIndex(e => e.AuditAssetId, "AuditImage_auditAssetId_idx");

            entity.HasIndex(e => e.OrganizationId, "AuditImage_organizationId_idx");

            entity.HasIndex(e => e.UploadedById, "AuditImage_uploadedById_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.ImageUrl).HasColumnName("imageUrl");
            entity.Property(e => e.ThumbnailUrl).HasColumnName("thumbnailUrl");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.AuditSessionId).HasColumnName("auditSessionId");
            entity.Property(e => e.AuditAssetId).HasColumnName("auditAssetId");
            entity.Property(e => e.UploadedById).HasColumnName("uploadedById");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.CapturedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("capturedAt");
            entity.Property(e => e.CaptureSource).HasColumnName("captureSource");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UpdatedAt)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("updatedAt");

            entity.HasOne(d => d.AuditSession).WithMany(p => p.Images)
                .HasForeignKey(d => d.AuditSessionId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditImage_auditSessionId_fkey");

            entity.HasOne(d => d.AuditAsset).WithMany(p => p.Images)
                .HasForeignKey(d => d.AuditAssetId)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AuditImage_auditAssetId_fkey");

            entity.HasOne(d => d.UploadedBy).WithMany(p => p.AuditImagesUploaded)
                .HasForeignKey(d => d.UploadedById)
                .OnDelete(DeleteBehavior.SetNull)
                .HasConstraintName("AuditImage_uploadedById_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.AuditImages)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("AuditImage_organizationId_fkey");
        });

        modelBuilder.Entity<RoleChangeLog>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("RoleChangeLog_pkey");

            entity.ToTable("RoleChangeLog");

            entity.HasIndex(e => new { e.UserId, e.OrganizationId }, "RoleChangeLog_userId_organizationId_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.CreatedAt }, "RoleChangeLog_organizationId_createdAt_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.PreviousRole).HasColumnName("previousRole");
            entity.Property(e => e.NewRole).HasColumnName("newRole");
            entity.Property(e => e.Source)
                .ValueGeneratedNever()
                .HasDefaultValue(RoleChangeSource.Manual)
                .HasColumnName("source");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("timestamp(3) without time zone")
                .HasColumnName("createdAt");
            entity.Property(e => e.UserId).HasColumnName("userId");
            entity.Property(e => e.ChangedById).HasColumnName("changedById");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");

            entity.HasOne(d => d.User).WithMany(p => p.RoleChangesReceived)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("RoleChangeLog_userId_fkey");

            entity.HasOne(d => d.ChangedBy).WithMany(p => p.RoleChangesPerformed)
                .HasForeignKey(d => d.ChangedById)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("RoleChangeLog_changedById_fkey");

            entity.HasOne(d => d.Organization).WithMany(p => p.RoleChangeLogs)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("RoleChangeLog_organizationId_fkey");
        });

        modelBuilder.Entity<ActivityEvent>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("ActivityEvent_pkey");

            entity.ToTable("ActivityEvent");

            entity.HasIndex(e => new { e.OrganizationId, e.OccurredAt }, "ActivityEvent_organizationId_occurredAt_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.Action, e.OccurredAt }, "ActivityEvent_organizationId_action_occurredAt_idx");

            entity.HasIndex(e => new { e.OrganizationId, e.EntityType, e.EntityId, e.OccurredAt }, "ActivityEvent_organizationId_entityType_entityId_occurredAt_idx");

            entity.HasIndex(e => new { e.ActorUserId, e.OccurredAt }, "ActivityEvent_actorUserId_occurredAt_idx");

            entity.HasIndex(e => new { e.AssetId, e.OccurredAt }, "ActivityEvent_assetId_occurredAt_idx");

            entity.HasIndex(e => new { e.BookingId, e.OccurredAt }, "ActivityEvent_bookingId_occurredAt_idx");

            entity.HasIndex(e => new { e.AuditSessionId, e.OccurredAt }, "ActivityEvent_auditSessionId_occurredAt_idx");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasColumnName("id");
            entity.Property(e => e.OrganizationId).HasColumnName("organizationId");
            entity.Property(e => e.OccurredAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasPrecision(3)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasConversion<UtcConverter>()
                .HasColumnName("occurredAt");
            entity.Property(e => e.ActorUserId).HasColumnName("actorUserId");
            entity.Property(e => e.ActorSnapshot)
                .HasColumnType("jsonb")
                .HasColumnName("actorSnapshot");
            entity.Property(e => e.Action).HasColumnName("action");
            entity.Property(e => e.EntityType).HasColumnName("entityType");
            entity.Property(e => e.EntityId).HasColumnName("entityId");
            entity.Property(e => e.AssetId).HasColumnName("assetId");
            entity.Property(e => e.BookingId).HasColumnName("bookingId");
            entity.Property(e => e.AuditSessionId).HasColumnName("auditSessionId");
            entity.Property(e => e.AuditAssetId).HasColumnName("auditAssetId");
            entity.Property(e => e.KitId).HasColumnName("kitId");
            entity.Property(e => e.LocationId).HasColumnName("locationId");
            entity.Property(e => e.TeamMemberId).HasColumnName("teamMemberId");
            entity.Property(e => e.TargetUserId).HasColumnName("targetUserId");
            entity.Property(e => e.Field).HasColumnName("field");
            entity.Property(e => e.FromValue)
                .HasColumnType("jsonb")
                .HasColumnName("fromValue");
            entity.Property(e => e.ToValue)
                .HasColumnType("jsonb")
                .HasColumnName("toValue");
            entity.Property(e => e.Meta)
                .HasColumnType("jsonb")
                .HasColumnName("meta");

            entity.HasOne(d => d.Organization).WithMany(p => p.ActivityEvents)
                .HasForeignKey(d => d.OrganizationId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ActivityEvent_organizationId_fkey");
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);

    public override int SaveChanges(bool acceptAllChangesOnSuccess)
    {
        StampUpdatedAt();
        return base.SaveChanges(acceptAllChangesOnSuccess);
    }

    public override Task<int> SaveChangesAsync(bool acceptAllChangesOnSuccess, CancellationToken cancellationToken = default)
    {
        StampUpdatedAt();
        return base.SaveChangesAsync(acceptAllChangesOnSuccess, cancellationToken);
    }

    private void StampUpdatedAt()
    {
        var now = DateTime.UtcNow;
        foreach (var entry in ChangeTracker.Entries<Image>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<User>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<UserContact>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<UserBusinessIntel>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Asset>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetFilterPreset>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetIndexSettings>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<PdfTemplate>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Category>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetModel>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<BookingModelRequest>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Tag>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Note>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<BookingNote>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<LocationNote>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Qr>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Barcode>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<PrintBatch>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<ReportFound>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Scan>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Location>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Role>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<TeamMember>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Custody>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Organization>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<UserOrganization>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<SsoDetails>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<UserScimExternalId>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Tier>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<TierLimit>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<CustomTierLimit>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<CustomField>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetCustomFieldValue>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetModelCustomFieldDefault>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Invite>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Announcement>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Booking>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<BookingSettings>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<PartialBookingCheckin>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<PartialBookingCheckout>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<PushToken>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<NotificationPreference>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Kit>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<KitCustody>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetKit>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetLocation>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetReminder>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<WorkingHours>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<WorkingHoursOverride>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<Update>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AuditNote>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<TeamMemberNote>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AuditSession>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AssetImport>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<LinkPreview>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AuditAssignment>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AuditAsset>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }

        foreach (var entry in ChangeTracker.Entries<AuditImage>())
        {
            if ((entry.State == EntityState.Added && entry.Entity.UpdatedAt == default)
                || (entry.State == EntityState.Modified && !entry.Property(e => e.UpdatedAt).IsModified))
            {
                entry.Entity.UpdatedAt = now;
            }
        }
    }
}

file sealed class CuidGenerator : ValueGenerator<string>
{
    public override bool GeneratesTemporaryValues => false;

#pragma warning disable VISLIB0001
    public override string Next(EntityEntry entry) => Cuid.NewCuid().ToString();
#pragma warning restore VISLIB0001
}

file sealed class UtcNowGenerator : ValueGenerator<DateTime>
{
    public override bool GeneratesTemporaryValues => false;

    public override DateTime Next(EntityEntry entry) => DateTime.UtcNow;
}

file sealed class UtcClockConverter : ValueConverter<DateTime, DateTime>
{
    public UtcClockConverter()
        : base(v => DateTime.SpecifyKind(v.Kind == DateTimeKind.Local ? v.ToUniversalTime() : v, DateTimeKind.Unspecified), v => DateTime.SpecifyKind(v, DateTimeKind.Utc))
    {
    }
}

file sealed class UtcConverter : ValueConverter<DateTime, DateTime>
{
    public UtcConverter()
        : base(v => DateTime.SpecifyKind(v.Kind == DateTimeKind.Local ? v.ToUniversalTime() : v, DateTimeKind.Utc), v => v)
    {
    }
}
