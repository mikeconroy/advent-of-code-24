mod day1;
mod day2;
mod utils;

fn main() {
    let input = utils::read_file("data/day1");
    println!(
        "Day 1\n\tPart 1: {}\n\tPart 2: {}",
        day1::part1(&input),
        day1::part2(&input)
    );
    println!(
        "Day 2\n\tPart 1: {}\n\tPart 2: {}",
        day2::part1(),
        day2::part2()
    );
}
