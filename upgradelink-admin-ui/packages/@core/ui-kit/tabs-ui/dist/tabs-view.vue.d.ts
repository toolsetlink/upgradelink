import type { TabsProps } from './types';
interface Props extends TabsProps {
}
declare const __VLS_export: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    close: (args_0: string) => any;
    unpin: (args_0: import("@vben-core/typings").TabDefinition) => any;
    sortTabs: (args_0: number, args_1: number) => any;
}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{
    onClose?: ((args_0: string) => any) | undefined;
    onUnpin?: ((args_0: import("@vben-core/typings").TabDefinition) => any) | undefined;
    onSortTabs?: ((args_0: number, args_1: number) => any) | undefined;
}>, {
    contentClass: string;
    styleType: import("@vben-core/typings").TabsStyleType;
    draggable: boolean;
    wheelable: boolean;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const _default: typeof __VLS_export;
export default _default;
