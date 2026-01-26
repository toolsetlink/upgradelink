import type { BaseFormComponentType, ExtendedFormApi, VbenFormProps } from './types';
export declare function useVbenForm<T extends BaseFormComponentType = BaseFormComponentType>(options: VbenFormProps<T>): readonly [import("vue").DefineSetupFnComponent<VbenFormProps<BaseFormComponentType>, {}, {}, VbenFormProps<BaseFormComponentType> & {}, import("vue").PublicProps>, ExtendedFormApi];
