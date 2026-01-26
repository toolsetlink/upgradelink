interface Props {
    /**
     * 折叠区域高度
     * @default 42
     */
    collapseHeight?: number;
    /**
     * 折叠宽度
     * @default 48
     */
    collapseWidth?: number;
    /**
     * 隐藏的dom是否可见
     * @default true
     */
    domVisible?: boolean;
    /**
     * 扩展区域宽度
     */
    extraWidth: number;
    /**
     * 固定扩展区域
     * @default false
     */
    fixedExtra?: boolean;
    /**
     * 头部高度
     */
    headerHeight: number;
    /**
     * 是否侧边混合模式
     * @default false
     */
    isSidebarMixed?: boolean;
    /**
     * 顶部margin
     * @default 60
     */
    marginTop?: number;
    /**
     * 混合菜单宽度
     * @default 80
     */
    mixedWidth?: number;
    /**
     * 顶部padding
     * @default 60
     */
    paddingTop?: number;
    /**
     * 是否显示
     * @default true
     */
    show?: boolean;
    /**
     * 显示折叠按钮
     * @default true
     */
    showCollapseButton?: boolean;
    /**
     * 显示固定按钮
     * @default true
     */
    showFixedButton?: boolean;
    /**
     * 主题
     */
    theme: string;
    /**
     * 宽度
     */
    width: number;
    /**
     * zIndex
     * @default 0
     */
    zIndex?: number;
}
type __VLS_Props = Props;
type __VLS_ModelProps = {
    'collapse'?: boolean;
    'extraCollapse'?: boolean;
    'expandOnHovering'?: boolean;
    'expandOnHover'?: boolean;
    'extraVisible'?: boolean;
};
type __VLS_PublicProps = __VLS_Props & __VLS_ModelProps;
declare var __VLS_6: {}, __VLS_14: {}, __VLS_31: {}, __VLS_39: {};
type __VLS_Slots = {} & {
    logo?: (props: typeof __VLS_6) => any;
} & {
    default?: (props: typeof __VLS_14) => any;
} & {
    'extra-title'?: (props: typeof __VLS_31) => any;
} & {
    extra?: (props: typeof __VLS_39) => any;
};
declare const __VLS_base: import("vue").DefineComponent<__VLS_PublicProps, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    "update:expandOnHover": (value: boolean | undefined) => any;
    leave: () => any;
    "update:collapse": (value: boolean | undefined) => any;
    "update:extraCollapse": (value: boolean | undefined) => any;
    "update:expandOnHovering": (value: boolean | undefined) => any;
    "update:extraVisible": (value: boolean | undefined) => any;
}, string, import("vue").PublicProps, Readonly<__VLS_PublicProps> & Readonly<{
    "onUpdate:expandOnHover"?: ((value: boolean | undefined) => any) | undefined;
    onLeave?: (() => any) | undefined;
    "onUpdate:collapse"?: ((value: boolean | undefined) => any) | undefined;
    "onUpdate:extraCollapse"?: ((value: boolean | undefined) => any) | undefined;
    "onUpdate:expandOnHovering"?: ((value: boolean | undefined) => any) | undefined;
    "onUpdate:extraVisible"?: ((value: boolean | undefined) => any) | undefined;
}>, {
    zIndex: number;
    show: boolean;
    paddingTop: number;
    marginTop: number;
    collapseHeight: number;
    collapseWidth: number;
    domVisible: boolean;
    fixedExtra: boolean;
    isSidebarMixed: boolean;
    mixedWidth: number;
    showCollapseButton: boolean;
    showFixedButton: boolean;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
