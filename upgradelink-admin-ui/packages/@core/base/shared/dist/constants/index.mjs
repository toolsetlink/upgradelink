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

/** @type {import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/packages/@core/base/shared/src/constants/index.js")} */
const _module = await jiti.import("/Users/songang/LinkProjects/open/upgradelink/upgradelink-admin-ui/packages/@core/base/shared/src/constants/index.ts");

export const CSS_VARIABLE_LAYOUT_CONTENT_HEIGHT = _module.CSS_VARIABLE_LAYOUT_CONTENT_HEIGHT;
export const CSS_VARIABLE_LAYOUT_CONTENT_WIDTH = _module.CSS_VARIABLE_LAYOUT_CONTENT_WIDTH;
export const CSS_VARIABLE_LAYOUT_HEADER_HEIGHT = _module.CSS_VARIABLE_LAYOUT_HEADER_HEIGHT;
export const CSS_VARIABLE_LAYOUT_FOOTER_HEIGHT = _module.CSS_VARIABLE_LAYOUT_FOOTER_HEIGHT;
export const ELEMENT_ID_MAIN_CONTENT = _module.ELEMENT_ID_MAIN_CONTENT;
export const DEFAULT_NAMESPACE = _module.DEFAULT_NAMESPACE;
export const VBEN_GITHUB_URL = _module.VBEN_GITHUB_URL;
export const VBEN_DOC_URL = _module.VBEN_DOC_URL;
export const VBEN_LOGO_URL = _module.VBEN_LOGO_URL;
export const VBEN_PREVIEW_URL = _module.VBEN_PREVIEW_URL;
export const VBEN_ELE_PREVIEW_URL = _module.VBEN_ELE_PREVIEW_URL;
export const VBEN_NAIVE_PREVIEW_URL = _module.VBEN_NAIVE_PREVIEW_URL;
export const VBEN_ANT_PREVIEW_URL = _module.VBEN_ANT_PREVIEW_URL;
export const VBEN_TD_PREVIEW_URL = _module.VBEN_TD_PREVIEW_URL;