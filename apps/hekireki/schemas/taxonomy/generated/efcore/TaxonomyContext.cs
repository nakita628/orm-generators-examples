#nullable enable

using System;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.ChangeTracking;
using Microsoft.EntityFrameworkCore.Metadata.Conventions;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using Microsoft.EntityFrameworkCore.ValueGeneration;
using Visus.Cuid;

namespace Taxonomy.Models;

public partial class TaxonomyContext : DbContext
{
    public TaxonomyContext(DbContextOptions<TaxonomyContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Account> Accounts { get; set; }

    public virtual DbSet<Session> Sessions { get; set; }

    public virtual DbSet<User> Users { get; set; }

    public virtual DbSet<VerificationToken> VerificationTokens { get; set; }

    public virtual DbSet<Post> Posts { get; set; }

    protected override void ConfigureConventions(ModelConfigurationBuilder configurationBuilder)
    {
        configurationBuilder.Conventions.Remove<ForeignKeyIndexConvention>();
        ConfigureConventionsPartial(configurationBuilder);
    }

    partial void ConfigureConventionsPartial(ModelConfigurationBuilder configurationBuilder);

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Account>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("accounts_pkey");

            entity.ToTable("accounts");

            entity.HasIndex(e => new { e.Provider, e.ProviderAccountId }, "accounts_provider_providerAccountId_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasMaxLength(191)
                .HasColumnName("id");
            entity.Property(e => e.UserId)
                .HasMaxLength(191)
                .HasColumnName("userId");
            entity.Property(e => e.Type)
                .HasMaxLength(191)
                .HasColumnName("type");
            entity.Property(e => e.Provider)
                .HasMaxLength(191)
                .HasColumnName("provider");
            entity.Property(e => e.ProviderAccountId)
                .HasMaxLength(191)
                .HasColumnName("providerAccountId");
            entity.Property(e => e.RefreshToken)
                .HasColumnType("text")
                .HasColumnName("refresh_token");
            entity.Property(e => e.AccessToken)
                .HasColumnType("text")
                .HasColumnName("access_token");
            entity.Property(e => e.ExpiresAt).HasColumnName("expires_at");
            entity.Property(e => e.TokenType)
                .HasMaxLength(191)
                .HasColumnName("token_type");
            entity.Property(e => e.Scope)
                .HasMaxLength(191)
                .HasColumnName("scope");
            entity.Property(e => e.IdToken)
                .HasColumnType("text")
                .HasColumnName("id_token");
            entity.Property(e => e.SessionState)
                .HasMaxLength(191)
                .HasColumnName("session_state");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP(3)")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("created_at");
            entity.Property(e => e.UpdatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP(3)")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("updated_at");

            entity.HasOne(d => d.User).WithMany(p => p.Accounts)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("accounts_userId_fkey");
        });

        modelBuilder.Entity<Session>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("sessions_pkey");

            entity.ToTable("sessions");

            entity.HasIndex(e => e.SessionToken, "sessions_sessionToken_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasMaxLength(191)
                .HasColumnName("id");
            entity.Property(e => e.SessionToken)
                .HasMaxLength(191)
                .HasColumnName("sessionToken");
            entity.Property(e => e.UserId)
                .HasMaxLength(191)
                .HasColumnName("userId");
            entity.Property(e => e.Expires)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("expires");

            entity.HasOne(d => d.User).WithMany(p => p.Sessions)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("sessions_userId_fkey");
        });

        modelBuilder.Entity<User>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("users_pkey");

            entity.ToTable("users");

            entity.HasIndex(e => e.Email, "users_email_key").IsUnique();

            entity.HasIndex(e => e.StripeCustomerId, "users_stripe_customer_id_key").IsUnique();

            entity.HasIndex(e => e.StripeSubscriptionId, "users_stripe_subscription_id_key").IsUnique();

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasMaxLength(191)
                .HasColumnName("id");
            entity.Property(e => e.Name)
                .HasMaxLength(191)
                .HasColumnName("name");
            entity.Property(e => e.Email)
                .HasMaxLength(191)
                .HasColumnName("email");
            entity.Property(e => e.EmailVerified)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("emailVerified");
            entity.Property(e => e.Image)
                .HasMaxLength(191)
                .HasColumnName("image");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP(3)")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("created_at");
            entity.Property(e => e.UpdatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP(3)")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("updated_at");
            entity.Property(e => e.StripeCustomerId)
                .HasMaxLength(191)
                .HasColumnName("stripe_customer_id");
            entity.Property(e => e.StripeSubscriptionId)
                .HasMaxLength(191)
                .HasColumnName("stripe_subscription_id");
            entity.Property(e => e.StripePriceId)
                .HasMaxLength(191)
                .HasColumnName("stripe_price_id");
            entity.Property(e => e.StripeCurrentPeriodEnd)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("stripe_current_period_end");
        });

        modelBuilder.Entity<VerificationToken>(entity =>
        {
            entity.HasKey(e => e.Token).HasName("verification_tokens_token_key");

            entity.ToTable("verification_tokens");

            entity.HasIndex(e => new { e.Identifier, e.Token }, "verification_tokens_identifier_token_key").IsUnique();

            entity.Property(e => e.Identifier)
                .HasMaxLength(191)
                .HasColumnName("identifier");
            entity.Property(e => e.Token)
                .HasMaxLength(191)
                .HasColumnName("token");
            entity.Property(e => e.Expires)
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("expires");
        });

        modelBuilder.Entity<Post>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("posts_pkey");

            entity.ToTable("posts");

            entity.Property(e => e.Id)
                .HasValueGenerator<CuidGenerator>()
                .HasMaxLength(191)
                .HasColumnName("id");
            entity.Property(e => e.Title)
                .HasMaxLength(191)
                .HasColumnName("title");
            entity.Property(e => e.Content)
                .HasColumnType("json")
                .HasColumnName("content");
            entity.Property(e => e.Published)
                .ValueGeneratedNever()
                .HasDefaultValue(false)
                .HasColumnName("published");
            entity.Property(e => e.CreatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP(3)")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("created_at");
            entity.Property(e => e.UpdatedAt)
                .HasValueGenerator<UtcNowGenerator>()
                .HasDefaultValueSql("CURRENT_TIMESTAMP(3)")
                .HasConversion<UtcClockConverter>()
                .HasColumnType("datetime(3)")
                .HasColumnName("updated_at");
            entity.Property(e => e.AuthorId)
                .HasMaxLength(191)
                .HasColumnName("authorId");

            entity.HasOne(d => d.Author).WithMany(p => p.Post)
                .HasForeignKey(d => d.AuthorId)
                .OnDelete(DeleteBehavior.Restrict)
                .HasConstraintName("posts_authorId_fkey");
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
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
