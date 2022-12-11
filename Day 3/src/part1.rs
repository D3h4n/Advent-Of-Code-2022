use crate::{calculate_priority, intersect_strings};
use std::{
    fs::File,
    io::{BufRead, BufReader},
};

pub fn part_1(file: &str) -> u32 {
    let mut total = 0;
    let reader = BufReader::new(File::open(file).unwrap());

    for line in reader.lines() {
        total += calculate_priority(find_first_shared_item(&line.unwrap()));
    }

    total
}

fn find_first_shared_item(data: &str) -> char {
    let (first_half, second_half) = data.split_at(data.len() / 2);
    let intersection = intersect_strings(first_half, second_half);
    match intersection.chars().next() {
        Some(c) => c,
        None => '\0',
    }
}

#[cfg(test)]
mod tests {
    use crate::part1::find_first_shared_item;
    use test_case::test_case;

    #[test_case("vJrwpWtwJgWrhcsFMMfFFhFp", 'p' ; "Case 1")]
    #[test_case("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L' ; "Case 2")]
    #[test_case("PmmdzqPrVvPwwTWBwg", 'P' ; "Case 3")]
    #[test_case("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v' ; "Case 4")]
    #[test_case("ttgJtRGJQctTZtZT", 't' ; "Case 5")]
    #[test_case("CrZsJsPPZsGzwwsLwLmpwMDw", 's' ; "Case 6")]
    fn can_find_shared_item(input: &str, expected_result: char) {
        // Act
        let result = find_first_shared_item(input);

        // Assert
        assert_eq!(result, expected_result);
    }
}
