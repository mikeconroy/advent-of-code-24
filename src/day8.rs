use crate::utils::grid::Grid;
use crate::Day;

pub struct DayEight;
impl Day for DayEight {
    fn part1(&self, input: &str) -> String {
        let grid = parse_input(input);

        return grid.len_to_string();
    }

    fn part2(&self, input: &str) -> String {
        let grid = parse_input(input);
        return grid.len_to_string();
    }
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
        assert_eq!(result, "31");
    }
}
