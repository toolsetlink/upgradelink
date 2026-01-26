/**
 * @copy https://github.com/element-plus/element-plus/blob/dev/packages/hooks/use-draggable/index.ts
 * 调整部分细节
 */
import type { ComputedRef, Ref } from 'vue';
export declare function useModalDraggable(targetRef: Ref<HTMLElement | undefined>, dragRef: Ref<HTMLElement | undefined>, draggable: ComputedRef<boolean>, containerSelector?: ComputedRef<string | undefined>, centered?: ComputedRef<boolean>): {
    dragging: Ref<boolean, boolean>;
    resetPosition: () => void;
    transform: {
        offsetX: number;
        offsetY: number;
    };
};
