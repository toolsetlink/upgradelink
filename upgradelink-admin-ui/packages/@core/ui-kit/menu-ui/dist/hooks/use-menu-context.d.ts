import type { MenuProvider, SubMenuProvider } from '../types';
/**
 * @zh_CN Provide menu context
 */
declare function createMenuContext(injectMenuData: MenuProvider): void;
/**
 * @zh_CN Provide menu context
 */
declare function createSubMenuContext(injectSubMenuData: SubMenuProvider): void;
/**
 * @zh_CN Inject menu context
 */
declare function useMenuContext(): MenuProvider;
/**
 * @zh_CN Inject menu context
 */
declare function useSubMenuContext(): SubMenuProvider;
export { createMenuContext, createSubMenuContext, useMenuContext, useSubMenuContext, };
