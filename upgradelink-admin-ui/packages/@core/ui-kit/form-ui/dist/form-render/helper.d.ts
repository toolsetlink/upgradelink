import type { AnyZodObject, ZodEffects, ZodTypeAny } from 'zod';
/**
 * Get the lowest level Zod type.
 * This will unpack optionals, refinements, etc.
 */
export declare function getBaseRules<ChildType extends AnyZodObject | ZodTypeAny = ZodTypeAny>(schema: ChildType | ZodEffects<ChildType>): ChildType | null;
/**
 * Search for a "ZodDefault" in the Zod stack and return its value.
 */
export declare function getDefaultValueInZodStack(schema: ZodTypeAny): any;
export declare function isEventObjectLike(obj: any): boolean;
