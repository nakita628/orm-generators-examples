# The identity this install presents to LangWatch (ADR-139, section 2).
#
# Exactly one row, minted on first use and never rewritten. It survives
# restarts, backups and hostname changes, and an install restored into
# another database mints another one, which is what the registry refuses as
# the wrong instance.
#
# It replaced an identity built from the organization name and id, which
# leaked the customer's name in cleartext, named one organization rather than
# one install, and could not be joined to a license row without splitting a
# string.
#
# id:: Fixed, so this table holds one row. A second insert loses the race on the
#   primary key and re-reads the winner's id.
# instanceId:: A random UUID. Carries nothing about the customer.
# lastReportAt:: When the usage report last reached LangWatch, and why the last attempt
#   did not. Both are what the checkup page reads to say whether reporting
#   works, which nothing recorded before: the sender logged success as soon
#   as `fetch` resolved and never looked at the response.
# optionalMetricsOptOut:: The optional category of the usage report, switched off from the checkup
#   page. With it off the report still says which release runs and how many
#   projects the install carries, and says nothing about what anyone does
#   with it (ADR-139, section 10).
# hostnameOptOut:: Hostname, which names the customer's own network and so has a switch of
#   its own rather than travelling with its category.
# upgradeDiagnosticsOptOut:: Sending a failed upgrade step's diagnostics to LangWatch on its own,
#   switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
class InstanceIdentity < ApplicationRecord
  self.table_name = "InstanceIdentity"

  attribute :id, default: "self"
  attribute :optionalMetricsOptOut, default: false
  attribute :hostnameOptOut, default: false
  attribute :upgradeDiagnosticsOptOut, default: false
  alias_attribute :created_at, :createdAt

  validates :instanceId, presence: true, uniqueness: true
end
