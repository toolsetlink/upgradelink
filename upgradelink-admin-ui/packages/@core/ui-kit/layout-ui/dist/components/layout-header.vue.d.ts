interface Props {
    /**
     * 横屏
     */
    fullWidth: boolean;
    /**
     * 高度
     */
    height: number;
    /**
     * 是否移动端
     */
    isMobile: boolean;
    /**
     * 是否显示
     */
    show: boolean;
    /**
     * 侧边菜单宽度
     */
    sidebarWidth: number;
    /**
     * 主题
     */
    theme: string | undefined;
    /**
     * 宽度
     */
    width: string;
    /**
     * zIndex
     */
    zIndex: number;
}
declare var __VLS_1: {}, __VLS_3: {}, __VLS_5: {};
type __VLS_Slots = {} & {
    logo?: (props: typeof __VLS_1) => any;
} & {
    'toggle-button'?: (props: typeof __VLS_3) => any;
} & {
    default?: (props: typeof __VLS_5) => any;
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
