use crate::Day;
use std::option::Option;

pub struct DayNine;
impl Day for DayNine {
    // 12345 -> 0..111....22222 -> 022111222......
    // 0*0 + 2*1 + 2*2 + 1*3 + 1*4 + 1*5 + 2*6...
    // total_ids = input.len
    // linked list
    // block {
    //  start_index -> The index for the block
    //  ID -> The ID of the block
    //  size -> Size - how many indexes this takes up
    // }
    // start_index=0, ID=0, size=1
    // start_index=1, ID=2, size=2
    // start_index=3, ID=1, size=3
    // start_index=6, ID=2, size=3
    //
    // Perhaps these should be held in a Vec instead...
    fn part1(&self, input: &str) -> String {
        let first_block = parse_input(input);
        println!("{:?}", first_block);
        return input.len().to_string();
    }

    fn part2(&self, input: &str) -> String {
        return input.len().to_string();
    }
}

fn parse_input(input: &str) -> Block {
    let mut first_block = Block {
        id: Some(0),
        size: input
            .chars()
            .nth(0)
            .unwrap()
            .to_string()
            .parse::<i64>()
            .unwrap(),
        next_block: None,
    };
    let mut index = 0;
    let mut prev_block = &mut first_block;
    for size in input.chars().skip(1) {
        if size == '\n' {
            break;
        }
        index += 1;
        let current_size = size.to_string().parse::<i64>().unwrap();

        let new_block = Box::new(Block {
            id: Some(index),
            size: current_size,
            next_block: None,
        });
        prev_block.next_block = Some(new_block);
        prev_block = prev_block.next_block.as_mut().unwrap();
    }
    return first_block;
}

#[derive(Debug)]
struct Block {
    id: Option<i64>,
    size: i64,
    next_block: Option<Box<Block>>,
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils::file::read_file;

    #[test]
    fn test_part1() {
        let test_input = read_file("data/day9_test");
        let day = DayNine {};
        assert_eq!("60", day.part1("12345"));
        let result = day.part1(&test_input);
        assert_eq!(result, "1928");
    }

    #[test]
    fn test_part2() {
        let test_input = read_file("data/day9_test");
        let day = DayNine {};
        let result = day.part2(&test_input);
        assert_eq!(result, "31");
    }
}
