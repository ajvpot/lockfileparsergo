import { build } from "esbuild";
import plugin from "node-stdlib-browser/helpers/esbuild/plugin";
import stdLibBrowser from "node-stdlib-browser";
import { createRequire } from "module";

const require = createRequire(import.meta.url);

stdLibBrowser.inherits = require.resolve("inherits/inherits_browser.js");
stdLibBrowser.fs = require.resolve("memfs");

build({
  entryPoints: ["index.js"],
  bundle: true,
  outfile: "dist/built.js",
  platform: "node",
  format: "cjs",
  external: ["snyk-config"],
  minify: true,
  inject: [
    require.resolve("node-stdlib-browser/helpers/esbuild/shim"),
    "./esbuild-shims.js",
  ],
  define: {
    global: "global",
    process: "process",
    Buffer: "Buffer",
  },
  plugins: [plugin(stdLibBrowser)],
}).catch((e) => {
  console.error(e);
  process.exit(1)
});
