use crate::Day;
use std::collections::HashMap;

pub struct DayFour;
impl Day for DayFour {
    fn part1(&self, input: &str) -> String {
        let (grid, cols, rows) = str_to_grid(input);
        let mut count = 0;
        for y in 0..rows {
            for x in 0..cols {
                count += count_xmas_at(&grid, x, y);
            }
        }
        return count.to_string();
    }

    fn part2(&self, input: &str) -> String {
        let (grid, cols, rows) = str_to_grid(input);
        let mut count = 0;
        for y in 0..rows {
            for x in 0..cols {
                if compare_at(&grid, x, y, 'A') == 1 {
                    if is_x(&grid, x, y) {
                        count += 1;
                    }
                }
            }
        }
        return count.to_string();
    }
}

fn is_x(grid: &HashMap<(i32, i32), char>, x: i32, y: i32) -> bool {
    let ne = *grid.get(&(x + 1, y - 1)).unwrap_or(&' ');
    let nw = *grid.get(&(x - 1, y - 1)).unwrap_or(&' ');
    let se = *grid.get(&(x + 1, y + 1)).unwrap_or(&' ');
    let sw = *grid.get(&(x - 1, y + 1)).unwrap_or(&' ');

    if !((ne == 'M' && sw == 'S') || (ne == 'S' && sw == 'M')) {
        return false;
    }

    if !((nw == 'M' && se == 'S') || (nw == 'S' && se == 'M')) {
        return false;
    }

    return true;
}

fn count_xmas_at(grid: &HashMap<(i32, i32), char>, x: i32, y: i32) -> i32 {
    let chars: Vec<char> = vec!['X', 'M', 'A', 'S'];
    let mut count = 0;
    let (
        mut north,
        mut south,
        mut east,
        mut west,
        mut north_east,
        mut north_west,
        mut south_east,
        mut south_west,
    ) = (0, 0, 0, 0, 0, 0, 0, 0);
    for index in 0..chars.len() {
        let delta = index as i32;
        north += compare_at(grid, x, y - delta, chars[index]);
        south += compare_at(grid, x, y + delta, chars[index]);
        east += compare_at(grid, x + delta, y, chars[index]);
        west += compare_at(grid, x - delta, y, chars[index]);
        north_east += compare_at(grid, x + delta, y - delta, chars[index]);
        north_west += compare_at(grid, x - delta, y - delta, chars[index]);
        south_east += compare_at(grid, x + delta, y + delta, chars[index]);
        south_west += compare_at(grid, x - delta, y + delta, chars[index]);
        // println!(
        // "N: {} S: {} E: {} W: {} NE: {} NW: {} SE: {} SW: {}",
        // north, south, east, west, north_east, north_west, south_east, south_west
        // );
    }

    if north == chars.len() as i32 {
        count += 1;
    }
    if south == chars.len() as i32 {
        count += 1;
    }
    if east == chars.len() as i32 {
        count += 1;
    }
    if west == chars.len() as i32 {
        count += 1;
    }
    if north_east == chars.len() as i32 {
        count += 1;
    }
    if north_west == chars.len() as i32 {
        count += 1;
    }
    if south_east == chars.len() as i32 {
        count += 1;
    }
    if south_west == chars.len() as i32 {
        count += 1;
    }
    return count;
}

fn compare_at(grid: &HashMap<(i32, i32), char>, x: i32, y: i32, target: char) -> i32 {
    let val = grid.get(&(x, y)).unwrap_or(&' ');
    // println!("X: {} Y: {} Val: {} Target: {}", x, y, val, target);
    if *val == target {
        return 1;
    } else {
        return 0;
    }
}

fn str_to_grid(input: &str) -> (HashMap<(i32, i32), char>, i32, i32) {
    let mut grid: HashMap<(i32, i32), char> = HashMap::new();
    let (mut x, mut y) = (0, 0);
    for line in input.lines() {
        x = 0;
        for char in line.chars() {
            grid.insert((x, y), char);
            x += 1;
        }
        y += 1;
    }
    return (grid, x, y);
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils;

    #[test]
    fn test_part1() {
        let day = DayFour {};
        let input = utils::read_file("data/day4_test");
        let result = day.part1(&input);
        assert_eq!(result, "18");
    }

    #[test]
    fn test_part2() {
        let day = DayFour {};
        let input = utils::read_file("data/day4_test");
        let result = day.part2(&input);
        assert_eq!(result, "9");
    }
}
