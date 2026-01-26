import type { ComputedRef } from 'vue';
import type { ExtendedFormApi, FormActions, VbenFormProps } from './types';
type ExtendFormProps = VbenFormProps & {
    formApi?: ExtendedFormApi;
};
export declare const injectFormProps: <T extends [ExtendFormProps | ComputedRef<ExtendFormProps>, FormActions] | null | undefined = [ExtendFormProps | ComputedRef<ExtendFormProps>, FormActions]>(fallback?: T | undefined) => T extends null ? [ExtendFormProps | ComputedRef<ExtendFormProps>, FormActions] | null : [ExtendFormProps | ComputedRef<ExtendFormProps>, FormActions], provideFormProps: (contextValue: [ExtendFormProps | ComputedRef<ExtendFormProps>, FormActions]) => [ExtendFormProps | ComputedRef<ExtendFormProps>, FormActions];
export declare const injectComponentRefMap: <T extends Map<string, unknown> | null | undefined = Map<string, unknown>>(fallback?: T | undefined) => T extends null ? Map<string, unknown> | null : Map<string, unknown>, provideComponentRefMap: (contextValue: Map<string, unknown>) => Map<string, unknown>;
export declare function useFormInitial(props: ComputedRef<VbenFormProps> | VbenFormProps): {
    delegatedSlots: ComputedRef<string[]>;
    form: import("vee-validate").FormContext<any, any>;
};
export {};
