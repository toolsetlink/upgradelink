import type { FormSchema, MaybeComponentProps } from '../types';
interface Props extends FormSchema {
}
type __VLS_Props = Props & {
    commonComponentProps: MaybeComponentProps;
};
declare var __VLS_31: any, __VLS_54: {};
type __VLS_Slots = {} & {
    default?: (props: typeof __VLS_31) => any;
} & {
    trigger?: (props: typeof __VLS_54) => any;
};
declare const __VLS_base: import("vue").DefineComponent<__VLS_Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").PublicProps, Readonly<__VLS_Props> & Readonly<{}>, {}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
