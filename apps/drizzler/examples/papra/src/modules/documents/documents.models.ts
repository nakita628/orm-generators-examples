// Stand-in: only what the tables import from this file.
import { generateId } from '../shared/random/ids';

export function generateDocumentId() {
  return generateId({ prefix: 'doc' });
}
