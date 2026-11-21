```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    django_ledger_entitymodel ||--o{ django_ledger_bankaccountmodel : "(uuid) - (entity_model_id)"
    django_ledger_accountmodel ||--o{ django_ledger_bankaccountmodel : "(uuid) - (account_model_id)"
    django_ledger_chartofaccountmodel ||--o{ django_ledger_accountmodel : "(uuid) - (coa_model_id)"
    django_ledger_entitymodel ||--o{ django_ledger_chartofaccountmodel : "(uuid) - (entity_id)"
    django_ledger_entitymodel ||--o{ django_ledger_customermodel : "(uuid) - (entity_model_id)"
    django_ledger_entitymodel ||--o{ django_ledger_unitofmeasuremodel : "(uuid) - (entity_id)"
    django_ledger_entityunitmodel |o--o{ django_ledger_itemtransactionmodel : "(uuid) - (entity_unit_id)"
    django_ledger_itemmodel ||--o{ django_ledger_itemtransactionmodel : "(uuid) - (item_model_id)"
    django_ledger_billmodel |o--o{ django_ledger_itemtransactionmodel : "(uuid) - (bill_model_id)"
    django_ledger_invoicemodel |o--o{ django_ledger_itemtransactionmodel : "(uuid) - (invoice_model_id)"
    django_ledger_purchaseordermodel |o--o{ django_ledger_itemtransactionmodel : "(uuid) - (po_model_id)"
    django_ledger_estimatemodel |o--o{ django_ledger_itemtransactionmodel : "(uuid) - (ce_model_id)"
    django_ledger_unitofmeasuremodel ||--o{ django_ledger_itemmodel : "(uuid) - (uom_id)"
    django_ledger_accountmodel |o--o{ django_ledger_itemmodel : "(uuid) - (inventory_account_id)"
    django_ledger_accountmodel |o--o{ django_ledger_itemmodel : "(uuid) - (cogs_account_id)"
    django_ledger_accountmodel |o--o{ django_ledger_itemmodel : "(uuid) - (earnings_account_id)"
    django_ledger_accountmodel |o--o{ django_ledger_itemmodel : "(uuid) - (expense_account_id)"
    django_ledger_entitymodel ||--o{ django_ledger_itemmodel : "(uuid) - (entity_id)"
    django_ledger_entitymodel ||--o{ django_ledger_ledgermodel : "(uuid) - (entity_id)"
    django_ledger_entitymodel ||--o{ django_ledger_entityunitmodel : "(uuid) - (entity_id)"
    django_ledger_entitymodel ||--o{ django_ledger_vendormodel : "(uuid) - (entity_model_id)"
    django_ledger_chartofaccountmodel |o--o| django_ledger_entitymodel : "(uuid) - (default_coa_id)"
    auth_user ||--o{ django_ledger_entitymodel : "(id) - (admin_id)"
    django_ledger_entitymodel ||--o{ django_ledger_entitystatemodel : "(uuid) - (entity_model_id)"
    django_ledger_entityunitmodel |o--o{ django_ledger_entitystatemodel : "(uuid) - (entity_unit_id)"
    django_ledger_entitymodel ||--o{ django_ledger_entitymanagementmodel : "(uuid) - (entity_id)"
    auth_user ||--o{ django_ledger_entitymanagementmodel : "(id) - (user_id)"
    django_ledger_ledgermodel ||--o| django_ledger_billmodel : "(uuid) - (ledger_id)"
    django_ledger_entitymodel |o--o{ django_ledger_billmodel : "(uuid) - (entity_model_id)"
    django_ledger_vendormodel ||--o{ django_ledger_billmodel : "(uuid) - (vendor_id)"
    django_ledger_accountmodel |o--o{ django_ledger_billmodel : "(uuid) - (cash_account_id)"
    django_ledger_accountmodel |o--o{ django_ledger_billmodel : "(uuid) - (prepaid_account_id)"
    django_ledger_accountmodel |o--o{ django_ledger_billmodel : "(uuid) - (unearned_account_id)"
    django_ledger_estimatemodel |o--o{ django_ledger_billmodel : "(uuid) - (ce_model_id)"
    django_ledger_ledgermodel ||--o| django_ledger_invoicemodel : "(uuid) - (ledger_id)"
    django_ledger_entitymodel |o--o{ django_ledger_invoicemodel : "(uuid) - (entity_model_id)"
    django_ledger_customermodel ||--o{ django_ledger_invoicemodel : "(uuid) - (customer_id)"
    django_ledger_accountmodel |o--o{ django_ledger_invoicemodel : "(uuid) - (cash_account_id)"
    django_ledger_accountmodel |o--o{ django_ledger_invoicemodel : "(uuid) - (prepaid_account_id)"
    django_ledger_accountmodel |o--o{ django_ledger_invoicemodel : "(uuid) - (unearned_account_id)"
    django_ledger_estimatemodel |o--o{ django_ledger_invoicemodel : "(uuid) - (ce_model_id)"
    django_ledger_journalentrymodel ||--o{ django_ledger_transactionmodel : "(uuid) - (journal_entry_id)"
    django_ledger_accountmodel ||--o{ django_ledger_transactionmodel : "(uuid) - (account_id)"
    django_ledger_entityunitmodel |o--o{ django_ledger_journalentrymodel : "(uuid) - (entity_unit_id)"
    django_ledger_ledgermodel ||--o{ django_ledger_journalentrymodel : "(uuid) - (ledger_id)"
    django_ledger_entitymodel ||--o{ django_ledger_purchaseordermodel : "(uuid) - (entity_id)"
    django_ledger_estimatemodel |o--o{ django_ledger_purchaseordermodel : "(uuid) - (ce_model_id)"
    django_ledger_entitymodel ||--o{ django_ledger_estimatemodel : "(uuid) - (entity_id)"
    django_ledger_customermodel ||--o{ django_ledger_estimatemodel : "(uuid) - (customer_id)"
    django_ledger_entitymodel ||--o{ django_ledger_closingentrymodel : "(uuid) - (entity_model_id)"
    django_ledger_ledgermodel ||--o| django_ledger_closingentrymodel : "(uuid) - (ledger_model_id)"
    django_ledger_closingentrymodel ||--o{ django_ledger_closingentrytransactionmodel : "(uuid) - (closing_entry_model_id)"
    django_ledger_accountmodel ||--o{ django_ledger_closingentrytransactionmodel : "(uuid) - (account_model_id)"
    django_ledger_entityunitmodel |o--o{ django_ledger_closingentrytransactionmodel : "(uuid) - (unit_model_id)"
    django_ledger_ledgermodel ||--o{ django_ledger_receiptmodel : "(uuid) - (ledger_model_id)"
    django_ledger_entityunitmodel |o--o{ django_ledger_receiptmodel : "(uuid) - (unit_model_id)"
    django_ledger_customermodel |o--o{ django_ledger_receiptmodel : "(uuid) - (customer_model_id)"
    django_ledger_vendormodel |o--o{ django_ledger_receiptmodel : "(uuid) - (vendor_model_id)"
    django_ledger_accountmodel ||--o{ django_ledger_receiptmodel : "(uuid) - (charge_account_id)"
    django_ledger_accountmodel |o--o{ django_ledger_receiptmodel : "(uuid) - (receipt_account_id)"
    django_ledger_stagedtransactionmodel |o--o| django_ledger_receiptmodel : "(uuid) - (staged_transaction_model_id)"
    django_ledger_bankaccountmodel ||--o{ django_ledger_importjobmodel : "(uuid) - (bank_account_model_id)"
    django_ledger_ledgermodel |o--o| django_ledger_importjobmodel : "(uuid) - (ledger_model_id)"
    django_ledger_stagedtransactionmodel |o--o{ django_ledger_stagedtransactionmodel : "(uuid) - (parent_id)"
    django_ledger_importjobmodel ||--o{ django_ledger_stagedtransactionmodel : "(uuid) - (import_job_id)"
    django_ledger_accountmodel |o--o{ django_ledger_stagedtransactionmodel : "(uuid) - (account_model_id)"
    django_ledger_entityunitmodel |o--o{ django_ledger_stagedtransactionmodel : "(uuid) - (unit_model_id)"
    django_ledger_transactionmodel |o--o| django_ledger_stagedtransactionmodel : "(uuid) - (transaction_model_id)"
    django_ledger_transactionmodel |o--o| django_ledger_stagedtransactionmodel : "(uuid) - (matched_transaction_model_id)"
    django_ledger_vendormodel |o--o{ django_ledger_stagedtransactionmodel : "(uuid) - (vendor_model_id)"
    django_ledger_customermodel |o--o{ django_ledger_stagedtransactionmodel : "(uuid) - (customer_model_id)"
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
    django_ledger_bankaccountmodel {
        DateTimeField created
        DateTimeField updated
        CharField financial_institution "Name of the financial institution (i.e. Bank Name)."
        CharField account_number
        CharField routing_number
        CharField aba_number
        CharField swift_number
        BankAccountModelAccountType account_type
        UUIDField uuid PK
        CharField name
        ForeignKey entity_model_id FK
        ForeignKey account_model_id FK "Account model be used to map transactions from financial institution"
        BooleanField active
        BooleanField hidden
    }
    django_ledger_accountmodel {
        DateTimeField created
        DateTimeField updated
        CharField path UK
        PositiveIntegerField depth
        PositiveIntegerField numchild
        AccountModelRole role
        BooleanField role_default
        UUIDField uuid PK
        CharField code
        CharField name
        AccountModelBalanceType balance_type
        BooleanField locked
        BooleanField active
        ForeignKey coa_model_id FK
    }
    django_ledger_chartofaccountmodel {
        SlugField slug UK
        CharField name
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        ForeignKey entity_id FK
        BooleanField active
        TextField description
    }
    django_ledger_customermodel {
        DateTimeField created
        DateTimeField updated
        CharField address_1
        CharField address_2
        CharField city
        CharField state
        CharField zip_code
        CharField country
        EmailField email
        URLField website
        CharField phone
        FloatField sales_tax_rate
        UUIDField uuid PK
        SlugField customer_code UK
        CharField customer_name
        CharField customer_number "System generated customer number."
        ForeignKey entity_model_id FK
        TextField description
        BooleanField active
        BooleanField hidden
        ImageField picture
        JSONField additional_info
    }
    django_ledger_unitofmeasuremodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        CharField name
        SlugField unit_abbr
        BooleanField is_active
        ForeignKey entity_id FK
    }
    django_ledger_itemtransactionmodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        ForeignKey entity_unit_id FK
        ForeignKey item_model_id FK
        ForeignKey bill_model_id FK
        ForeignKey invoice_model_id FK
        FloatField quantity
        FloatField unit_cost
        DecimalField total_amount
        ForeignKey po_model_id FK
        FloatField po_quantity "Authorized item quantity for purchasing."
        FloatField po_unit_cost "Purchase Order unit cost."
        DecimalField po_total_amount "Maximum authorized cost per Purchase Order."
        ItemTransactionModelPoItemStatus po_item_status
        ForeignKey ce_model_id FK
        FloatField ce_quantity
        FloatField ce_unit_cost_estimate
        DecimalField ce_cost_estimate
        FloatField ce_unit_revenue_estimate
        DecimalField ce_revenue_estimate
        CharField item_notes
    }
    django_ledger_itemmodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        CharField name
        CharField item_id
        CharField item_number
        ItemModelItemRole item_role
        ItemModelItemType item_type
        ForeignKey uom_id FK
        CharField sku
        CharField upc
        BooleanField is_active
        DecimalField default_amount
        BooleanField for_inventory "It is an item you require for your inventory."
        BooleanField is_product_or_service "Is a product or service you sell or provide to customers."
        BooleanField sold_as_unit
        ForeignKey inventory_account_id FK "Inventory account where cost will be capitalized."
        DecimalField inventory_received
        DecimalField inventory_received_value
        ForeignKey cogs_account_id FK "COGS account where cost will be recognized on Income Statement."
        ForeignKey earnings_account_id FK "Earnings account where revenue will be recognized on Income Statement."
        ForeignKey expense_account_id FK "Expense account where cost will be recognized on Income Statement."
        JSONField additional_info
        ForeignKey entity_id FK
    }
    django_ledger_ledgermodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        SlugField ledger_xid "User Defined Ledger ID"
        CharField name
        ForeignKey entity_id FK
        BooleanField posted
        BooleanField locked
        BooleanField hidden
        JSONField additional_info
    }
    django_ledger_entityunitmodel {
        CharField name
        DateTimeField created
        DateTimeField updated
        CharField path UK
        PositiveIntegerField depth
        PositiveIntegerField numchild
        UUIDField uuid PK
        SlugField slug
        ForeignKey entity_id FK
        CharField document_prefix
        BooleanField active
        BooleanField hidden
    }
    django_ledger_vendormodel {
        DateTimeField created
        DateTimeField updated
        CharField address_1
        CharField address_2
        CharField city
        CharField state
        CharField zip_code
        CharField country
        EmailField email
        URLField website
        CharField phone
        CharField financial_institution "Name of the financial institution (i.e. Bank Name)."
        CharField account_number
        CharField routing_number
        CharField aba_number
        CharField swift_number
        VendorModelAccountType account_type
        CharField tax_id_number
        UUIDField uuid PK
        SlugField vendor_code
        CharField vendor_number "System generated vendor number."
        CharField vendor_name
        ForeignKey entity_model_id FK
        TextField description
        BooleanField active
        BooleanField hidden
        ImageField picture
        JSONField additional_info
    }
    django_ledger_entitymodel {
        SlugField slug UK
        DateTimeField created
        DateTimeField updated
        CharField address_1
        CharField address_2
        CharField city
        CharField state
        CharField zip_code
        CharField country
        EmailField email
        URLField website
        CharField phone
        CharField path UK
        PositiveIntegerField depth
        PositiveIntegerField numchild
        UUIDField uuid PK
        CharField name
        OneToOneField default_coa_id FK, UK
        ForeignKey admin_id FK
        BooleanField hidden
        BooleanField accrual_method
        EntityModelFyStartMonth fy_start_month
        DateField last_closing_date
        ImageField picture
        JSONField meta
    }
    django_ledger_entitystatemodel {
        UUIDField uuid PK
        ForeignKey entity_model_id FK
        ForeignKey entity_unit_id FK
        SmallIntegerField fiscal_year
        EntityStateModelKey key
        BigIntegerField sequence
    }
    django_ledger_entitymanagementmodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        ForeignKey entity_id FK
        ForeignKey user_id FK
        EntityManagementModelPermissionLevel permission_level
    }
    django_ledger_billmodel {
        DateTimeField created
        DateTimeField updated
        DecimalField amount_due
        DecimalField amount_paid
        DecimalField amount_receivable
        DecimalField amount_unearned
        DecimalField amount_earned
        BooleanField accrue
        DecimalField progress
        OneToOneField ledger_id FK, UK
        BillModelTerms terms
        DateField date_due
        TextField markdown_notes
        UUIDField uuid PK
        ForeignKey entity_model_id FK
        SlugField bill_number
        BillModelBillStatus bill_status
        SlugField xref
        ForeignKey vendor_id FK
        ForeignKey cash_account_id FK
        ForeignKey prepaid_account_id FK
        ForeignKey unearned_account_id FK
        JSONField additional_info
        ForeignKey ce_model_id FK
        DateField date_draft
        DateField date_in_review
        DateField date_approved
        DateField date_paid
        DateField date_void
        DateField date_canceled
    }
    django_ledger_invoicemodel {
        DateTimeField created
        DateTimeField updated
        DecimalField amount_due
        DecimalField amount_paid
        DecimalField amount_receivable
        DecimalField amount_unearned
        DecimalField amount_earned
        BooleanField accrue
        DecimalField progress
        OneToOneField ledger_id FK, UK
        InvoiceModelTerms terms
        DateField date_due
        TextField markdown_notes
        UUIDField uuid PK
        ForeignKey entity_model_id FK
        SlugField invoice_number
        InvoiceModelInvoiceStatus invoice_status
        ForeignKey customer_id FK
        ForeignKey cash_account_id FK
        ForeignKey prepaid_account_id FK
        ForeignKey unearned_account_id FK
        JSONField additional_info
        ForeignKey ce_model_id FK
        DateField date_draft
        DateField date_in_review
        DateField date_approved
        DateField date_paid
        DateField date_void
        DateField date_canceled
    }
    django_ledger_transactionmodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        TransactionModelTxType tx_type
        ForeignKey journal_entry_id FK "Journal Entry to be associated with this transaction."
        ForeignKey account_id FK "Account from Chart of Accounts to be associated with this transaction."
        DecimalField amount "Amount of the transaction."
        CharField description "A description to be included with this individual transaction."
        BooleanField cleared
        BooleanField reconciled
    }
    django_ledger_journalentrymodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        SlugField je_number
        DateTimeField timestamp
        CharField description
        ForeignKey entity_unit_id FK
        JournalEntryModelActivity activity
        CharField origin
        BooleanField posted
        BooleanField locked
        BooleanField is_closing_entry
        ForeignKey ledger_id FK
    }
    django_ledger_purchaseordermodel {
        DateTimeField created
        DateTimeField updated
        TextField markdown_notes
        UUIDField uuid PK
        SlugField po_number
        CharField po_title
        PurchaseOrderModelPoStatus po_status
        DecimalField po_amount
        DecimalField po_amount_received
        ForeignKey entity_id FK
        DateField date_draft
        DateField date_in_review
        DateField date_approved
        DateField date_void
        DateField date_fulfilled
        DateField date_canceled
        ForeignKey ce_model_id FK
    }
    django_ledger_estimatemodel {
        DateTimeField created
        DateTimeField updated
        TextField markdown_notes
        UUIDField uuid PK
        SlugField estimate_number
        ForeignKey entity_id FK
        ForeignKey customer_id FK
        EstimateModelTerms terms
        CharField title
        EstimateModelStatus status
        DateField date_draft
        DateField date_in_review
        DateField date_approved
        DateField date_completed
        DateField date_canceled
        DateField date_void
        DecimalField revenue_estimate "Estimated cost to complete the quoted work."
        DecimalField labor_estimate "Estimated labor cost to complete the quoted work."
        DecimalField material_estimate "Estimated material cost to complete the quoted work."
        DecimalField equipment_estimate "Estimated equipment cost to complete the quoted work."
        DecimalField other_estimate "Estimated equipment cost to complete the quoted work."
    }
    django_ledger_closingentrymodel {
        DateTimeField created
        DateTimeField updated
        TextField markdown_notes
        UUIDField uuid PK
        ForeignKey entity_model_id FK
        OneToOneField ledger_model_id FK, UK
        DateField closing_date
        BooleanField posted
    }
    django_ledger_closingentrytransactionmodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        ForeignKey closing_entry_model_id FK
        ForeignKey account_model_id FK
        ForeignKey unit_model_id FK
        ClosingEntryTransactionModelActivity activity
        ClosingEntryTransactionModelTxType tx_type
        DecimalField balance
    }
    django_ledger_receiptmodel {
        DateTimeField created
        DateTimeField updated
        TextField markdown_notes
        UUIDField uuid PK
        CharField receipt_number
        DateField receipt_date
        ReceiptModelReceiptType receipt_type
        ForeignKey ledger_model_id FK
        ForeignKey unit_model_id FK "Helps segregate receipts and transactions into different classes or departments."
        ForeignKey customer_model_id FK
        ForeignKey vendor_model_id FK
        ForeignKey charge_account_id FK "The financial account (cash or credit) where this transaction was made."
        ForeignKey receipt_account_id FK "The income or expense account where this transaction will be reflected"
        DecimalField amount "Amount of the receipt."
        OneToOneField staged_transaction_model_id FK, UK "The staged transaction associated with the receipt from bank feeds."
    }
    django_ledger_importjobmodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        CharField description
        ForeignKey bank_account_model_id FK
        OneToOneField ledger_model_id FK, UK
        BooleanField completed
    }
    django_ledger_stagedtransactionmodel {
        DateTimeField created
        DateTimeField updated
        UUIDField uuid PK
        ForeignKey parent_id FK
        ForeignKey import_job_id FK
        CharField fit_id
        DateField date_posted
        BooleanField bundle_split
        StagedTransactionModelActivity activity
        DecimalField amount
        DecimalField amount_split
        CharField name
        CharField memo
        ForeignKey account_model_id FK
        ForeignKey unit_model_id FK
        OneToOneField transaction_model_id FK, UK
        OneToOneField matched_transaction_model_id FK, UK
        BooleanField matched_transaction "Indicates whether this transaction is matched to another."
        StagedTransactionModelReceiptType receipt_type "The receipt type of the transaction."
        ForeignKey vendor_model_id FK "The Vendor associated with the transaction."
        ForeignKey customer_model_id FK "The Customer associated with the transaction."
        TextField notes
    }
```
