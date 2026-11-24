<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasVersion4Uuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

class LiteLLMBudgetTable extends Model
{
    use HasVersion4Uuids;
    use PrismaDates;

    protected $table = 'LiteLLM_BudgetTable';

    protected $primaryKey = 'budget_id';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'rollover_max_budget',
        'name',
        'max_budget',
        'soft_budget',
        'max_parallel_requests',
        'tpm_limit',
        'rpm_limit',
        'tpd_limit',
        'model_max_budget',
        'model_tpm_limit',
        'model_rpm_limit',
        'budget_duration',
        'budget_limits',
        'cap_policy',
        'budget_reset_at',
        'budget_reset_count',
        'allowed_models',
        'temp_budget_increase',
        'temp_budget_expiry',
        'created_by',
        'updated_by',
    ];

    protected $attributes = [
        'budget_reset_count' => 0,
    ];

    protected $casts = [
        'rollover_max_budget' => 'float',
        'max_budget' => 'float',
        'soft_budget' => 'float',
        'max_parallel_requests' => 'integer',
        'tpm_limit' => 'integer',
        'rpm_limit' => 'integer',
        'tpd_limit' => 'integer',
        'model_max_budget' => 'array',
        'model_tpm_limit' => 'array',
        'model_rpm_limit' => 'array',
        'budget_limits' => 'array',
        'budget_reset_at' => 'datetime',
        'budget_reset_count' => 'integer',
        'allowed_models' => AsPrismaList::class,
        'temp_budget_increase' => 'float',
        'temp_budget_expiry' => 'datetime',
    ];

    public function apps(): HasMany
    {
        return $this->hasMany(LiteLLMAppTable::class, 'budget_id');
    }

    public function app_groups(): HasMany
    {
        return $this->hasMany(LiteLLMAppGroupTable::class, 'budget_id');
    }

    public function app_group_org_budgets(): HasMany
    {
        return $this->hasMany(LiteLLMAppGroupCustomerOrgTable::class, 'budget_id');
    }

    public function app_default_tiers(): HasMany
    {
        return $this->hasMany(LiteLLMAppTable::class, 'default_tier_budget_id');
    }

    public function app_default_members(): HasMany
    {
        return $this->hasMany(LiteLLMAppTable::class, 'default_customer_member_budget_id');
    }

    public function app_org_tiers(): HasMany
    {
        return $this->hasMany(LiteLLMAppCustomerOrgTable::class, 'tier_budget_id');
    }

    public function app_org_budgets(): HasMany
    {
        return $this->hasMany(LiteLLMAppCustomerOrgTable::class, 'budget_id');
    }

    public function app_member_budgets(): HasMany
    {
        return $this->hasMany(LiteLLMAppCustomerMemberTable::class, 'budget_id');
    }

    public function app_org_overrides(): HasMany
    {
        return $this->hasMany(LiteLLMAppCustomerOrgTable::class, 'budget_override_id');
    }

    public function app_org_members(): HasMany
    {
        return $this->hasMany(LiteLLMAppCustomerOrgTable::class, 'customer_member_budget_id');
    }

    public function app_member_overrides(): HasMany
    {
        return $this->hasMany(LiteLLMAppCustomerMemberTable::class, 'budget_override_id');
    }

    public function customer_organization_defaults(): HasMany
    {
        return $this->hasMany(LiteLLMCustomerOrganizationTable::class, 'default_budget_template_id');
    }

    public function customer_organizations(): HasMany
    {
        return $this->hasMany(LiteLLMCustomerOrganizationTable::class, 'budget_id');
    }

    public function customer_teams(): HasMany
    {
        return $this->hasMany(LiteLLMCustomerTeamTable::class, 'budget_id');
    }

    public function customer_members(): HasMany
    {
        return $this->hasMany(LiteLLMCustomerMemberTable::class, 'budget_id');
    }

    public function customer_member_overrides(): HasMany
    {
        return $this->hasMany(LiteLLMCustomerMemberTable::class, 'budget_override_id');
    }

    public function customer_team_overrides(): HasMany
    {
        return $this->hasMany(LiteLLMCustomerTeamTable::class, 'budget_override_id');
    }

    public function organization(): HasMany
    {
        return $this->hasMany(LiteLLMOrganizationTable::class, 'budget_id');
    }

    public function projects(): HasMany
    {
        return $this->hasMany(LiteLLMProjectTable::class, 'budget_id');
    }

    public function keys(): HasMany
    {
        return $this->hasMany(LiteLLMVerificationToken::class, 'budget_id');
    }

    public function end_users(): HasMany
    {
        return $this->hasMany(LiteLLMEndUserTable::class, 'budget_id');
    }

    public function tags(): HasMany
    {
        return $this->hasMany(LiteLLMTagTable::class, 'budget_id');
    }

    public function model_access_groups(): HasMany
    {
        return $this->hasMany(LiteLLMModelAccessGroupBudgetTable::class, 'budget_id');
    }

    public function team_membership(): HasMany
    {
        return $this->hasMany(LiteLLMTeamMembership::class, 'budget_id');
    }

    public function organization_membership(): HasMany
    {
        return $this->hasMany(LiteLLMOrganizationMembership::class, 'budget_id');
    }
}