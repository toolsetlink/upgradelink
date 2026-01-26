import { computed, getCurrentInstance } from "vue";
import { findComponentUpward } from "../utils/index.mjs";
function useMenu() {
  const instance = getCurrentInstance();
  if (!instance) {
    throw new Error("instance is required");
  }
  const parentPaths = computed(() => {
    let parent = instance.parent;
    const paths = [instance.props.path];
    while (parent?.type.name !== "Menu") {
      if (parent?.props.path) {
        paths.unshift(parent.props.path);
      }
      parent = parent?.parent ?? null;
    }
    return paths;
  });
  const parentMenu = computed(() => {
    return findComponentUpward(instance, ["Menu", "SubMenu"]);
  });
  return {
    parentMenu,
    parentPaths
  };
}
function useMenuStyle(menu) {
  const subMenuStyle = computed(() => {
    return {
      "--menu-level": menu ? menu?.level ?? 0 + 1 : 0
    };
  });
  return subMenuStyle;
}
export { useMenu, useMenuStyle };
