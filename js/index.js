import { buildDepTreeFromFiles } from "snyk-nodejs-lockfile-parser";

import fs from "fs";

function loadFile(path, data) {
  return fs.writeFileSync(path, data);
}

export { loadFile, buildDepTreeFromFiles };
