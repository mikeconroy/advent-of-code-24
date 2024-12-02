use crate::traits::Day;

pub struct DayTwo;

impl Day for DayTwo {
    fn part1(&self, input: &str) -> String {
        let mut safe_count = 0;
        for report in input.lines() {
            let mut is_safe = true;
            let levels = report.split(" ").collect::<Vec<&str>>();
            let is_increasing: bool;
            if levels[0].parse::<i32>().unwrap() > levels[1].parse::<i32>().unwrap() {
                is_increasing = false;
            } else {
                is_increasing = true;
            }
            let mut prev_level: i32 = levels[0].parse().unwrap();
            for index in 1..levels.len() {
                let current_level: i32 = levels[index].parse().unwrap();
                let diff: i32;
                if is_increasing {
                    diff = current_level - prev_level;
                } else {
                    diff = prev_level - current_level;
                }
                if diff <= 0 || diff > 3 {
                    is_safe = false;
                }
                prev_level = current_level;
            }
            if is_safe {
                safe_count += 1
            }
        }
        return safe_count.to_string();
    }

    fn part2(&self, input: &str) -> String {
        let mut safe_count = 0;
        for line in input.lines() {
            let report: Vec<i32> = line
                .split_whitespace()
                .filter_map(|s| s.parse::<i32>().ok())
                .collect();
            if is_safe(&report) {
                safe_count += 1;
                continue;
            }

            for index in 0..report.len() {
                let mut dampened_report = report.clone();
                dampened_report.remove(index);
                if is_safe(&dampened_report) {
                    safe_count += 1;
                    break;
                }
            }
        }
        return safe_count.to_string();
    }
}

fn is_safe(report: &Vec<i32>) -> bool {
    let is_increasing: bool;
    if report[0] > report[1] {
        is_increasing = false;
    } else {
        is_increasing = true;
    }
    let mut prev_level: i32 = report[0];
    for index in 1..report.len() {
        let current_level = report[index];
        let diff: i32;
        if is_increasing {
            diff = current_level - prev_level;
        } else {
            diff = prev_level - current_level;
        }
        if diff <= 0 || diff > 3 {
            return false;
        }
        prev_level = current_level;
    }
    return true;
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils;

    #[test]
    fn test_part1() {
        let test_input = utils::read_file("data/day2_test");
        let day = DayTwo {};
        let mut result = day.part1("7 10 12 14");
        assert_eq!(result, "1");
        result = day.part1(&test_input);
        assert_eq!(result, "2");
    }

    #[test]
    fn test_part2() {
        let test_input = utils::read_file("data/day2_test");
        let day = DayTwo {};
        let mut result = day.part2(&test_input);
        assert_eq!(result, "4");
        result = day.part2("10 17 9 8 7");
        assert_eq!(result, "1");
        result = day.part2("10 11 9 8 7");
        assert_eq!(result, "1");
    }
}
