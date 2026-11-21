```mermaid
erDiagram
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    person_user ||--o| authtoken_token : "(id) - (user_id)"
    event_event |o--o{ common_activitylog : "(id) - (event_id)"
    person_user |o--o{ common_activitylog : "(id) - (person_id)"
    django_content_type ||--o{ common_activitylog : "(id) - (content_type_id)"
    event_organiser |o--o{ event_event : "(id) - (organiser_id)"
    event_event ||--o{ event_event_settingsstore : "(id) - (object_id)"
    event_event ||--o{ event_eventextralink : "(id) - (event_id)"
    event_organiser ||--o{ event_team : "(id) - (organiser_id)"
    event_team }o--o{ person_user : "(members) - (teams)"
    event_team }o--o{ event_event : "(limit_events) - (team_set)"
    event_team }o--o{ submission_track : "(limit_tracks) - (limit_teams)"
    event_team ||--o{ event_teaminvite : "(id) - (team_id)"
    event_event |o--o{ mail_queuedmail : "(id) - (event_id)"
    mail_mailtemplate |o--o{ mail_queuedmail : "(id) - (template_id)"
    mail_queuedmail }o--o{ person_user : "(to_users) - (mails)"
    mail_queuedmail }o--o{ submission_submission : "(submissions) - (mails)"
    event_event ||--o{ mail_mailtemplate : "(id) - (event_id)"
    person_user ||--o{ person_attendeeprofile : "(id) - (user_id)"
    event_event ||--o{ person_attendeeprofile : "(id) - (event_id)"
    person_user ||--o{ person_userapitoken : "(id) - (user_id)"
    person_userapitoken }o--o{ event_event : "(limit_events) - (_person_userapitoken_limit_events_+)"
    event_event ||--o{ person_speakerinformation : "(id) - (event_id)"
    person_speakerinformation }o--o{ submission_track : "(limit_tracks) - (speakerinformation_set)"
    person_speakerinformation }o--o{ submission_submissiontype : "(limit_types) - (speakerinformation_set)"
    person_user ||--o{ person_profilepicture : "(id) - (user_id)"
    person_user ||--o{ person_usereventpreferences : "(id) - (user_id)"
    event_event |o--o{ person_usereventpreferences : "(id) - (event_id)"
    person_user |o--o{ person_speakerprofile : "(id) - (user_id)"
    event_event ||--o{ person_speakerprofile : "(id) - (event_id)"
    person_profilepicture |o--o{ person_speakerprofile : "(id) - (profile_picture_id)"
    person_profilepicture |o--o{ person_user : "(id) - (profile_picture_id)"
    person_user }o--o{ auth_group : "(groups) - (user_set)"
    person_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    event_event ||--o{ schedule_availability : "(id) - (event_id)"
    person_speakerprofile |o--o{ schedule_availability : "(id) - (person_id)"
    schedule_room |o--o{ schedule_availability : "(id) - (room_id)"
    event_event ||--o{ schedule_room : "(id) - (event_id)"
    event_event ||--o{ schedule_schedule : "(id) - (event_id)"
    submission_submission |o--o{ schedule_talkslot : "(id) - (submission_id)"
    schedule_room |o--o{ schedule_talkslot : "(id) - (room_id)"
    schedule_schedule ||--o{ schedule_talkslot : "(id) - (schedule_id)"
    event_event ||--o{ submission_submitteraccesscode : "(id) - (event_id)"
    submission_submitteraccesscode }o--o{ submission_track : "(tracks) - (submitter_access_codes)"
    submission_submitteraccesscode }o--o{ submission_submissiontype : "(submission_types) - (submitter_access_codes)"
    event_event ||--o| submission_cfp : "(id) - (event_id)"
    submission_submissiontype ||--o{ submission_cfp : "(id) - (default_type_id)"
    submission_submission ||--o{ submission_submissioncomment : "(id) - (submission_id)"
    person_user ||--o{ submission_submissioncomment : "(id) - (user_id)"
    submission_submissioncomment |o--o{ submission_submissioncomment : "(id) - (reply_to_id)"
    submission_submission ||--o{ submission_feedback : "(id) - (talk_id)"
    person_speakerprofile |o--o{ submission_feedback : "(id) - (speaker_id)"
    event_event ||--o{ submission_question : "(id) - (event_id)"
    submission_question }o--o{ submission_track : "(tracks) - (questions)"
    submission_question }o--o{ submission_submissiontype : "(submission_types) - (questions)"
    submission_question }o--o{ event_team : "(limit_teams) - (_submission_question_limit_teams_+)"
    submission_question ||--o{ submission_answeroption : "(id) - (question_id)"
    submission_question ||--o{ submission_answer : "(id) - (question_id)"
    submission_submission |o--o{ submission_answer : "(id) - (submission_id)"
    person_speakerprofile |o--o{ submission_answer : "(id) - (speaker_id)"
    submission_review |o--o{ submission_answer : "(id) - (review_id)"
    submission_answer }o--o{ submission_answeroption : "(options) - (answers)"
    submission_submission ||--o{ submission_resource : "(id) - (submission_id)"
    event_event ||--o{ submission_reviewscorecategory : "(id) - (event_id)"
    submission_reviewscorecategory }o--o{ submission_track : "(limit_tracks) - (reviewscorecategory_set)"
    submission_reviewscorecategory ||--o{ submission_reviewscore : "(id) - (category_id)"
    submission_submission ||--o{ submission_review : "(id) - (submission_id)"
    person_user ||--o{ submission_review : "(id) - (user_id)"
    submission_review }o--o{ submission_reviewscore : "(scores) - (reviews)"
    event_event ||--o{ submission_reviewphase : "(id) - (event_id)"
    submission_submission ||--o{ submission_attendeesignup : "(id) - (submission_id)"
    person_attendeeprofile ||--o{ submission_attendeesignup : "(id) - (attendee_id)"
    submission_submission ||--o{ submission_speakerrole : "(id) - (submission_id)"
    person_speakerprofile ||--o{ submission_speakerrole : "(id) - (speaker_id)"
    event_event ||--o{ submission_submission : "(id) - (event_id)"
    submission_submissiontype ||--o{ submission_submission : "(id) - (submission_type_id)"
    submission_track |o--o{ submission_submission : "(id) - (track_id)"
    submission_submitteraccesscode |o--o{ submission_submission : "(id) - (access_code_id)"
    submission_submission }o--o{ submission_tag : "(tags) - (submissions)"
    submission_submission }o--o{ person_user : "(assigned_reviewers) - (assigned_reviews)"
    person_user ||--o{ submission_submissionfavourite : "(id) - (user_id)"
    submission_submission ||--o{ submission_submissionfavourite : "(id) - (submission_id)"
    submission_submission ||--o{ submission_submissioninvitation : "(id) - (submission_id)"
    event_event ||--o{ submission_tag : "(id) - (event_id)"
    event_event ||--o{ submission_track : "(id) - (event_id)"
    event_event ||--o{ submission_submissiontype : "(id) - (event_id)"
    auth_permission {
        AutoField id PK
        CharField name
        ForeignKey content_type_id FK
        CharField codename
    }
    auth_group {
        AutoField id PK
        CharField name UK
    }
    django_content_type {
        AutoField id PK
        CharField app_label
        CharField model
    }
    django_session {
        CharField session_key PK
        TextField session_data
        DateTimeField expire_date
    }
    authtoken_token {
        CharField key PK
        OneToOneField user_id FK, UK
        DateTimeField created
    }
    common_activitylog {
        AutoField id PK
        ForeignKey event_id FK
        ForeignKey person_id FK
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
        DateTimeField timestamp
        CharField action_type
        JSONField data
        BooleanField is_orga_action
    }
    common_cachedfile {
        UUIDField id PK
        DateTimeField expires
        DateTimeField timestamp
        CharField filename
        CharField content_type
        TextField session_key
        FileField file
    }
    common_globalsettings_settingsstore {
        AutoField id PK
        CharField key
        TextField value
    }
    event_event {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        I18nCharField name
        SlugField slug UK "The slug may only contain letters, numbers, dots and dashes."
        ForeignKey organiser_id FK
        BooleanField is_public
        DateField date_from
        DateField date_to
        EventTimezone timezone "All event dates will be localised and interpreted to be in this timezone."
        EmailField email "Will be used as Reply-To in emails."
        URLField custom_domain "Enter a custom domain, such as https://my.event.example.org"
        JSONField feature_flags
        JSONField display_settings
        JSONField review_settings
        JSONField mail_settings
        JSONField attendee_signup_settings
        CharField primary_color "Provide a hex value like #00ff00 if you want to style pretalx in your event’s colour scheme."
        FileField custom_css "Upload a custom CSS file if changing the primary colour is not sufficient for you."
        ImageField logo "If you provide a logo image, your event’s name will not be shown in the event header. The logo will be scaled down to a height of 140px."
        ImageField header_image "If you provide a header image, it will be displayed instead of your event’s color and/or header pattern at the top of all event pages. It will be center-aligned, so when the window shrinks, the center parts will continue to be displayed, and not stretched."
        ImageField og_image "This image will be shown as a preview when links to your event are shared on social media or messaging apps. For best results, use an image at least 1200x630 pixels. If not set, the logo or header image will be used instead."
        JSONField locales
        JSONField content_locales
        EventLocale locale
        I18nTextField landing_page_text "This text will be shown on the landing page, alongside with links to the CfP and schedule, if appropriate. You can use <a href='https://docs.pretalx.org/user/markdown/' target='_blank' rel='noopener'>Markdown</a> here."
        I18nTextField featured_sessions_text "This text will be shown at the top of the featured sessions page instead of the default text. You can use <a href='https://docs.pretalx.org/user/markdown/' target='_blank' rel='noopener'>Markdown</a> here."
        TextField plugins
    }
    event_event_settingsstore {
        AutoField id PK
        CharField key
        TextField value
        ForeignKey object_id FK
    }
    event_eventextralink {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        I18nCharField label
        URLField url
        EventExtraLinkRole role
    }
    event_organiser {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        I18nCharField name
        SlugField slug UK "Should be short, only contain lowercase letters and numbers, and must be unique, as it is used in URLs."
    }
    event_team {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey organiser_id FK
        CharField name
        BooleanField all_events
        BooleanField can_create_events
        BooleanField can_change_teams
        BooleanField can_change_organiser_settings
        BooleanField can_change_event_settings
        BooleanField can_change_submissions
        BooleanField is_reviewer
        BooleanField force_hide_speaker_names "Normally, anonymisation is configured in the event review settings. This setting will <strong>override the event settings</strong> and always hide speaker names for this team."
    }
    event_teaminvite {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey team_id FK
        EmailField email
        CharField token UK
    }
    mail_queuedmail {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        ForeignKey template_id FK
        CharField to "One email address or several addresses separated by commas."
        CharField reply_to "By default, the organiser address is used as Reply-To."
        CharField cc "One email address or several addresses separated by commas."
        CharField bcc "One email address or several addresses separated by commas."
        CharField subject
        TextField text
        TextField text_html
        DateTimeField sent
        QueuedMailStates state
        JSONField error_data
        DateTimeField error_timestamp
        CharField locale
        JSONField attachments
    }
    mail_mailtemplate {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        MailTemplateRoles role
        I18nCharField subject
        I18nTextField text
        CharField reply_to "Change the Reply-To address if you do not want to use the default organiser address"
        CharField bcc "Enter comma separated addresses. Will receive a blind copy of every email sent from this template. This may be a LOT!"
        BooleanField is_auto_created
    }
    person_attendeeprofile {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey user_id FK
        ForeignKey event_id FK
    }
    person_userapitoken {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        CharField name
        CharField token UK
        ForeignKey user_id FK
        BooleanField all_events
        DateTimeField expires
        JSONField endpoints
        CharField version
        DateTimeField last_used
    }
    person_speakerinformation {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        SpeakerInformationTargetGroup target_group
        I18nCharField title
        I18nTextField text "You can use <a href='https://docs.pretalx.org/user/markdown/' target='_blank' rel='noopener'>Markdown</a> here."
        FileField resource "Please try to keep your upload small, preferably below 16 MB."
    }
    person_profilepicture {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey user_id FK
        ImageField avatar
        ImageField avatar_thumbnail
        ImageField avatar_thumbnail_tiny
    }
    person_usereventpreferences {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey user_id FK
        ForeignKey event_id FK
        JSONField preferences
    }
    person_speakerprofile {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey user_id FK
        ForeignKey event_id FK
        CharField name
        CharField code
        MarkdownField biography
        BooleanField has_arrived
        TextField internal_notes "Internal notes for other organisers/reviewers. Not visible to the speakers or the public."
        ForeignKey profile_picture_id FK
    }
    person_user {
        AutoField id PK
        CharField password
        DateTimeField last_login
        CharField code UK
        CharField nick
        CharField name "Please enter the name you wish to be displayed publicly. This name will be used for all events you are participating in on this server."
        EmailField email UK "Your email address will be used for password resets and notification about your event/proposals."
        DateTimeField created
        BooleanField is_active "Inactive users are not allowed to log in."
        BooleanField is_staff "A default Django flag. Not in use in pretalx."
        BooleanField is_administrator "Should only be ``True`` for people with administrative access to the server pretalx runs on."
        BooleanField is_superuser "Never set this flag to ``True``, since it short-circuits all authorisation mechanisms."
        UserLocale locale
        UserTimezone timezone
        ForeignKey profile_picture_id FK
        CharField pw_reset_token
        DateTimeField pw_reset_time
    }
    schedule_availability {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        ForeignKey person_id FK
        ForeignKey room_id FK
        DateTimeField start
        DateTimeField end
    }
    schedule_room {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        I18nCharField name
        UUIDField guid "Unique identifier (UUID) to help external tools identify the room."
        I18nCharField description "A description for attendees, for example directions."
        I18nCharField speaker_info "Information relevant for speakers scheduled in this room, for example room size, special directions, available adaptors for video input …"
        PositiveIntegerField capacity "How many people can fit in the room?"
        PositiveIntegerField position
    }
    schedule_schedule {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        CharField version
        DateTimeField published
        I18nTextField comment "This text will be shown in the public changelog and the RSS feed. You can use <a href='https://docs.pretalx.org/user/markdown/' target='_blank' rel='noopener'>Markdown</a> here."
    }
    schedule_talkslot {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey submission_id FK
        ForeignKey room_id FK "The room this session is scheduled in, if any"
        ForeignKey schedule_id FK
        BooleanField is_visible
        SlotType slot_type "For non-submission slots: 'break' for public breaks, 'blocker' for hidden blockers"
        DateTimeField start "When the session starts, if it is currently scheduled"
        DateTimeField end "When the session ends, if it is currently scheduled"
        I18nCharField description
    }
    submission_submitteraccesscode {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        CharField code
        DateTimeField valid_until "You can set or change this date later to invalidate the access code."
        PositiveIntegerField maximum_uses "Number of times this access code can be used to submit a proposal. Leave empty for no limit."
        PositiveIntegerField redeemed
        TextField internal_notes "Internal notes for other organisers/reviewers. Not visible to the speakers or the public."
    }
    submission_cfp {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        OneToOneField event_id FK, UK
        I18nCharField headline
        I18nTextField text "You can use <a href='https://docs.pretalx.org/user/markdown/' target='_blank' rel='noopener'>Markdown</a> here."
        ForeignKey default_type_id FK
        DateTimeField opening "Please put in the date you want to start accepting proposals from users. Leave the date empty to start accepting proposals immediately once the event is live."
        DateTimeField deadline "Please put in the last date you want to accept proposals from users."
        JSONField settings
        JSONField fields
    }
    submission_submissioncomment {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey submission_id FK
        ForeignKey user_id FK
        MarkdownField text
        ForeignKey reply_to_id FK
    }
    submission_feedback {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey talk_id FK
        ForeignKey speaker_id FK
        IntegerField rating
        MarkdownField review
    }
    submission_question {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        QuestionVariant variant
        QuestionTarget target "Do you require an answer from every speaker or for every session?"
        DateTimeField deadline "Set a deadline to make this field required after the given date."
        DateTimeField freeze_after "Set a deadline to stop changes to responses after the given date."
        QuestionRequired question_required
        I18nCharField question
        I18nCharField help_text "Will appear just like this text below the custom input field. You can use <a href='https://docs.pretalx.org/user/markdown/' target='_blank' rel='noopener'>Markdown</a> here."
        TextField default_answer
        IntegerField position
        CharField identifier "You can enter any value here to make it easier to match the data with other sources. If you do not input one, we will generate one automatically."
        BooleanField active "Inactive fields will no longer be shown."
        BooleanField contains_personal_data "If a user deletes their account, responses containing personal data will be removed, too."
        PositiveIntegerField min_length "Minimum text length in characters or words (set in CfP settings)."
        PositiveIntegerField max_length "Maximum text length in characters or words (set in CfP settings)."
        DecimalField min_number
        DecimalField max_number
        DateField min_date
        DateField max_date
        DateTimeField min_datetime
        DateTimeField max_datetime
        BooleanField is_public "Responses will be shown on session or speaker pages as appropriate."
        BooleanField is_visible_to_reviewers "Should responses to this field be shown to reviewers? This is helpful if you want to collect personal information, but use anonymous reviews."
        QuestionIcon icon "Custom URL fields that are shown publicly can use an icon when displaying the link."
    }
    submission_answeroption {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey question_id FK
        I18nCharField answer
        IntegerField position
        CharField identifier "You can enter any value here to make it easier to match the data with other sources. If you do not input one, we will generate one automatically."
    }
    submission_answer {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey question_id FK
        ForeignKey submission_id FK
        ForeignKey speaker_id FK
        ForeignKey review_id FK
        TextField answer
        FileField answer_file
    }
    submission_resource {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey submission_id FK
        FileField resource
        URLField link
        CharField description
        BooleanField is_public
    }
    submission_reviewscorecategory {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        I18nCharField name
        DecimalField weight
        BooleanField required
        BooleanField active
        BooleanField is_independent "Independent scores are not part of the total score. Instead they are shown in a separate column in the review dashboard."
    }
    submission_reviewscore {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey category_id FK
        DecimalField value
        CharField label
    }
    submission_review {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey submission_id FK
        ForeignKey user_id FK
        MarkdownField text
        DecimalField score
    }
    submission_reviewphase {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        CharField name
        DateTimeField start
        DateTimeField end
        BooleanField is_active
        BooleanField can_review
        ReviewPhaseProposalVisibility proposal_visibility "If you select “all”, reviewers can review all proposals that their teams have access to (so either all, or specific tracks). In this mode, assigned proposals will be highlighted and will be shown first in the review workflow."
        ReviewPhaseCanSeeOtherReviews can_see_other_reviews
        BooleanField can_see_speaker_names
        BooleanField can_see_reviewer_names
        BooleanField can_change_submission_state
        ReviewPhaseCanTagSubmissions can_tag_submissions
        BooleanField speakers_can_change_submissions "By default, modification of proposals is locked after the CfP ends, and is re-enabled once the proposal was accepted."
    }
    submission_attendeesignup {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey submission_id FK
        ForeignKey attendee_id FK
        AttendeeSignupStates state
        PositiveIntegerField position
    }
    submission_speakerrole {
        AutoField id PK
        ForeignKey submission_id FK
        ForeignKey speaker_id FK
        PositiveIntegerField position
    }
    submission_submission {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        CharField code UK
        ForeignKey event_id FK
        CharField title
        ForeignKey submission_type_id FK
        ForeignKey track_id FK
        SubmissionStates state
        SubmissionStates pending_state
        MarkdownField abstract
        MarkdownField description
        MarkdownField notes "These notes are meant for the organisers and won’t be made public."
        TextField internal_notes "Internal notes for other organisers/reviewers. Not visible to the speakers or the public."
        PositiveIntegerField duration "Duration in minutes"
        IntegerField slot_count "How often this session takes place."
        BooleanField attendee_signup_required "Override whether attendees must sign up to attend this session."
        PositiveIntegerField attendee_signup_capacity "Override the room capacity for this session."
        CharField content_locale
        BooleanField is_featured
        BooleanField do_not_record
        ImageField image "Use this if you want an illustration to go with your proposal."
        CharField invitation_token
        ForeignKey access_code_id FK
        CharField review_code UK
        JSONField anonymised
        JSONField draft_additional_speakers
    }
    submission_submissionfavourite {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey user_id FK
        ForeignKey submission_id FK
    }
    submission_submissioninvitation {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey submission_id FK
        EmailField email
        CharField token UK
    }
    submission_tag {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        CharField tag
        I18nTextField description
        CharField color
        BooleanField is_public "Public tags can be selected by submitters in the CfP if you activate the tags field. In a future release of pretalx, public tags may also be shown in the public schedule."
    }
    submission_track {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        I18nCharField name
        I18nTextField description
        CharField color
        PositiveIntegerField position "The position field is used to determine the order that tracks are displayed in (lowest first)."
        BooleanField requires_access_code "This track will only be shown to submitters with a matching access code."
        BooleanField attendee_signup_required "Sessions will require attendee signup by default. You can always override this setting for individual sessions."
    }
    submission_submissiontype {
        AutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey event_id FK
        I18nCharField name
        PositiveIntegerField default_duration "Duration in minutes"
        DateTimeField deadline "If you want a different deadline than the global deadline for this session type, enter it here."
        BooleanField requires_access_code "This session type will only be shown to submitters with a matching access code."
        BooleanField attendee_signup_required "Sessions will require attendee signup by default. You can always override this setting for individual sessions."
    }
```
