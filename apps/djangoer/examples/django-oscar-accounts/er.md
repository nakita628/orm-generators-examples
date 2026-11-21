```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    django_flatpage }o--o{ django_site : "(sites) - (flatpage_set)"
    catalogue_product ||--o| analytics_productrecord : "(id) - (product_id)"
    auth_user ||--o| analytics_userrecord : "(id) - (user_id)"
    auth_user ||--o{ analytics_userproductview : "(id) - (user_id)"
    catalogue_product ||--o{ analytics_userproductview : "(id) - (product_id)"
    auth_user ||--o{ analytics_usersearch : "(id) - (user_id)"
    address_country ||--o{ address_useraddress : "(iso_3166_1_a2) - (country_id)"
    auth_user ||--o{ address_useraddress : "(id) - (user_id)"
    shipping_orderanditemcharges }o--o{ address_country : "(countries) - (orderanditemcharges_set)"
    shipping_weightbased }o--o{ address_country : "(countries) - (weightbased_set)"
    shipping_weightbased ||--o{ shipping_weightband : "(id) - (method_id)"
    catalogue_productclass }o--o{ catalogue_option : "(options) - (productclass_set)"
    catalogue_product ||--o{ catalogue_productcategory : "(id) - (product_id)"
    catalogue_category ||--o{ catalogue_productcategory : "(id) - (category_id)"
    catalogue_product |o--o{ catalogue_product : "(id) - (parent_id)"
    catalogue_productclass |o--o{ catalogue_product : "(id) - (product_class_id)"
    catalogue_product }o--o{ catalogue_option : "(product_options) - (product_set)"
    catalogue_product ||--o{ catalogue_productrecommendation : "(id) - (primary_id)"
    catalogue_product ||--o{ catalogue_productrecommendation : "(id) - (recommendation_id)"
    catalogue_productclass |o--o{ catalogue_productattribute : "(id) - (product_class_id)"
    catalogue_attributeoptiongroup |o--o{ catalogue_productattribute : "(id) - (option_group_id)"
    catalogue_productattribute ||--o{ catalogue_productattributevalue : "(id) - (attribute_id)"
    catalogue_product ||--o{ catalogue_productattributevalue : "(id) - (product_id)"
    catalogue_attributeoption |o--o{ catalogue_productattributevalue : "(id) - (value_option_id)"
    django_content_type |o--o{ catalogue_productattributevalue : "(id) - (entity_content_type_id)"
    catalogue_productattributevalue }o--o{ catalogue_attributeoption : "(value_multi_option) - (multi_valued_attribute_values)"
    catalogue_attributeoptiongroup ||--o{ catalogue_attributeoption : "(id) - (group_id)"
    catalogue_attributeoptiongroup |o--o{ catalogue_option : "(id) - (option_group_id)"
    catalogue_product ||--o{ catalogue_productimage : "(id) - (product_id)"
    catalogue_product |o--o{ reviews_productreview : "(id) - (product_id)"
    auth_user |o--o{ reviews_productreview : "(id) - (user_id)"
    reviews_productreview ||--o{ reviews_vote : "(id) - (review_id)"
    auth_user ||--o{ reviews_vote : "(id) - (user_id)"
    auth_user |o--o{ communication_email : "(id) - (user_id)"
    auth_user ||--o{ communication_notification : "(id) - (recipient_id)"
    auth_user |o--o{ communication_notification : "(id) - (sender_id)"
    partner_partner }o--o{ auth_user : "(users) - (partners)"
    address_country ||--o{ partner_partneraddress : "(iso_3166_1_a2) - (country_id)"
    partner_partner ||--o{ partner_partneraddress : "(id) - (partner_id)"
    catalogue_product ||--o{ partner_stockrecord : "(id) - (product_id)"
    partner_partner ||--o{ partner_stockrecord : "(id) - (partner_id)"
    partner_stockrecord ||--o{ partner_stockalert : "(id) - (stockrecord_id)"
    auth_user |o--o{ basket_basket : "(id) - (owner_id)"
    basket_basket }o--o{ voucher_voucher : "(vouchers) - (basket_set)"
    basket_basket ||--o{ basket_line : "(id) - (basket_id)"
    catalogue_product ||--o{ basket_line : "(id) - (product_id)"
    partner_stockrecord ||--o{ basket_line : "(id) - (stockrecord_id)"
    basket_line ||--o{ basket_lineattribute : "(id) - (line_id)"
    catalogue_option ||--o{ basket_lineattribute : "(id) - (option_id)"
    payment_source ||--o{ payment_transaction : "(id) - (source_id)"
    order_order ||--o{ payment_source : "(id) - (order_id)"
    payment_sourcetype ||--o{ payment_source : "(id) - (source_type_id)"
    auth_user ||--o{ payment_bankcard : "(id) - (user_id)"
    offer_condition ||--o{ offer_conditionaloffer : "(id) - (condition_id)"
    offer_benefit ||--o{ offer_conditionaloffer : "(id) - (benefit_id)"
    offer_conditionaloffer }o--o{ offer_conditionaloffer : "(combinations) - (in_combination)"
    offer_range |o--o{ offer_benefit : "(id) - (range_id)"
    offer_range |o--o{ offer_condition : "(id) - (range_id)"
    offer_range }o--o{ catalogue_product : "(excluded_products) - (excludes)"
    offer_range }o--o{ catalogue_productclass : "(classes) - (classes)"
    offer_range }o--o{ catalogue_category : "(included_categories) - (includes)"
    offer_range }o--o{ catalogue_category : "(excluded_categories) - (excludes)"
    offer_range ||--o{ offer_rangeproduct : "(id) - (range_id)"
    catalogue_product ||--o{ offer_rangeproduct : "(id) - (product_id)"
    offer_range ||--o{ offer_rangeproductfileupload : "(id) - (range_id)"
    auth_user ||--o{ offer_rangeproductfileupload : "(id) - (uploaded_by_id)"
    order_paymentevent ||--o{ order_paymenteventquantity : "(id) - (event_id)"
    order_line ||--o{ order_paymenteventquantity : "(id) - (line_id)"
    django_site |o--o{ order_order : "(id) - (site_id)"
    basket_basket |o--o{ order_order : "(id) - (basket_id)"
    auth_user |o--o{ order_order : "(id) - (user_id)"
    order_billingaddress |o--o{ order_order : "(id) - (billing_address_id)"
    order_shippingaddress |o--o{ order_order : "(id) - (shipping_address_id)"
    order_order ||--o{ order_ordernote : "(id) - (order_id)"
    auth_user |o--o{ order_ordernote : "(id) - (user_id)"
    order_order ||--o{ order_orderstatuschange : "(id) - (order_id)"
    order_order ||--o{ order_communicationevent : "(id) - (order_id)"
    communication_communicationeventtype ||--o{ order_communicationevent : "(id) - (event_type_id)"
    address_country ||--o{ order_shippingaddress : "(iso_3166_1_a2) - (country_id)"
    address_country ||--o{ order_billingaddress : "(iso_3166_1_a2) - (country_id)"
    order_order ||--o{ order_line : "(id) - (order_id)"
    partner_partner |o--o{ order_line : "(id) - (partner_id)"
    partner_stockrecord |o--o{ order_line : "(id) - (stockrecord_id)"
    catalogue_product |o--o{ order_line : "(id) - (product_id)"
    order_order ||--o{ order_lineprice : "(id) - (order_id)"
    order_line ||--o{ order_lineprice : "(id) - (line_id)"
    order_line ||--o{ order_lineattribute : "(id) - (line_id)"
    catalogue_option |o--o{ order_lineattribute : "(id) - (option_id)"
    order_order ||--o{ order_shippingevent : "(id) - (order_id)"
    order_shippingeventtype ||--o{ order_shippingevent : "(id) - (event_type_id)"
    order_shippingevent ||--o{ order_shippingeventquantity : "(id) - (event_id)"
    order_line ||--o{ order_shippingeventquantity : "(id) - (line_id)"
    order_order ||--o{ order_paymentevent : "(id) - (order_id)"
    order_paymenteventtype ||--o{ order_paymentevent : "(id) - (event_type_id)"
    order_shippingevent |o--o{ order_paymentevent : "(id) - (shipping_event_id)"
    order_order ||--o{ order_orderdiscount : "(id) - (order_id)"
    order_line ||--o{ order_orderlinediscount : "(id) - (line_id)"
    order_orderdiscount ||--o{ order_orderlinediscount : "(id) - (order_discount_id)"
    order_order ||--o{ order_surcharge : "(id) - (order_id)"
    catalogue_product ||--o{ customer_productalert : "(id) - (product_id)"
    auth_user |o--o{ customer_productalert : "(id) - (user_id)"
    voucher_voucherset |o--o{ voucher_voucher : "(id) - (voucher_set_id)"
    voucher_voucher }o--o{ offer_conditionaloffer : "(offers) - (vouchers)"
    voucher_voucher ||--o{ voucher_voucherapplication : "(id) - (voucher_id)"
    auth_user |o--o{ voucher_voucherapplication : "(id) - (user_id)"
    order_order ||--o{ voucher_voucherapplication : "(id) - (order_id)"
    auth_user ||--o{ wishlists_wishlist : "(id) - (owner_id)"
    wishlists_wishlist ||--o{ wishlists_line : "(id) - (wishlist_id)"
    catalogue_product |o--o{ wishlists_line : "(id) - (product_id)"
    wishlists_wishlist ||--o{ wishlists_wishlistsharedemail : "(id) - (wishlist_id)"
    oscar_accounts_accounttype |o--o{ oscar_accounts_account : "(id) - (account_type_id)"
    auth_user |o--o{ oscar_accounts_account : "(id) - (primary_user_id)"
    offer_range |o--o{ oscar_accounts_account : "(id) - (product_range_id)"
    oscar_accounts_account }o--o{ auth_user : "(secondary_users) - (account_set)"
    oscar_accounts_account ||--o{ oscar_accounts_transfer : "(id) - (source_id)"
    oscar_accounts_account ||--o{ oscar_accounts_transfer : "(id) - (destination_id)"
    oscar_accounts_transfer |o--o{ oscar_accounts_transfer : "(id) - (parent_id)"
    auth_user |o--o{ oscar_accounts_transfer : "(id) - (user_id)"
    oscar_accounts_transfer ||--o{ oscar_accounts_transaction : "(id) - (transfer_id)"
    oscar_accounts_account ||--o{ oscar_accounts_transaction : "(id) - (account_id)"
    django_admin_log {
        AutoField id PK
        DateTimeField action_time
        ForeignKey user_id FK
        ForeignKey content_type_id FK
        TextField object_id
        CharField object_repr
        LogEntryActionFlag action_flag
        TextField change_message
    }
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
    auth_user {
        AutoField id PK
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        CharField username UK "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
        CharField first_name
        CharField last_name
        EmailField email
        BooleanField is_staff "Designates whether the user can log into this admin site."
        BooleanField is_active "Designates whether this user should be treated as active. Unselect this instead of deleting accounts."
        DateTimeField date_joined
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
    django_site {
        AutoField id PK
        CharField domain UK
        CharField name
    }
    django_flatpage {
        AutoField id PK
        CharField url
        CharField title
        TextField content
        BooleanField enable_comments
        CharField template_name "Example: “flatpages/contact_page.html”. If this isn’t provided, the system will use “flatpages/default.html”."
        BooleanField registration_required "If this is checked, only logged-in users will be able to view the page."
    }
    analytics_productrecord {
        BigAutoField id PK
        OneToOneField product_id FK, UK
        PositiveIntegerField num_views
        PositiveIntegerField num_basket_additions
        PositiveIntegerField num_purchases
        FloatField score
    }
    analytics_userrecord {
        BigAutoField id PK
        OneToOneField user_id FK, UK
        PositiveIntegerField num_product_views
        PositiveIntegerField num_basket_additions
        PositiveIntegerField num_orders
        PositiveIntegerField num_order_lines
        PositiveIntegerField num_order_items
        DecimalField total_spent
        DateTimeField date_last_order
    }
    analytics_userproductview {
        BigAutoField id PK
        ForeignKey user_id FK
        ForeignKey product_id FK
        DateTimeField date_created
    }
    analytics_usersearch {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField query
        DateTimeField date_created
    }
    address_useraddress {
        BigAutoField id PK
        NullCharField code UK
        UserAddressTitle title
        CharField first_name
        CharField last_name
        CharField line1
        CharField line2
        CharField line3
        CharField line4
        CharField state
        UppercaseCharField postcode
        ForeignKey country_id FK
        TextField search_text
        PhoneNumberField phone_number "In case we need to call you about your order"
        TextField notes "Tell us anything we should know when delivering your order."
        ForeignKey user_id FK
        BooleanField is_default_for_shipping
        BooleanField is_default_for_billing
        PositiveIntegerField num_orders_as_shipping_address
        PositiveIntegerField num_orders_as_billing_address
        CharField hash
        DateTimeField date_created
    }
    address_country {
        CharField iso_3166_1_a2 PK
        CharField iso_3166_1_a3
        CharField iso_3166_1_numeric
        CharField printable_name
        CharField name
        PositiveSmallIntegerField display_order "Higher the number, higher the country in the list."
        BooleanField is_shipping_country
    }
    shipping_orderanditemcharges {
        BigAutoField id PK
        AutoSlugField code UK
        CharField name UK
        TextField description
        DecimalField price_per_order
        DecimalField price_per_item
        DecimalField free_shipping_threshold
    }
    shipping_weightbased {
        BigAutoField id PK
        AutoSlugField code UK
        CharField name UK
        TextField description
        DecimalField default_weight "Default product weight in kg when no weight attribute is defined"
    }
    shipping_weightband {
        BigAutoField id PK
        ForeignKey method_id FK
        DecimalField upper_limit "Enter upper limit of this weight band in kg. The lower limit will be determined by the other weight bands."
        DecimalField charge
    }
    catalogue_productclass {
        BigAutoField id PK
        CharField name
        AutoSlugField slug UK
        BooleanField requires_shipping
        BooleanField track_stock
    }
    catalogue_category {
        BigAutoField id PK
        CharField path UK
        PositiveIntegerField depth
        PositiveIntegerField numchild
        CharField name
        NullCharField code UK
        TextField description
        TextField long_description
        CharField meta_title
        TextField meta_description
        ImageField image
        SlugField slug
        BooleanField exclude_from_menu "Exclude this category from the menu."
        BooleanField is_public "Show this category in search results and catalogue listings."
        BooleanField ancestors_are_public "The ancestors of this category are public"
    }
    catalogue_productcategory {
        BigAutoField id PK
        ForeignKey product_id FK
        ForeignKey category_id FK
    }
    catalogue_product {
        BigAutoField id PK
        ProductStructure structure
        BooleanField is_public "Show this product in search results and catalogue listings."
        NullCharField upc UK "Universal Product Code (UPC) is an identifier for a product which is not specific to a particular  supplier. Eg an ISBN for a book."
        ForeignKey parent_id FK "Only choose a parent product if you're creating a child product.  For example if this is a size 4 of a particular t-shirt.  Leave blank if this is a stand-alone product (i.e. there is only one version of this product)."
        CharField title
        SlugField slug
        TextField description
        CharField meta_title
        TextField meta_description
        SmallIntegerField priority "The highest priority products are shown first"
        ForeignKey product_class_id FK "Choose what type of product this is"
        FloatField rating
        DateTimeField date_created
        DateTimeField date_updated
        BooleanField is_discountable "This flag indicates if this product can be used in an offer or not"
        CharField code UK
    }
    catalogue_productrecommendation {
        BigAutoField id PK
        ForeignKey primary_id FK
        ForeignKey recommendation_id FK
        PositiveSmallIntegerField ranking "Determines order of the products. A product with a higher value will appear before one with a lower ranking."
    }
    catalogue_productattribute {
        BigAutoField id PK
        ForeignKey product_class_id FK
        CharField name
        SlugField code
        ProductAttributeType type
        ForeignKey option_group_id FK "Select an option group if using type 'Option' or 'Multi Option'"
        BooleanField required
    }
    catalogue_productattributevalue {
        BigAutoField id PK
        ForeignKey attribute_id FK
        ForeignKey product_id FK
        TextField value_text
        IntegerField value_integer
        BooleanField value_boolean
        FloatField value_float
        TextField value_richtext
        DateField value_date
        DateTimeField value_datetime
        ForeignKey value_option_id FK
        FileField value_file
        ImageField value_image
        ForeignKey entity_content_type_id FK
        PositiveIntegerField entity_object_id
    }
    catalogue_attributeoptiongroup {
        BigAutoField id PK
        CharField name
        NullCharField code UK
    }
    catalogue_attributeoption {
        BigAutoField id PK
        ForeignKey group_id FK
        CharField option
        NullCharField code UK
    }
    catalogue_option {
        BigAutoField id PK
        CharField name
        AutoSlugField code UK
        OptionType type
        BooleanField required
        ForeignKey option_group_id FK "Select an option group if using type 'Option' or 'Multi Option'"
        CharField help_text "Help text shown to the user on the add to cart form"
        IntegerField order "Controls the ordering of product options on product detail pages"
    }
    catalogue_productimage {
        BigAutoField id PK
        ForeignKey product_id FK
        NullCharField code UK
        ImageField original
        CharField caption
        PositiveIntegerField display_order "An image with a display order of zero will be the primary image for a product"
        DateTimeField date_created
    }
    catalogue_product_category_hierarchy {
        CharField id PK
        IntegerField product_id
        IntegerField category_id
    }
    reviews_productreview {
        BigAutoField id PK
        ForeignKey product_id FK
        ProductReviewScore score
        CharField title
        TextField body
        ForeignKey user_id FK
        CharField name
        EmailField email
        URLField homepage
        ProductReviewStatus status
        NullCharField code UK
        IntegerField total_votes
        IntegerField delta_votes
        DateTimeField date_created
    }
    reviews_vote {
        BigAutoField id PK
        ForeignKey review_id FK
        ForeignKey user_id FK
        VoteDelta delta
        DateTimeField date_created
    }
    communication_email {
        BigAutoField id PK
        ForeignKey user_id FK
        EmailField email
        TextField subject
        TextField body_text
        TextField body_html
        DateTimeField date_sent
    }
    communication_communicationeventtype {
        BigAutoField id PK
        AutoSlugField code UK "Code used for looking up this event programmatically"
        CharField name
        CommunicationEventTypeCategory category
        CharField email_subject_template
        TextField email_body_template
        TextField email_body_html_template "HTML template"
        CharField sms_template "SMS template"
        DateTimeField date_created
        DateTimeField date_updated
    }
    communication_notification {
        BigAutoField id PK
        ForeignKey recipient_id FK
        ForeignKey sender_id FK
        CharField subject
        TextField body
        NotificationLocation location
        DateTimeField date_sent
        DateTimeField date_read
    }
    partner_partner {
        BigAutoField id PK
        AutoSlugField code UK
        CharField name
    }
    partner_partneraddress {
        BigAutoField id PK
        NullCharField code UK
        PartnerAddressTitle title
        CharField first_name
        CharField last_name
        CharField line1
        CharField line2
        CharField line3
        CharField line4
        CharField state
        UppercaseCharField postcode
        ForeignKey country_id FK
        TextField search_text
        ForeignKey partner_id FK
    }
    partner_stockrecord {
        BigAutoField id PK
        ForeignKey product_id FK
        ForeignKey partner_id FK
        CharField partner_sku
        CharField price_currency
        DecimalField price
        PositiveIntegerField num_in_stock
        IntegerField num_allocated
        PositiveIntegerField low_stock_threshold
        DateTimeField date_created
        DateTimeField date_updated
    }
    partner_stockalert {
        BigAutoField id PK
        ForeignKey stockrecord_id FK
        PositiveIntegerField threshold
        StockAlertStatus status
        DateTimeField date_created
        DateTimeField date_closed
    }
    basket_basket {
        BigAutoField id PK
        ForeignKey owner_id FK
        BasketStatus status
        DateTimeField date_created
        DateTimeField date_merged
        DateTimeField date_submitted
    }
    basket_line {
        BigAutoField id PK
        ForeignKey basket_id FK
        SlugField line_reference
        ForeignKey product_id FK
        ForeignKey stockrecord_id FK
        PositiveIntegerField quantity
        CharField price_currency
        DecimalField price_excl_tax
        DecimalField price_incl_tax
        CharField tax_code
        DateTimeField date_created
        DateTimeField date_updated
    }
    basket_lineattribute {
        BigAutoField id PK
        ForeignKey line_id FK
        ForeignKey option_id FK
        JSONField value
    }
    payment_transaction {
        BigAutoField id PK
        ForeignKey source_id FK
        CharField txn_type
        DecimalField amount
        CharField reference
        CharField status
        DateTimeField date_created
    }
    payment_source {
        BigAutoField id PK
        ForeignKey order_id FK
        ForeignKey source_type_id FK
        CharField currency
        DecimalField amount_allocated
        DecimalField amount_debited
        DecimalField amount_refunded
        CharField reference
        CharField label
    }
    payment_sourcetype {
        BigAutoField id PK
        CharField name
        AutoSlugField code UK "This is used within forms to identify this source type"
    }
    payment_bankcard {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField card_type
        CharField name
        CharField number
        DateField expiry_date
        CharField partner_reference
    }
    offer_conditionaloffer {
        BigAutoField id PK
        CharField name UK "This is displayed within the customer's cart"
        AutoSlugField slug UK
        TextField description "This is displayed on the offer browsing page"
        ConditionalOfferOfferType offer_type
        BooleanField exclusive "Exclusive offers cannot be combined on the same items"
        CharField status
        ForeignKey condition_id FK
        ForeignKey benefit_id FK
        IntegerField priority "The highest priority offers are applied first"
        DateTimeField start_datetime "Offers are active from the start date. Leave this empty if the offer has no start date."
        DateTimeField end_datetime "Offers are active until the end date. Leave this empty if the offer has no expiry date."
        PositiveIntegerField max_global_applications "The number of times this offer can be used before it is unavailable"
        PositiveIntegerField max_user_applications "The number of times a single user can use this offer"
        PositiveIntegerField max_basket_applications "The number of times this offer can be applied to a cart (and order)"
        DecimalField max_discount "When an offer has given more discount to orders than this threshold, then the offer becomes unavailable"
        DecimalField total_discount
        PositiveIntegerField num_applications
        PositiveIntegerField num_orders
        ExtendedURLField redirect_url
        DateTimeField date_created
    }
    offer_benefit {
        BigAutoField id PK
        ForeignKey range_id FK
        BenefitType type
        PositiveDecimalField value
        PositiveIntegerField max_affected_items "Set this to prevent the discount consuming all items within the range that are in the cart."
        NullCharField proxy_class
    }
    offer_condition {
        BigAutoField id PK
        ForeignKey range_id FK
        ConditionType type
        PositiveDecimalField value
        NullCharField proxy_class
    }
    offer_range {
        BigAutoField id PK
        CharField name UK
        AutoSlugField slug UK
        TextField description
        BooleanField is_public "Public ranges have a customer-facing page"
        BooleanField includes_all_products
        NullCharField proxy_class UK
        DateTimeField date_created
    }
    offer_rangeproduct {
        BigAutoField id PK
        ForeignKey range_id FK
        ForeignKey product_id FK
        IntegerField display_order
    }
    offer_rangeproductfileupload {
        BigAutoField id PK
        ForeignKey range_id FK
        CharField filepath
        PositiveIntegerField size
        ForeignKey uploaded_by_id FK
        DateTimeField date_uploaded
        RangeProductFileUploadUploadType upload_type
        RangeProductFileUploadStatus status
        CharField error_message
        DateTimeField date_processed
        PositiveIntegerField num_new_skus
        PositiveIntegerField num_unknown_skus
        PositiveIntegerField num_duplicate_skus
    }
    order_paymenteventquantity {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey line_id FK
        PositiveIntegerField quantity
    }
    order_order {
        BigAutoField id PK
        CharField number UK
        ForeignKey site_id FK
        ForeignKey basket_id FK
        ForeignKey user_id FK
        ForeignKey billing_address_id FK
        CharField currency
        DecimalField total_incl_tax
        DecimalField total_excl_tax
        DecimalField shipping_incl_tax
        DecimalField shipping_excl_tax
        CharField shipping_tax_code
        ForeignKey shipping_address_id FK
        CharField shipping_method
        CharField shipping_code
        CharField status
        EmailField guest_email
        DateTimeField date_placed
        BooleanField analytics_tracked
    }
    order_ordernote {
        BigAutoField id PK
        ForeignKey order_id FK
        ForeignKey user_id FK
        CharField note_type
        TextField message
        DateTimeField date_created
        DateTimeField date_updated
    }
    order_orderstatuschange {
        BigAutoField id PK
        ForeignKey order_id FK
        CharField old_status
        CharField new_status
        DateTimeField date_created
    }
    order_communicationevent {
        BigAutoField id PK
        ForeignKey order_id FK
        ForeignKey event_type_id FK
        DateTimeField date_created
    }
    order_shippingaddress {
        BigAutoField id PK
        NullCharField code UK
        ShippingAddressTitle title
        CharField first_name
        CharField last_name
        CharField line1
        CharField line2
        CharField line3
        CharField line4
        CharField state
        UppercaseCharField postcode
        ForeignKey country_id FK
        TextField search_text
        PhoneNumberField phone_number "In case we need to call you about your order"
        TextField notes "Tell us anything we should know when delivering your order."
    }
    order_billingaddress {
        BigAutoField id PK
        NullCharField code UK
        BillingAddressTitle title
        CharField first_name
        CharField last_name
        CharField line1
        CharField line2
        CharField line3
        CharField line4
        CharField state
        UppercaseCharField postcode
        ForeignKey country_id FK
        TextField search_text
    }
    order_line {
        BigAutoField id PK
        ForeignKey order_id FK
        ForeignKey partner_id FK
        CharField partner_name
        CharField partner_sku
        CharField partner_line_reference "This is the item number that the partner uses within their system"
        TextField partner_line_notes
        ForeignKey stockrecord_id FK
        ForeignKey product_id FK
        CharField title
        CharField upc
        PositiveIntegerField quantity
        DecimalField line_price_incl_tax
        DecimalField line_price_excl_tax
        DecimalField line_price_before_discounts_incl_tax
        DecimalField line_price_before_discounts_excl_tax
        DecimalField unit_price_incl_tax
        DecimalField unit_price_excl_tax
        CharField tax_code
        CharField status
        PositiveIntegerField num_allocated
        BooleanField allocation_cancelled
    }
    order_lineprice {
        BigAutoField id PK
        ForeignKey order_id FK
        ForeignKey line_id FK
        PositiveIntegerField quantity
        DecimalField price_incl_tax
        DecimalField price_excl_tax
        DecimalField shipping_incl_tax
        DecimalField shipping_excl_tax
        CharField tax_code
    }
    order_lineattribute {
        BigAutoField id PK
        ForeignKey line_id FK
        ForeignKey option_id FK
        CharField type
        JSONField value
    }
    order_shippingevent {
        BigAutoField id PK
        ForeignKey order_id FK
        ForeignKey event_type_id FK
        TextField notes "This could be the dispatch reference, or a tracking number"
        DateTimeField date_created
    }
    order_shippingeventquantity {
        BigAutoField id PK
        ForeignKey event_id FK
        ForeignKey line_id FK
        PositiveIntegerField quantity
    }
    order_shippingeventtype {
        BigAutoField id PK
        CharField name UK
        AutoSlugField code UK
    }
    order_paymentevent {
        BigAutoField id PK
        ForeignKey order_id FK
        DecimalField amount
        CharField reference
        ForeignKey event_type_id FK
        ForeignKey shipping_event_id FK
        DateTimeField date_created
    }
    order_paymenteventtype {
        BigAutoField id PK
        CharField name UK
        AutoSlugField code UK
    }
    order_orderdiscount {
        BigAutoField id PK
        ForeignKey order_id FK
        OrderDiscountCategory category
        PositiveIntegerField offer_id
        CharField offer_name
        PositiveIntegerField voucher_id
        CharField voucher_code
        PositiveIntegerField frequency
        DecimalField amount
        TextField message
    }
    order_orderlinediscount {
        BigAutoField id PK
        ForeignKey line_id FK
        ForeignKey order_discount_id FK
        BooleanField is_incl_tax
        DecimalField amount
    }
    order_surcharge {
        BigAutoField id PK
        ForeignKey order_id FK
        CharField name
        CharField code
        DecimalField incl_tax
        DecimalField excl_tax
        CharField tax_code
    }
    customer_productalert {
        BigAutoField id PK
        ForeignKey product_id FK
        ForeignKey user_id FK
        EmailField email
        CharField key
        ProductAlertStatus status
        DateTimeField date_created
        DateTimeField date_confirmed
        DateTimeField date_cancelled
        DateTimeField date_closed
    }
    voucher_voucherset {
        BigAutoField id PK
        CharField name UK
        PositiveIntegerField count
        IntegerField code_length
        TextField description
        DateTimeField date_created
        DateTimeField start_datetime
        DateTimeField end_datetime
    }
    voucher_voucher {
        BigAutoField id PK
        CharField name UK "This will be shown in the checkout and basket once the coupon is entered"
        CharField code UK "Case insensitive / No spaces allowed"
        VoucherUsage usage
        DateTimeField start_datetime
        DateTimeField end_datetime
        PositiveIntegerField num_basket_additions
        PositiveIntegerField num_orders
        DecimalField total_discount
        ForeignKey voucher_set_id FK
        DateTimeField date_created
    }
    voucher_voucherapplication {
        BigAutoField id PK
        ForeignKey voucher_id FK
        ForeignKey user_id FK
        ForeignKey order_id FK
        DateTimeField date_created
    }
    wishlists_wishlist {
        BigAutoField id PK
        ForeignKey owner_id FK
        CharField name
        CharField key UK
        WishListVisibility visibility
        DateTimeField date_created
    }
    wishlists_line {
        BigAutoField id PK
        ForeignKey wishlist_id FK
        ForeignKey product_id FK
        PositiveIntegerField quantity
        CharField title
    }
    wishlists_wishlistsharedemail {
        BigAutoField id PK
        ForeignKey wishlist_id FK
        EmailField email
    }
    thumbnail_kvstore {
        CharField key PK
        TextField value
    }
    oscar_accounts_accounttype {
        BigAutoField id PK
        CharField path UK
        PositiveIntegerField depth
        PositiveIntegerField numchild
        CharField code UK
        CharField name
    }
    oscar_accounts_account {
        BigAutoField id PK
        CharField name UK
        TextField description "This text is shown to customers during checkout"
        ForeignKey account_type_id FK
        CharField code UK
        ForeignKey primary_user_id FK
        CharField status
        DecimalField credit_limit
        DecimalField balance
        DateTimeField start_date
        DateTimeField end_date
        ForeignKey product_range_id FK
        BooleanField can_be_used_for_non_products "Whether this account can be used to pay for shipping and other charges"
        DateTimeField date_created
    }
    oscar_accounts_transfer {
        BigAutoField id PK
        CharField reference UK
        ForeignKey source_id FK
        ForeignKey destination_id FK
        DecimalField amount
        ForeignKey parent_id FK
        CharField merchant_reference
        CharField description
        ForeignKey user_id FK
        CharField username
        DateTimeField date_created
    }
    oscar_accounts_transaction {
        BigAutoField id PK
        ForeignKey transfer_id FK
        ForeignKey account_id FK
        DecimalField amount
        DateTimeField date_created
    }
    oscar_accounts_ipaddressrecord {
        BigAutoField id PK
        GenericIPAddressField ip_address UK
        PositiveIntegerField total_failures
        PositiveIntegerField consecutive_failures
        DateTimeField date_created
        DateTimeField date_last_failure
    }
```
