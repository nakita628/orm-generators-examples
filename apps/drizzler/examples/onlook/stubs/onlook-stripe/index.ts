// Stand-in for @onlook/stripe: the enums the schema imports, copied as they are.

// From packages/stripe/src/constants.ts
export enum PriceKey {
    PRO_MONTHLY_TIER_1 = 'PRO_MONTHLY_TIER_1',
    PRO_MONTHLY_TIER_2 = 'PRO_MONTHLY_TIER_2',
    PRO_MONTHLY_TIER_3 = 'PRO_MONTHLY_TIER_3',
    PRO_MONTHLY_TIER_4 = 'PRO_MONTHLY_TIER_4',
    PRO_MONTHLY_TIER_5 = 'PRO_MONTHLY_TIER_5',
    PRO_MONTHLY_TIER_6 = 'PRO_MONTHLY_TIER_6',
    PRO_MONTHLY_TIER_7 = 'PRO_MONTHLY_TIER_7',
    PRO_MONTHLY_TIER_8 = 'PRO_MONTHLY_TIER_8',
    PRO_MONTHLY_TIER_9 = 'PRO_MONTHLY_TIER_9',
    PRO_MONTHLY_TIER_10 = 'PRO_MONTHLY_TIER_10',
    PRO_MONTHLY_TIER_11 = 'PRO_MONTHLY_TIER_11',
}

// From packages/stripe/src/types.ts
export enum ProductType {
    FREE = 'free',
    PRO = 'pro',
}

// From packages/stripe/src/types.ts
export enum ScheduledSubscriptionAction {
    PRICE_CHANGE = 'price_change',
    CANCELLATION = 'cancellation',
}

// From packages/stripe/src/types.ts
export enum SubscriptionStatus {
    ACTIVE = 'active',
    CANCELED = 'canceled',
}
