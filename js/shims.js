// runs in v8 context, no esmodules
function require(moduleName) {
  switch (moduleName) {
    case "snyk-config":
      return {
        loadConfig: () => ({
          NPM_TREE_SIZE_LIMIT: 6.0e6,
          YARN_TREE_SIZE_LIMIT: 6.0e6,
        }),
      };
    case "v8":
      return {
        serialize: (v) => JSON.stringify(v),
        deserialize: (v) => JSON.parse(v),
      };
    default:
      throw new Error("dont know how to handle " + moduleName);
  }
}

// lol
module = {};
