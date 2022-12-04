// esbuild shims
// https://esbuild.github.io/api/#inject

// TypeError: Cannot read properties of undefined (reading 'split')\n    at node_modules/@nodelib/fs.scandir/out/constants.js
import process from "process";
process.versions.node = "12.0";
process.stdout = {isTTY: false};

const __dirname = "/";
const __filename = "index.js";

export {__dirname, __filename};
