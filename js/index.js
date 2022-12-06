import { buildDepTreeFromFiles as buildJavascriptDepTreeFromFiles } from "snyk-nodejs-lockfile-parser";
import { buildDepGraph as buildPoetryDepTree } from "snyk-poetry-lockfile-parser";
import { buildDepTreeFromFiles as buildComposerDepTreeFromFiles } from "@snyk/composer-lockfile-parser";

import fs from "fs";

export { fs, buildJavascriptDepTreeFromFiles, buildPoetryDepTree, buildComposerDepTreeFromFiles };
