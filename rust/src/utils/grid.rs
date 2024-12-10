use std::collections::HashMap;
#[derive(Hash, Eq, PartialEq, Debug, Clone)]
pub struct Point {
    pub x: i64,
    pub y: i64,
}

pub struct Grid {
    pub cells: HashMap<Point, char>,
    pub row_count: i64,
    pub col_count: i64,
}
impl Grid {
    pub fn new() -> Self {
        Grid {
            cells: HashMap::new(),
            row_count: 0,
            col_count: 0,
        }
    }

    pub fn insert(&mut self, point: Point, value: char) {
        self.cells.insert(point, value);
    }

    pub fn get_char(&self, point: &Point) -> &char {
        return self.cells.get(point).unwrap_or(&' ');
    }

    pub fn load(&mut self, input: &str) {
        let mut x = 0;
        let mut y = 0;
        for line in input.lines() {
            x = 0;
            for val in line.chars() {
                self.insert(Point { x, y }, val);
                x += 1;
            }
            y += 1;
        }
        self.col_count = x;
        self.row_count = y;
    }
}
