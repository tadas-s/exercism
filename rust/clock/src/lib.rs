use std::fmt;
use std::fmt::Formatter;

#[derive(Debug, PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut hours = (hours + minutes / 60) % 24;
        let mut minutes = minutes % 60;
        if minutes < 0 { minutes += 60; hours -= 1 };
        if hours < 0 { hours += 24 };

        Self { hours, minutes }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Self::new(self.hours, self.minutes + minutes)
    }
}

impl std::fmt::Display for Clock {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}