import type { SubMenuProvider } from '../types';
declare function useMenu(): {
    parentMenu: import("vue").ComputedRef<import("vue").ComponentInternalInstance | null>;
    parentPaths: import("vue").ComputedRef<string[]>;
};
declare function useMenuStyle(menu?: SubMenuProvider): import("vue").ComputedRef<{
    '--menu-level': number;
}>;
export { useMenu, useMenuStyle };
