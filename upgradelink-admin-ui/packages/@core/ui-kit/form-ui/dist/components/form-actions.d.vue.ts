declare function handleSubmit(e: Event): Promise<void>;
declare function handleReset(e: Event): Promise<void>;
declare const __VLS_defaultModels: {
    modelValue: boolean;
};
type __VLS_ModelProps = {
    modelValue?: typeof __VLS_defaultModels['modelValue'];
};
declare var __VLS_1: {}, __VLS_11: {}, __VLS_21: {}, __VLS_31: {}, __VLS_39: {};
type __VLS_Slots = {} & {
    'submit-before'?: (props: typeof __VLS_1) => any;
} & {
    'reset-before'?: (props: typeof __VLS_11) => any;
} & {
    'submit-before'?: (props: typeof __VLS_21) => any;
} & {
    'expand-before'?: (props: typeof __VLS_31) => any;
} & {
    'expand-after'?: (props: typeof __VLS_39) => any;
};
declare const __VLS_base: import("vue").DefineComponent<__VLS_ModelProps, {
    handleReset: typeof handleReset;
    handleSubmit: typeof handleSubmit;
}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    "update:modelValue": (value: boolean) => any;
}, string, import("vue").PublicProps, Readonly<__VLS_ModelProps> & Readonly<{
    "onUpdate:modelValue"?: ((value: boolean) => any) | undefined;
}>, {}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
export default _default;
type __VLS_WithSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
