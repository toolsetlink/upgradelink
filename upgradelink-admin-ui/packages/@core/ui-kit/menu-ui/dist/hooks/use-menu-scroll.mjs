import { watch } from "vue";
import { useDebounceFn } from "@vueuse/core";
export function useMenuScroll(activePath, options = {}) {
  const { enable = true, delay = 320 } = options;
  function scrollToActiveItem() {
    const isEnabled = typeof enable === "boolean" ? enable : enable.value;
    if (!isEnabled) return;
    const activeElement = document.querySelector(
      `aside li[role=menuitem].is-active`
    );
    if (activeElement) {
      activeElement.scrollIntoView({
        behavior: "smooth",
        block: "center",
        inline: "center"
      });
    }
  }
  const debouncedScroll = useDebounceFn(scrollToActiveItem, delay);
  watch(activePath, () => {
    const isEnabled = typeof enable === "boolean" ? enable : enable.value;
    if (!isEnabled) return;
    debouncedScroll();
  });
  return {
    scrollToActiveItem
  };
}
