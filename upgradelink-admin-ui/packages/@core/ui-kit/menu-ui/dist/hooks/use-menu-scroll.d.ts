import type { Ref } from 'vue';
interface UseMenuScrollOptions {
    delay?: number;
    enable?: boolean | Ref<boolean>;
}
export declare function useMenuScroll(activePath: Ref<string | undefined>, options?: UseMenuScrollOptions): {
    scrollToActiveItem: () => void;
};
export {};
