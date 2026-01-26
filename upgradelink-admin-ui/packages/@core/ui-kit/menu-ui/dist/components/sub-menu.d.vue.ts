import type { SubMenuProps } from '../types';
interface Props extends SubMenuProps {
    isSubMenuMore?: boolean;
}
declare var __VLS_17: {}, __VLS_19: {}, __VLS_29: {}, __VLS_32: {}, __VLS_40: {};
type __VLS_Slots = {} & {
    title?: (props: typeof __VLS_17) => any;
} & {
    default?: (props: typeof __VLS_19) => any;
} & {
    content?: (props: typeof __VLS_29) => any;
} & {
    title?: (props: typeof __VLS_32) => any;
} & {
    default?: (props: typeof __VLS_40) => any;
};
declare const __VLS_base: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{}>, {
    disabled: boolean;
    isSubMenuMore: boolean;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
