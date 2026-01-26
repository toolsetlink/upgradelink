import type { Component } from 'vue';
import type { BaseFormComponentType, FormCommonConfig, VbenFormAdapterOptions } from './types';
export declare const DEFAULT_FORM_COMMON_CONFIG: FormCommonConfig;
export declare const COMPONENT_MAP: Record<BaseFormComponentType, Component>;
export declare const COMPONENT_BIND_EVENT_MAP: Partial<Record<BaseFormComponentType, string>>;
export declare function setupVbenForm<T extends BaseFormComponentType = BaseFormComponentType>(options: VbenFormAdapterOptions<T>): void;
