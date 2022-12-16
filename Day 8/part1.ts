async function main() {
  const text = await Deno.readTextFile("puzzle_input.txt");
  const trees = text
    .trim()
    .split("\n")
    .map((line) => line.split(""))
    .map((row) => row.map(Number))
  const count = countVisibleTreesInGrid(trees);
  console.log(count);
}

export function countVisibleTreesInGrid(trees: number[][]): number {
  return numPerimTrees(trees) + numVisibleInnerTrees(trees);
}

function numPerimTrees(trees: number[][]): number {
  return 2 * (trees.length + trees[0].length - 2);
}

function numVisibleInnerTrees(trees: number[][]): number {
  let count = 0;

  for (let y = 1; y < trees.length - 1; y++) {
    for (let x = 1; x < trees[y].length - 1; x++) {
      if (treeIsVisible(x, y, trees)) {
        count++;
      }
    }
  }

  return count;
}

function treeIsVisible(x: number, y: number, trees: number[][]): boolean {
  return (
    treeIsVisibleFromWest(x, y, trees) ||
    treeIsVisibleFromEast(x, y, trees) ||
    treeIsVisibleFromNorth(x, y, trees) ||
    treeIsVisibleFromSouth(x, y, trees)
  );
}

function treeIsVisibleFromWest(
  x: number,
  y: number,
  trees: number[][],
): boolean {
  const row = trees[y];
  const tree = row[x];

  for (let i = x - 1; i >= 0; i--) {
    if (row[i] >= tree) {
      return false;
    }
  }

  return true;
}

function treeIsVisibleFromEast(
  x: number,
  y: number,
  trees: number[][],
): boolean {
  const row = trees[y];
  const tree = row[x];

  for (let i = x + 1; i < row.length; i++) {
    if (row[i] >= tree) {
      return false;
    }
  }

  return true;
}

function treeIsVisibleFromNorth(
  x: number,
  y: number,
  trees: number[][],
): boolean {
  const tree = trees[y][x];

  for (let i = y - 1; i >= 0; i--) {
    if (trees[i][x] >= tree) {
      return false;
    }
  }

  return true;
}

function treeIsVisibleFromSouth(
  x: number,
  y: number,
  trees: number[][],
): boolean {
  const tree = trees[y][x];

  for (let i = y + 1; i < trees.length; i++) {
    if (trees[i][x] >= tree) {
      return false;
    }
  }

  return true;
}

if (import.meta.main) {
  main();
}
