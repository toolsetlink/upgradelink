import type { FormState, GenericObject, ResetFormOpts, ValidationOptions } from 'vee-validate';
import type { ComponentPublicInstance } from 'vue';
import type { Recordable } from '@vben-core/typings';
import type { FormActions, FormSchema, VbenFormProps } from './types';
import { Store } from '@vben-core/shared/store';
import { StateHandler } from '@vben-core/shared/utils';
export declare class FormApi {
    form: FormActions;
    isMounted: boolean;
    state: null | VbenFormProps;
    stateHandler: StateHandler;
    store: Store<VbenFormProps>;
    /**
     * 组件实例映射
     */
    private componentRefMap;
    private latestSubmissionValues;
    private prevState;
    constructor(options?: VbenFormProps);
    /**
     * 获取字段组件实例
     * @param fieldName 字段名
     * @returns 组件实例
     */
    getFieldComponentRef<T = ComponentPublicInstance>(fieldName: string): T | undefined;
    /**
     * 获取当前聚焦的字段，如果没有聚焦的字段则返回undefined
     */
    getFocusedField(): any;
    getLatestSubmissionValues(): Recordable<any>;
    getState(): VbenFormProps<import("./types").BaseFormComponentType> | null;
    getValues<T = Recordable<any>>(): Promise<T>;
    isFieldValid(fieldName: string): Promise<boolean>;
    merge(formApi: FormApi): any;
    mount(formActions: FormActions, componentRefMap: Map<string, unknown>): void;
    /**
     * 根据字段名移除表单项
     * @param fields
     */
    removeSchemaByFields(fields: string[]): Promise<void>;
    /**
     * 重置表单
     */
    resetForm(state?: Partial<FormState<GenericObject>> | undefined, opts?: Partial<ResetFormOpts>): Promise<void>;
    resetValidate(): Promise<void>;
    /**
     * 滚动到第一个错误字段
     * @param errors 验证错误对象
     */
    scrollToFirstError(errors: Record<string, any> | string): void;
    setFieldValue(field: string, value: any, shouldValidate?: boolean): Promise<void>;
    setLatestSubmissionValues(values: null | Recordable<any>): void;
    setState(stateOrFn: ((prev: VbenFormProps) => Partial<VbenFormProps>) | Partial<VbenFormProps>): void;
    /**
     * 设置表单值
     * @param fields record
     * @param filterFields 过滤不在schema中定义的字段 默认为true
     * @param shouldValidate
     */
    setValues(fields: Record<string, any>, filterFields?: boolean, shouldValidate?: boolean): Promise<void>;
    submitForm(e?: Event): Promise<Recordable<any>>;
    unmount(): void;
    updateSchema(schema: Partial<FormSchema>[]): void;
    validate(opts?: Partial<ValidationOptions>): Promise<import("vee-validate").FormValidationResult<GenericObject, GenericObject>>;
    validateAndSubmitForm(): Promise<Recordable<any> | undefined>;
    validateField(fieldName: string, opts?: Partial<ValidationOptions>): Promise<import("vee-validate").ValidationResult<any>>;
    private getForm;
    private handleMultiFields;
    private handleRangeTimeValue;
    private processFields;
    private updateState;
}
