import type { VbenFormProps } from './types';
interface Props extends VbenFormProps {
}
declare var __VLS_10: string, __VLS_11: any, __VLS_14: {
    shapes: import("./types").FormShape[];
};
type __VLS_Slots = {} & {
    [K in NonNullable<typeof __VLS_10>]?: (props: typeof __VLS_11) => any;
} & {
    default?: (props: typeof __VLS_14) => any;
};
declare const __VLS_base: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{}>, {
    layout: import("./types").FormLayout;
    collapsed: boolean;
    collapsedRows: number;
    commonConfig: import("./types").FormCommonConfig;
    showCollapseButton: boolean;
    wrapperClass: import("./types").WrapperClassType;
    actionWrapperClass: import("@vben-core/typings").ClassType;
    handleReset: import("./types").HandleResetFn;
    handleSubmit: import("./types").HandleSubmitFn;
    resetButtonOptions: import("./types").ActionButtonOptions;
    showDefaultActions: boolean;
    submitButtonOptions: import("./types").ActionButtonOptions;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
