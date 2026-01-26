import type { ExtendedModalApi, ModalApiOptions, ModalProps } from './modal';
export declare function setDefaultModalProps(props: Partial<ModalProps>): void;
export declare function useVbenModal<TParentModalProps extends ModalProps = ModalProps>(options?: ModalApiOptions): readonly [import("vue").DefineSetupFnComponent<ModalProps, {}, {}, ModalProps & {}, import("vue").PublicProps>, ExtendedModalApi];
