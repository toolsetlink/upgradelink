import type { MenuItemProps, MenuItemRegistered } from '../types';
interface Props extends MenuItemProps {
}
declare var __VLS_13: {}, __VLS_15: {}, __VLS_17: {}, __VLS_29: {}, __VLS_31: {};
type __VLS_Slots = {} & {
    default?: (props: typeof __VLS_13) => any;
} & {
    title?: (props: typeof __VLS_15) => any;
} & {
    title?: (props: typeof __VLS_17) => any;
} & {
    default?: (props: typeof __VLS_29) => any;
} & {
    title?: (props: typeof __VLS_31) => any;
};
declare const __VLS_base: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    click: (args_0: MenuItemRegistered) => any;
}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{
    onClick?: ((args_0: MenuItemRegistered) => any) | undefined;
}>, {
    disabled: boolean;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
