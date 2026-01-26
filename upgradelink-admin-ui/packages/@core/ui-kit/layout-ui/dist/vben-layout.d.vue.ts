import type { VbenLayoutProps } from './vben-layout';
interface Props extends VbenLayoutProps {
}
type __VLS_Props = Props;
type __VLS_ModelProps = {
    'sidebarCollapse'?: boolean;
    'sidebarExtraVisible'?: boolean;
    'sidebarExtraCollapse'?: boolean;
    'sidebarExpandOnHover'?: boolean;
    'sidebarEnable'?: boolean;
};
type __VLS_PublicProps = __VLS_Props & __VLS_ModelProps;
declare var __VLS_10: {}, __VLS_12: {}, __VLS_14: {}, __VLS_17: {}, __VLS_20: {}, __VLS_29: {}, __VLS_50: {}, __VLS_58: {}, __VLS_66: {}, __VLS_69: {}, __VLS_77: {}, __VLS_79: {};
type __VLS_Slots = {} & {
    logo?: (props: typeof __VLS_10) => any;
} & {
    'mixed-menu'?: (props: typeof __VLS_12) => any;
} & {
    menu?: (props: typeof __VLS_14) => any;
} & {
    'side-extra'?: (props: typeof __VLS_17) => any;
} & {
    'side-extra-title'?: (props: typeof __VLS_20) => any;
} & {
    logo?: (props: typeof __VLS_29) => any;
} & {
    header?: (props: typeof __VLS_50) => any;
} & {
    tabbar?: (props: typeof __VLS_58) => any;
} & {
    content?: (props: typeof __VLS_66) => any;
} & {
    'content-overlay'?: (props: typeof __VLS_69) => any;
} & {
    footer?: (props: typeof __VLS_77) => any;
} & {
    extra?: (props: typeof __VLS_79) => any;
};
declare const __VLS_base: import("vue").DefineComponent<__VLS_PublicProps, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    sideMouseLeave: () => any;
    toggleSidebar: () => any;
    "update:sidebarCollapse": (value: boolean) => any;
    "update:sidebarExtraVisible": (value: boolean | undefined) => any;
    "update:sidebarExtraCollapse": (value: boolean) => any;
    "update:sidebarExpandOnHover": (value: boolean) => any;
    "update:sidebarEnable": (value: boolean) => any;
}, string, import("vue").PublicProps, Readonly<__VLS_PublicProps> & Readonly<{
    onSideMouseLeave?: (() => any) | undefined;
    onToggleSidebar?: (() => any) | undefined;
    "onUpdate:sidebarCollapse"?: ((value: boolean) => any) | undefined;
    "onUpdate:sidebarExtraVisible"?: ((value: boolean | undefined) => any) | undefined;
    "onUpdate:sidebarExtraCollapse"?: ((value: boolean) => any) | undefined;
    "onUpdate:sidebarExpandOnHover"?: ((value: boolean) => any) | undefined;
    "onUpdate:sidebarEnable"?: ((value: boolean) => any) | undefined;
}>, {
    zIndex: number;
    layout: import("@vben-core/typings").LayoutType;
    contentCompact: import("@vben-core/typings").ContentCompactType;
    contentCompactWidth: number;
    isMobile: boolean;
    sidebarWidth: number;
    headerHeight: number;
    contentPadding: number;
    contentPaddingBottom: number;
    contentPaddingLeft: number;
    contentPaddingRight: number;
    contentPaddingTop: number;
    footerEnable: boolean;
    footerFixed: boolean;
    footerHeight: number;
    headerHidden: boolean;
    headerMode: import("@vben-core/typings").LayoutHeaderModeType;
    headerToggleSidebarButton: boolean;
    headerVisible: boolean;
    sidebarCollapsedButton: boolean;
    sidebarCollapseShowTitle: boolean;
    sidebarExtraCollapsedWidth: number;
    sidebarFixedButton: boolean;
    sidebarHidden: boolean;
    sidebarMixedWidth: number;
    sidebarTheme: import("@vben-core/typings").ThemeModeType;
    sideCollapseWidth: number;
    tabbarEnable: boolean;
    tabbarHeight: number;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
