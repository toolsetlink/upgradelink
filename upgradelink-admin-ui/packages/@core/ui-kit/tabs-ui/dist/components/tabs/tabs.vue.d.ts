import type { TabDefinition } from '@vben-core/typings';
import type { TabsProps } from '../../types';
interface Props extends TabsProps {
}
type __VLS_Props = Props;
type __VLS_ModelProps = {
    'active'?: string;
};
type __VLS_PublicProps = __VLS_Props & __VLS_ModelProps;
declare const __VLS_export: import("vue").DefineComponent<__VLS_PublicProps, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    close: (args_0: string) => any;
    unpin: (args_0: TabDefinition) => any;
    "update:active": (value: string | undefined) => any;
}, string, import("vue").PublicProps, Readonly<__VLS_PublicProps> & Readonly<{
    onClose?: ((args_0: string) => any) | undefined;
    onUnpin?: ((args_0: TabDefinition) => any) | undefined;
    "onUpdate:active"?: ((value: string | undefined) => any) | undefined;
}>, {
    contentClass: string;
    tabs: TabDefinition[];
    contextMenus: (data: any) => import("@vben-core/shadcn-ui").IContextMenuItem[];
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const _default: typeof __VLS_export;
export default _default;
