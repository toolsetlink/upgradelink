import type { MenuRecordRaw } from '@vben-core/typings';
import type { NormalMenuProps } from './normal-menu';
interface Props extends NormalMenuProps {
}
declare const __VLS_export: import("vue").DefineComponent<Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    select: (args_0: MenuRecordRaw) => any;
    enter: (args_0: MenuRecordRaw) => any;
}, string, import("vue").PublicProps, Readonly<Props> & Readonly<{
    onSelect?: ((args_0: MenuRecordRaw) => any) | undefined;
    onEnter?: ((args_0: MenuRecordRaw) => any) | undefined;
}>, {
    menus: MenuRecordRaw[];
    theme: "dark" | "light";
    collapse: boolean;
    activePath: string;
}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
declare const _default: typeof __VLS_export;
export default _default;
