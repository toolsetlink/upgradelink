import type { FormRenderProps } from '../types';
/**
 * 动态计算行数
 */
export declare function useExpandable(props: FormRenderProps): {
    isCalculated: import("vue").Ref<boolean, boolean>;
    keepFormItemIndex: import("vue").ComputedRef<number>;
    wrapperRef: Readonly<import("vue").ShallowRef<HTMLElement | null>>;
};
