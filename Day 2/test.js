const { expect } = require('@jest/globals')
const { runPart1Game, runPart2Game, readGamesFromFile } = require('.')

test("Should correctly read games from file", () => {
  // Act
  let games = readGamesFromFile('test_input.txt')

  // Assert
  expect(games).toEqual(['A Y', 'B X', 'C Z'])
})

/**
 * A - Rock - 1 Point
 * B - Paper - 2 Points
 * C - Scissors - 3 Points
 *
 * X - Rock - A
 * Y - Paper - B
 * Z - Scissors - C
 * 
 * Win - 6 Points
 * Draw - 3 Points
 * Loss - 1 Point 
 */

test.each`
game     | expectedPoints
${'A Y'} | ${8}
${'B X'} | ${1}
${'C Z'} | ${6}
`("Part 1 - Should return $expectedPoints points for game '$game'", ({game, expectedPoints}) => {
  const config = {
    mapping: {
      'X': 'A',
      'Y': 'B',
      'Z': 'C'
    },
    rules: {
      'A': 'C',
      'B': 'A',
      'C': 'B'
    },
    shapeValue: {
      'A': 1,
      'B': 2,
      'C': 3
    },
    resultValue: {
      'win': 6,
      'draw': 3,
      'loss': 0
    }
  }

  // Act
  let points = runPart1Game(game, config)

  // Assert
  expect(points).toEqual(expectedPoints)
})


/**
 * A - Rock - 1 Point
 * B - Paper - 2 Points
 * C - Scissors - 3 Points
 *
 * Win - 6 Points
 * Draw - 3 Points
 * Loss - 1 Point 
 *
 * X - Lose
 * Y - Draw
 * Z - Win
 */

test.each`
game     | expectedPoints
${'A Y'} | ${4}
${'B X'} | ${1}
${'C Z'} | ${7}
`("Part 2 - Should return $expectedPoints points for game '$game'", ({game, expectedPoints}) => {
  const config = {
    mapping: {
      'X': 'loss',
      'Y': 'draw',
      'Z': 'win'
    },
    rules: {
      'A': 'C',
      'B': 'A',
      'C': 'B'
    },
    shapeValue: {
      'A': 1,
      'B': 2,
      'C': 3
    },
    resultValue: {
      'win': 6,
      'draw': 3,
      'loss': 0
    }
  }

  // Act
  let points = runPart2Game(game, config)

  // Assert
  expect(points).toEqual(expectedPoints)
})
