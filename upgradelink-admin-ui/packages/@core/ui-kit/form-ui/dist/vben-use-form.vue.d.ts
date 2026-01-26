import type { ExtendedFormApi, VbenFormProps } from './types';
interface Props extends VbenFormProps {
    formApi?: ExtendedFormApi;
}
declare var __VLS_12: string, __VLS_13: any, __VLS_16: {
    shapes: import("./types").FormShape[];
}, __VLS_27: {}, __VLS_30: {}, __VLS_33: {}, __VLS_36: {};
type __VLS_Slots = {} & {
    [K in NonNullable<typeof __VLS_12>]?: (props: typeof __VLS_13) => any;
} & {
    default?: (props: typeof __VLS_16) => any;
} & {
    'reset-before'?: (props: typeof __VLS_27) => any;
} & {
    'submit-before'?: (props: typeof __VLS_30) => any;
} & {
    'expand-before'?: (props: typeof __VLS_33) => any;
} & {
    'expand-after'?: (props: typeof __VLS_36) => any;
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
