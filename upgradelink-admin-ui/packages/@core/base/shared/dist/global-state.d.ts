interface ComponentsState {
    [key: string]: any;
}
interface MessageState {
    copyPreferencesSuccess?: (title: string, content?: string) => void;
}
interface IGlobalSharedState {
    components: ComponentsState;
    message: MessageState;
}
declare class GlobalShareState {
    #private;
    defineMessage({ copyPreferencesSuccess }: MessageState): void;
    getComponents(): ComponentsState;
    getMessage(): MessageState;
    setComponents(value: ComponentsState): void;
}
declare const globalShareState: GlobalShareState;

export { globalShareState };
export type { IGlobalSharedState };
