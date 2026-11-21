"""Write the list of examples into README.md, between the intro and "## How".

    uv run scripts/readme.py           # write it
    uv run scripts/readme.py --check   # fail when README.md is not what it would write

Every directory of examples/ must be in GROUPS, and every entry of GROUPS a directory.
"""

from __future__ import annotations

import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
GH = "https://github.com/"
GROUPS: dict[str, dict[str, tuple[str, str]]] = {
    "Tutorials (models written here)": {
        "tutorial-polls": ("Django tutorial: polls", "https://docs.djangoproject.com/en/stable/intro/tutorial02/"),
        "docs-relations": (
            "Django docs: many-to-one, many-to-many, one-to-one",
            "https://docs.djangoproject.com/en/stable/topics/db/examples/",
        ),
        "tutorial-django-girls": ("Django Girls: blog", "https://tutorial.djangogirls.org/en/django_models/"),
        "tutorial-local-library": (
            "MDN: LocalLibrary",
            "https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Server-side/Django/Models",
        ),
        "tutorial-drf-snippets": (
            "Django REST framework tutorial: snippets",
            "https://www.django-rest-framework.org/tutorial/1-serialization/",
        ),
        "docs-models": (
            "Django docs: models (Musician, Membership, inheritance)",
            "https://docs.djangoproject.com/en/stable/topics/db/models/",
        ),
        "docs-queries": (
            "Django docs: making queries and aggregation",
            "https://docs.djangoproject.com/en/stable/topics/db/queries/",
        ),
    },
    "Domains (models written here)": {
        "domain-shop": (
            "Online shop: catalog, carts, orders, payments",
            "https://docs.djangoproject.com/en/stable/topics/db/models/",
        ),
        "domain-saas": (
            "SaaS: organizations, roles, plans, subscriptions, invoices",
            "https://docs.djangoproject.com/en/stable/topics/auth/customizing/",
        ),
        "domain-booking": (
            "Hotel bookings on PostgreSQL: no overlapping stays",
            "https://docs.djangoproject.com/en/stable/ref/contrib/postgres/constraints/",
        ),
        "domain-lms": (
            "Learning platform: courses, lessons, quizzes, certificates",
            "https://docs.djangoproject.com/en/stable/topics/db/models/",
        ),
        "domain-social": (
            "Social network: follows, posts, likes, notifications",
            "https://docs.djangoproject.com/en/stable/ref/contrib/contenttypes/",
        ),
        "domain-tracker": (
            "Issue tracker: workflows, issues, sprints, work logs",
            "https://docs.djangoproject.com/en/stable/topics/db/models/",
        ),
        "domain-chat": (
            "Chat: workspaces, channels, threads, reactions",
            "https://docs.djangoproject.com/en/stable/topics/db/models/",
        ),
        "domain-inventory": (
            "Inventory: warehouses, stock, movements, purchase orders",
            "https://docs.djangoproject.com/en/stable/ref/models/constraints/",
        ),
        "domain-hr": (
            "HR: departments, employees, salaries, leave, reviews",
            "https://docs.djangoproject.com/en/stable/topics/db/models/",
        ),
    },
    "Django features (models written here)": {
        "features-fields": ("Every field type", "https://docs.djangoproject.com/en/stable/ref/models/fields/"),
        "features-relations": (
            "Every kind of relation",
            "https://docs.djangoproject.com/en/stable/topics/db/models/#relationships",
        ),
        "features-inheritance": (
            "Inheritance and a custom user",
            "https://docs.djangoproject.com/en/stable/topics/db/models/#model-inheritance",
        ),
        "features-constraints": (
            "Constraints, indexes, table options",
            "https://docs.djangoproject.com/en/stable/ref/models/constraints/",
        ),
        "features-postgres": (
            "django.contrib.postgres",
            "https://docs.djangoproject.com/en/stable/ref/contrib/postgres/",
        ),
        "features-mysql": (
            "MySQL: its types, collations and comments",
            "https://docs.djangoproject.com/en/stable/ref/databases/#mysql-notes",
        ),
    },
    "Applications": {
        "nautobot": ("Nautobot", GH + "nautobot/nautobot"),
        "pulpcore": ("Pulp", GH + "pulp/pulpcore"),
        "pulp-plugins": ("Pulp with its container, Python and Ansible plugins", GH + "pulp/pulp_container"),
        "review-board": ("Review Board", GH + "reviewboard/reviewboard"),
        "weblate": ("Weblate", GH + "WeblateOrg/weblate"),
        "mayan-edms": ("Mayan EDMS", "https://gitlab.com/mayan-edms/mayan-edms"),
        "kiwitcms": ("Kiwi TCMS", GH + "kiwitcms/Kiwi"),
        "misago": ("Misago", GH + "rafalp/Misago"),
        "allianceauth": ("Alliance Auth", "https://gitlab.com/allianceauth/allianceauth"),
        "devilry": ("Devilry", GH + "devilry/devilry-django"),
        "django-crm": ("Django CRM", GH + "DjangoCRM/django-crm"),
        "pretix": ("pretix", GH + "pretix/pretix"),
        "pretalx": ("pretalx", GH + "pretalx/pretalx"),
        "label-studio": ("Label Studio", GH + "HumanSignal/label-studio"),
        "mailman-web": ("Mailman 3 web (Postorius, HyperKitty)", "https://gitlab.com/mailman/mailman-web"),
        "django-oscar": ("Django Oscar", GH + "django-oscar/django-oscar"),
        "django-ledger": ("Django Ledger", GH + "arrobalytics/django-ledger"),
        "django-helpdesk": ("django-helpdesk", GH + "django-helpdesk/django-helpdesk"),
        "django-wiki": ("django-wiki", GH + "django-wiki/django-wiki"),
        "django-machina": ("django-machina", GH + "ellmetha/django-machina"),
        "spirit": ("Spirit", GH + "nitely/Spirit"),
        "django-sql-explorer": ("SQL Explorer", GH + "explorerhq/sql-explorer"),
        "openwisp-users": ("OpenWISP Users", GH + "openwisp/openwisp-users"),
        "django-survey-and-report": ("django-survey", GH + "Pierre-Sassoulas/django-survey"),
        "openedx-core": ("Open edX Core (openedx-learning)", GH + "openedx/openedx-core"),
        "edx-organizations": ("edx-organizations", GH + "openedx/edx-organizations"),
        "openwisp-radius": ("OpenWISP RADIUS", GH + "openwisp/openwisp-radius"),
        "openwisp-ipam": ("OpenWISP IPAM", GH + "openwisp/openwisp-ipam"),
        "openwisp-notifications": ("OpenWISP Notifications", GH + "openwisp/openwisp-notifications"),
        "django-appointment": ("django-appointment", GH + "adamspd/django-appointment"),
        "edx-completion": ("edx-completion", GH + "openedx/completion"),
        "edx-milestones": ("edx-milestones", GH + "openedx/edx-milestones"),
        "edx-when": ("edx-when", GH + "openedx/edx-when"),
        "edx-submissions": ("edx-submissions", GH + "openedx/edx-submissions"),
        "openedx-ledger": ("openedx-ledger", GH + "openedx/openedx-ledger"),
        "edx-name-affirmation": ("edx-name-affirmation", GH + "openedx/edx-name-affirmation"),
        "edx-celeryutils": ("edx-celeryutils", GH + "openedx/edx-celeryutils"),
        "django-todo": ("django-todo", GH + "shacker/django-todo"),
        "django-report-builder": ("Django Report Builder", "https://gitlab.com/burke-software/django-report-builder"),
        "data-wizard": ("Django Data Wizard", GH + "wq/django-data-wizard"),
        "django-ai-assistant": ("Django AI Assistant", GH + "vintasoftware/django-ai-assistant"),
    },
    "CMS": {
        "wagtail": ("Wagtail", GH + "wagtail/wagtail"),
        "wagtail-localize": ("Wagtail Localize", GH + "wagtail/wagtail-localize"),
        "wagtailmenus": ("wagtailmenus", GH + "jazzband/wagtailmenus"),
        "wagtail-ab-testing": ("Wagtail A/B Testing", GH + "wagtail-nest/wagtail-ab-testing"),
        "wagtailmedia": ("wagtailmedia", GH + "torchbox/wagtailmedia"),
        "wagtail-transfer": ("Wagtail Transfer", GH + "wagtail/wagtail-transfer"),
        "wagtailstreamforms": ("wagtailstreamforms", GH + "labd/wagtailstreamforms"),
        "wagtail-headless-preview": ("wagtail-headless-preview", GH + "torchbox/wagtail-headless-preview"),
        "wagtail-newsletter": ("wagtail-newsletter", GH + "wagtail/wagtail-newsletter"),
        "wagtail-personalisation": ("wagtail-personalisation", GH + "wagtail/wagtail-personalisation"),
        "wagtail-localize-git": ("wagtail-localize-git", GH + "wagtail/wagtail-localize-git"),
        "wagtail-flags": ("Wagtail-Flags", GH + "cfpb/wagtail-flags"),
        "wagtail-2fa": ("wagtail-2fa", GH + "labd/wagtail-2fa"),
        "coderedcms": ("Wagtail CRX (CodeRed CMS)", GH + "coderedcorp/coderedcms"),
        "puput": ("Puput", GH + "APSL/puput"),
        "django-cms": ("django CMS", GH + "django-cms/django-cms"),
        "djangocms-blog": ("djangocms-blog", GH + "django-cms/djangocms-blog"),
        "djangocms-versioning": ("djangocms-versioning", GH + "django-cms/djangocms-versioning"),
        "djangocms-plugins": (
            "django CMS plugins (alias, snippet, picture, file, video)",
            GH + "django-cms/djangocms-alias",
        ),
        "django-cms-qe": ("django CMS QE", GH + "CZ-NIC/django-cms-qe"),
        "djangocms-moderation": ("djangocms-moderation", GH + "django-cms/djangocms-moderation"),
        "djangocms-stories": ("djangocms-stories", GH + "django-cms/djangocms-stories"),
        "feincms": ("FeinCMS", GH + "feincms/feincms"),
        "django-flatblocks": ("django-flatblocks", GH + "funkybob/django-flatblocks"),
        "django-moderation": ("django-moderation", GH + "dominno/django-moderation"),
        "django-fobi": ("django-fobi", GH + "barseghyanartur/django-fobi"),
        "django-filer": ("django-filer", GH + "django-cms/django-filer"),
        "mezzanine": ("Mezzanine", GH + "stephenmcd/mezzanine"),
        "django-dbtemplates": ("django-dbtemplates", GH + "jazzband/django-dbtemplates"),
        "django-summernote": ("django-summernote", GH + "lqez/django-summernote"),
        "django-photologue": ("django-photologue", GH + "richardbarran/django-photologue"),
        "sorl-thumbnail": ("sorl-thumbnail", GH + "jazzband/sorl-thumbnail"),
        "easy-thumbnails": ("easy-thumbnails", GH + "SmileyChris/easy-thumbnails"),
    },
    "Django": {
        "django-contrib": ("Django (contrib apps)", "https://docs.djangoproject.com/en/stable/ref/contrib/"),
    },
    "Authentication and security": {
        "django-allauth": ("django-allauth", GH + "pennersr/django-allauth"),
        "social-auth-app-django": ("Python Social Auth (Django)", GH + "python-social-auth/social-app-django"),
        "django-oauth-toolkit": ("Django OAuth Toolkit", GH + "jazzband/django-oauth-toolkit"),
        "django-oidc-provider": ("django-oidc-provider", GH + "juanifioren/django-oidc-provider"),
        "django-cas-server": ("django-cas-server", GH + "nitmir/django-cas-server"),
        "django-lti": ("django-lti", GH + "academic-innovation/django-lti"),
        "django-otp": ("django-otp / django-two-factor-auth", GH + "django-otp/django-otp"),
        "django-otp-webauthn": ("django-otp-webauthn", GH + "Stormbase/django-otp-webauthn"),
        "django-otp-yubikey": ("django-otp-yubikey", GH + "django-otp/django-otp-yubikey"),
        "django-otp-twilio": ("django-otp-twilio", GH + "django-otp/django-otp-twilio"),
        "django-trench": ("django-trench", GH + "merixstudio/django-trench"),
        "django-cas-ng": ("django-cas-ng", GH + "django-cas-ng/django-cas-ng"),
        "django-simple-sso": ("django-simple-sso", GH + "django-cms/django-simple-sso"),
        "django-microsoft-auth": ("django-microsoft-auth", GH + "AngellusMortis/django_microsoft_auth"),
        "django-scim2": ("django-scim2", GH + "15five/django-scim2"),
        "django-esi": ("django-esi (EVE Online)", "https://gitlab.com/allianceauth/django-esi"),
        "django-user-visit": ("django-user-visit", GH + "yunojuno/django-user-visit"),
        "django-user-accounts": ("django-user-accounts (Pinax)", GH + "pinax/django-user-accounts"),
        "django-password-validators": ("django-password-validators", GH + "fizista/django-password-validators"),
        "django-admin-honeypot": ("django-admin-honeypot", GH + "dmpayton/django-admin-honeypot"),
        "django-simple-captcha": ("django-simple-captcha", GH + "mbi/django-simple-captcha"),
        "django-phone-verify": ("django-phone-verify", GH + "CuriousLearner/django-phone-verify"),
        "django-termsandconditions": ("django-termsandconditions", GH + "cyface/django-termsandconditions"),
        "django-mfa3": ("django-mfa3", GH + "xi/django-mfa3"),
        "django-passkeys": ("django-passkeys", GH + "mkalioby/django-passkeys"),
        "django-registration-redux": ("django-registration-redux", GH + "macropin/django-registration"),
        "django-guardian": ("django-guardian", GH + "django-guardian/django-guardian"),
        "django-axes": ("django-axes", GH + "jazzband/django-axes"),
        "django-defender": ("django-defender", GH + "jazzband/django-defender"),
        "django-impersonate": ("django-impersonate", "https://codeberg.org/iamdefinitelyahuman/django-impersonate"),
        "django-user-sessions": ("django-user-sessions", GH + "jazzband/django-user-sessions"),
        "django-invitations": ("django-invitations", GH + "jazzband/django-invitations"),
        "django-organizations": ("django-organizations", GH + "bennylope/django-organizations"),
        "django-gdpr-assist": ("django-gdpr-assist", GH + "wildfish/django-gdpr-assist"),
    },
    "APIs": {
        "django-rest-framework": ("Django REST framework", GH + "encode/django-rest-framework"),
        "djangorestframework-simplejwt": ("Simple JWT", GH + "jazzband/djangorestframework-simplejwt"),
        "djangorestframework-api-key": (
            "Django REST Framework API Key",
            GH + "florimondmanca/djangorestframework-api-key",
        ),
        "django-rest-knox": ("django-rest-knox", GH + "jazzband/django-rest-knox"),
        "django-rest-passwordreset": ("django-rest-passwordreset", GH + "anexia-it/django-rest-passwordreset"),
        "drfpasswordless": ("drfpasswordless", GH + "aaronn/django-rest-framework-passwordless"),
        "django-rest-durin": ("django-rest-durin", GH + "eshaan7/django-rest-durin"),
        "django-ninja-jwt": ("Django Ninja JWT", GH + "eadwinCode/django-ninja-jwt"),
        "django-graphql-jwt": ("Django GraphQL JWT", GH + "flavors/django-graphql-jwt"),
        "drf-api-tracking": ("drf-api-tracking", GH + "lingster/drf-api-tracking"),
        "drf-api-logger": ("DRF API Logger", GH + "vishalanandl177/DRF-API-Logger"),
        "django-webhook": ("django-webhook", GH + "danihodovic/django-webhook"),
        "django-eventstream": ("django-eventstream", GH + "fanout/django-eventstream"),
    },
    "Payments and subscriptions": {
        "dj-stripe": ("dj-stripe", GH + "dj-stripe/dj-stripe"),
        "django-paypal": ("django-paypal", GH + "spookylukey/django-paypal"),
        "django-plans": ("django-plans", GH + "django-getpaid/django-plans"),
        "django-salesman": ("Salesman", GH + "dinoperovic/django-salesman"),
        "django-oscar-accounts": ("django-oscar-accounts", GH + "django-oscar/django-oscar-accounts"),
        "django-hordak": ("Hordak", GH + "adamcharnock/django-hordak"),
        "django-flexible-subscriptions": (
            "django-flexible-subscriptions",
            GH + "studybuffalo/django-flexible-subscriptions",
        ),
        "drf-stripe-subscription": ("drf-stripe-subscription", GH + "oscarychen/drf-stripe-subscription"),
    },
    "Tasks and workflows": {
        "django-celery": ("django-celery-beat / django-celery-results", GH + "celery/django-celery-beat"),
        "django-q2": ("Django Q2", GH + "django-q2/django-q2"),
        "django-tasks": ("django-tasks (DEP 14)", GH + "RealOrangeOne/django-tasks"),
        "procrastinate": ("Procrastinate", GH + "procrastinate-org/procrastinate"),
        "django-tasks-scheduler": ("django-tasks-scheduler", GH + "django-commons/django-tasks-scheduler"),
        "django-chroniker": ("django-chroniker", GH + "chrisspen/django-chroniker"),
        "django-huey-monitor": ("django-huey-monitor", GH + "boxine/django-huey-monitor"),
        "django-db-queue": ("django-db-queue", GH + "dabapps/django-db-queue"),
        "django-import-export-celery": ("django-import-export-celery", GH + "auto-mat/django-import-export-celery"),
        "django-pgpubsub": ("django-pgpubsub", GH + "PaulGilmartin/django-pgpubsub"),
        "django-dramatiq": ("django_dramatiq", GH + "Bogdanp/django_dramatiq"),
        "django-background-tasks": ("django-background-tasks", GH + "arteria/django-background-tasks"),
        "django-viewflow": ("Viewflow", GH + "viewflow/viewflow"),
        "django-fsm-log": ("django-fsm-log", GH + "jazzband/django-fsm-log"),
        "django-scheduler": ("django-scheduler", GH + "llazzaro/django-scheduler"),
    },
    "History and auditing": {
        "django-reversion": ("django-reversion", GH + "etianen/django-reversion"),
        "django-simple-history": ("django-simple-history", GH + "jazzband/django-simple-history"),
        "django-pghistory": ("django-pghistory", GH + "Opus10/django-pghistory"),
        "django-field-history": ("django-field-history", GH + "grantmcconnaughey/django-field-history"),
        "django-auditlog": ("django-auditlog", GH + "jazzband/django-auditlog"),
        "django-easy-audit": ("django-easy-audit", GH + "soynatan/django-easy-audit"),
        "django-softdelete": ("django-softdelete", GH + "scoursen/django-softdelete"),
        "django-request": ("django-request", GH + "django-request/django-request"),
        "django-tracking2": ("django-tracking2", GH + "bruth/django-tracking2"),
        "django-silk": ("django-silk", GH + "jazzband/django-silk"),
        "django-request-profiler": ("django-request-profiler", GH + "yunojuno/django-request-profiler"),
        "django-db-logger": ("django-db-logger", GH + "CiCiApp/django-db-logger"),
        "django-log-outgoing-requests": (
            "django-log-outgoing-requests",
            GH + "maykinmedia/django-log-outgoing-requests",
        ),
        "django-sql-dashboard": ("Django SQL Dashboard", GH + "simonw/django-sql-dashboard"),
    },
    "Content and community": {
        "django-taggit": ("django-taggit", GH + "jazzband/django-taggit"),
        "django-categories": ("django-categories", GH + "jazzband/django-categories"),
        "django-comments": ("django-contrib-comments", GH + "django/django-contrib-comments"),
        "django-comments-xtd": ("django-comments-xtd", GH + "danirus/django-comments-xtd"),
        "django-star-ratings": ("django-star-ratings", GH + "wildfish/django-star-ratings"),
        "django-vote": ("django-vote", GH + "shanbay/django-vote"),
        "django-attachments": ("django-attachments", GH + "jazzband/django-attachments"),
        "django-file-form": ("django-file-form", GH + "mbraak/django-file-form"),
        "django-drf-filepond": ("django-drf-filepond", GH + "ImperialCollegeLondon/django-drf-filepond"),
        "django-db-file-storage": ("django-db-file-storage", GH + "victor-o-silva/db_file_storage"),
        "pinax": ("Pinax (announcements, documents, events, images, invitations, teams)", GH + "pinax"),
        "django-private-chat2": ("django-private-chat2", GH + "Bearle/django_private_chat2"),
        "django-activity-stream": ("django-activity-stream", GH + "justquick/django-activity-stream"),
        "django-friendship": ("django-friendship", GH + "revsys/django-friendship"),
        "django-postman": ("django-postman", "https://gitlab.com/psam/django-postman"),
        "django-newsletter": ("django-newsletter", GH + "jazzband/django-newsletter"),
        "django-avatar": ("django-avatar", GH + "jazzband/django-avatar"),
        "django-hitcount": ("django-hitcount", GH + "thornomad/django-hitcount"),
        "django-watson": ("django-watson", GH + "etianen/django-watson"),
        "django-sitetree": ("django-sitetree", GH + "idlesign/django-sitetree"),
        "django-robots": ("django-robots", GH + "jazzband/django-robots"),
        "django-address": ("django-address", GH + "furious-luke/django-address"),
        "django-cities-light": ("django-cities-light", GH + "yourlabs/django-cities-light"),
        "django-countries-plus": ("django-countries-plus", GH + "cordery/django-countries-plus"),
        "django-languages-plus": ("django-languages-plus", GH + "cordery/django-languages-plus"),
    },
    "Admin": {
        "django-admin-interface": ("django-admin-interface", GH + "fabiocaccamo/django-admin-interface"),
        "django-admin-tools": ("django-admin-tools", GH + "django-admin-tools/django-admin-tools"),
        "django-jet-reboot": ("Django JET Reboot", GH + "assem-ch/django-jet-reboot"),
    },
    "Configuration and messaging": {
        "django-constance": ("django-constance", GH + "jazzband/django-constance"),
        "django-dynamic-preferences": ("django-dynamic-preferences", GH + "agateblue/django-dynamic-preferences"),
        "django-waffle": ("django-waffle", GH + "django-waffle/django-waffle"),
        "django-flags": ("django-flags", GH + "cfpb/django-flags"),
        "django-cookie-consent": ("django-cookie-consent", GH + "django-commons/django-cookie-consent"),
        "django-post-office": ("django-post-office", GH + "ui/django-post-office"),
        "django-mailer": ("django-mailer", GH + "pinax/django-mailer"),
        "django-mail-queue": ("django-mail-queue", GH + "dstegelman/django-mail-queue"),
        "django-ses": ("django-ses", GH + "django-ses/django-ses"),
        "django-herald": ("django-herald", GH + "worthwhile/django-herald"),
        "django-mailbox": ("django-mailbox", GH + "coddingtonbear/django-mailbox"),
        "django-push-notifications": ("django-push-notifications", GH + "jazzband/django-push-notifications"),
        "fcm-django": ("FCM Django", GH + "xtrinch/fcm-django"),
        "django-webpush": ("django-webpush", GH + "safwanrahman/django-webpush"),
        "django-twilio": ("django-twilio", GH + "rdegges/django-twilio"),
    },
}
# A sample's models: the file the Models column links to.
SAMPLES = {
    "tutorial-polls": "polls/models.py",
    "tutorial-django-girls": "blog/models.py",
    "tutorial-local-library": "catalog/models.py",
    "tutorial-drf-snippets": "snippets/models.py",
    "docs-relations": "many_to_one/models.py",
    "features-fields": "fields/models.py",
    "features-relations": "relations/models.py",
    "features-inheritance": "inheritance/models.py",
    "features-constraints": "constraints/models.py",
    "features-postgres": "pg/models.py",
    "features-mysql": "store/models.py",
    "docs-models": "musicians/models.py",
    "docs-queries": "queries/models.py",
    "domain-shop": "catalog/models.py",
    "domain-saas": "accounts/models.py",
    "domain-booking": "hotels/models.py",
    "domain-lms": "courses/models.py",
    "domain-social": "social/models.py",
    "domain-tracker": "tracker/models.py",
    "domain-chat": "chat/models.py",
    "domain-inventory": "inventory/models.py",
    "domain-hr": "people/models.py",
}
ISSUE_POSTGRES = "https://github.com/nakita628/djangoer/issues/12"


