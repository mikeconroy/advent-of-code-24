use std::collections::HashMap;

fn parse_input(input: &String) -> (Vec<i32>, Vec<i32>) {
    let mut l1: Vec<i32> = Vec::new();
    let mut l2: Vec<i32> = Vec::new();

    for line in input.lines() {
        let split = line.split("   ").collect::<Vec<&str>>();
        l1.push(split[0].parse().unwrap());
        l2.push(split[1].parse().unwrap());
    }
    return (l1, l2);
}

pub fn part1(input: &String) -> String {
    let (mut l1, mut l2) = parse_input(input);
    l1.sort();
    l2.sort();
    let mut result = 0;
    for i in 0..l1.len() {
        if l1[i] >= l2[i] {
            result += l1[i] - l2[i];
        } else {
            result += l2[i] - l1[i];
        }
    }

    return result.to_string();
}

pub fn part2(input: &String) -> String {
    let (l1, l2) = parse_input(input);
    let mut l2_counts: HashMap<i32, i32> = HashMap::new();
    for el in l2 {
        let current_count = l2_counts.get(&el);
        let new_count = current_count.unwrap_or(&0) + 1;
        l2_counts.insert(el, new_count);
    }

    let mut result = 0;
    for el in l1 {
        let rh_count = l2_counts.get(&el);
        result += el * rh_count.unwrap_or(&0);
    }

    return result.to_string();
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils;

    #[test]
    fn test_part1() {
        let test_input = utils::read_file("data/day1_test");
        let result = part1(&test_input);
        assert_eq!(result, "11");
    }

    #[test]
    fn test_part2() {
        let test_input = utils::read_file("data/day1_test");
        let result = part2(&test_input);
        assert_eq!(result, "31");
    }
}
