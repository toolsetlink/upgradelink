import { createJiti } from "../../../node_modules/.pnpm/jiti@2.6.1/node_modules/jiti/lib/jiti.mjs";

const jiti = createJiti(import.meta.url, {
  "interopDefault": true,
  "alias": {
    "@vben/vsh": "/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/scripts/vsh"
  },
  "transformOptions": {
    "babel": {
      "plugins": []
    }
  }
})

/** @type {import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/scripts/vsh/src/index.js")} */
const _module = await jiti.import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/scripts/vsh/src/index.ts");

export default _module?.default ?? _module;