def main() -> None:
    names = [name for group in GROUPS.values() for name in group]
    on_disk = sorted(p.name for p in Path(ROOT, "examples").iterdir() if p.is_dir())
    missing, extra = set(on_disk) - set(names), set(names) - set(on_disk)
    assert not missing and not extra, f"not in GROUPS: {missing}; no directory: {extra}"
    out = []
    for title, rows in GROUPS.items():
        out += [
            f"### {title}",
            "",
            "| Project | Models | ER diagram | Docs | DBML |",
            "| --- | --- | --- | --- | --- |",
        ]
        for d, (name, url) in rows.items():
            e = f"examples/{d}"
            docs = (
                f"[docs]({e}/docs/README.md)"
                if Path(ROOT, e, "docs/README.md").exists()
                else f"— ([#12]({ISSUE_POSTGRES}))"
            )
            models = f"[models]({e}/models)" if Path(ROOT, e, "models").is_dir() else f"[models.py]({e}/{SAMPLES[d]})"
            out.append(
                f"| [{name}]({url}) | {models} | [er.md]({e}/er.md) · [er.svg]({e}/er.svg) | {docs} | [schema.dbml]({e}/schema.dbml) |"
            )
        out.append("")
    readme = ROOT / "README.md"
    text = readme.read_text()
    start, end = text.index("projects:\n"), text.index("## How")
    start = text.rindex("\n", 0, start) + 1
    new = text[:start] + f"{len(names)} projects:\n\n" + "\n".join(out) + "\n" + text[end:]
    if "--check" in sys.argv[1:]:
        if new != text:
            sys.exit("README.md is stale: run `uv run scripts/readme.py` in apps/djangoer")
        print(f"{len(names)} projects, README.md is up to date")
        return
    readme.write_text(new)
    print(f"{len(names)} projects")


if __name__ == "__main__":
    main()
