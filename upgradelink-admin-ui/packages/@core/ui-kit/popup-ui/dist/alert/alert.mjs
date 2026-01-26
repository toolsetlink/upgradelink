import { createContext } from "@vben-core/shadcn-ui";
export const [injectAlertContext, provideAlertContext] = createContext("VbenAlertContext");
export function useAlertContext() {
  const context = injectAlertContext();
  if (!context) {
    throw new Error("useAlertContext must be used within an AlertProvider");
  }
  return context;
}
