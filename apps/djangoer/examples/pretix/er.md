```mermaid
erDiagram
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    pretixbase_event ||--o{ pretixbase_taxrule : "(id) - (event_id)"
    pretixbase_user }o--o{ auth_group : "(groups) - (user_set)"
    pretixbase_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    pretixbase_user ||--o{ pretixbase_userknownloginsource : "(id) - (user_id)"
    pretixbase_user ||--o{ pretixbase_staffsession : "(id) - (user_id)"
    pretixbase_staffsession ||--o{ pretixbase_staffsessionauditlog : "(id) - (session_id)"
    pretixbase_user |o--o{ pretixbase_staffsessionauditlog : "(id) - (impersonating_id)"
    pretixbase_user ||--o{ pretixbase_u2fdevice : "(id) - (user_id)"
    pretixbase_user ||--o{ pretixbase_webauthndevice : "(id) - (user_id)"
    pretixbase_user ||--o{ pretixbase_historicpassword : "(id) - (user_id)"
    pretixbase_event ||--o{ pretixbase_checkinlist : "(id) - (event_id)"
    pretixbase_subevent |o--o{ pretixbase_checkinlist : "(id) - (subevent_id)"
    pretixbase_checkinlist }o--o{ pretixbase_item : "(limit_products) - (checkinlist_set)"
    pretixbase_checkinlist }o--o{ pretixbase_gate : "(gates) - (checkinlist_set)"
    pretixbase_orderposition |o--o{ pretixbase_checkin : "(id) - (position_id)"
    pretixbase_item |o--o{ pretixbase_checkin : "(id) - (raw_item_id)"
    pretixbase_itemvariation |o--o{ pretixbase_checkin : "(id) - (raw_variation_id)"
    pretixbase_subevent |o--o{ pretixbase_checkin : "(id) - (raw_subevent_id)"
    pretixbase_checkinlist ||--o{ pretixbase_checkin : "(id) - (list_id)"
    pretixbase_device |o--o{ pretixbase_checkin : "(id) - (device_id)"
    pretixbase_gate |o--o{ pretixbase_checkin : "(id) - (gate_id)"
    pretixbase_organizer ||--o{ pretixbase_giftcardacceptance : "(id) - (issuer_id)"
    pretixbase_organizer ||--o{ pretixbase_giftcardacceptance : "(id) - (acceptor_id)"
    pretixbase_organizer ||--o{ pretixbase_giftcard : "(id) - (issuer_id)"
    pretixbase_orderposition |o--o{ pretixbase_giftcard : "(id) - (issued_in_id)"
    pretixbase_orderposition |o--o{ pretixbase_giftcard : "(id) - (owner_ticket_id)"
    pretixbase_customer |o--o{ pretixbase_giftcard : "(id) - (customer_id)"
    pretixbase_giftcard ||--o{ pretixbase_giftcardtransaction : "(id) - (card_id)"
    pretixbase_order |o--o{ pretixbase_giftcardtransaction : "(id) - (order_id)"
    pretixbase_orderpayment |o--o{ pretixbase_giftcardtransaction : "(id) - (payment_id)"
    pretixbase_orderrefund |o--o{ pretixbase_giftcardtransaction : "(id) - (refund_id)"
    pretixbase_organizer |o--o{ pretixbase_giftcardtransaction : "(id) - (acceptor_id)"
    pretixbase_organizer ||--o{ pretixbase_organizer_settingsstore : "(id) - (object_id)"
    pretixbase_organizer ||--o{ pretixbase_team : "(id) - (organizer_id)"
    pretixbase_team }o--o{ pretixbase_user : "(members) - (teams)"
    pretixbase_team }o--o{ pretixbase_event : "(limit_events) - (team_set)"
    pretixbase_team ||--o{ pretixbase_teaminvite : "(id) - (team_id)"
    pretixbase_team ||--o{ pretixbase_teamapitoken : "(id) - (team_id)"
    pretixbase_organizer ||--o{ pretixbase_organizerfooterlink : "(id) - (organizer_id)"
    pretixbase_organizer ||--o{ pretixbase_saleschannel : "(id) - (organizer_id)"
    pretixbase_organizer ||--o{ pretixbase_customerssoprovider : "(id) - (organizer_id)"
    pretixbase_organizer ||--o{ pretixbase_customer : "(id) - (organizer_id)"
    pretixbase_customerssoprovider |o--o{ pretixbase_customer : "(id) - (provider_id)"
    pretixbase_customer ||--o{ pretixbase_attendeeprofile : "(id) - (customer_id)"
    pretixbase_organizer ||--o{ pretixbase_customerssoclient : "(id) - (organizer_id)"
    pretixbase_customerssoclient ||--o{ pretixbase_customerssogrant : "(id) - (client_id)"
    pretixbase_customer ||--o{ pretixbase_customerssogrant : "(id) - (customer_id)"
    pretixbase_customerssoclient ||--o{ pretixbase_customerssoaccesstoken : "(id) - (client_id)"
    pretixbase_customer ||--o{ pretixbase_customerssoaccesstoken : "(id) - (customer_id)"
    pretixbase_organizer ||--o{ pretixbase_event : "(id) - (organizer_id)"
    pretixbase_seatingplan |o--o{ pretixbase_event : "(id) - (seating_plan_id)"
    pretixbase_event }o--o{ pretixbase_saleschannel : "(limit_sales_channels) - (event_set)"
    pretixbase_event ||--o{ pretixbase_event_settingsstore : "(id) - (object_id)"
    pretixbase_event ||--o{ pretixbase_subevent : "(id) - (event_id)"
    pretixbase_seatingplan |o--o{ pretixbase_subevent : "(id) - (seating_plan_id)"
    pretixbase_organizer ||--o{ pretixbase_eventmetaproperty : "(id) - (organizer_id)"
    pretixbase_event ||--o{ pretixbase_eventmetavalue : "(id) - (event_id)"
    pretixbase_eventmetaproperty ||--o{ pretixbase_eventmetavalue : "(id) - (property_id)"
    pretixbase_subevent ||--o{ pretixbase_subeventmetavalue : "(id) - (subevent_id)"
    pretixbase_eventmetaproperty ||--o{ pretixbase_subeventmetavalue : "(id) - (property_id)"
    pretixbase_event ||--o{ pretixbase_eventfooterlink : "(id) - (event_id)"
    pretixbase_organizer ||--o{ pretixbase_gate : "(id) - (organizer_id)"
    pretixbase_organizer ||--o{ pretixbase_device : "(id) - (organizer_id)"
    pretixbase_gate |o--o{ pretixbase_device : "(id) - (gate_id)"
    pretixbase_device }o--o{ pretixbase_event : "(limit_events) - (device_set)"
    pretixbase_device ||--o| pretixbase_devicelastseen : "(id) - (device_id)"
    pretixbase_event ||--o{ pretixbase_discount : "(id) - (event_id)"
    pretixbase_discount }o--o{ pretixbase_saleschannel : "(limit_sales_channels) - (discount_set)"
    pretixbase_discount }o--o{ pretixbase_item : "(condition_limit_products) - (discount_set)"
    pretixbase_discount }o--o{ pretixbase_item : "(benefit_limit_products) - (benefit_discounts)"
    pretixbase_user ||--o{ pretixbase_scheduledeventexport : "(id) - (owner_id)"
    pretixbase_event ||--o{ pretixbase_scheduledeventexport : "(id) - (event_id)"
    pretixbase_user ||--o{ pretixbase_scheduledorganizerexport : "(id) - (owner_id)"
    pretixbase_organizer ||--o{ pretixbase_scheduledorganizerexport : "(id) - (organizer_id)"
    pretixbase_order ||--o{ pretixbase_invoice : "(id) - (order_id)"
    pretixbase_organizer ||--o{ pretixbase_invoice : "(id) - (organizer_id)"
    pretixbase_event ||--o{ pretixbase_invoice : "(id) - (event_id)"
    pretixbase_invoice |o--o{ pretixbase_invoice : "(id) - (refers_id)"
    pretixbase_invoice ||--o{ pretixbase_invoiceline : "(id) - (invoice_id)"
    pretixbase_subevent |o--o{ pretixbase_invoiceline : "(id) - (subevent_id)"
    pretixbase_item |o--o{ pretixbase_invoiceline : "(id) - (item_id)"
    pretixbase_itemvariation |o--o{ pretixbase_invoiceline : "(id) - (variation_id)"
    pretixbase_event ||--o{ pretixbase_itemcategory : "(id) - (event_id)"
    pretixbase_itemcategory }o--o{ pretixbase_item : "(cross_selling_match_products) - (matched_by_cross_selling_categories)"
    pretixbase_subevent ||--o{ pretixbase_subeventitem : "(id) - (subevent_id)"
    pretixbase_item ||--o{ pretixbase_subeventitem : "(id) - (item_id)"
    pretixbase_subevent ||--o{ pretixbase_subeventitemvariation : "(id) - (subevent_id)"
    pretixbase_itemvariation ||--o{ pretixbase_subeventitemvariation : "(id) - (variation_id)"
    pretixbase_event ||--o{ pretixbase_item : "(id) - (event_id)"
    pretixbase_itemcategory |o--o{ pretixbase_item : "(id) - (category_id)"
    pretixbase_taxrule |o--o{ pretixbase_item : "(id) - (tax_rule_id)"
    pretixbase_quota |o--o{ pretixbase_item : "(id) - (hidden_if_available_id)"
    pretixbase_item |o--o{ pretixbase_item : "(id) - (hidden_if_item_available_id)"
    pretixbase_membershiptype |o--o{ pretixbase_item : "(id) - (grant_membership_type_id)"
    pretixbase_item }o--o{ pretixbase_saleschannel : "(limit_sales_channels) - (item_set)"
    pretixbase_item }o--o{ pretixbase_membershiptype : "(require_membership_types) - (item_set)"
    pretixbase_item ||--o{ pretixbase_itemvariation : "(id) - (item_id)"
    pretixbase_itemvariation }o--o{ pretixbase_membershiptype : "(require_membership_types) - (itemvariation_set)"
    pretixbase_itemvariation }o--o{ pretixbase_saleschannel : "(limit_sales_channels) - (itemvariation_set)"
    pretixbase_item ||--o{ pretixbase_itemaddon : "(id) - (base_item_id)"
    pretixbase_itemcategory ||--o{ pretixbase_itemaddon : "(id) - (addon_category_id)"
    pretixbase_item ||--o{ pretixbase_itembundle : "(id) - (base_item_id)"
    pretixbase_item ||--o{ pretixbase_itembundle : "(id) - (bundled_item_id)"
    pretixbase_itemvariation |o--o{ pretixbase_itembundle : "(id) - (bundled_variation_id)"
    pretixbase_event ||--o{ pretixbase_question : "(id) - (event_id)"
    pretixbase_question |o--o{ pretixbase_question : "(id) - (dependency_question_id)"
    pretixbase_question }o--o{ pretixbase_item : "(items) - (questions)"
    pretixbase_question ||--o{ pretixbase_questionoption : "(id) - (question_id)"
    pretixbase_event ||--o{ pretixbase_quota : "(id) - (event_id)"
    pretixbase_subevent |o--o{ pretixbase_quota : "(id) - (subevent_id)"
    pretixbase_quota }o--o{ pretixbase_item : "(items) - (quotas)"
    pretixbase_quota }o--o{ pretixbase_itemvariation : "(variations) - (quotas)"
    pretixbase_event ||--o{ pretixbase_itemmetaproperty : "(id) - (event_id)"
    pretixbase_item ||--o{ pretixbase_itemmetavalue : "(id) - (item_id)"
    pretixbase_itemmetaproperty ||--o{ pretixbase_itemmetavalue : "(id) - (property_id)"
    pretixbase_itemvariation ||--o{ pretixbase_itemvariationmetavalue : "(id) - (variation_id)"
    pretixbase_itemmetaproperty ||--o{ pretixbase_itemvariationmetavalue : "(id) - (property_id)"
    pretixbase_item ||--o{ pretixbase_itemprogramtime : "(id) - (item_id)"
    django_content_type ||--o{ pretixbase_logentry : "(id) - (content_type_id)"
    pretixbase_user |o--o{ pretixbase_logentry : "(id) - (user_id)"
    pretixbase_teamapitoken |o--o{ pretixbase_logentry : "(id) - (api_token_id)"
    pretixbase_device |o--o{ pretixbase_logentry : "(id) - (device_id)"
    pretixapi_oauthapplication |o--o{ pretixbase_logentry : "(id) - (oauth_application_id)"
    pretixbase_event |o--o{ pretixbase_logentry : "(id) - (event_id)"
    pretixbase_organizer |o--o{ pretixbase_logentry : "(id) - (organizer_link_id)"
    pretixbase_organizer |o--o{ pretixbase_outgoingmail : "(id) - (organizer_id)"
    pretixbase_event |o--o{ pretixbase_outgoingmail : "(id) - (event_id)"
    pretixbase_order |o--o{ pretixbase_outgoingmail : "(id) - (order_id)"
    pretixbase_orderposition |o--o{ pretixbase_outgoingmail : "(id) - (orderposition_id)"
    pretixbase_customer |o--o{ pretixbase_outgoingmail : "(id) - (customer_id)"
    pretixbase_user |o--o{ pretixbase_outgoingmail : "(id) - (user_id)"
    pretixbase_outgoingmail }o--o{ pretixbase_invoice : "(should_attach_invoices) - (outgoing_mails)"
    pretixbase_outgoingmail }o--o{ pretixbase_cachedfile : "(should_attach_cached_files) - (outgoing_mails)"
    pretixbase_organizer ||--o{ pretixbase_order : "(id) - (organizer_id)"
    pretixbase_event ||--o{ pretixbase_order : "(id) - (event_id)"
    pretixbase_customer |o--o{ pretixbase_order : "(id) - (customer_id)"
    pretixbase_saleschannel ||--o{ pretixbase_order : "(id) - (sales_channel_id)"
    pretixbase_orderposition |o--o{ pretixbase_questionanswer : "(id) - (orderposition_id)"
    pretixbase_cartposition |o--o{ pretixbase_questionanswer : "(id) - (cartposition_id)"
    pretixbase_order |o--o{ pretixbase_questionanswer : "(id) - (order_id)"
    pretixbase_checkoutsession |o--o{ pretixbase_questionanswer : "(id) - (checkoutsession_id)"
    pretixbase_question ||--o{ pretixbase_questionanswer : "(id) - (question_id)"
    pretixbase_questionanswer }o--o{ pretixbase_questionoption : "(options) - (answers)"
    pretixbase_order ||--o{ pretixbase_orderpayment : "(id) - (order_id)"
    pretixbase_orderfee |o--o{ pretixbase_orderpayment : "(id) - (fee_id)"
    pretixbase_order ||--o{ pretixbase_orderrefund : "(id) - (order_id)"
    pretixbase_orderpayment |o--o{ pretixbase_orderrefund : "(id) - (payment_id)"
    pretixbase_order ||--o{ pretixbase_orderfee : "(id) - (order_id)"
    pretixbase_taxrule |o--o{ pretixbase_orderfee : "(id) - (tax_rule_id)"
    pretixbase_subevent |o--o{ pretixbase_orderposition : "(id) - (subevent_id)"
    pretixbase_item ||--o{ pretixbase_orderposition : "(id) - (item_id)"
    pretixbase_itemvariation |o--o{ pretixbase_orderposition : "(id) - (variation_id)"
    pretixbase_voucher |o--o{ pretixbase_orderposition : "(id) - (voucher_id)"
    pretixbase_membership |o--o{ pretixbase_orderposition : "(id) - (used_membership_id)"
    pretixbase_orderposition |o--o{ pretixbase_orderposition : "(id) - (addon_to_id)"
    pretixbase_seat |o--o{ pretixbase_orderposition : "(id) - (seat_id)"
    pretixbase_discount |o--o{ pretixbase_orderposition : "(id) - (discount_id)"
    pretixbase_organizer ||--o{ pretixbase_orderposition : "(id) - (organizer_id)"
    pretixbase_order ||--o{ pretixbase_orderposition : "(id) - (order_id)"
    pretixbase_taxrule |o--o{ pretixbase_orderposition : "(id) - (tax_rule_id)"
    pretixbase_order ||--o{ pretixbase_transaction : "(id) - (order_id)"
    pretixbase_item |o--o{ pretixbase_transaction : "(id) - (item_id)"
    pretixbase_itemvariation |o--o{ pretixbase_transaction : "(id) - (variation_id)"
    pretixbase_subevent |o--o{ pretixbase_transaction : "(id) - (subevent_id)"
    pretixbase_taxrule |o--o{ pretixbase_transaction : "(id) - (tax_rule_id)"
    pretixbase_event ||--o{ pretixbase_checkoutsession : "(id) - (event_id)"
    pretixbase_customer |o--o{ pretixbase_checkoutsession : "(id) - (customer_id)"
    pretixbase_saleschannel ||--o{ pretixbase_checkoutsession : "(id) - (sales_channel_id)"
    pretixbase_subevent |o--o{ pretixbase_cartposition : "(id) - (subevent_id)"
    pretixbase_item ||--o{ pretixbase_cartposition : "(id) - (item_id)"
    pretixbase_itemvariation |o--o{ pretixbase_cartposition : "(id) - (variation_id)"
    pretixbase_voucher |o--o{ pretixbase_cartposition : "(id) - (voucher_id)"
    pretixbase_membership |o--o{ pretixbase_cartposition : "(id) - (used_membership_id)"
    pretixbase_cartposition |o--o{ pretixbase_cartposition : "(id) - (addon_to_id)"
    pretixbase_seat |o--o{ pretixbase_cartposition : "(id) - (seat_id)"
    pretixbase_discount |o--o{ pretixbase_cartposition : "(id) - (discount_id)"
    pretixbase_event ||--o{ pretixbase_cartposition : "(id) - (event_id)"
    pretixbase_checkoutsession |o--o| pretixbase_invoiceaddress : "(id) - (checkout_session_id)"
    pretixbase_order |o--o| pretixbase_invoiceaddress : "(id) - (order_id)"
    pretixbase_customer |o--o{ pretixbase_invoiceaddress : "(id) - (customer_id)"
    pretixbase_orderposition ||--o{ pretixbase_cachedticket : "(id) - (order_position_id)"
    pretixbase_order ||--o{ pretixbase_cachedcombinedticket : "(id) - (order_id)"
    pretixbase_order ||--o{ pretixbase_cancellationrequest : "(id) - (order_id)"
    pretixbase_event ||--o{ pretixbase_revokedticketsecret : "(id) - (event_id)"
    pretixbase_orderposition |o--o{ pretixbase_revokedticketsecret : "(id) - (position_id)"
    pretixbase_event ||--o{ pretixbase_blockedticketsecret : "(id) - (event_id)"
    pretixbase_orderposition |o--o{ pretixbase_blockedticketsecret : "(id) - (position_id)"
    pretixbase_orderposition ||--o{ pretixbase_printlog : "(id) - (position_id)"
    pretixbase_device |o--o{ pretixbase_printlog : "(id) - (device_id)"
    pretixbase_user |o--o{ pretixbase_printlog : "(id) - (user_id)"
    pretixbase_teamapitoken |o--o{ pretixbase_printlog : "(id) - (api_token_id)"
    pretixapi_oauthapplication |o--o{ pretixbase_printlog : "(id) - (oauth_application_id)"
    pretixbase_organizer ||--o{ pretixbase_reusablemedium : "(id) - (organizer_id)"
    pretixbase_customer |o--o{ pretixbase_reusablemedium : "(id) - (customer_id)"
    pretixbase_giftcard |o--o{ pretixbase_reusablemedium : "(id) - (linked_giftcard_id)"
    pretixbase_reusablemedium }o--o{ pretixbase_orderposition : "(linked_orderpositions) - (linked_media)"
    pretixbase_organizer ||--o{ pretixbase_mediumkeyset : "(id) - (organizer_id)"
    pretixbase_organizer ||--o{ pretixbase_membershiptype : "(id) - (organizer_id)"
    pretixbase_customer ||--o{ pretixbase_membership : "(id) - (customer_id)"
    pretixbase_membershiptype ||--o{ pretixbase_membership : "(id) - (membership_type_id)"
    pretixbase_orderposition |o--o{ pretixbase_membership : "(id) - (granted_in_id)"
    pretixbase_user ||--o{ pretixbase_notificationsetting : "(id) - (user_id)"
    pretixbase_event |o--o{ pretixbase_notificationsetting : "(id) - (event_id)"
    pretixbase_organizer ||--o{ pretixbase_seatingplan : "(id) - (organizer_id)"
    pretixbase_event ||--o{ pretixbase_seatcategorymapping : "(id) - (event_id)"
    pretixbase_subevent |o--o{ pretixbase_seatcategorymapping : "(id) - (subevent_id)"
    pretixbase_item ||--o{ pretixbase_seatcategorymapping : "(id) - (product_id)"
    pretixbase_event ||--o{ pretixbase_seat : "(id) - (event_id)"
    pretixbase_subevent |o--o{ pretixbase_seat : "(id) - (subevent_id)"
    pretixbase_item |o--o{ pretixbase_seat : "(id) - (product_id)"
    pretixbase_event ||--o{ pretixbase_voucher : "(id) - (event_id)"
    pretixbase_subevent |o--o{ pretixbase_voucher : "(id) - (subevent_id)"
    pretixbase_item |o--o{ pretixbase_voucher : "(id) - (item_id)"
    pretixbase_itemvariation |o--o{ pretixbase_voucher : "(id) - (variation_id)"
    pretixbase_quota |o--o{ pretixbase_voucher : "(id) - (quota_id)"
    pretixbase_seat |o--o{ pretixbase_voucher : "(id) - (seat_id)"
    pretixbase_event ||--o{ pretixbase_waitinglistentry : "(id) - (event_id)"
    pretixbase_subevent |o--o{ pretixbase_waitinglistentry : "(id) - (subevent_id)"
    pretixbase_voucher |o--o{ pretixbase_waitinglistentry : "(id) - (voucher_id)"
    pretixbase_item ||--o{ pretixbase_waitinglistentry : "(id) - (item_id)"
    pretixbase_itemvariation |o--o{ pretixbase_waitinglistentry : "(id) - (variation_id)"
    pretixbase_order ||--o{ pretixbase_ordersyncqueue : "(id) - (order_id)"
    pretixbase_event ||--o{ pretixbase_ordersyncqueue : "(id) - (event_id)"
    pretixbase_order ||--o{ pretixbase_ordersyncresult : "(id) - (order_id)"
    pretixbase_orderposition |o--o{ pretixbase_ordersyncresult : "(id) - (order_position_id)"
    pretixbase_organizer |o--o{ pretixmultidomain_knowndomain : "(id) - (organizer_id)"
    pretixbase_event |o--o| pretixmultidomain_knowndomain : "(id) - (event_id)"
    pretixmultidomain_knowndomain ||--o{ pretixmultidomain_alternativedomainassignment : "(domainname) - (domain_id)"
    pretixbase_event ||--o| pretixmultidomain_alternativedomainassignment : "(id) - (event_id)"
    pretixbase_user |o--o{ pretixapi_oauthapplication : "(id) - (user_id)"
    pretixbase_user ||--o{ pretixapi_oauthgrant : "(id) - (user_id)"
    pretixapi_oauthapplication ||--o{ pretixapi_oauthgrant : "(id) - (application_id)"
    pretixapi_oauthgrant }o--o{ pretixbase_organizer : "(organizers) - (oauthgrant_set)"
    pretixbase_user |o--o{ pretixapi_oauthidtoken : "(id) - (user_id)"
    pretixapi_oauthapplication ||--o{ pretixapi_oauthidtoken : "(id) - (application_id)"
    pretixapi_oauthidtoken }o--o{ pretixbase_organizer : "(organizers) - (oauthidtoken_set)"
    pretixbase_user |o--o{ pretixapi_oauthaccesstoken : "(id) - (user_id)"
    pretixapi_oauthrefreshtoken |o--o| pretixapi_oauthaccesstoken : "(id) - (source_refresh_token_id)"
    pretixapi_oauthidtoken |o--o| pretixapi_oauthaccesstoken : "(id) - (id_token_id)"
    pretixapi_oauthapplication |o--o{ pretixapi_oauthaccesstoken : "(id) - (application_id)"
    pretixapi_oauthaccesstoken }o--o{ pretixbase_organizer : "(organizers) - (oauthaccesstoken_set)"
    pretixbase_user ||--o{ pretixapi_oauthrefreshtoken : "(id) - (user_id)"
    pretixapi_oauthapplication ||--o{ pretixapi_oauthrefreshtoken : "(id) - (application_id)"
    pretixapi_oauthaccesstoken |o--o| pretixapi_oauthrefreshtoken : "(id) - (access_token_id)"
    pretixbase_organizer ||--o{ pretixapi_webhook : "(id) - (organizer_id)"
    pretixapi_webhook }o--o{ pretixbase_event : "(limit_events) - (webhook_set)"
    pretixapi_webhook ||--o{ pretixapi_webhookeventlistener : "(id) - (webhook_id)"
    pretixapi_webhook ||--o{ pretixapi_webhookcall : "(id) - (webhook_id)"
    pretixapi_webhook ||--o{ pretixapi_webhookcallretry : "(id) - (webhook_id)"
    pretixbase_logentry ||--o{ pretixapi_webhookcallretry : "(id) - (logentry_id)"
    pretixbase_event |o--o{ banktransfer_bankimportjob : "(id) - (event_id)"
    pretixbase_organizer |o--o{ banktransfer_bankimportjob : "(id) - (organizer_id)"
    pretixbase_event |o--o{ banktransfer_banktransaction : "(id) - (event_id)"
    pretixbase_organizer |o--o{ banktransfer_banktransaction : "(id) - (organizer_id)"
    banktransfer_bankimportjob ||--o{ banktransfer_banktransaction : "(id) - (import_job_id)"
    pretixbase_order |o--o{ banktransfer_banktransaction : "(id) - (order_id)"
    pretixbase_event |o--o{ banktransfer_refundexport : "(id) - (event_id)"
    pretixbase_organizer |o--o{ banktransfer_refundexport : "(id) - (organizer_id)"
    pretixbase_order ||--o{ stripe_referencedstripeobject : "(id) - (order_id)"
    pretixbase_orderpayment |o--o{ stripe_referencedstripeobject : "(id) - (payment_id)"
    pretixbase_order ||--o{ paypal_referencedpaypalobject : "(id) - (order_id)"
    pretixbase_orderpayment |o--o{ paypal_referencedpaypalobject : "(id) - (payment_id)"
    pretixbase_event ||--o{ ticketoutputpdf_ticketlayout : "(id) - (event_id)"
    pretixbase_item |o--o{ ticketoutputpdf_ticketlayoutitem : "(id) - (item_id)"
    ticketoutputpdf_ticketlayout ||--o{ ticketoutputpdf_ticketlayoutitem : "(id) - (layout_id)"
    pretixbase_saleschannel ||--o{ ticketoutputpdf_ticketlayoutitem : "(id) - (sales_channel_id)"
    sendmail_rule ||--o{ sendmail_scheduledmail : "(id) - (rule_id)"
    pretixbase_subevent |o--o{ sendmail_scheduledmail : "(id) - (subevent_id)"
    pretixbase_event ||--o{ sendmail_scheduledmail : "(id) - (event_id)"
    pretixbase_event ||--o{ sendmail_rule : "(id) - (event_id)"
    pretixbase_subevent |o--o{ sendmail_rule : "(id) - (subevent_id)"
    sendmail_rule }o--o{ pretixbase_item : "(limit_products) - (rule_set)"
    pretixbase_event ||--o{ badges_badgelayout : "(id) - (event_id)"
    pretixbase_item |o--o| badges_badgeitem : "(id) - (item_id)"
    badges_badgelayout |o--o{ badges_badgeitem : "(id) - (layout_id)"
    pretixbase_event ||--o{ autocheckin_autocheckinrule : "(id) - (event_id)"
    pretixbase_checkinlist |o--o{ autocheckin_autocheckinrule : "(id) - (list_id)"
    autocheckin_autocheckinrule }o--o{ pretixbase_saleschannel : "(limit_sales_channels) - (autocheckinrule_set)"
    autocheckin_autocheckinrule }o--o{ pretixbase_item : "(limit_products) - (autocheckinrule_set)"
    autocheckin_autocheckinrule }o--o{ pretixbase_itemvariation : "(limit_variations) - (autocheckinrule_set)"
    pretixbase_user ||--o{ otp_totp_totpdevice : "(id) - (user_id)"
    pretixbase_user ||--o{ otp_static_staticdevice : "(id) - (user_id)"
    otp_static_staticdevice ||--o{ otp_static_statictoken : "(id) - (device_id)"
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
    pretixbase_cachedfile {
        UUIDField id PK
        DateTimeField expires
        DateTimeField date
        CharField filename
        CharField type
        FileField file
        BooleanField web_download
        TextField session_key
    }
    pretixbase_taxrule {
        BigAutoField id PK
        ForeignKey event_id FK
        CharField internal_name
        I18nCharField name "Should be short, e.g. 'VAT'"
        TaxRuleCode code "If you help us understand what this tax rules legally is, we can use this information for eInvoices, exporting to accounting system, etc."
        NormalizedDecimalField rate
        BooleanField price_includes_tax
        BooleanField keep_gross_if_rate_changes
        BooleanField eu_reverse_charge "<span class='label label-warning' data-toggle='tooltip' title='This feature will be removed in the future as it does not handle VAT for non-business customers in other EU countries in a way that works for all organizers. Use custom rules instead.'>DEPRECATED</span> Not recommended. Most events will NOT be qualified for reverse charge since the place of taxation is the location of the event. This option disables charging VAT for all customers outside the EU and for business customers in different EU countries who entered a valid EU VAT ID. Only enable this option after consulting a tax counsel. No warranty given for correct tax calculation. USE AT YOUR OWN RISK."
        TaxRuleHomeCountry home_country "Your country of residence. This is the country the EU reverse charge rule will not apply in, if configured above."
        TextField custom_rules
        BooleanField default
    }
    pretixbase_globalsettingsobject_settingsstore {
        BigAutoField id PK
        CharField key
        TextField value
    }
    pretixbase_user {
        BigAutoField id PK
        CharField password
        DateTimeField last_login
        EmailField email UK
        BooleanField is_verified
        CharField fullname
        BooleanField is_active
        BooleanField is_staff
        DateTimeField date_joined
        BooleanField needs_password_change
        UserLocale locale
        CharField timezone
        BooleanField require_2fa
        BooleanField notifications_send "If turned off, you will not get any notifications."
        CharField notifications_token
        CharField auth_backend
        CharField auth_backend_identifier
        CharField session_token
    }
    pretixbase_userknownloginsource {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField agent_type
        CharField device_type
        CharField os_type
        UserKnownLoginSourceCountry country
        DateTimeField last_seen
    }
    pretixbase_staffsession {
        BigAutoField id PK
        ForeignKey user_id FK
        DateTimeField date_start
        DateTimeField date_end
        CharField session_key
        TextField comment
    }
    pretixbase_staffsessionauditlog {
        BigAutoField id PK
        ForeignKey session_id FK
        DateTimeField datetime
        CharField url
        CharField method
        ForeignKey impersonating_id FK
    }
    pretixbase_u2fdevice {
        BigAutoField id PK
        ForeignKey user_id FK "The user that this device belongs to."
        CharField name "The human-readable name of this device."
        BooleanField confirmed "Is this device ready for use?"
        TextField json_data
    }
    pretixbase_webauthndevice {
        BigAutoField id PK
        ForeignKey user_id FK "The user that this device belongs to."
        CharField name "The human-readable name of this device."
        BooleanField confirmed "Is this device ready for use?"
        CharField credential_id
        CharField rp_id
        CharField icon_url
        TextField ukey
        TextField pub_key
        IntegerField sign_count
    }
    pretixbase_historicpassword {
        BigAutoField id PK
        ForeignKey user_id FK
        DateTimeField created
        CharField password
    }
    pretixbase_checkinlist {
        BigAutoField id PK
        ForeignKey event_id FK
        CharField name
        BooleanField all_products
        ForeignKey subevent_id FK "If you choose 'all dates', tickets will be considered part of this list and valid for check-in regardless of which date they are purchased for. You can limit their validity through the advanced check-in rules, though."
        BooleanField ignore_in_statistics
        BooleanField consider_tickets_used "This is relevant in various situations, e.g. for deciding if a ticket can still be canceled by the customer."
        BooleanField include_pending "With this option, people will be able to check in even if the order has not been paid."
        BooleanField addon_match "A scan will only be possible if the check-in list is configured such that there is always exactly one matching add-on ticket. Ambiguous scans will be rejected.."
        BooleanField allow_entry_after_exit
        BooleanField allow_multiple_entries "Use this option to turn off warnings if a ticket is scanned a second time."
        DateTimeField exit_all_at
        JSONField rules
    }
    pretixbase_checkin {
        BigAutoField id PK
        BooleanField successful
        CheckinErrorReason error_reason
        TextField error_explanation
        ForeignKey position_id FK
        TextField raw_barcode
        CheckinRawSourceType raw_source_type
        ForeignKey raw_item_id FK
        ForeignKey raw_variation_id FK
        ForeignKey raw_subevent_id FK
        DateTimeField datetime
        DateTimeField created
        ForeignKey list_id FK
        CheckinType type
        CharField nonce
        BooleanField force_sent
        BooleanField forced
        ForeignKey device_id FK
        ForeignKey gate_id FK
        BooleanField auto_checked_in
    }
    pretixbase_exchangerate {
        BigAutoField id PK
        CharField source
        DateField source_date
        DateTimeField updated
        CharField source_currency
        CharField other_currency
        DecimalField rate
    }
    pretixbase_giftcardacceptance {
        BigAutoField id PK
        ForeignKey issuer_id FK
        ForeignKey acceptor_id FK
        BooleanField active
        BooleanField reusable_media
    }
    pretixbase_giftcard {
        BigAutoField id PK
        ForeignKey issuer_id FK
        ForeignKey issued_in_id FK
        ForeignKey owner_ticket_id FK
        ForeignKey customer_id FK
        DateTimeField issuance
        CharField secret
        BooleanField testmode
        DateTimeField expires
        TextField conditions
        GiftCardCurrency currency
    }
    pretixbase_giftcardtransaction {
        BigAutoField id PK
        ForeignKey card_id FK
        DateTimeField datetime
        DecimalField value
        ForeignKey order_id FK
        ForeignKey payment_id FK
        ForeignKey refund_id FK
        TextField text
        JSONField info
        ForeignKey acceptor_id FK
    }
    pretixbase_organizer {
        BigAutoField id PK
        CharField name
        CharField slug UK "Should be short, only contain lowercase letters, numbers, dots, and dashes. Every slug can only be used once. This is being used in URLs to refer to your organizer accounts and your events."
        TextField plugins
    }
    pretixbase_organizer_settingsstore {
        BigAutoField id PK
        CharField key
        TextField value
        ForeignKey object_id FK
    }
    pretixbase_team {
        BigAutoField id PK
        ForeignKey organizer_id FK
        CharField name
        BooleanField require_2fa "If you turn this on, all members of the team will be required to either set up two-factor authentication or leave the team. The setting may take a few minutes to become effective for all users."
        BooleanField all_events
        BooleanField all_event_permissions
        JSONField limit_event_permissions
        BooleanField all_organizer_permissions
        JSONField limit_organizer_permissions
    }
    pretixbase_teaminvite {
        BigAutoField id PK
        ForeignKey team_id FK
        EmailField email
        CharField token
    }
    pretixbase_teamapitoken {
        BigAutoField id PK
        ForeignKey team_id FK
        CharField name
        BooleanField active
        CharField token
    }
    pretixbase_organizerfooterlink {
        BigAutoField id PK
        ForeignKey organizer_id FK
        I18nCharField label
        URLField url
    }
    pretixbase_saleschannel {
        BigAutoField id PK
        ForeignKey organizer_id FK
        I18nCharField label
        CharField identifier
        CharField type
        PositiveIntegerField position
        JSONField configuration
    }
    pretixbase_customerssoprovider {
        BigAutoField id PK
        ForeignKey organizer_id FK
        I18nCharField name
        BooleanField is_active
        I18nCharField button_label
        CustomerSSOProviderMethod method
        JSONField configuration
    }
    pretixbase_customer {
        BigAutoField id PK
        ForeignKey organizer_id FK
        ForeignKey provider_id FK
        CharField identifier "You can enter any value here to make it easier to match the data with other sources. If you do not input one, we will generate one automatically."
        EmailField email
        PhoneNumberField phone
        CharField password
        CharField name_cached
        JSONField name_parts
        BooleanField is_active
        BooleanField is_verified
        DateTimeField last_login
        DateTimeField date_joined
        CustomerLocale locale
        DateTimeField last_modified
        CharField external_identifier
        TextField notes
    }
    pretixbase_attendeeprofile {
        BigAutoField id PK
        ForeignKey customer_id FK
        CharField attendee_name_cached
        JSONField attendee_name_parts
        EmailField attendee_email
        CharField company
        TextField street
        CharField zipcode
        CharField city
        AttendeeProfileCountry country
        CharField state
        JSONField answers
    }
    pretixbase_customerssoclient {
        BigAutoField id PK
        ForeignKey organizer_id FK
        CharField name
        BooleanField is_active
        CharField client_id UK
        CharField client_secret
        CustomerSSOClientClientType client_type
        CustomerSSOClientAuthorizationGrantType authorization_grant_type
        BooleanField require_pkce
        TextField redirect_uris "Allowed URIs list, space separated"
        MultiStringField allowed_scopes "Separate multiple values with spaces"
    }
    pretixbase_customerssogrant {
        BigAutoField id PK
        ForeignKey client_id FK
        ForeignKey customer_id FK
        CharField code UK
        CharField nonce
        IntegerField auth_time
        DateTimeField expires
        TextField redirect_uri
        TextField scope
        TextField code_challenge
        CharField code_challenge_method
    }
    pretixbase_customerssoaccesstoken {
        BigAutoField id PK
        ForeignKey client_id FK
        ForeignKey customer_id FK
        CharField from_code
        CharField token UK
        DateTimeField expires
        TextField scope
    }
    pretixbase_event {
        BigAutoField id PK
        ForeignKey organizer_id FK
        BooleanField testmode
        I18nCharField name
        CharField slug "Should be short, only contain lowercase letters, numbers, dots, and dashes, and must be unique among your events. We recommend some kind of abbreviation or a date with less than 10 characters that can be easily remembered, but you can also choose to use a random value. This will be used in URLs, order codes, invoice numbers, and bank transfer references."
        BooleanField live
        EventCurrency currency
        DateTimeField date_from
        DateTimeField date_to
        DateTimeField date_admission
        BooleanField is_public "If selected, this event will show up publicly on the list of events for your organizer account."
        DateTimeField presale_end "Optional. No products will be sold after this date. If you do not set this value, the presale will end after the end date of your event."
        DateTimeField presale_start "Optional. No products will be sold before this date."
        I18nTextField location
        BooleanField is_remote "This will be used to let users know if the event is in a different timezone, and to let us calculate the local time of a user."
        FloatField geo_lat
        FloatField geo_lon
        TextField plugins
        TextField comment
        BooleanField has_subevents
        ForeignKey seating_plan_id FK
        DateTimeField last_modified
        BooleanField all_sales_channels
    }
    pretixbase_event_settingsstore {
        BigAutoField id PK
        CharField key
        TextField value
        ForeignKey object_id FK
    }
    pretixbase_subevent {
        BigAutoField id PK
        ForeignKey event_id FK
        BooleanField active "Only with this checkbox enabled, this date is visible in the frontend to users."
        BooleanField is_public "If selected, this event will show up publicly on the list of dates for your event."
        I18nCharField name
        DateTimeField date_from
        DateTimeField date_to
        DateTimeField date_admission
        DateTimeField presale_end "Optional. No products will be sold after this date. If you do not set this value, the presale will end after the end date of your event."
        DateTimeField presale_start "Optional. No products will be sold before this date."
        I18nTextField location
        FloatField geo_lat
        FloatField geo_lon
        I18nTextField frontpage_text
        ForeignKey seating_plan_id FK
        TextField comment
        DateTimeField last_modified
    }
    pretixbase_eventlock {
        CharField event PK
        DateTimeField date
        UUIDField token
    }
    pretixbase_eventmetaproperty {
        BigAutoField id PK
        ForeignKey organizer_id FK
        CharField name "Can not contain spaces or special characters except underscores"
        TextField default
        BooleanField protected
        BooleanField required "If checked, an event can only be taken live if the property is set. In event series, its always optional to set a value for individual dates"
        JSONField choices
        BooleanField filter_public "This field will be shown to filter events in the public event list and calendar."
        I18nCharField public_label
        BooleanField filter_allowed "This field will be shown to filter events or reports in the backend, and it can also be used for hidden filter parameters in the frontend (e.g. using the widget)."
        IntegerField position
    }
    pretixbase_eventmetavalue {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey property_id FK
        TextField value
    }
    pretixbase_subeventmetavalue {
        BigAutoField id PK
        ForeignKey subevent_id FK
        ForeignKey property_id FK
        TextField value
    }
    pretixbase_eventfooterlink {
        BigAutoField id PK
        ForeignKey event_id FK
        I18nCharField label
        URLField url
    }
    pretixbase_gate {
        BigAutoField id PK
        ForeignKey organizer_id FK
        CharField name
        CharField identifier "You can enter any value here to make it easier to match the data with other sources. If you do not input one, we will generate one automatically."
    }
    pretixbase_device {
        BigAutoField id PK
        ForeignKey organizer_id FK
        ForeignKey gate_id FK
        PositiveIntegerField device_id
        CharField unique_serial UK
        CharField initialization_token UK
        CharField api_token UK
        BooleanField all_events
        BooleanField revoked
        CharField name
        DateTimeField created
        DateTimeField initialized
        CharField hardware_brand
        CharField hardware_model
        CharField os_name
        CharField os_version
        CharField software_brand
        CharField software_version
        CharField security_profile
        TextField rsa_pubkey
        JSONField info
    }
    pretixbase_devicelastseen {
        BigAutoField id PK
        OneToOneField device_id FK, UK
        DateTimeField last_seen
    }
    pretixbase_discount {
        BigAutoField id PK
        ForeignKey event_id FK
        BooleanField active
        CharField internal_name
        PositiveIntegerField position
        BooleanField all_sales_channels
        DateTimeField available_from
        DateTimeField available_until
        DiscountSubeventMode subevent_mode
        BooleanField condition_all_products
        BooleanField condition_apply_to_addons "Discounts never apply to bundled products"
        BooleanField condition_ignore_voucher_discounted "If this option is checked, products that already received a discount through a voucher will not be considered for this discount. However, products that use a voucher only to e.g. unlock a hidden product or gain access to sold-out quota will still be considered."
        PositiveIntegerField condition_min_count
        DecimalField condition_min_value
        BooleanField benefit_same_products "By default, the discount is applied across the same selection of products than the condition for the discount given above. If you want, you can however also select a different selection of products."
        DecimalField benefit_discount_matching_percent
        PositiveIntegerField benefit_only_apply_to_cheapest_n_matches "This option allows you to create discounts of the type 'buy X get Y reduced/for free'. For example, if you set 'Minimum number of matching products' to four and this value to two, the customer's cart will be split into groups of four tickets and the cheapest two tickets within every group will be discounted. If you want to grant the discount on all matching products, keep this field empty."
        BooleanField benefit_apply_to_addons "Discounts never apply to bundled products"
        BooleanField benefit_ignore_voucher_discounted "If this option is checked, products that already received a discount through a voucher will not be discounted. However, products that use a voucher only to e.g. unlock a hidden product or gain access to sold-out quota will still receive the discount."
        DateTimeField subevent_date_from
        DateTimeField subevent_date_until
    }
    pretixbase_scheduledeventexport {
        BigAutoField id PK
        CharField export_identifier
        JSONField export_form_data
        ForeignKey owner_id FK
        CharField locale
        TextField mail_additional_recipients "You can specify multiple recipients separated by commas."
        TextField mail_additional_recipients_cc "You can specify multiple recipients separated by commas."
        TextField mail_additional_recipients_bcc "You can specify multiple recipients separated by commas."
        CharField mail_subject
        TextField mail_template
        TextField schedule_rrule
        TimeField schedule_rrule_time "The actual start time might be delayed depending on system load."
        DateTimeField schedule_next_run
        IntegerField error_counter
        TextField error_last_message
        ForeignKey event_id FK
    }
    pretixbase_scheduledorganizerexport {
        BigAutoField id PK
        CharField export_identifier
        JSONField export_form_data
        ForeignKey owner_id FK
        CharField locale
        TextField mail_additional_recipients "You can specify multiple recipients separated by commas."
        TextField mail_additional_recipients_cc "You can specify multiple recipients separated by commas."
        TextField mail_additional_recipients_bcc "You can specify multiple recipients separated by commas."
        CharField mail_subject
        TextField mail_template
        TextField schedule_rrule
        TimeField schedule_rrule_time "The actual start time might be delayed depending on system load."
        DateTimeField schedule_next_run
        IntegerField error_counter
        TextField error_last_message
        ForeignKey organizer_id FK
        CharField timezone
    }
    pretixbase_invoice {
        BigAutoField id PK
        ForeignKey order_id FK
        ForeignKey organizer_id FK
        ForeignKey event_id FK
        CharField prefix
        CharField invoice_no
        CharField full_invoice_no
        BooleanField is_cancellation
        ForeignKey refers_id FK
        TextField invoice_from
        CharField invoice_from_name
        CharField invoice_from_zipcode
        CharField invoice_from_city
        CharField invoice_from_state
        InvoiceInvoiceFromCountry invoice_from_country
        CharField invoice_from_tax_id
        CharField invoice_from_vat_id
        TextField invoice_to
        TextField invoice_to_company
        BooleanField invoice_to_is_business
        TextField invoice_to_name
        TextField invoice_to_street
        CharField invoice_to_zipcode
        TextField invoice_to_city
        CharField invoice_to_state
        InvoiceInvoiceToCountry invoice_to_country
        TextField invoice_to_vat_id
        TextField invoice_to_beneficiary
        JSONField invoice_to_transmission_info
        TextField internal_reference
        CharField custom_field
        DateTimeField created
        DateField date
        CharField locale
        TextField introductory_text
        TextField additional_text
        BooleanField reverse_charge
        TextField payment_provider_text
        CharField payment_provider_stamp
        TextField footer_text
        CharField foreign_currency_display
        DecimalField foreign_currency_rate
        DateField foreign_currency_rate_date
        CharField foreign_currency_source
        BooleanField shredded
        BooleanField sent_to_organizer
        CharField transmission_type
        CharField transmission_provider
        InvoiceTransmissionStatus transmission_status
        DateTimeField transmission_date
        JSONField transmission_info
        FileField file
        JSONField plugin_data
    }
    pretixbase_invoiceline {
        BigAutoField id PK
        ForeignKey invoice_id FK
        PositiveIntegerField position
        TextField description
        DecimalField gross_value
        DecimalField tax_value
        NormalizedDecimalField tax_rate
        CharField tax_name
        CharField tax_code
        ForeignKey subevent_id FK
        DateTimeField period_start
        DateTimeField period_end
        TextField event_location
        ForeignKey item_id FK
        ForeignKey variation_id FK
        TextField attendee_name
        CharField fee_type
        CharField fee_internal_type
    }
    pretixbase_itemcategory {
        BigAutoField id PK
        ForeignKey event_id FK
        I18nCharField name
        CharField internal_name "If you set this, this will be used instead of the public name in the backend."
        I18nTextField description
        IntegerField position
        BooleanField is_addon "If selected, the products belonging to this category are not for sale on their own. They can only be bought in combination with a product that has this category configured as a possible source for add-ons."
        ItemCategoryCrossSellingMode cross_selling_mode
        ItemCategoryCrossSellingCondition cross_selling_condition
    }
    pretixbase_subeventitem {
        BigAutoField id PK
        ForeignKey subevent_id FK
        ForeignKey item_id FK
        DecimalField price
        BooleanField disabled
        DateTimeField available_from "This product will not be sold before the given date."
        DateTimeField available_until "This product will not be sold after the given date."
    }
    pretixbase_subeventitemvariation {
        BigAutoField id PK
        ForeignKey subevent_id FK
        ForeignKey variation_id FK
        DecimalField price
        BooleanField disabled
        DateTimeField available_from "This product will not be sold before the given date."
        DateTimeField available_until "This product will not be sold after the given date."
    }
    pretixbase_item {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey category_id FK "If you have many products, you can optionally sort them into categories to keep things organized."
        I18nCharField name
        CharField internal_name "If you set this, this will be used instead of the public name in the backend."
        BooleanField active
        I18nTextField description "This is shown below the product name in lists."
        DecimalField default_price "If this product has multiple variations, you can set different prices for each of the variations. If a variation does not have a special price or if you do not have variations, this price will be used."
        BooleanField free_price "If this option is active, your users can choose the price themselves. The price configured above is then interpreted as the minimum price a user has to enter. You could use this e.g. to collect additional donations for your event."
        DecimalField free_price_suggestion "This price will be used as the default value of the input field. The user can choose a lower value, but not lower than the price this product would have without the free price option. This will be ignored if a voucher is used that lowers the price."
        ForeignKey tax_rule_id FK
        BooleanField admission "Whether or not buying this product allows a person to enter your event"
        BooleanField personalized "Whether or not buying this product allows to enter attendee information"
        BooleanField generate_tickets
        BooleanField allow_waitinglist "This will only work if waiting lists are enabled for this event."
        BooleanField show_quota_left "Publicly show how many tickets are still available."
        IntegerField position
        ImageField picture
        DateTimeField available_from "This product will not be sold before the given date."
        ItemAvailableFromMode available_from_mode
        DateTimeField available_until "This product will not be sold after the given date."
        ItemAvailableUntilMode available_until_mode
        ForeignKey hidden_if_available_id FK "If you select a quota here, this product will only be shown when that quota is unavailable. If combined with the option to hide sold-out products, this allows you to swap out products for more expensive ones once they are sold out. There might be a short period in which both products are visible while all tickets in the referenced quota are reserved, but not yet sold."
        ForeignKey hidden_if_item_available_id FK "If you select a product here, this product will only be shown when that product is no longer available. This will happen either because the other product has sold out or because the time is outside of the sales window for the other product. If combined with the option to hide sold-out products, this allows you to swap out products for more expensive ones once the cheaper option is sold out. There might be a short period in which both products are visible while all tickets of the referenced product are reserved, but not yet sold."
        ItemHiddenIfItemAvailableMode hidden_if_item_available_mode
        BooleanField require_voucher "To buy this product, the user needs a voucher that applies to this product either directly or via a quota."
        BooleanField require_approval "If this product is part of an order, the order will be put into an 'approval' state and will need to be confirmed by you before it can be paid and completed. You can use this e.g. for discounted tickets that are only available to specific groups."
        BooleanField hide_without_voucher "This product will be hidden from the event page until the user enters a voucher that unlocks this product."
        BooleanField require_bundling "If this option is set, the product will only be sold as part of bundle products. Do <strong>not</strong> check this option if you want to use this product as an add-on product, but only for fixed bundles!"
        BooleanField allow_cancel "If this is checked, the usual cancellation and order change settings of this event apply. If this is unchecked, orders containing this product can not be canceled by users but only by you."
        IntegerField min_per_order "This product can only be bought if it is added to the cart at least this many times. If you keep the field empty or set it to 0, there is no special limit for this product."
        IntegerField max_per_order "This product can only be bought at most this many times within one order. If you keep the field empty or set it to 0, there is no special limit for this product. The limit for the maximum number of items in the whole order applies regardless."
        BooleanField checkin_attention "If you set this, the check-in app will show a visible warning that this ticket requires special attention. You can use this for example for student tickets to indicate to the person at check-in that the student ID card still needs to be checked."
        TextField checkin_text "This text will be shown by the check-in app if a ticket of this type is scanned."
        DecimalField original_price "If set, this will be displayed next to the current price to show that the current price is a discounted one. This is just a cosmetic setting and will not actually impact pricing."
        BooleanField all_sales_channels
        BooleanField issue_giftcard "When a customer buys this product, they will get a gift card with a value corresponding to the product price."
        BooleanField require_membership
        BooleanField require_membership_hidden "Do not show this unless the customer is logged in and has a valid membership. Be aware that this means it will never be visible in the widget."
        ForeignKey grant_membership_type_id FK
        BooleanField grant_membership_duration_like_event
        IntegerField grant_membership_duration_days
        IntegerField grant_membership_duration_months
        ItemValidityMode validity_mode "When setting up a regular event, or an event series with time slots, you typically do NOT need to change this value. The default setting means that the validity time of tickets will not be decided by the product, but by the event and check-in configuration. Only use the other options if you need them to realize e.g. a booking of a year-long ticket with a dynamic start date. Note that the validity will be stored with the ticket, so if you change the settings here later, existing tickets will not be affected by the change but keep their current validity."
        DateTimeField validity_fixed_from
        DateTimeField validity_fixed_until
        PositiveIntegerField validity_dynamic_duration_minutes
        PositiveIntegerField validity_dynamic_duration_hours
        PositiveIntegerField validity_dynamic_duration_days
        PositiveIntegerField validity_dynamic_duration_months
        BooleanField validity_dynamic_start_choice "If not selected, the validity always starts at the time of purchase."
        PositiveIntegerField validity_dynamic_start_choice_day_limit "The selected start date may only be this many days in the future."
        ItemMediaPolicy media_policy "If this product should be stored on a reusable physical medium, you can attach a physical media policy. This is not required for regular tickets, which just use a one-time barcode, but only for products like renewable season tickets or re-chargeable gift card wristbands. This is an advanced feature that also requires specific configuration of ticketing and printing settings."
        ItemMediaType media_type "Select the type of physical medium that should be used for this product. Note that not all media types support all types of products, and not all media types are supported across all sales channels or check-in processes."
    }
    pretixbase_itemvariation {
        BigAutoField id PK
        ForeignKey item_id FK
        I18nCharField value
        BooleanField active
        I18nTextField description "This is shown below the variation name in lists."
        PositiveIntegerField position
        DecimalField default_price
        DecimalField original_price "If set, this will be displayed next to the current price to show that the current price is a discounted one. This is just a cosmetic setting and will not actually impact pricing."
        DecimalField free_price_suggestion "This price will be used as the default value of the input field. The user can choose a lower value, but not lower than the price this product would have without the free price option. This will be ignored if a voucher is used that lowers the price."
        BooleanField require_approval "If this variation is part of an order, the order will be put into an 'approval' state and will need to be confirmed by you before it can be paid and completed. You can use this e.g. for discounted tickets that are only available to specific groups."
        BooleanField require_membership
        BooleanField require_membership_hidden "Do not show this unless the customer is logged in and has a valid membership. Be aware that this means it will never be visible in the widget."
        DateTimeField available_from "This variation will not be sold before the given date."
        ItemVariationAvailableFromMode available_from_mode
        DateTimeField available_until "This variation will not be sold after the given date."
        ItemVariationAvailableUntilMode available_until_mode
        BooleanField all_sales_channels
        BooleanField hide_without_voucher "This variation will be hidden from the event page until the user enters a voucher that unlocks this variation."
        BooleanField checkin_attention "If you set this, the check-in app will show a visible warning that this ticket requires special attention. You can use this for example for student tickets to indicate to the person at check-in that the student ID card still needs to be checked."
        TextField checkin_text "This text will be shown by the check-in app if a ticket of this type is scanned."
    }
    pretixbase_itemaddon {
        BigAutoField id PK
        ForeignKey base_item_id FK
        ForeignKey addon_category_id FK
        PositiveIntegerField min_count
        PositiveIntegerField max_count
        BooleanField price_included "If selected, adding add-ons to this ticket is free, even if the add-ons would normally cost money individually."
        BooleanField multi_allowed
        PositiveIntegerField position
    }
    pretixbase_itembundle {
        BigAutoField id PK
        ForeignKey base_item_id FK
        ForeignKey bundled_item_id FK
        ForeignKey bundled_variation_id FK
        PositiveIntegerField count
        DecimalField designated_price "If set, it will be shown that this bundled item is responsible for the given value of the total gross price. This might be important in cases of mixed taxation, but can be kept blank otherwise. This value will NOT be added to the base item's price."
    }
    pretixbase_question {
        BigAutoField id PK
        ForeignKey event_id FK
        QuestionContainerType container_type
        I18nTextField question
        CharField identifier "You can enter any value here to make it easier to match the data with other sources. If you do not input one, we will generate one automatically."
        I18nTextField help_text "If the question needs to be explained or clarified, do it here!"
        QuestionType type
        BooleanField required
        PositiveIntegerField position
        BooleanField ask_during_checkin "Not supported by all check-in apps for all question types."
        BooleanField show_during_checkin "Not supported by all check-in apps for all question types."
        BooleanField hidden "This question will only show up in the backend."
        BooleanField print_on_invoice
        ForeignKey dependency_question_id FK
        MultiStringField dependency_values
        DecimalField valid_number_min "Currently not supported in our apps and during check-in"
        DecimalField valid_number_max "Currently not supported in our apps and during check-in"
        DateField valid_date_min "Currently not supported in our apps and during check-in"
        DateField valid_date_max "Currently not supported in our apps and during check-in"
        DateTimeField valid_datetime_min "Currently not supported in our apps and during check-in"
        DateTimeField valid_datetime_max "Currently not supported in our apps and during check-in"
        PositiveIntegerField valid_string_length_min "Currently not supported in our apps and during check-in"
        PositiveIntegerField valid_string_length_max "Currently not supported in our apps and during check-in"
        BooleanField valid_file_portrait "If checked, files must be images with an aspect ratio of 3:4. This is commonly used for photos printed on badges."
    }
    pretixbase_questionoption {
        BigAutoField id PK
        ForeignKey question_id FK
        CharField identifier "You can enter any value here to make it easier to match the data with other sources. If you do not input one, we will generate one automatically."
        I18nCharField answer
        IntegerField position
    }
    pretixbase_quota {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey subevent_id FK
        CharField name
        PositiveIntegerField size "Leave empty for an unlimited number of tickets."
        BooleanField ignore_for_event_availability "If you enable this, this quota will be ignored when determining event availability in your event calendar. This is useful e.g. for merchandise that is added to each event but should not stop the event from being shown as sold out."
        BooleanField close_when_sold_out "If you enable this, when the quota is sold out once, no more tickets will be sold, even if tickets become available again through cancellations or expiring orders. Of course, you can always re-open it manually."
        BooleanField closed
        BooleanField release_after_exit "With this option, quota will be released as soon as people are scanned at an exit of your event. This will only happen if they have been scanned both at an entry and at an exit and the exit is the more recent scan. It does not matter which check-in list either of the scans was on, but check-in lists are ignored if they are set to 'Allow re-entering after an exit scan' to prevent accidental overbooking."
    }
    pretixbase_itemmetaproperty {
        BigAutoField id PK
        ForeignKey event_id FK
        CharField name "Can not contain spaces or special characters except underscores"
        TextField default
        BooleanField required "If checked, this property must be set in each product. Does not apply if a default value is set."
        TextField allowed_values "If you keep this empty, any value is allowed. Otherwise, enter one possible value per line."
    }
    pretixbase_itemmetavalue {
        BigAutoField id PK
        ForeignKey item_id FK
        ForeignKey property_id FK
        TextField value
    }
    pretixbase_itemvariationmetavalue {
        BigAutoField id PK
        ForeignKey variation_id FK
        ForeignKey property_id FK
        TextField value
    }
    pretixbase_itemprogramtime {
        BigAutoField id PK
        ForeignKey item_id FK
        DateTimeField start
        DateTimeField end
        I18nTextField location
    }
    pretixbase_logentry {
        BigAutoField id PK
        ForeignKey content_type_id FK
        PositiveBigIntegerField object_id
        DateTimeField datetime
        ForeignKey user_id FK
        ForeignKey api_token_id FK
        ForeignKey device_id FK
        ForeignKey oauth_application_id FK
        ForeignKey event_id FK
        ForeignKey organizer_link_id FK
        CharField action_type
        TextField data
        BooleanField visible
        BooleanField shredded
    }
    pretixbase_outgoingmail {
        BigAutoField id PK
        UUIDField guid
        OutgoingMailStatus status
        DateTimeField created
        DateTimeField sent
        DateTimeField inflight_since
        DateTimeField retry_after
        TextField error
        TextField error_detail
        ForeignKey organizer_id FK
        ForeignKey event_id FK
        ForeignKey order_id FK
        ForeignKey orderposition_id FK
        ForeignKey customer_id FK
        ForeignKey user_id FK
        BooleanField sensitive
        TextField subject
        TextField body_plain
        TextField body_html
        CharField sender
        JSONField headers
        JSONField to
        JSONField cc
        JSONField bcc
        IntegerField recipient_count
        BooleanField should_attach_tickets
        BooleanField should_attach_ical
        JSONField should_attach_other_files
        JSONField actual_attachments
    }
    pretixbase_order {
        BigAutoField id PK
        CharField code
        OrderStatus status
        BooleanField valid_if_pending
        BooleanField testmode
        ForeignKey organizer_id FK
        ForeignKey event_id FK
        ForeignKey customer_id FK
        EmailField email
        PhoneNumberField phone
        CharField locale
        CharField secret
        CharField internal_secret
        DateTimeField datetime
        DateTimeField cancellation_date
        DateTimeField expires
        DecimalField total
        TextField comment "The text entered in this field will not be visible to the user and is available for your convenience."
        DateField custom_followup_at "We'll show you this order to be due for a follow-up on this day."
        BooleanField checkin_attention "If you set this, the check-in app will show a visible warning that tickets of this order require special attention. This will not show any details or custom message, so you need to brief your check-in staff how to handle these cases."
        TextField checkin_text "This text will be shown by the check-in app if a ticket of this order is scanned."
        BooleanField expiry_reminder_sent
        BooleanField download_reminder_sent
        TextField meta_info
        JSONField api_meta
        DateTimeField last_modified
        BooleanField require_approval
        ForeignKey sales_channel_id FK
        BooleanField email_known_to_work
        BooleanField invoice_dirty
        OrderTaxRoundingMode tax_rounding_mode
    }
    pretixbase_questionanswer {
        BigAutoField id PK
        ForeignKey orderposition_id FK
        ForeignKey cartposition_id FK
        ForeignKey order_id FK
        ForeignKey checkoutsession_id FK
        ForeignKey question_id FK
        TextField answer
        FileField file
    }
    pretixbase_orderpayment {
        BigAutoField id PK
        PositiveIntegerField local_id
        OrderPaymentState state
        DecimalField amount
        ForeignKey order_id FK
        DateTimeField created
        DateTimeField payment_date
        CharField provider
        TextField info
        ForeignKey fee_id FK
        BooleanField migrated
        BooleanField process_initiated
    }
    pretixbase_orderrefund {
        BigAutoField id PK
        PositiveIntegerField local_id
        OrderRefundState state
        OrderRefundSource source
        DecimalField amount
        ForeignKey order_id FK
        ForeignKey payment_id FK
        DateTimeField created
        DateTimeField execution_date
        CharField provider
        TextField comment "May be shown to the end user or used e.g. as part of a payment reference."
        TextField info
    }
    pretixbase_orderfee {
        BigAutoField id PK
        DecimalField value
        DecimalField value_includes_rounding_correction
        ForeignKey order_id FK
        OrderFeeFeeType fee_type
        CharField description
        CharField internal_type
        NormalizedDecimalField tax_rate
        ForeignKey tax_rule_id FK
        CharField tax_code
        DecimalField tax_value
        DecimalField tax_value_includes_rounding_correction
        BooleanField canceled
    }
    pretixbase_orderposition {
        BigAutoField id PK
        ForeignKey subevent_id FK
        ForeignKey item_id FK
        ForeignKey variation_id FK
        DecimalField price
        DecimalField price_includes_rounding_correction
        CharField attendee_name_cached "Empty, if this product is not an admission ticket"
        JSONField attendee_name_parts
        EmailField attendee_email "Empty, if this product is not an admission ticket"
        ForeignKey voucher_id FK
        ForeignKey used_membership_id FK
        ForeignKey addon_to_id FK
        TextField meta_info
        ForeignKey seat_id FK
        BooleanField is_bundled
        ForeignKey discount_id FK
        CharField company
        TextField street
        CharField zipcode
        CharField city
        OrderPositionCountry country
        CharField state
        PositiveIntegerField positionid
        ForeignKey organizer_id FK
        ForeignKey order_id FK
        DecimalField voucher_budget_use
        NormalizedDecimalField tax_rate
        ForeignKey tax_rule_id FK
        CharField tax_code
        DecimalField tax_value
        DecimalField tax_value_includes_rounding_correction
        CharField secret
        CharField web_secret
        CharField pseudonymization_id UK
        BooleanField canceled
        JSONField blocked
        BooleanField ignore_from_quota_while_blocked
        DateTimeField valid_from
        DateTimeField valid_until
    }
    pretixbase_transaction {
        BigAutoField id PK
        ForeignKey order_id FK
        DateTimeField created
        DateTimeField datetime
        BooleanField migrated
        PositiveIntegerField positionid
        IntegerField count
        ForeignKey item_id FK
        ForeignKey variation_id FK
        ForeignKey subevent_id FK
        DecimalField price
        DecimalField price_includes_rounding_correction
        NormalizedDecimalField tax_rate
        ForeignKey tax_rule_id FK
        CharField tax_code
        DecimalField tax_value
        DecimalField tax_value_includes_rounding_correction
        TransactionFeeType fee_type
        CharField internal_type
    }
    pretixbase_checkoutsession {
        BigAutoField id PK
        ForeignKey event_id FK
        CharField cart_id UK
        DateTimeField created
        ForeignKey customer_id FK
        ForeignKey sales_channel_id FK
        BooleanField testmode
        JSONField session_data
    }
    pretixbase_cartposition {
        BigAutoField id PK
        ForeignKey subevent_id FK
        ForeignKey item_id FK
        ForeignKey variation_id FK
        DecimalField price
        DecimalField price_includes_rounding_correction
        CharField attendee_name_cached "Empty, if this product is not an admission ticket"
        JSONField attendee_name_parts
        EmailField attendee_email "Empty, if this product is not an admission ticket"
        ForeignKey voucher_id FK
        ForeignKey used_membership_id FK
        ForeignKey addon_to_id FK
        TextField meta_info
        ForeignKey seat_id FK
        BooleanField is_bundled
        ForeignKey discount_id FK
        CharField company
        TextField street
        CharField zipcode
        CharField city
        CartPositionCountry country
        CharField state
        ForeignKey event_id FK
        CharField cart_id
        DateTimeField datetime
        DateTimeField expires
        DateTimeField max_extend
        NormalizedDecimalField tax_rate
        CharField tax_code
        DecimalField tax_value_includes_rounding_correction
        DecimalField listed_price
        DecimalField price_after_voucher
        DecimalField custom_price_input
        BooleanField custom_price_input_is_net
        DecimalField line_price_gross
        DateTimeField requested_valid_from
    }
    pretixbase_invoiceaddress {
        BigAutoField id PK
        DateTimeField last_modified
        OneToOneField checkout_session_id FK, UK
        OneToOneField order_id FK, UK
        ForeignKey customer_id FK
        BooleanField is_business
        CharField company
        CharField name_cached
        JSONField name_parts
        TextField street
        CharField zipcode
        CharField city
        CharField country_old
        InvoiceAddressCountry country
        CharField state
        CharField vat_id
        BooleanField vat_id_validated
        CharField custom_field
        TextField internal_reference "This reference will be printed on your invoice for your convenience."
        TextField beneficiary
        CharField transmission_type
        JSONField transmission_info
    }
    pretixbase_cachedticket {
        BigAutoField id PK
        ForeignKey order_position_id FK
        CharField provider
        CharField type
        CharField extension
        FileField file
        DateTimeField created
    }
    pretixbase_cachedcombinedticket {
        BigAutoField id PK
        ForeignKey order_id FK
        CharField provider
        CharField type
        CharField extension
        FileField file
        DateTimeField created
    }
    pretixbase_cancellationrequest {
        BigAutoField id PK
        ForeignKey order_id FK
        DateTimeField created
        DecimalField cancellation_fee
        BooleanField refund_as_giftcard
    }
    pretixbase_revokedticketsecret {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey position_id FK
        TextField secret
        DateTimeField created
    }
    pretixbase_blockedticketsecret {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey position_id FK
        TextField secret
        BooleanField blocked
        DateTimeField updated
    }
    pretixbase_printlog {
        BigAutoField id PK
        ForeignKey position_id FK
        BooleanField successful
        DateTimeField datetime
        DateTimeField created
        ForeignKey device_id FK
        ForeignKey user_id FK
        ForeignKey api_token_id FK
        ForeignKey oauth_application_id FK
        CharField source
        PrintLogType type
        JSONField info
    }
    pretixbase_reusablemedium {
        BigAutoField id PK
        DateTimeField created
        DateTimeField updated
        ForeignKey organizer_id FK
        ReusableMediumType type
        CharField identifier
        CharField claim_token
        CharField label
        BooleanField active
        DateTimeField expires
        ForeignKey customer_id FK
        ForeignKey linked_giftcard_id FK
        JSONField info
        TextField notes
    }
    pretixbase_mediumkeyset {
        BigAutoField id PK
        ForeignKey organizer_id FK
        BigIntegerField public_id UK
        CharField media_type
        BooleanField active
        BinaryField uid_key
        BinaryField diversification_key
    }
    pretixbase_membershiptype {
        BigAutoField id PK
        ForeignKey organizer_id FK
        I18nCharField name
        BooleanField transferable "If this is selected, the membership can be used to purchase tickets for multiple persons. If not, the attendee name always needs to stay the same."
        BooleanField allow_parallel_usage "If this is selected, the membership can be used to purchase tickets for events happening at the same time. Note that this will only check for an identical start time of the events, not for any overlap between events. An overlap check will be performed if there is a product-level validity of the ticket."
        PositiveIntegerField max_usages "Number of times this membership can be used in a purchase."
    }
    pretixbase_membership {
        BigAutoField id PK
        BooleanField testmode
        BooleanField canceled
        ForeignKey customer_id FK
        ForeignKey membership_type_id FK
        ForeignKey granted_in_id FK
        DateTimeField date_start
        DateTimeField date_end
        JSONField attendee_name_parts
    }
    pretixbase_notificationsetting {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField action_type
        ForeignKey event_id FK
        NotificationSettingMethod method
        BooleanField enabled
    }
    pretixbase_seatingplan {
        BigAutoField id PK
        CharField name
        ForeignKey organizer_id FK
        TextField layout
    }
    pretixbase_seatcategorymapping {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey subevent_id FK
        CharField layout_category
        ForeignKey product_id FK
    }
    pretixbase_seat {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey subevent_id FK
        CharField zone_name
        CharField row_name
        CharField row_label
        CharField seat_number
        CharField seat_label
        CharField seat_guid
        ForeignKey product_id FK
        BooleanField blocked
        BigIntegerField sorting_rank
        FloatField x
        FloatField y
    }
    pretixbase_voucher {
        BigAutoField id PK
        DateTimeField created
        ForeignKey event_id FK
        ForeignKey subevent_id FK
        CharField code
        PositiveIntegerField max_usages "Number of times this voucher can be redeemed."
        PositiveIntegerField redeemed
        PositiveIntegerField min_usages "If set to more than one, the voucher must be redeemed for this many products when it is used for the first time. On later usages, it can also be used for lower numbers of products. Note that this means that the total number of usages in some cases can be lower than this limit, e.g. in case of cancellations."
        DecimalField budget "This is the maximum monetary amount that will be discounted using this voucher across all usages. If this is sum reached, the voucher can no longer be used."
        DateTimeField valid_until
        BooleanField block_quota "If activated, this voucher will be substracted from the affected product's quotas, such that it is guaranteed that anyone with this voucher code does receive a ticket."
        BooleanField allow_ignore_quota "If activated, a holder of this voucher code can buy tickets, even if there are none left."
        VoucherPriceMode price_mode
        DecimalField value
        ForeignKey item_id FK "This product is added to the user's cart if the voucher is redeemed. Instead of a specific product, you can also select a quota. In this case, all products assigned to this quota can be selected."
        ForeignKey variation_id FK "This variation of the product select above is being used."
        ForeignKey quota_id FK "If enabled, the voucher is valid for any product affected by this quota."
        ForeignKey seat_id FK
        CharField tag "You can use this field to group multiple vouchers together. If you enter the same value for multiple vouchers, you can get statistics on how many of them have been redeemed etc."
        TextField comment "The text entered in this field will not be visible to the user and is available for your convenience."
        BooleanField show_hidden_items
        BooleanField all_addons_included
        BooleanField all_bundles_included
    }
    pretixbase_waitinglistentry {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey subevent_id FK
        DateTimeField created
        CharField name_cached
        JSONField name_parts
        EmailField email
        PhoneNumberField phone
        ForeignKey voucher_id FK
        ForeignKey item_id FK "The product the user waits for."
        ForeignKey variation_id FK "The variation of the product selected above."
        CharField locale
        IntegerField priority
    }
    pretixbase_ordersyncqueue {
        BigAutoField id PK
        ForeignKey order_id FK
        ForeignKey event_id FK
        CharField sync_provider
        CharField triggered_by
        DateTimeField triggered
        PositiveIntegerField failed_attempts
        DateTimeField not_before
        OrderSyncQueueNeedManualRetry need_manual_retry
        BooleanField in_flight
        DateTimeField in_flight_since
    }
    pretixbase_ordersyncresult {
        BigAutoField id PK
        ForeignKey order_id FK
        CharField sync_provider
        ForeignKey order_position_id FK
        IntegerField mapping_id
        CharField external_object_type
        CharField external_id_field
        CharField id_value
        CharField external_link_href
        CharField external_link_display_name
        DateTimeField transmitted
        JSONField sync_info
    }
    pretixmultidomain_knowndomain {
        CharField domainname PK
        KnownDomainMode mode
        ForeignKey organizer_id FK
        OneToOneField event_id FK, UK
    }
    pretixmultidomain_alternativedomainassignment {
        BigAutoField id PK
        ForeignKey domain_id FK
        OneToOneField event_id FK, UK
    }
    pretixapi_oauthapplication {
        BigAutoField id PK
        ForeignKey user_id FK
        OAuthApplicationClientType client_type
        OAuthApplicationAuthorizationGrantType authorization_grant_type
        BooleanField skip_authorization
        DateTimeField created
        DateTimeField updated
        OAuthApplicationAlgorithm algorithm
        CharField name
        TextField redirect_uris "Allowed URIs list, space separated"
        TextField post_logout_redirect_uris "Allowed Post Logout URIs list, space separated"
        CharField client_id UK
        ClientSecretField client_secret
        BooleanField active
    }
    pretixapi_oauthgrant {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField code UK
        DateTimeField expires
        TextField scope
        DateTimeField created
        DateTimeField updated
        CharField code_challenge
        OAuthGrantCodeChallengeMethod code_challenge_method
        CharField nonce
        TextField claims
        ForeignKey application_id FK
        CharField redirect_uri
    }
    pretixapi_oauthidtoken {
        BigAutoField id PK
        ForeignKey user_id FK
        UUIDField jti UK
        DateTimeField expires
        TextField scope
        DateTimeField created
        DateTimeField updated
        ForeignKey application_id FK
    }
    pretixapi_oauthaccesstoken {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField token UK
        DateTimeField expires
        TextField scope
        DateTimeField created
        DateTimeField updated
        OneToOneField source_refresh_token_id FK, UK
        OneToOneField id_token_id FK, UK
        ForeignKey application_id FK
    }
    pretixapi_oauthrefreshtoken {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField token
        DateTimeField created
        DateTimeField updated
        DateTimeField revoked
        ForeignKey application_id FK
        OneToOneField access_token_id FK, UK
    }
    pretixapi_webhook {
        BigAutoField id PK
        ForeignKey organizer_id FK
        BooleanField enabled
        URLField target_url
        BooleanField all_events
        CharField comment
    }
    pretixapi_webhookeventlistener {
        BigAutoField id PK
        ForeignKey webhook_id FK
        CharField action_type
    }
    pretixapi_webhookcall {
        BigAutoField id PK
        ForeignKey webhook_id FK
        DateTimeField datetime
        URLField target_url
        CharField action_type
        BooleanField is_retry
        FloatField execution_time
        PositiveIntegerField return_code
        BooleanField success
        TextField payload
        TextField response_body
    }
    pretixapi_webhookcallretry {
        BigAutoField id PK
        ForeignKey webhook_id FK
        ForeignKey logentry_id FK
        DateTimeField retry_not_before
        PositiveIntegerField retry_count
        CharField action_type
    }
    pretixapi_apicall {
        BigAutoField id PK
        CharField idempotency_key
        CharField auth_hash
        DateTimeField created
        DateTimeField locked
        CharField request_method
        CharField request_path
        PositiveIntegerField response_code
        TextField response_headers
        BinaryField response_body
    }
    pretixhelpers_thumbnail {
        BigAutoField id PK
        CharField source
        CharField size
        FileField thumb
        DateTimeField created
    }
    banktransfer_bankimportjob {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey organizer_id FK
        CharField currency
        DateTimeField created
        BankImportJobState state
    }
    banktransfer_banktransaction {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey organizer_id FK
        ForeignKey import_job_id FK
        CharField currency
        BankTransactionState state
        TextField message
        CharField external_id
        CharField checksum
        TextField payer
        TextField reference
        DecimalField amount
        CharField date
        DateField date_parsed
        CharField iban
        CharField bic
        ForeignKey order_id FK
        TextField comment
    }
    banktransfer_refundexport {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey organizer_id FK
        CharField currency
        DateTimeField datetime
        BooleanField testmode
        TextField rows
        BooleanField downloaded
    }
    stripe_referencedstripeobject {
        BigAutoField id PK
        CharField reference UK
        ForeignKey order_id FK
        ForeignKey payment_id FK
    }
    stripe_registeredapplepaydomain {
        BigAutoField id PK
        CharField domain
        CharField account
    }
    paypal_referencedpaypalobject {
        BigAutoField id PK
        CharField reference UK
        ForeignKey order_id FK
        ForeignKey payment_id FK
    }
    ticketoutputpdf_ticketlayout {
        BigAutoField id PK
        ForeignKey event_id FK
        BooleanField default
        CharField name
        TextField layout
        FileField background
    }
    ticketoutputpdf_ticketlayoutitem {
        BigAutoField id PK
        ForeignKey item_id FK
        ForeignKey layout_id FK
        ForeignKey sales_channel_id FK
    }
    sendmail_scheduledmail {
        BigAutoField id PK
        ForeignKey rule_id FK
        ForeignKey subevent_id FK
        ForeignKey event_id FK
        DateTimeField last_computed
        DateTimeField computed_datetime
        ScheduledMailState state
        BigIntegerField last_successful_order_id
    }
    sendmail_rule {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey subevent_id FK
        I18nCharField subject
        I18nTextField template
        BooleanField all_products
        MultiStringField restrict_to_status
        RuleCheckedInStatus checked_in_status
        BooleanField attach_ical
        DateTimeField send_date
        IntegerField send_offset_days
        TimeField send_offset_time
        BooleanField date_is_absolute
        BooleanField offset_to_event_end
        BooleanField offset_is_after
        RuleSendTo send_to
        BooleanField enabled "Only enabled rules are actually sent"
    }
    badges_badgelayout {
        BigAutoField id PK
        ForeignKey event_id FK
        BooleanField default
        CharField name
        TextField layout
        FileField background
    }
    badges_badgeitem {
        BigAutoField id PK
        OneToOneField item_id FK, UK
        ForeignKey layout_id FK
    }
    autocheckin_autocheckinrule {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey list_id FK "If you keep this empty, all lists that match the purchased product will be used."
        AutoCheckinRuleMode mode
        BooleanField all_sales_channels
        BooleanField all_products
        BooleanField all_payment_methods
        MultiStringField limit_payment_methods
    }
    otp_totp_totpdevice {
        AutoField id PK
        ForeignKey user_id FK "The user that this device belongs to."
        CharField name "The human-readable name of this device."
        BooleanField confirmed "Is this device ready for use?"
        DateTimeField throttling_failure_timestamp "A timestamp of the last failed verification attempt. Null if last attempt succeeded."
        PositiveIntegerField throttling_failure_count "Number of successive failed attempts."
        DateTimeField created_at "The date and time when this device was initially created in the system."
        DateTimeField last_used_at "The most recent date and time this device was used."
        CharField key "A hex-encoded secret key of up to 40 bytes."
        PositiveSmallIntegerField step "The time step in seconds."
        BigIntegerField t0 "The Unix time at which to begin counting steps."
        TOTPDeviceDigits digits "The number of digits to expect in a token."
        PositiveSmallIntegerField tolerance "The number of time steps in the past or future to allow."
        SmallIntegerField drift "The number of time steps the prover is known to deviate from our clock."
        BigIntegerField last_t "The t value of the latest verified token. The next token must be at a higher time step."
    }
    otp_static_staticdevice {
        AutoField id PK
        ForeignKey user_id FK "The user that this device belongs to."
        CharField name "The human-readable name of this device."
        BooleanField confirmed "Is this device ready for use?"
        DateTimeField throttling_failure_timestamp "A timestamp of the last failed verification attempt. Null if last attempt succeeded."
        PositiveIntegerField throttling_failure_count "Number of successive failed attempts."
        DateTimeField created_at "The date and time when this device was initially created in the system."
        DateTimeField last_used_at "The most recent date and time this device was used."
    }
    otp_static_statictoken {
        AutoField id PK
        ForeignKey device_id FK
        CharField token
    }
```
