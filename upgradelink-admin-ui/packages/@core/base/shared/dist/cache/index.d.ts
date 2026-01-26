type StorageType = 'localStorage' | 'sessionStorage';
interface StorageManagerOptions {
    prefix?: string;
    storageType?: StorageType;
}
declare class StorageManager {
    private prefix;
    private storage;
    constructor({ prefix, storageType, }?: StorageManagerOptions);
    clear(): void;
    clearExpiredItems(): void;
    getItem<T>(key: string, defaultValue?: null | T): null | T;
    removeItem(key: string): void;
    setItem<T>(key: string, value: T, ttl?: number): void;
    private getFullKey;
}

export { StorageManager };
