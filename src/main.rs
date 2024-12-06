mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod traits;
mod utils;

use crate::day1::DayOne;
use crate::day2::DayTwo;
use crate::day3::DayThree;
use crate::day4::DayFour;
use crate::day5::DayFive;
use crate::day6::DaySix;
use crate::traits::Day;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let days: Vec<Box<dyn Day>> = vec![
        Box::new(DayOne),
        Box::new(DayTwo),
        Box::new(DayThree),
        Box::new(DayFour),
        Box::new(DayFive),
        Box::new(DaySix),
    ];

    if args.len() > 1 {
        let day = &args[1];
        let day_index = day.parse::<usize>().unwrap_or(1) - 1;
        if day_index < days.len() {
            let day_instance = &days[day_index];
            let input = utils::read_file(format!("data/day{}", day).as_str());
            print_day(day, day_instance.part1(&input), day_instance.part2(&input));
        } else {
            println!("Day {} Not Found", day);
        }
    } else {
        for (day, day_instance) in days.iter().enumerate() {
            let input = utils::read_file(format!("data/day{}", day + 1).as_str());
            print_day(
                (day + 1).to_string().as_str(),
                day_instance.part1(&input),
                day_instance.part2(&input),
            );
        }
    }
}

fn print_day(day: &str, p1: String, p2: String) {
    println!("Day {}\n\tPart 1: {}\n\tPart 2: {}", day, p1, p2);
}
