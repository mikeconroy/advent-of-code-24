pub fn part1(input: &String) -> String {
    println!("{}", input);
    return "d1p1".to_string();
}

pub fn part2(input: &String) -> String {
    return "d1p2".to_string();
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils;

    #[test]
    fn test_part1() {
        let test_input = utils::read_file("data/day1_test");
        let result = part1(&test_input);
        assert_eq!(result, "3");
    }

    #[test]
    fn test_part2() {
        let result = part2(&"test".to_string());
        assert_eq!(result, "d1p2");
    }
}
