defmodule Shelf.Booking do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          status: atom(),
          description: String.t() | nil,
          active_scheduler_reference: String.t() | nil,
          from: DateTime.t(),
          to: DateTime.t(),
          original_from: DateTime.t() | nil,
          original_to: DateTime.t() | nil,
          auto_archived_at: DateTime.t() | nil,
          archived_without_checkin: boolean(),
          cancellation_reason: String.t() | nil,
          approved_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          creator: Shelf.User.t() | nil,
          custodian_user: Shelf.User.t() | nil,
          custodian_team_member: Shelf.TeamMember.t() | nil,
          organization: Shelf.Organization.t() | nil,
          approved_by: Shelf.User.t() | nil,
          booking_assets: [Shelf.BookingAsset.t()],
          consumption_logs: [Shelf.ConsumptionLog.t()],
          notes: [Shelf.BookingNote.t()],
          model_requests: [Shelf.BookingModelRequest.t()],
          partial_checkins: [Shelf.PartialBookingCheckin.t()],
          partial_checkouts: [Shelf.PartialBookingCheckout.t()],
          tags: [Shelf.Tag.t()],
          notification_recipients: [Shelf.TeamMember.t()]
        }

  schema "Booking" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:status, Ecto.Enum, values: [:DRAFT, :RESERVED, :ONGOING, :OVERDUE, :COMPLETE, :ARCHIVED, :CANCELLED], default: :DRAFT)
    field(:description, :string, default: "")
    field(:active_scheduler_reference, :string, source: :activeSchedulerReference)
    field(:from, Shelf.PrismaDateTime)
    field(:to, Shelf.PrismaDateTime)
    field(:original_from, Shelf.PrismaDateTime, source: :originalFrom)
    field(:original_to, Shelf.PrismaDateTime, source: :originalTo)
    field(:auto_archived_at, Shelf.PrismaDateTime, source: :autoArchivedAt)
    field(:archived_without_checkin, :boolean, default: false, source: :archivedWithoutCheckin)
    field(:cancellation_reason, :string, source: :cancellationReason)
    field(:approved_at, Shelf.PrismaDateTime, source: :approvedAt)
    field(:creator_id, :string, source: :creatorId)
    field(:custodian_user_id, :string, source: :custodianUserId)
    field(:custodian_team_member_id, :string, source: :custodianTeamMemberId)
    field(:organization_id, :string, source: :organizationId)
    field(:approved_by_id, :string, source: :approvedById)
    belongs_to(:creator, Shelf.User, foreign_key: :creator_id, define_field: false, type: :string)
    belongs_to(:custodian_user, Shelf.User, foreign_key: :custodian_user_id, define_field: false, type: :string)
    belongs_to(:custodian_team_member, Shelf.TeamMember, foreign_key: :custodian_team_member_id, define_field: false, type: :string)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:approved_by, Shelf.User, foreign_key: :approved_by_id, define_field: false, type: :string)
    has_many(:booking_assets, Shelf.BookingAsset, foreign_key: :booking_id, references: :id)
    has_many(:consumption_logs, Shelf.ConsumptionLog, foreign_key: :booking_id, references: :id)
    has_many(:notes, Shelf.BookingNote, foreign_key: :booking_id, references: :id)
    has_many(:model_requests, Shelf.BookingModelRequest, foreign_key: :booking_id, references: :id)
    has_many(:partial_checkins, Shelf.PartialBookingCheckin, foreign_key: :booking_id, references: :id)
    has_many(:partial_checkouts, Shelf.PartialBookingCheckout, foreign_key: :booking_id, references: :id)
    many_to_many(:tags, Shelf.Tag, join_through: "_BookingToTag", join_keys: [A: :id, B: :id])
    many_to_many(:notification_recipients, Shelf.TeamMember, join_through: "_BookingNotificationRecipients", join_keys: [A: :id, B: :id])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end