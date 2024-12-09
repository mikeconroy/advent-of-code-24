use crate::utils::grid::Grid;
use crate::utils::grid::Point;
use crate::Day;
use std::collections::HashMap;
use std::collections::HashSet;
use std::io;
use std::io::Write;

pub struct DayEight;
impl Day for DayEight {
    fn part1(&self, input: &str) -> String {
        let grid = parse_input(input);
        let freqs = get_freqs(&grid);
        let mut antinodes: HashSet<Point> = HashSet::new();
        for freq in freqs {
            antinodes.extend(locate_antinodes(&freq.1, &grid));
        }
        return antinodes.len().to_string();
    }

    // This is calculating every single line (antinode) between every single new point.
    // We should only calculate all antinodes in a line from the original 2.
    // I.E. don't add Antinodes then calculate new lines.
    // Change locate_antinodes function so when it adds new antinode points it loops around
    // increasing X&Y until we reach the end of the map space.
    // Then change this function to not 'loop' repeatedly - just call each 'freq pair' once.
    // 869 is too low
    // 1011 is too high
    fn part2(&self, input: &str) -> String {
        let grid = parse_input(input);
        let freqs = get_freqs(&grid);

        let mut nodes: HashSet<Point> = HashSet::new();
        for (_key, val) in &freqs.clone() {
            nodes.extend(get_lines(val, &grid));
        }
        return nodes.len().to_string();
    }
}

fn get_lines(freq: &HashSet<Point>, grid: &Grid) -> HashSet<Point> {
    let mut nodes: HashSet<Point> = HashSet::new();
    for (idx, i) in freq.clone().iter().enumerate() {
        for j in freq.clone().iter().skip(idx + 1) {
            if i == j {
                continue;
            }
            let dy = i.y - j.y;
            let dx = i.x - j.x;
            // let m: f64 = ((i.y as f64 - j.y as f64) / (i.x as f64 - j.x as f64)) as f64;
            let c: f64 = i.y as f64 - ((dy * i.x) as f64 / dx as f64);
            // println!("y=({}/{})x + {}", dy, dx, c);

            for x in 0..grid.col_count {
                let mx: f64 = (dy * x) as f64 / dx as f64;
                let y: f64 = mx + c;
                if y >= 0.0 && y < grid.row_count as f64 {
                    if y.fract() <= 0.01 {
                        nodes.insert(Point { x, y: y as i64 });
                    } else if y.fract() >= 0.999 {
                        nodes.insert(Point {
                            x,
                            y: (y as i64) + 1,
                        });
                    }
                }
            }
        }
    }
    return nodes;
}

fn locate_antinodes(freq: &HashSet<Point>, grid: &Grid) -> HashSet<Point> {
    let mut antinodes: HashSet<Point> = HashSet::new();
    for i in freq.iter() {
        for j in freq.iter() {
            if i == j {
                continue;
            }

            let mut new_antinode = Point {
                x: (i.x - j.x) + i.x,
                y: (i.y - j.y) + i.y,
            };
            if !grid.get_char(&new_antinode).eq(&' ') {
                antinodes.insert(new_antinode);
            }

            new_antinode = Point {
                x: (j.x - i.x) + j.x,
                y: (j.y - i.y) + j.y,
            };
            if !grid.get_char(&new_antinode).eq(&' ') {
                antinodes.insert(new_antinode);
            }
        }
    }
    return antinodes;
}

fn get_freqs(grid: &Grid) -> HashMap<char, HashSet<Point>> {
    let mut freqs: HashMap<char, HashSet<Point>> = HashMap::new();
    for key in grid.cells.keys() {
        let freq = *grid.get_char(key);
        if freq != '.' {
            let set = freqs.entry(freq).or_insert(HashSet::new());
            set.insert(key.clone());
        }
    }
    // println!("{:?}", freqs);
    return freqs;
}

fn parse_input(input: &str) -> Grid {
    let mut grid: Grid = Grid::new();
    grid.load(input);
    return grid;
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils::file::read_file;

    #[test]
    fn test_part1() {
        let test_input = read_file("data/day8_test");
        let day = DayEight {};
        let result = day.part1(&test_input);
        assert_eq!(result, "14");
    }

    #[test]
    fn test_part2() {
        let test_input = read_file("data/day8_test");
        let day = DayEight {};
        let result = day.part2(&test_input);
        assert_eq!(result, "34");
    }
}
