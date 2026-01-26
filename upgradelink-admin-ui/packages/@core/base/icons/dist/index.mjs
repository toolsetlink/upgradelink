import { defineComponent, h } from 'vue';
import { Icon } from '@iconify/vue';
export { Icon as IconifyIcon, addCollection, addIcon, listIcons } from '@iconify/vue';
export { ArrowDown, ArrowLeft, ArrowLeftToLine, ArrowRightLeft, ArrowRightToLine, ArrowUp, ArrowUpToLine, Bell, BookOpenText, Check, ChevronDown, ChevronLeft, ChevronRight, ChevronsLeft, ChevronsRight, Circle, CircleAlert, CircleCheckBig, CircleHelp, CircleX, Copy, CornerDownLeft, Ellipsis, Expand, ExternalLink, Eye, EyeOff, FoldHorizontal, Fullscreen, Github, Grip, GripVertical, Menu as IconDefault, Inbox, Info, InspectionPanel, Languages, LoaderCircle, LockKeyhole, LogOut, MailCheck, Maximize, ArrowRightFromLine as MdiMenuClose, ArrowLeftFromLine as MdiMenuOpen, Menu, Minimize, Minimize2, MoonStar, Palette, PanelLeft, PanelRight, Pin, PinOff, Plus, RotateCw, Search, SearchX, Settings, Shrink, Square, SquareCheckBig, SquareMinus, Sun, SunMoon, SwatchBook, UserRoundPen, X } from 'lucide-vue-next';

function createIconifyIcon(icon) {
  return defineComponent({
    name: `Icon-${icon}`,
    setup(props, { attrs }) {
      return () => h(Icon, { icon, ...props, ...attrs });
    }
  });
}

export { createIconifyIcon };
