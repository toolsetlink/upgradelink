import type { ComponentInternalInstance, VNode, VNodeChild, VNodeNormalizedChildren } from 'vue';
type VNodeChildAtom = Exclude<VNodeChild, Array<any>>;
type RawSlots = Exclude<VNodeNormalizedChildren, Array<any> | null | string>;
type FlattenVNodes = Array<RawSlots | VNodeChildAtom>;
/**
 * @zh_CN Find the parent component upward
 * @param instance
 * @param parentNames
 */
declare function findComponentUpward(instance: ComponentInternalInstance, parentNames: string[]): ComponentInternalInstance | null;
declare const flattedChildren: (children: FlattenVNodes | VNode | VNodeNormalizedChildren) => FlattenVNodes;
export { findComponentUpward, flattedChildren };
