import { nextTick, onMounted, onUnmounted, ref, watch } from "vue";
import { useDebounceFn } from "@vueuse/core";
export function useTabsViewScroll(props) {
  let resizeObserver = null;
  let mutationObserver = null;
  let tabItemCount = 0;
  const scrollbarRef = ref(null);
  const scrollViewportEl = ref(null);
  const showScrollButton = ref(false);
  const scrollIsAtLeft = ref(true);
  const scrollIsAtRight = ref(false);
  function getScrollClientWidth() {
    const scrollbarEl = scrollbarRef.value?.$el;
    if (!scrollbarEl || !scrollViewportEl.value) return {};
    const scrollbarWidth = scrollbarEl.clientWidth;
    const scrollViewWidth = scrollViewportEl.value.clientWidth;
    return {
      scrollbarWidth,
      scrollViewWidth
    };
  }
  function scrollDirection(direction, distance = 150) {
    const { scrollbarWidth, scrollViewWidth } = getScrollClientWidth();
    if (!scrollbarWidth || !scrollViewWidth) return;
    if (scrollbarWidth > scrollViewWidth) return;
    scrollViewportEl.value?.scrollBy({
      behavior: "smooth",
      left: direction === "left" ? -(scrollbarWidth - distance) : +(scrollbarWidth - distance)
    });
  }
  async function initScrollbar() {
    await nextTick();
    const scrollbarEl = scrollbarRef.value?.$el;
    if (!scrollbarEl) {
      return;
    }
    const viewportEl = scrollbarEl?.querySelector(
      "div[data-reka-scroll-area-viewport]"
    );
    scrollViewportEl.value = viewportEl;
    calcShowScrollbarButton();
    await nextTick();
    scrollToActiveIntoView();
    resizeObserver?.disconnect();
    resizeObserver = new ResizeObserver(
      useDebounceFn((_entries) => {
        calcShowScrollbarButton();
        scrollToActiveIntoView();
      }, 100)
    );
    resizeObserver.observe(viewportEl);
    tabItemCount = props.tabs?.length || 0;
    mutationObserver?.disconnect();
    mutationObserver = new MutationObserver(() => {
      const count = viewportEl.querySelectorAll(
        `div[data-tab-item="true"]`
      ).length;
      if (count > tabItemCount) {
        scrollToActiveIntoView();
      }
      if (count !== tabItemCount) {
        calcShowScrollbarButton();
        tabItemCount = count;
      }
    });
    mutationObserver.observe(viewportEl, {
      attributes: false,
      childList: true,
      subtree: true
    });
  }
  async function scrollToActiveIntoView() {
    if (!scrollViewportEl.value) {
      return;
    }
    await nextTick();
    const viewportEl = scrollViewportEl.value;
    const { scrollbarWidth } = getScrollClientWidth();
    const { scrollWidth } = viewportEl;
    if (scrollbarWidth >= scrollWidth) {
      return;
    }
    requestAnimationFrame(() => {
      const activeItem = viewportEl?.querySelector(".is-active");
      activeItem?.scrollIntoView({ behavior: "smooth", inline: "start" });
    });
  }
  async function calcShowScrollbarButton() {
    if (!scrollViewportEl.value) {
      return;
    }
    const { scrollbarWidth } = getScrollClientWidth();
    showScrollButton.value = scrollViewportEl.value.scrollWidth > scrollbarWidth;
  }
  const handleScrollAt = useDebounceFn(({ left, right }) => {
    scrollIsAtLeft.value = left;
    scrollIsAtRight.value = right;
  }, 100);
  function handleWheel({ deltaY }) {
    scrollViewportEl.value?.scrollBy({
      // behavior: 'smooth',
      left: deltaY * 3
    });
  }
  watch(
    () => props.active,
    async () => {
      scrollToActiveIntoView();
    },
    {
      flush: "post"
    }
  );
  watch(
    () => props.styleType,
    () => {
      initScrollbar();
    }
  );
  onMounted(initScrollbar);
  onUnmounted(() => {
    resizeObserver?.disconnect();
    mutationObserver?.disconnect();
    resizeObserver = null;
    mutationObserver = null;
  });
  return {
    handleScrollAt,
    handleWheel,
    initScrollbar,
    scrollbarRef,
    scrollDirection,
    scrollIsAtLeft,
    scrollIsAtRight,
    showScrollButton
  };
}
