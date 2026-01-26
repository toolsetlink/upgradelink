import type { AlertProps, PromptProps } from './alert';
export declare function vbenAlert(options: AlertProps): Promise<void>;
export declare function vbenAlert(message: string, options?: Partial<AlertProps>): Promise<void>;
export declare function vbenAlert(message: string, title?: string, options?: Partial<AlertProps>): Promise<void>;
export declare function vbenConfirm(options: AlertProps): Promise<void>;
export declare function vbenConfirm(message: string, options?: Partial<AlertProps>): Promise<void>;
export declare function vbenConfirm(message: string, title?: string, options?: Partial<AlertProps>): Promise<void>;
export declare function vbenPrompt<T = any>(options: PromptProps<T>): Promise<T | undefined>;
export declare function clearAllAlerts(): void;
