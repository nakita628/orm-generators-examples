import os
import tempfile

os.environ.setdefault("PRETALX_DATA_DIR", tempfile.mkdtemp())
from pretalx.settings import *  # noqa
