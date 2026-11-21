// Stand-in: only what the tables import from this file.
import { generateId } from '../shared/random/ids';

const SHARE_LINK_ID_PREFIX = 'dsl';

export function generateShareLinkId() {
  return generateId({ prefix: SHARE_LINK_ID_PREFIX });
}
