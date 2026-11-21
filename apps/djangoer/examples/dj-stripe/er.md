```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    djstripe_account |o--o{ djstripe_apikey : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_account : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_accountv2 : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_coupon : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_promotioncode : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_discount : "(id) - (djstripe_owner_account_id)"
    djstripe_customer |o--o{ djstripe_discount : "(id) - (customer_id)"
    djstripe_invoice |o--o{ djstripe_discount : "(id) - (invoice_id)"
    djstripe_subscription |o--o{ djstripe_discount : "(id) - (subscription_id)"
    djstripe_account |o--o{ djstripe_invoice : "(id) - (djstripe_owner_account_id)"
    djstripe_charge |o--o| djstripe_invoice : "(djstripe_id) - (charge_id)"
    djstripe_customer ||--o{ djstripe_invoice : "(id) - (customer_id)"
    djstripe_paymentmethod |o--o{ djstripe_invoice : "(id) - (default_payment_method_id)"
    djstripe_paymentintent |o--o| djstripe_invoice : "(djstripe_id) - (payment_intent_id)"
    djstripe_subscription |o--o{ djstripe_invoice : "(id) - (subscription_id)"
    djstripe_invoice }o--o{ djstripe_taxrate : "(default_tax_rates) - (_djstripe_invoice_default_tax_rates_+)"
    djstripe_account |o--o{ djstripe_upcominginvoice : "(id) - (djstripe_owner_account_id)"
    djstripe_charge |o--o| djstripe_upcominginvoice : "(djstripe_id) - (charge_id)"
    djstripe_customer ||--o{ djstripe_upcominginvoice : "(id) - (customer_id)"
    djstripe_paymentmethod |o--o{ djstripe_upcominginvoice : "(id) - (default_payment_method_id)"
    djstripe_paymentintent |o--o| djstripe_upcominginvoice : "(djstripe_id) - (payment_intent_id)"
    djstripe_subscription |o--o{ djstripe_upcominginvoice : "(id) - (subscription_id)"
    djstripe_account |o--o{ djstripe_invoiceitem : "(id) - (djstripe_owner_account_id)"
    djstripe_customer ||--o{ djstripe_invoiceitem : "(id) - (customer_id)"
    djstripe_invoice |o--o{ djstripe_invoiceitem : "(id) - (invoice_id)"
    djstripe_price |o--o{ djstripe_invoiceitem : "(djstripe_id) - (price_id)"
    djstripe_subscription |o--o{ djstripe_invoiceitem : "(id) - (subscription_id)"
    djstripe_invoiceitem }o--o{ djstripe_taxrate : "(tax_rates) - (_djstripe_invoiceitem_tax_rates_+)"
    djstripe_account |o--o{ djstripe_lineitem : "(id) - (djstripe_owner_account_id)"
    djstripe_invoiceitem |o--o{ djstripe_lineitem : "(id) - (invoice_item_id)"
    djstripe_subscription |o--o{ djstripe_lineitem : "(id) - (subscription_id)"
    djstripe_subscriptionitem |o--o{ djstripe_lineitem : "(id) - (subscription_item_id)"
    djstripe_account |o--o{ djstripe_subscription : "(id) - (djstripe_owner_account_id)"
    djstripe_customer ||--o{ djstripe_subscription : "(id) - (customer_id)"
    djstripe_account |o--o{ djstripe_subscriptionitem : "(id) - (djstripe_owner_account_id)"
    djstripe_price |o--o{ djstripe_subscriptionitem : "(djstripe_id) - (price_id)"
    djstripe_subscription ||--o{ djstripe_subscriptionitem : "(id) - (subscription_id)"
    djstripe_subscriptionitem }o--o{ djstripe_taxrate : "(tax_rates) - (_djstripe_subscriptionitem_tax_rates_+)"
    djstripe_account |o--o{ djstripe_subscriptionschedule : "(id) - (djstripe_owner_account_id)"
    djstripe_customer ||--o{ djstripe_subscriptionschedule : "(djstripe_id) - (customer_id)"
    djstripe_subscription |o--o{ djstripe_subscriptionschedule : "(djstripe_id) - (released_subscription_id)"
    djstripe_subscription |o--o{ djstripe_subscriptionschedule : "(djstripe_id) - (subscription_id)"
    djstripe_account |o--o{ djstripe_shippingrate : "(id) - (djstripe_owner_account_id)"
    djstripe_taxcode |o--o{ djstripe_shippingrate : "(id) - (tax_code_id)"
    djstripe_account |o--o{ djstripe_taxcode : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_taxid : "(id) - (djstripe_owner_account_id)"
    djstripe_customer |o--o{ djstripe_taxid : "(id) - (customer_id)"
    djstripe_account |o--o{ djstripe_taxrate : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_session : "(id) - (djstripe_owner_account_id)"
    djstripe_customer |o--o{ djstripe_session : "(id) - (customer_id)"
    djstripe_paymentintent |o--o{ djstripe_session : "(id) - (payment_intent_id)"
    djstripe_setupintent |o--o{ djstripe_session : "(id) - (setup_intent_id)"
    djstripe_subscription |o--o{ djstripe_session : "(id) - (subscription_id)"
    djstripe_account |o--o{ djstripe_applicationfee : "(id) - (djstripe_owner_account_id)"
    djstripe_account ||--o{ djstripe_applicationfee : "(id) - (account_id)"
    djstripe_balancetransaction ||--o{ djstripe_applicationfee : "(id) - (balance_transaction_id)"
    djstripe_charge ||--o{ djstripe_applicationfee : "(id) - (charge_id)"
    djstripe_account |o--o{ djstripe_applicationfeerefund : "(id) - (djstripe_owner_account_id)"
    djstripe_balancetransaction ||--o{ djstripe_applicationfeerefund : "(id) - (balance_transaction_id)"
    djstripe_applicationfee ||--o{ djstripe_applicationfeerefund : "(id) - (fee_id)"
    djstripe_account |o--o{ djstripe_transfer : "(id) - (djstripe_owner_account_id)"
    djstripe_balancetransaction |o--o{ djstripe_transfer : "(id) - (balance_transaction_id)"
    djstripe_account |o--o{ djstripe_transferreversal : "(id) - (djstripe_owner_account_id)"
    djstripe_balancetransaction |o--o{ djstripe_transferreversal : "(id) - (balance_transaction_id)"
    djstripe_transfer ||--o{ djstripe_transferreversal : "(id) - (transfer_id)"
    djstripe_account |o--o{ djstripe_balancetransaction : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_charge : "(id) - (djstripe_owner_account_id)"
    djstripe_balancetransaction |o--o{ djstripe_charge : "(id) - (balance_transaction_id)"
    djstripe_customer |o--o{ djstripe_charge : "(id) - (customer_id)"
    djstripe_invoice |o--o{ djstripe_charge : "(id) - (invoice_id)"
    djstripe_paymentintent |o--o{ djstripe_charge : "(id) - (payment_intent_id)"
    djstripe_paymentmethod |o--o{ djstripe_charge : "(id) - (payment_method_id)"
    djstripe_djstripepaymentmethod |o--o{ djstripe_charge : "(id) - (source_id)"
    djstripe_account |o--o{ djstripe_mandate : "(id) - (djstripe_owner_account_id)"
    djstripe_paymentmethod ||--o{ djstripe_mandate : "(id) - (payment_method_id)"
    djstripe_account |o--o{ djstripe_product : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_customer : "(id) - (djstripe_owner_account_id)"
    djstripe_paymentmethod |o--o{ djstripe_customer : "(id) - (default_payment_method_id)"
    auth_user |o--o{ djstripe_customer : "(id) - (subscriber_id)"
    djstripe_account |o--o{ djstripe_dispute : "(id) - (djstripe_owner_account_id)"
    djstripe_balancetransaction |o--o{ djstripe_dispute : "(id) - (balance_transaction_id)"
    djstripe_charge |o--o{ djstripe_dispute : "(id) - (charge_id)"
    djstripe_paymentintent |o--o{ djstripe_dispute : "(id) - (payment_intent_id)"
    djstripe_account |o--o{ djstripe_event : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_file : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_filelink : "(id) - (djstripe_owner_account_id)"
    djstripe_file ||--o{ djstripe_filelink : "(id) - (file_id)"
    djstripe_account |o--o{ djstripe_paymentintent : "(id) - (djstripe_owner_account_id)"
    djstripe_customer |o--o{ djstripe_paymentintent : "(id) - (customer_id)"
    djstripe_account |o--o{ djstripe_paymentintent : "(id) - (on_behalf_of_id)"
    djstripe_paymentmethod |o--o{ djstripe_paymentintent : "(id) - (payment_method_id)"
    djstripe_account |o--o{ djstripe_setupintent : "(id) - (djstripe_owner_account_id)"
    djstripe_customer |o--o{ djstripe_setupintent : "(id) - (customer_id)"
    djstripe_account |o--o{ djstripe_setupintent : "(id) - (on_behalf_of_id)"
    djstripe_paymentmethod |o--o{ djstripe_setupintent : "(id) - (payment_method_id)"
    djstripe_account |o--o{ djstripe_payout : "(id) - (djstripe_owner_account_id)"
    djstripe_balancetransaction |o--o{ djstripe_payout : "(id) - (balance_transaction_id)"
    djstripe_bankaccount |o--o{ djstripe_payout : "(id) - (destination_id)"
    djstripe_balancetransaction |o--o{ djstripe_payout : "(id) - (failure_balance_transaction_id)"
    djstripe_payout |o--o| djstripe_payout : "(djstripe_id) - (original_payout_id)"
    djstripe_payout |o--o| djstripe_payout : "(djstripe_id) - (reversed_by_id)"
    djstripe_account |o--o{ djstripe_price : "(id) - (djstripe_owner_account_id)"
    djstripe_product ||--o{ djstripe_price : "(id) - (product_id)"
    djstripe_account |o--o{ djstripe_refund : "(id) - (djstripe_owner_account_id)"
    djstripe_balancetransaction |o--o{ djstripe_refund : "(id) - (balance_transaction_id)"
    djstripe_charge ||--o{ djstripe_refund : "(id) - (charge_id)"
    djstripe_balancetransaction |o--o{ djstripe_refund : "(id) - (failure_balance_transaction_id)"
    djstripe_account |o--o{ djstripe_feature : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_productfeature : "(id) - (djstripe_owner_account_id)"
    djstripe_product ||--o{ djstripe_productfeature : "(id) - (product_id)"
    djstripe_feature ||--o{ djstripe_productfeature : "(id) - (entitlement_feature_id)"
    djstripe_account |o--o{ djstripe_activeentitlement : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_verificationsession : "(id) - (djstripe_owner_account_id)"
    djstripe_verificationreport |o--o{ djstripe_verificationsession : "(id) - (last_verification_report_id)"
    djstripe_account |o--o{ djstripe_verificationreport : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_issuing_authorization : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_issuing_card : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_issuing_cardholder : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_issuing_dispute : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_issuing_transaction : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_bankaccount : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_bankaccount : "(id) - (account_id)"
    djstripe_account |o--o{ djstripe_card : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_card : "(id) - (account_id)"
    djstripe_account |o--o{ djstripe_source : "(id) - (djstripe_owner_account_id)"
    djstripe_customer |o--o{ djstripe_source : "(id) - (customer_id)"
    djstripe_account |o--o{ djstripe_sourcetransaction : "(id) - (djstripe_owner_account_id)"
    djstripe_source |o--o{ djstripe_sourcetransaction : "(id) - (source_id)"
    djstripe_account |o--o{ djstripe_paymentmethod : "(id) - (djstripe_owner_account_id)"
    djstripe_customer |o--o{ djstripe_paymentmethod : "(id) - (customer_id)"
    djstripe_account |o--o{ djstripe_earlyfraudwarning : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_review : "(id) - (djstripe_owner_account_id)"
    djstripe_account |o--o{ djstripe_scheduledqueryrun : "(id) - (djstripe_owner_account_id)"
    djstripe_file |o--o{ djstripe_scheduledqueryrun : "(id) - (file_id)"
    djstripe_account |o--o{ djstripe_webhookendpoint : "(id) - (djstripe_owner_account_id)"
    djstripe_event |o--o{ djstripe_webhookeventtrigger : "(id) - (event_id)"
    djstripe_account |o--o{ djstripe_webhookeventtrigger : "(id) - (stripe_trigger_account_id)"
    djstripe_webhookendpoint |o--o{ djstripe_webhookeventtrigger : "(id) - (webhook_endpoint_id)"
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
    djstripe_idempotencykey {
        UUIDField uuid PK
        CharField action
        BooleanField livemode "Whether the key was used in live or test mode."
        DateTimeField created
    }
    djstripe_apikey {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        StripeDateTimeField created "The datetime this object was created in stripe."
        CharField id
        APIKeyType type
        CharField name "An optional name to identify the key."
        CharField secret UK "The value of the key."
        BooleanField livemode "Whether the key is valid for live or test mode."
    }
    djstripe_account {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_accountv2 {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_coupon {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeIdField id
    }
    djstripe_promotioncode {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_discount {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey customer_id FK "The ID of the customer associated with this discount."
        StripeForeignKey invoice_id FK "The invoice that the discount’s coupon was applied to, if it was applied directly to a particular invoice."
        CharField promotion_code "The promotion code applied to create this discount."
        StripeForeignKey subscription_id FK "The subscription that this coupon is applied to, if it is applied to a particular subscription."
    }
    djstripe_invoice {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        OneToOneField charge_id FK, UK "The latest charge generated for this invoice, if any."
        StripeForeignKey customer_id FK "The customer associated with this invoice."
        StripeForeignKey default_payment_method_id FK "Default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription's default payment method, if any, or to the default payment method in the customer's invoice settings."
        OneToOneField payment_intent_id FK, UK "The PaymentIntent associated with this invoice. The PaymentIntent is generated when the invoice is finalized, and can then be used to pay the invoice.Note that voiding an invoice will cancel the PaymentIntent"
        StripeForeignKey subscription_id FK "The subscription that this invoice was prepared for, if any."
    }
    djstripe_upcominginvoice {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        OneToOneField charge_id FK, UK "The latest charge generated for this invoice, if any."
        StripeForeignKey customer_id FK "The customer associated with this invoice."
        StripeForeignKey default_payment_method_id FK "Default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription's default payment method, if any, or to the default payment method in the customer's invoice settings."
        OneToOneField payment_intent_id FK, UK "The PaymentIntent associated with this invoice. The PaymentIntent is generated when the invoice is finalized, and can then be used to pay the invoice.Note that voiding an invoice will cancel the PaymentIntent"
        StripeForeignKey subscription_id FK "The subscription that this invoice was prepared for, if any."
    }
    djstripe_invoiceitem {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey customer_id FK "The customer associated with this invoiceitem."
        StripeForeignKey invoice_id FK "The invoice to which this invoiceitem is attached."
        ForeignKey price_id FK "If the invoice item is a proration, the price of the subscription for which the proration was computed."
        StripeForeignKey subscription_id FK "The subscription that this invoice item has been created for, if any."
    }
    djstripe_lineitem {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey invoice_item_id FK "The ID of the invoice item associated with this line item if any."
        StripeForeignKey subscription_id FK "The subscription that the invoice item pertains to, if any."
        StripeForeignKey subscription_item_id FK "The subscription item that generated this invoice item. Left empty if the line item is not an explicit result of a subscription."
    }
    djstripe_subscription {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey customer_id FK "The customer associated with this subscription."
    }
    djstripe_subscriptionitem {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        ForeignKey price_id FK "The price the customer is subscribed to."
        StripeForeignKey subscription_id FK "The subscription this subscription item belongs to."
    }
    djstripe_subscriptionschedule {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeDateTimeField canceled_at "Time at which the subscription schedule was canceled."
        StripeDateTimeField completed_at "Time at which the subscription schedule was completed."
        JSONField current_phase "Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`."
        ForeignKey customer_id FK "The customer who owns the subscription schedule."
        JSONField default_settings "Object representing the subscription schedule's default settings."
        SubscriptionScheduleEndBehavior end_behavior "Behavior of the subscription schedule and underlying subscription when it ends."
        JSONField phases "Configuration for the subscription schedule's phases."
        StripeDateTimeField released_at "Time at which the subscription schedule was released."
        ForeignKey released_subscription_id FK "The subscription once managed by this subscription schedule (if it is released)."
        SubscriptionScheduleStatus status "The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`."
        ForeignKey subscription_id FK "ID of the subscription managed by the subscription schedule."
    }
    djstripe_shippingrate {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey tax_code_id FK "The shipping tax code"
    }
    djstripe_taxcode {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
    }
    djstripe_taxid {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        StripeForeignKey customer_id FK
    }
    djstripe_taxrate {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_session {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey customer_id FK "Customer this Checkout is for if one exists."
        StripeForeignKey payment_intent_id FK "PaymentIntent created if SKUs or line items were provided."
        StripeForeignKey setup_intent_id FK "The ID of the SetupIntent for Checkout Sessions in setup mode."
        StripeForeignKey subscription_id FK "Subscription created if one or more plans were provided."
    }
    djstripe_applicationfee {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey account_id FK "ID of the Stripe account this fee was taken from."
        StripeForeignKey balance_transaction_id FK "Balance transaction that describes the impact on your account balance."
        StripeForeignKey charge_id FK "The charge that the application fee was taken from."
    }
    djstripe_applicationfeerefund {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey balance_transaction_id FK "Balance transaction that describes the impact on your account balance."
        StripeForeignKey fee_id FK "The application fee that was refunded"
    }
    djstripe_countryspec {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        CharField id PK
        StripeCurrencyCodeField default_currency "The default currency for this country. This applies to both payment methods and bank accounts."
        JSONField supported_bank_account_currencies "Currencies that can be accepted in the specific country (for transfers)."
        JSONField supported_payment_currencies "Currencies that can be accepted in the specified country (for payments)."
        JSONField supported_payment_methods "Payment methods available in the specified country."
        JSONField supported_transfer_countries "Countries that can accept transfers from the specified country."
        JSONField verification_fields "Lists the types of verification data needed to keep an account open."
    }
    djstripe_transfer {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey balance_transaction_id FK "Balance transaction that describes the impact on your account balance."
    }
    djstripe_transferreversal {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey balance_transaction_id FK "Balance transaction that describes the impact on your account balance."
        StripeForeignKey transfer_id FK "The transfer that was reversed."
    }
    djstripe_balancetransaction {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeQuantumCurrencyAmountField amount "Gross amount of the transaction, in cents."
        StripeCurrencyCodeField currency "Three-letter ISO currency code"
        StripeIdField source
        BalanceTransactionType type
    }
    djstripe_charge {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeDecimalCurrencyAmountField amount "Amount charged (as decimal)."
        StripeCurrencyCodeField currency "The currency in which the charge was made."
        StripeForeignKey balance_transaction_id FK "The balance transaction that describes the impact of this charge on your account balance (not including refunds or disputes)."
        StripeForeignKey customer_id FK "The customer associated with this charge."
        StripeForeignKey invoice_id FK "The invoice this charge is for if one exists."
        StripeForeignKey payment_intent_id FK "PaymentIntent associated with this charge, if one exists."
        StripeForeignKey payment_method_id FK "PaymentMethod used in this charge."
        PaymentMethodForeignKey source_id FK "The source used for this charge."
        ChargeStatus status "The status of the payment."
    }
    djstripe_mandate {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey payment_method_id FK
    }
    djstripe_product {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        TextField name "The product's name, meant to be displayable to the customer. Applicable to both `service` and `good` types."
        BooleanField active "Whether the product is currently available for purchase. Only applicable to products of `type=good`."
    }
    djstripe_customer {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        TextField email
        StripeForeignKey default_payment_method_id FK "default payment method used for subscriptions and invoices for the customer."
        ForeignKey subscriber_id FK
        DateTimeField date_purged
    }
    djstripe_dispute {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey balance_transaction_id FK "Balance transaction that describes the impact on your account balance."
        StripeForeignKey charge_id FK "The charge that was disputed"
        StripeForeignKey payment_intent_id FK "The PaymentIntent that was disputed"
    }
    djstripe_event {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        CharField api_version "the API version at which the event data was rendered. Blank for old entries only, all new entries will have this value"
        JSONField data "data received at webhook. data should be considered to be garbage until validity check is run and valid flag is set"
        CharField request_id "Information about the request that triggered this event, for traceability purposes. If empty string then this is an old entry without that data. If Null then this is not an old entry, but a Stripe 'automated' event with no associated request."
        TextField idempotency_key
        CharField type "Stripe's event description code"
    }
    djstripe_file {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_filelink {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey file_id FK
    }
    djstripe_paymentintent {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey customer_id FK "Customer this PaymentIntent is for if one exists."
        StripeForeignKey on_behalf_of_id FK "The account (if any) for which the funds of the PaymentIntent are intended."
        StripeForeignKey payment_method_id FK "Payment method used in this PaymentIntent."
    }
    djstripe_setupintent {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey customer_id FK "Customer this SetupIntent belongs to, if one exists."
        StripeForeignKey on_behalf_of_id FK "The account (if any) for which the setup is intended."
        StripeForeignKey payment_method_id FK "Payment method used in this PaymentIntent."
    }
    djstripe_payout {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey balance_transaction_id FK "Balance transaction that describes the impact on your account balance."
        StripeCurrencyCodeField currency "Three-letter ISO currency code"
        StripeForeignKey destination_id FK "Bank account or card the payout was sent to."
        StripeForeignKey failure_balance_transaction_id FK "If the payout failed or was canceled, this will be the balance transaction that reversed the initial balance transaction, and puts the funds from the failed payout back in your balance."
        OneToOneField original_payout_id FK, UK "If this payout reverses another, this is the ID of the original payout."
        OneToOneField reversed_by_id FK, UK "If this payout was reversed, this is the ID of the payout that reverses this payout."
    }
    djstripe_price {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        BooleanField active "Whether the price can be used for new purchases."
        StripeCurrencyCodeField currency "Three-letter ISO currency code"
        CharField nickname "A brief description of the plan, hidden from customers."
        StripeForeignKey product_id FK "The product this price is associated with."
        CharField lookup_key "A lookup key used to retrieve prices dynamically from a static string."
    }
    djstripe_refund {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeQuantumCurrencyAmountField amount "Amount, in cents."
        StripeForeignKey balance_transaction_id FK "Balance transaction that describes the impact on your account balance."
        StripeForeignKey charge_id FK "The charge that was refunded"
        StripeCurrencyCodeField currency "Three-letter ISO currency code"
        StripeForeignKey failure_balance_transaction_id FK "If the refund failed, this balance transaction describes the adjustment made on your account balance that reverses the initial balance transaction."
    }
    djstripe_feature {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_productfeature {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey product_id FK "The product this feature is attached to."
        StripeForeignKey entitlement_feature_id FK "The feature that this product feature references."
    }
    djstripe_activeentitlement {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_verificationsession {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        JSONField last_error
        StripeForeignKey last_verification_report_id FK
        JSONField redaction
        JSONField verified_outputs
        VerificationSessionStatus status
        VerificationSessionType type
    }
    djstripe_verificationreport {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        JSONField document
        JSONField id_number
        JSONField options
        JSONField selfie
        VerificationReportType type
    }
    djstripe_issuing_authorization {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_issuing_card {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_issuing_cardholder {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_issuing_dispute {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_issuing_transaction {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_djstripepaymentmethod {
        CharField id PK
        CharField type
    }
    djstripe_bankaccount {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey account_id FK "The external account the charge was made on behalf of. Null here indicates that this value was never set."
        CharField fingerprint "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same."
    }
    djstripe_card {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey account_id FK "The external account the charge was made on behalf of. Null here indicates that this value was never set."
        CharField fingerprint "Uniquely identifies this particular card number."
    }
    djstripe_source {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey customer_id FK
    }
    djstripe_sourcetransaction {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField ach_credit_transfer "The data corresponding to the ach_credit_transfer type."
        StripeDecimalCurrencyAmountField amount "Amount (as decimal) associated with the ACH Credit Transfer. This is the amount your customer has sent for which the source will be chargeable once ready."
        StripeCurrencyCodeField currency "Three-letter ISO currency code"
        JSONField customer_data "Customer defined string used to initiate the ACH Credit Transfer."
        StripeForeignKey source_id FK
        SourceTransactionStatus status "The status of the ACH Credit Transfer. Only `chargeable` sources can be used to create a charge."
    }
    djstripe_paymentmethod {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeForeignKey customer_id FK "Customer to which this PaymentMethod is saved. This will not be set when the PaymentMethod has not been saved to a Customer."
    }
    djstripe_earlyfraudwarning {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_review {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
    }
    djstripe_scheduledqueryrun {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        StripeDateTimeField data_load_time "When the query was run, Sigma contained a snapshot of your Stripe data at this time."
        JSONField error "If the query run was not succeesful, contains information about the failure."
        StripeForeignKey file_id FK "The file object representing the results of the query."
        StripeDateTimeField result_available_until "Time at which the result expires and is no longer available for download."
        TextField sql "SQL for the query."
        ScheduledQueryRunStatus status "The query's execution status."
        TextField title "Title of the query."
    }
    djstripe_webhookendpoint {
        DateTimeField djstripe_created
        DateTimeField djstripe_updated
        JSONField stripe_data
        BigAutoField djstripe_id PK
        StripeIdField id UK
        StripeForeignKey djstripe_owner_account_id FK "The Stripe Account this object belongs to."
        BooleanField livemode "Null here indicates that the livemode status is unknown or was previously unrecorded. Otherwise, this field indicates whether this record comes from Stripe test mode or live mode operation."
        StripeDateTimeField created "The datetime this object was created in stripe."
        JSONField metadata "A set of key/value pairs that you can attach to an object. It can be useful for storing additional information about an object in a structured format."
        CharField api_version "The API version events are rendered as for this webhook endpoint. Defaults to the configured Stripe API Version."
        JSONField enabled_events "The list of events to enable for this endpoint. ['*'] indicates that all events are enabled, except those that require explicit selection."
        CharField secret "The endpoint's secret, used to generate webhook signatures."
        WebhookEndpointStatus status "The status of the webhook. It can be enabled or disabled."
        URLField url "The URL of the webhook endpoint."
        CharField application "The ID of the associated Connect application."
        UUIDField djstripe_uuid UK "A UUID specific to dj-stripe generated for the endpoint"
        PositiveSmallIntegerField djstripe_tolerance "Controls the milliseconds tolerance which wards against replay attacks. Leave this to its default value unless you know what you're doing."
        WebhookEndpointDjstripeValidationMethod djstripe_validation_method "Controls the webhook validation method."
    }
    djstripe_webhookeventtrigger {
        BigAutoField id PK
        GenericIPAddressField remote_ip "IP address of the request client."
        JSONField headers
        TextField body
        BooleanField valid "Whether or not the webhook event has passed validation"
        BooleanField processed "Whether or not the webhook event has been successfully processed"
        CharField exception
        TextField traceback "Traceback if an exception was thrown during processing"
        StripeForeignKey event_id FK "Event object contained in the (valid) Webhook"
        CharField djstripe_version "The version of dj-stripe when the webhook was received"
        DateTimeField created
        DateTimeField updated
        StripeForeignKey stripe_trigger_account_id FK "The Stripe Account this object belongs to."
        StripeForeignKey webhook_endpoint_id FK "The endpoint this webhook was received on"
    }
```
