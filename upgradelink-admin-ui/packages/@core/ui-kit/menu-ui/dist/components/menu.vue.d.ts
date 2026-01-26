import type { MenuProps } from '../types';
interface Props extends MenuProps {
}
declare var __VLS_23: {};
type __VLS_Slots = {} & {
    default?: (props: typeof __VLS_23) => any;
};
declare const __VLS_base: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    select: (args_0: string, args_1: string[]) => any;
    open: (args_0: string, args_1: string[]) => any;
    close: (args_0: string, args_1: string[]) => any;
}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{
    onSelect?: ((args_0: string, args_1: string[]) => any) | undefined;
    onOpen?: ((args_0: string, args_1: string[]) => any) | undefined;
    onClose?: ((args_0: string, args_1: string[]) => any) | undefined;
}>, {
    mode: "horizontal" | "vertical";
    theme: import("@vben-core/typings").ThemeModeType;
    rounded: boolean;
    accordion: boolean;
    collapse: boolean;
    scrollToActive: boolean;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
