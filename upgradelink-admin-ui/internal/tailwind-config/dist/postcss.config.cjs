const { createJiti } = require("../../../node_modules/.pnpm/jiti@2.6.1/node_modules/jiti/lib/jiti.cjs")

const jiti = createJiti(__filename, {
  "interopDefault": true,
  "alias": {
    "@vben/tailwind-config": "/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/internal/tailwind-config"
  },
  "transformOptions": {
    "babel": {
      "plugins": []
    }
  }
})

/** @type {import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/internal/tailwind-config/src/postcss.config.js")} */
module.exports = jiti("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/internal/tailwind-config/src/postcss.config.ts")