import type { FormRenderProps } from '../types';
export declare const injectRenderFormProps: <T extends FormRenderProps<import("../types").BaseFormComponentType> | null | undefined = FormRenderProps<import("../types").BaseFormComponentType>>(fallback?: T | undefined) => T extends null ? FormRenderProps<import("../types").BaseFormComponentType> | null : FormRenderProps<import("../types").BaseFormComponentType>, provideFormRenderProps: (contextValue: FormRenderProps<import("../types").BaseFormComponentType>) => FormRenderProps<import("../types").BaseFormComponentType>;
export declare const useFormContext: () => {
    componentBindEventMap: import("vue").ComputedRef<Partial<Record<import("../types").BaseFormComponentType, string>> | undefined>;
    componentMap: import("vue").ComputedRef<Record<import("../types").BaseFormComponentType, import("vue").Component>>;
    isVertical: import("vue").ComputedRef<boolean>;
};
