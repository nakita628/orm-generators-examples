// Stand-in: only what the tables import from this file.
import { generateId } from '../shared/random/ids';
import { ORGANIZATION_ID_PREFIX } from './organizations.constants';

export function generateOrganizationId() {
  return generateId({ prefix: ORGANIZATION_ID_PREFIX });
}
