const fs = require("node:fs");

/**
 * A - Rock - 1 Point
 * B - Paper - 2 Points
 * C - Scissors - 3 Points
 *
 * X - Rock / Lose
 * Y - Paper / Draw
 * Z - Scissors / Win
 *
 * Win - 6 Points
 * Draw - 3 Points
 * Loss - 1 Point
 */

/**
 * @typedef {Object} GameConfig
 * @prop {{A: string, B: string, C: string}} rules - Maps a shape to the shape that loses to it. E.g. A -> C represents Rock beats Scissors.
 * @prop {{X: string, Y: string, Z: string}} mapping - Maps X,Y,Z to A,B,C or win, loss, draw.
 * @prop {{A: number, B: number, C: number}} shapeValue - The number of points associated with a shape.
 * @prop {{win: number, draw: number, loss: number}} resultValue - The number of points associated with a result.
 */

function main() {
  /** @type GameConfig */
  const part1Config = {
    rules: {
      A: "C",
      B: "A",
      C: "B",
    },
    mapping: {
      X: "A",
      Y: "B",
      Z: "C",
    },
    shapeValue: {
      A: 1,
      B: 2,
      C: 3,
    },
    resultValue: {
      win: 6,
      draw: 3,
      loss: 0,
    },
  };

  /** @type GameConfig */
  const part2Config = {
    ...part1Config,
    mapping: {
      X: "loss",
      Y: "draw",
      Z: "win",
    },
  };

  let part1Points = 0;
  let part2Points = 0;

  readGamesFromFile(process.argv[2]).forEach((game) => {
    part1Points += runPart1Game(game, part1Config);
    part2Points += runPart2Game(game, part2Config);
  });

  console.log(`Part 1: ${part1Points}`);
  console.log(`Part 2: ${part2Points}`);
}

/**
 * @param {string} game
 * @param {GameConfig} config
 * @returns points earned from game
 */
function runPart1Game(game, config) {
  let [opponentShape, shape] = game.split(" ");
  shape = config.mapping[shape];
  let result;

  if (opponentShape === shape) {
    result = "draw";
  } else if (opponentShape === config.rules[shape]) {
    result = "win";
  } else {
    result = "loss";
  }

  return config.resultValue[result] + config.shapeValue[shape];
}

/**
 * @param {string} game
 * @param {GameConfig} config
 * @returns points earned from game
 */
function runPart2Game(game, config) {
  let [opponentShape, result] = game.split(" ");
  result = config.mapping[result];
  let shape;

  switch (result) {
    case "draw":
      shape = opponentShape;
      break;

    case "win":
      Object.entries(config.rules).forEach(([winningShape, losingShape]) => {
        if (opponentShape === losingShape) {
          shape = winningShape;
        }
      });
      break;

    case "loss":
      shape = config.rules[opponentShape];
      break;
  }

  return config.resultValue[result] + config.shapeValue[shape];
}

/**
 * @param {string} filename
 */
function readGamesFromFile(filename) {
  return fs.readFileSync(filename).toString().trim().split("\n");
}

if (require.main === module) {
  main();
}

module.exports = {
  runPart1Game,
  runPart2Game,
  readGamesFromFile,
};
