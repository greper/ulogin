export function eachTree(tree: any[], callback: (item: any) => void) {
  tree.forEach((item) => {
    callback(item);
    if (item.children) {
      eachTree(item.children, callback);
    }
  });
}

export const treeUtils = {
  eachTree
};
