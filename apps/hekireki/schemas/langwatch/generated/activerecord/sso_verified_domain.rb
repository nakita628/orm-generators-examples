# better-auth's single sign-on plugin's provider table (D09), and a
# PROJECTION of the connection log rather than a second source of truth.
#
# The plugin owns the columns and the reads: it is the only thing that
# speaks OpenID Connect discovery or reads a SAML assertion, and it dials
# what it finds here. What it never does is decide that a row should exist —
# that is folded from `SsoConnection`'s events by the same projection that
# writes the connection head, in the same transaction, so a replay rebuilds
# both together and a row can never outlive the connection that justified
# it. Nothing else writes here.
#
# `providerId` is the CONNECTION ID rather than the name the customer typed.
# The plugin's column is globally unique, and two organizations both calling
# their provider `okta` must both be able to.
#
# `oidcConfig` and `samlConfig` are the plugin's own JSON documents, rebuilt
# on every fold from `SsoCredential` — the vault is where the material
# lives, and this row is a derived copy the engine can read synchronously.
#
# SEALED AT REST, under the same key the vault uses. The document carries
# the client secret the derivation reads OUT of the vault, so a plaintext
# copy here would hand every customer's live identity-provider credential
# to anybody holding a database dump — the exact threat `SsoCredential`
# exists to defeat, reopened one table over. A stored value announces its
# form (`enc:v1:`), so a row written before the seal existed still opens;
# the one reader is the identity storage adapter, on the path that dials.
# The reader tolerates both forms, so this needs no backfill — but a pod
# without it cannot read a sealed row, so it must not be rolled back past.
# WHO OWNS A PROVED DOMAIN, as a row the database can refuse twice.
#
# "First verifier owns" was enforced by reading `findDomainOwner` and then
# writing the fact — two different aggregates, so the per-connection queue
# never ordered them against each other, and `SsoConnection.verifiedDomains`
# is a `String[]` with only a GIN index behind it, which enforces nothing.
# Two organizations proving `acme.com` at the same moment both read "nobody
# owns it" and both committed; the sign-in router then picked between two
# live connections by planner order, so a member of one company could be
# handed to another company's identity provider — and with `arrivalPolicy`
# on `admit`, provisioned into it.
#
# The legacy single-domain model had this: `Organization.ssoDomain` is
# `@unique`. The multi-domain rewrite dropped the guarantee, and self-serve
# proof is what made the race reachable by customers rather than by two
# operators working at once.
#
# `domain` is the primary key, so a second organization loses on the
# constraint rather than on a check it raced. Connection holders are
# separate because an exact predecessor/replacement pair must coexist
# during migration. Both are projected with the connection head.
class SsoVerifiedDomain < ApplicationRecord
  self.table_name = "SsoVerifiedDomain"
  self.primary_key = "domain"

  validates :domain, uniqueness: { scope: :organizationId }
  validates :organizationId, presence: true

  has_many :holders, class_name: "SsoVerifiedDomainHolder", foreign_key: [ :domain, :organizationId ],
                     primary_key: [ :domain, :organizationId ], inverse_of: :ownership, dependent: :destroy
end
