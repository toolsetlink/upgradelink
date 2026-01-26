import type { FormCommonConfig, FormRenderProps, FormShape } from '../types';
interface Props extends FormRenderProps {
}
type __VLS_Props = Props & {
    globalCommonConfig?: FormCommonConfig;
};
declare var __VLS_16: string, __VLS_17: any, __VLS_19: {
    shapes: FormShape[];
};
type __VLS_Slots = {} & {
    [K in NonNullable<typeof __VLS_16>]?: (props: typeof __VLS_17) => any;
} & {
    default?: (props: typeof __VLS_19) => any;
};
declare const __VLS_base: import("vue").DefineComponent<__VLS_Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    submit: (event: any) => any;
}, string, import("vue").PublicProps, Readonly<__VLS_Props> & Readonly<{
    onSubmit?: ((event: any) => any) | undefined;
}>, {
    collapsedRows: number;
    commonConfig: FormCommonConfig;
    showCollapseButton: boolean;
    wrapperClass: import("../types").WrapperClassType;
    globalCommonConfig: FormCommonConfig;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
