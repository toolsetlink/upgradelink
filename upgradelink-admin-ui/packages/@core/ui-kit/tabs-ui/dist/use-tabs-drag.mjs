import { nextTick, onMounted, onUnmounted, ref, watch } from "vue";
import { useIsMobile, useSortable } from "@vben-core/composables";
function findParentElement(element) {
  const parentCls = "group";
  return element.classList.contains(parentCls) ? element : element.closest(`.${parentCls}`);
}
export function useTabsDrag(props, emit) {
  const sortableInstance = ref(null);
  async function initTabsSortable() {
    await nextTick();
    const el = document.querySelectorAll(
      `.${props.contentClass}`
    )?.[0];
    if (!el) {
      console.warn("Element not found for sortable initialization");
      return;
    }
    const resetElState = async () => {
      el.style.cursor = "default";
      el.querySelector(".draggable")?.classList.remove("dragging");
    };
    const { initializeSortable } = useSortable(el, {
      filter: (_evt, target) => {
        const parent = findParentElement(target);
        const draggable = parent?.classList.contains("draggable");
        return !draggable || !props.draggable;
      },
      onEnd(evt) {
        const { newIndex, oldIndex } = evt;
        const { srcElement } = evt.originalEvent;
        if (!srcElement) {
          resetElState();
          return;
        }
        const srcParent = findParentElement(srcElement);
        if (!srcParent) {
          resetElState();
          return;
        }
        if (!srcParent.classList.contains("draggable")) {
          resetElState();
          return;
        }
        if (oldIndex !== void 0 && newIndex !== void 0 && !Number.isNaN(oldIndex) && !Number.isNaN(newIndex) && oldIndex !== newIndex) {
          emit("sortTabs", oldIndex, newIndex);
        }
        resetElState();
      },
      onMove(evt) {
        const parent = findParentElement(evt.related);
        if (parent?.classList.contains("draggable") && props.draggable) {
          const isCurrentAffix = evt.dragged.classList.contains("affix-tab");
          const isRelatedAffix = evt.related.classList.contains("affix-tab");
          return isCurrentAffix === isRelatedAffix;
        } else {
          return false;
        }
      },
      onStart: () => {
        el.style.cursor = "grabbing";
        el.querySelector(".draggable")?.classList.add("dragging");
      }
    });
    sortableInstance.value = await initializeSortable();
  }
  async function init() {
    const { isMobile } = useIsMobile();
    if (isMobile.value) {
      return;
    }
    await nextTick();
    initTabsSortable();
  }
  onMounted(init);
  watch(
    () => props.styleType,
    () => {
      sortableInstance.value?.destroy();
      init();
    }
  );
  onUnmounted(() => {
    sortableInstance.value?.destroy();
  });
}
