use std::collections::{HashMap, HashSet};

use crate::Day;

pub struct DaySix;
impl Day for DaySix {
    fn part1(&self, input: &str) -> String {
        let (grid, start) = parse_input(input);
        let mut visited: HashSet<(i32, i32)> = HashSet::new();
        visited.insert(start);
        let mut dir = DIR::UP;
        let (mut new_pos, mut new_val) = (start, '.');
        while new_val != ' ' {
            let old_pos = new_pos;
            (new_pos, new_val) = next_pos(&grid, &dir, new_pos);
            if new_val == '#' {
                dir = turn_clockwise(&dir);
                new_pos = old_pos;
            } else if new_val != ' ' {
                visited.insert(new_pos);
            }
        }
        // println!("{:?}", visited);
        let count = visited.len();
        return count.to_string();
    }

    fn part2(&self, input: &str) -> String {
        let (mut grid, start) = parse_input(input);
        let mut visited: HashSet<(i32, i32)> = HashSet::new();
        let mut dir = DIR::UP;
        let (mut new_pos, mut new_val) = (start, '.');
        while new_val != ' ' {
            let old_pos = new_pos;
            (new_pos, new_val) = next_pos(&grid, &dir, new_pos);
            if new_val == '#' {
                dir = turn_clockwise(&dir);
                new_pos = old_pos;
            } else if new_val != ' ' {
                visited.insert(new_pos);
            }
        }
        // println!("Visited: {:?}", visited);
        let mut obstacles: HashSet<(i32, i32)> = HashSet::new();

        // let total = visited.len();
        // let mut count = 0;
        let mut prev_pos = start;
        for pos in &visited {
            // println!("{}/{}", count, total);
            // count += 1;
            // let mut updated_grid: Grid = grid.clone();
            grid.insert(prev_pos, '.');
            grid.insert(*pos, '#');
            // updated_grid.insert((pos.0, pos.1), '#');
            if is_loop(&grid, start) {
                obstacles.insert((pos.0, pos.1));
            }
            prev_pos = *pos;
        }

        return obstacles.len().to_string();
    }
}

fn is_loop(grid: &Grid, pos: (i32, i32)) -> bool {
    let mut dir = DIR::UP;
    let (mut new_pos, mut new_val) = (pos, '.');
    let mut visited: HashSet<((i32, i32), DIR)> = HashSet::new();
    while new_val != ' ' {
        let old_pos = new_pos;
        (new_pos, new_val) = next_pos(&grid, &dir, new_pos);
        if new_val == '#' {
            dir = turn_clockwise(&dir);
            new_pos = old_pos;
            if has_visited(&visited, new_pos, dir.clone()) {
                return true;
            }
            visited.insert((new_pos, dir.clone()));
        }
    }
    return false;
}

fn has_visited(visited: &HashSet<((i32, i32), DIR)>, pos: (i32, i32), dir: DIR) -> bool {
    if visited.contains(&(pos, dir)) {
        return true;
    }
    return false;
}

fn turn_clockwise(dir: &DIR) -> DIR {
    match dir {
        DIR::UP => DIR::RIGHT,
        DIR::RIGHT => DIR::DOWN,
        DIR::DOWN => DIR::LEFT,
        DIR::LEFT => DIR::UP,
    }
}

#[derive(Eq, Hash, PartialEq, Clone, Debug)]
enum DIR {
    UP,
    DOWN,
    LEFT,
    RIGHT,
}

fn next_pos(grid: &Grid, dir: &DIR, pos: (i32, i32)) -> ((i32, i32), char) {
    let new_pos: (i32, i32);
    match dir {
        DIR::UP => new_pos = (pos.0, pos.1 - 1),
        DIR::DOWN => new_pos = (pos.0, pos.1 + 1),
        DIR::LEFT => new_pos = (pos.0 - 1, pos.1),
        DIR::RIGHT => new_pos = (pos.0 + 1, pos.1),
    }
    return (new_pos, *grid.get(&new_pos).unwrap_or(&' '));
}

type Grid = HashMap<(i32, i32), char>;

fn parse_input(input: &str) -> (Grid, (i32, i32)) {
    let mut grid: Grid = HashMap::new();
    let (mut x, mut y) = (0, 0);
    let mut start: (i32, i32) = (0, 0);
    for line in input.lines() {
        for cell in line.chars() {
            grid.insert((x, y), cell);
            if cell == '^' {
                start = (x, y);
            }
            x += 1;
        }
        x = 0;
        y += 1;
    }
    return (grid, start);
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils::file::read_file;

    #[test]
    fn test_part1() {
        let day = DaySix {};
        let input = read_file("data/day6_test");
        let result = day.part1(&input);
        assert_eq!(result, "41");
    }

    #[test]
    fn test_part2() {
        let day = DaySix {};
        let input = read_file("data/day6_test");
        let result = day.part2(&input);
        assert_eq!(result, "6");
    }
}
