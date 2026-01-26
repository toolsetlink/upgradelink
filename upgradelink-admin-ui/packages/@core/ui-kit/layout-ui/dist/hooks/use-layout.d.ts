import type { LayoutType } from '@vben-core/typings';
import type { VbenLayoutProps } from '../vben-layout';
export declare function useLayout(props: VbenLayoutProps): {
    currentLayout: import("vue").ComputedRef<LayoutType>;
    isFullContent: import("vue").ComputedRef<boolean>;
    isHeaderMixedNav: import("vue").ComputedRef<boolean>;
    isHeaderNav: import("vue").ComputedRef<boolean>;
    isMixedNav: import("vue").ComputedRef<boolean>;
    isSidebarMixedNav: import("vue").ComputedRef<boolean>;
};
