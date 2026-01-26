import type { ContentCompactType } from '@vben-core/typings';
interface Props {
    /**
     * 内容区域定宽
     */
    contentCompact: ContentCompactType;
    /**
     * 定宽布局宽度
     */
    contentCompactWidth: number;
    padding: number;
    paddingBottom: number;
    paddingLeft: number;
    paddingRight: number;
    paddingTop: number;
}
declare var __VLS_7: {}, __VLS_9: {};
type __VLS_Slots = {} & {
    overlay?: (props: typeof __VLS_7) => any;
} & {
    default?: (props: typeof __VLS_9) => any;
};
declare const __VLS_base: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{}>, {}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
