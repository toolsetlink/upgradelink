import type { DrawerApiOptions, DrawerProps, ExtendedDrawerApi } from './drawer';
export declare function setDefaultDrawerProps(props: Partial<DrawerProps>): void;
export declare function useVbenDrawer<TParentDrawerProps extends DrawerProps = DrawerProps>(options?: DrawerApiOptions): readonly [import("vue").DefineSetupFnComponent<DrawerProps, {}, {}, DrawerProps & {}, import("vue").PublicProps>, ExtendedDrawerApi];
