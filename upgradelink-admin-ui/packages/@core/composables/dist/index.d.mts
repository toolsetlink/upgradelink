import * as vue from 'vue';
import { CSSProperties, Ref, ComputedRef } from 'vue';
import { VisibleDomRect } from '@vben-core/shared/utils';
import Sortable, { SortableOptions } from 'sortablejs';
export { default as Sortable } from 'sortablejs';
export { useEmitAsProps, useForwardExpose, useForwardProps, useForwardPropsEmits } from 'reka-ui';

declare function useIsMobile(): {
    isMobile: vue.ComputedRef<boolean>;
};

declare function useLayoutContentStyle(): {
    contentElement: vue.Ref<HTMLDivElement | null, HTMLDivElement | null>;
    overlayStyle: vue.ComputedRef<CSSProperties>;
    visibleDomRect: vue.Ref<{
        bottom: number;
        height: number;
        left: number;
        right: number;
        top: number;
        width: number;
    } | null, VisibleDomRect | {
        bottom: number;
        height: number;
        left: number;
        right: number;
        top: number;
        width: number;
    } | null>;
};
declare function useLayoutHeaderStyle(): {
    getLayoutHeaderHeight: () => number;
    setLayoutHeaderHeight: (height: number) => void;
};
declare function useLayoutFooterStyle(): {
    getLayoutFooterHeight: () => number;
    setLayoutFooterHeight: (height: number) => void;
};

declare const useNamespace: (block: string) => {
    b: (blockSuffix?: string) => string;
    be: (blockSuffix?: string, element?: string) => string;
    bem: (blockSuffix?: string, element?: string, modifier?: string) => string;
    bm: (blockSuffix?: string, modifier?: string) => string;
    cssVar: (object: Record<string, string>) => Record<string, string>;
    cssVarBlock: (object: Record<string, string>) => Record<string, string>;
    cssVarBlockName: (name: string) => string;
    cssVarName: (name: string) => string;
    e: (element?: string) => string;
    em: (element?: string, modifier?: string) => string;
    is: {
        (name: string): string;
        (name: string, state: boolean | undefined): string;
    };
    m: (modifier?: string) => string;
    namespace: string;
};
type UseNamespaceReturn = ReturnType<typeof useNamespace>;

declare function usePriorityValue<T extends Record<string, any>, S extends Record<string, any>, K extends keyof T = keyof T>(key: K, props: T, state: Readonly<Ref<NoInfer<S>>> | undefined): ComputedRef<T[K]>;
declare function usePriorityValues<T extends Record<string, any>, S extends Ref<Record<string, any>> = Readonly<Ref<NoInfer<T>, NoInfer<T>>>>(props: T, state: S | undefined): { [K in keyof T]: ComputedRef<T[K]>; };
declare function useForwardPriorityValues<T extends Record<string, any>, S extends Ref<Record<string, any>> = Readonly<Ref<NoInfer<T>, NoInfer<T>>>>(props: T, state: S | undefined): ComputedRef<{ [K in keyof T]: T[K]; }>;

declare const SCROLL_FIXED_CLASS = "_scroll__fixed_";
declare function useScrollLock(): void;

type Locale = 'en-US' | 'zh-CN';

declare const useSimpleLocale: () => {
    $t: vue.ComputedRef<(key: string) => string>;
    currentLocale: vue.Ref<Locale, Locale>;
    setSimpleLocale: (locale: Locale) => void;
};

declare function useSortable<T extends HTMLElement>(sortableContainer: T, options?: SortableOptions): {
    initializeSortable: () => Promise<Sortable>;
};

export { SCROLL_FIXED_CLASS, useForwardPriorityValues, useIsMobile, useLayoutContentStyle, useLayoutFooterStyle, useLayoutHeaderStyle, useNamespace, usePriorityValue, usePriorityValues, useScrollLock, useSimpleLocale, useSortable };
export type { UseNamespaceReturn };
