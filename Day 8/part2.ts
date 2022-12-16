async function main() {
  const text = await Deno.readTextFile("puzzle_input.txt");
  const trees = text
    .trim()
    .split("\n")
    .map((line) => line.split(""))
    .map((row) => row.map(Number));
  const max = findMaxScenicScore(trees);
  console.log(max);
}

function findMaxScenicScore(trees: number[][]): number {
  let max = -1;

  for (let i = 1; i < trees.length - 1; i++) {
    for (let j = 1; j < trees[i].length - 1; j++) {
      const scenicScore = calculateScenicScore(i, j, trees);

      if (scenicScore > max) {
        max = scenicScore;
      }
    }
  }

  return max;
}

function calculateScenicScore(i: number, j: number, trees: number[][]): number {
  return calculateScenicScoreNorth(i, j, trees) *
    calculateScenicScoreSouth(i, j, trees) *
    calculateScenicScoreEast(i, j, trees) *
    calculateScenicScoreWest(i, j, trees);
}

function calculateScenicScoreNorth(
  i: number,
  j: number,
  trees: number[][],
): number {
  let count = 0;
  const tree = trees[i][j];

  for (let y = i - 1; y >= 0; y--) {
    count++;
    if (trees[y][j] >= tree) {
      break;
    }
  }

  return count;
}

function calculateScenicScoreSouth(
  i: number,
  j: number,
  trees: number[][],
): number {
  let count = 0;
  const tree = trees[i][j];

  for (let y = i + 1; y < trees.length; y++) {
    count++;
    if (trees[y][j] >= tree) {
      break;
    }
  }

  return count;
}

function calculateScenicScoreEast(
  i: number,
  j: number,
  trees: number[][],
): number {
  let count = 0;
  const row = trees[i];
  const tree = row[j];

  for (let x = j + 1; x < row.length; x++) {
    count++;
    if (row[x] >= tree) {
      break;
    }
  }

  return count;
}

function calculateScenicScoreWest(
  i: number,
  j: number,
  trees: number[][],
): number {
  let count = 0;
  const row = trees[i];
  const tree = row[j];

  for (let x = j - 1; x >= 0; x--) {
    count++;
    if (row[x] >= tree) {
      break;
    }
  }

  return count;
}

if (import.meta.main) {
  main();
}
