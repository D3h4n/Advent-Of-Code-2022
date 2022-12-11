mod part1;
mod part2;
use std::collections::HashSet;

use crate::part1::part_1;
use crate::part2::part_2;

fn main() {
    println!("Part 1: {}", part_1(".\\puzzle_input.txt"));
    println!("Part 2: {}", part_2(".\\puzzle_input.txt"));
}

pub fn intersect_strings(str1: &str, str2: &str) -> String {
    let chars: HashSet<char> = str1.chars().collect();
    let mut intersection: Vec<char> = vec![];

    for c in str2.chars() {
        if chars.contains(&c) {
            intersection.push(c);
        }
    }

    intersection.sort();
    intersection.into_iter().collect::<String>()
}

pub fn calculate_priority(c: char) -> u32 {
    if !c.is_ascii_alphabetic() {
        panic!("Received a non-alphabetic character");
    }
    match c.is_ascii_uppercase() {
        true => c as u32 - 'A' as u32 + 27,
        false => c as u32 - 'a' as u32 + 1,
    }
}

#[cfg(test)]
mod tests {
    use crate::{calculate_priority, intersect_strings};
    use test_case::test_case;

    #[test_case('p', 16 ; "Lower Case Letter 1")]
    #[test_case('t', 20 ; "Lower Case Letter 2")]
    #[test_case('P', 42 ; "Upper Case Letter 1")]
    #[test_case('L', 38 ; "Upper Case Letter 2")]
    fn calculates_priority(input: char, expected_result: u32) {
        // Act
        let result = calculate_priority(input);

        // Assert
        assert_eq!(result, expected_result);
    }

    #[test_case("aPUI", "BacCzY", "a" ; "With 1 Shared Character")]
    #[test_case("aPUIB", "BacCzY", "Ba" ; "With 2 Shared Characters")]
    #[test_case("xyz", "abc", "" ; "With No Shared Characters")]
    fn can_find_the_intersection_of_strings(str1: &str, str2: &str, expected_result: &str) {
        // Act
        let result = intersect_strings(str1, str2);

        // Assert
        assert_eq!(result, expected_result);
    }
}
