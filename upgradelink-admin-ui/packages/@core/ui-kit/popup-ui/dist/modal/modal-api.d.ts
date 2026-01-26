import type { ModalApiOptions, ModalState } from './modal';
import { Store } from '@vben-core/shared/store';
export declare class ModalApi {
    sharedData: Record<'payload', any>;
    store: Store<ModalState>;
    private api;
    private state;
    constructor(options?: ModalApiOptions);
    /**
     * 关闭弹窗
     * @description 关闭弹窗时会调用 onBeforeClose 钩子函数，如果 onBeforeClose 返回 false，则不关闭弹窗
     */
    close(): Promise<void>;
    getData<T extends object = Record<string, any>>(): T;
    /**
     * 锁定弹窗状态（用于提交过程中的等待状态）
     * @description 锁定状态将禁用默认的取消按钮，使用spinner覆盖弹窗内容，隐藏关闭按钮，阻止手动关闭弹窗，将默认的提交按钮标记为loading状态
     * @param isLocked 是否锁定
     */
    lock(isLocked?: boolean): this;
    /**
     * 取消操作
     */
    onCancel(): void;
    /**
     * 弹窗关闭动画播放完毕后的回调
     */
    onClosed(): void;
    /**
     * 确认操作
     */
    onConfirm(): void;
    /**
     * 弹窗打开动画播放完毕后的回调
     */
    onOpened(): void;
    open(): void;
    setData<T>(payload: T): this;
    setState(stateOrFn: ((prev: ModalState) => Partial<ModalState>) | Partial<ModalState>): this;
    /**
     * 解除弹窗的锁定状态
     * @description 解除由lock方法设置的锁定状态，是lock(false)的别名
     */
    unlock(): this;
}
