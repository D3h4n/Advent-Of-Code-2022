use crate::{calculate_priority, intersect_strings};
use std::{
    fs::File,
    io::{BufRead, BufReader},
};

pub fn part_2(file: &str) -> u32 {
    let mut total = 0;
    let mut intersection = String::new();
    let reader = BufReader::new(File::open(file).unwrap());

    for (i, line) in reader.lines().enumerate() {
        let ruck_sack = line.unwrap();

        match i % 3 {
            0 => intersection = ruck_sack.clone(),
            1 => intersection = intersect_strings(&intersection, &ruck_sack),
            2 => {
                intersection = intersect_strings(&intersection, &ruck_sack);
                total += calculate_priority(intersection.chars().next().unwrap());
            }
            _ => panic!("Unhandled Case"),
        };
    }

    total
}
