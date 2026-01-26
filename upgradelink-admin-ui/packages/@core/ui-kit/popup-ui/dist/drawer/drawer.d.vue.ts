import type { DrawerProps, ExtendedDrawerApi } from './drawer';
interface Props extends DrawerProps {
    drawerApi?: ExtendedDrawerApi;
}
declare var __VLS_37: {}, __VLS_61: {}, __VLS_75: {}, __VLS_93: {}, __VLS_101: {}, __VLS_130: {}, __VLS_143: {}, __VLS_145: {}, __VLS_155: {}, __VLS_157: {}, __VLS_167: {}, __VLS_169: {};
type __VLS_Slots = {} & {
    'close-icon'?: (props: typeof __VLS_37) => any;
} & {
    title?: (props: typeof __VLS_61) => any;
} & {
    description?: (props: typeof __VLS_75) => any;
} & {
    extra?: (props: typeof __VLS_93) => any;
} & {
    'close-icon'?: (props: typeof __VLS_101) => any;
} & {
    default?: (props: typeof __VLS_130) => any;
} & {
    'prepend-footer'?: (props: typeof __VLS_143) => any;
} & {
    footer?: (props: typeof __VLS_145) => any;
} & {
    cancelText?: (props: typeof __VLS_155) => any;
} & {
    'center-footer'?: (props: typeof __VLS_157) => any;
} & {
    confirmText?: (props: typeof __VLS_167) => any;
} & {
    'append-footer'?: (props: typeof __VLS_169) => any;
};
declare const __VLS_base: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{}>, {
    zIndex: number;
    appendToMain: boolean;
    closeIconPlacement: import("./drawer").CloseIconPlacement;
    destroyOnClose: boolean;
    submitting: boolean;
    drawerApi: ExtendedDrawerApi;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
