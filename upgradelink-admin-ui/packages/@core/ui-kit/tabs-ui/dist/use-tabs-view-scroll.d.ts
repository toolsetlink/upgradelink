import type { TabsProps } from './types';
export declare function useTabsViewScroll(props: TabsProps): {
    handleScrollAt: import("@vueuse/core").UseDebounceFnReturn<({ left, right }: any) => void>;
    handleWheel: ({ deltaY }: WheelEvent) => void;
    initScrollbar: () => Promise<void>;
    scrollbarRef: any;
    scrollDirection: (direction: "left" | "right", distance?: number) => void;
    scrollIsAtLeft: import("vue").Ref<boolean, boolean>;
    scrollIsAtRight: import("vue").Ref<boolean, boolean>;
    showScrollButton: import("vue").Ref<boolean, boolean>;
};
