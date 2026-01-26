import type { ExtendedModalApi, ModalProps } from './modal';
interface Props extends ModalProps {
    modalApi?: ExtendedModalApi;
}
declare var __VLS_41: {}, __VLS_43: {}, __VLS_57: {}, __VLS_75: {}, __VLS_108: {}, __VLS_110: {}, __VLS_120: {}, __VLS_122: {}, __VLS_132: {}, __VLS_134: {};
type __VLS_Slots = {} & {
    title?: (props: typeof __VLS_41) => any;
} & {
    titleTooltip?: (props: typeof __VLS_43) => any;
} & {
    description?: (props: typeof __VLS_57) => any;
} & {
    default?: (props: typeof __VLS_75) => any;
} & {
    'prepend-footer'?: (props: typeof __VLS_108) => any;
} & {
    footer?: (props: typeof __VLS_110) => any;
} & {
    cancelText?: (props: typeof __VLS_120) => any;
} & {
    'center-footer'?: (props: typeof __VLS_122) => any;
} & {
    confirmText?: (props: typeof __VLS_132) => any;
} & {
    'append-footer'?: (props: typeof __VLS_134) => any;
};
declare const __VLS_base: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{}>, {
    appendToMain: boolean;
    destroyOnClose: boolean;
    modalApi: ExtendedModalApi;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
