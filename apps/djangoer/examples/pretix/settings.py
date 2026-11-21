import os
import tempfile

os.environ.setdefault("PRETIX_DATA_DIR", tempfile.mkdtemp())
from pretix.settings import *  # noqa
