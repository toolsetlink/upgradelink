import { computed } from "vue";
export function useLayout(props) {
  const currentLayout = computed(
    () => props.isMobile ? "sidebar-nav" : props.layout
  );
  const isFullContent = computed(() => currentLayout.value === "full-content");
  const isSidebarMixedNav = computed(
    () => currentLayout.value === "sidebar-mixed-nav"
  );
  const isHeaderNav = computed(() => currentLayout.value === "header-nav");
  const isMixedNav = computed(
    () => currentLayout.value === "mixed-nav" || currentLayout.value === "header-sidebar-nav"
  );
  const isHeaderMixedNav = computed(
    () => currentLayout.value === "header-mixed-nav"
  );
  return {
    currentLayout,
    isFullContent,
    isHeaderMixedNav,
    isHeaderNav,
    isMixedNav,
    isSidebarMixedNav
  };
}
