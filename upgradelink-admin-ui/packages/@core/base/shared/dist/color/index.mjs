import { createJiti } from "../../../../../../node_modules/.pnpm/jiti@2.6.1/node_modules/jiti/lib/jiti.mjs";

const jiti = createJiti(import.meta.url, {
  "interopDefault": true,
  "alias": {
    "@vben-core/shared": "/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/packages/@core/base/shared"
  },
  "transformOptions": {
    "babel": {
      "plugins": []
    }
  }
})

/** @type {import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/packages/@core/base/shared/src/color/index.js")} */
const _module = await jiti.import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/packages/@core/base/shared/src/color/index.ts");

export const isDarkColor = _module.isDarkColor;
export const isLightColor = _module.isLightColor;
export const convertToHsl = _module.convertToHsl;
export const convertToHslCssVar = _module.convertToHslCssVar;
export const convertToRgb = _module.convertToRgb;
export const isValidColor = _module.isValidColor;
export const TinyColor = _module.TinyColor;
export const generatorColorVariables = _module.generatorColorVariables;