import { getCurrentInstance, inject, provide } from "vue";
import { findComponentUpward } from "../utils/index.mjs";
const menuContextKey = Symbol("menuContext");
function createMenuContext(injectMenuData) {
  provide(menuContextKey, injectMenuData);
}
function createSubMenuContext(injectSubMenuData) {
  const instance = getCurrentInstance();
  provide(`subMenu:${instance?.uid}`, injectSubMenuData);
}
function useMenuContext() {
  const instance = getCurrentInstance();
  if (!instance) {
    throw new Error("instance is required");
  }
  const rootMenu = inject(menuContextKey);
  return rootMenu;
}
function useSubMenuContext() {
  const instance = getCurrentInstance();
  if (!instance) {
    throw new Error("instance is required");
  }
  const parentMenu = findComponentUpward(instance, ["Menu", "SubMenu"]);
  const subMenu = inject(`subMenu:${parentMenu?.uid}`);
  return subMenu;
}
export {
  createMenuContext,
  createSubMenuContext,
  useMenuContext,
  useSubMenuContext
};
