#[derive(Debug)]
pub struct Duration {
    seconds: u64
}

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Self { seconds: s }
    }
}

pub trait Planet {
    const EARTH_YEAR_MULTIPLIER: f64;

    fn years_during(d: &Duration) -> f64 {
        (d.seconds as f64) / (Self::EARTH_YEAR_MULTIPLIER * (31557600 as f64))
    }
}

pub struct Mercury;
pub struct Venus;
pub struct Earth;
pub struct Mars;
pub struct Jupiter;
pub struct Saturn;
pub struct Uranus;
pub struct Neptune;

impl Planet for Mercury {
    const EARTH_YEAR_MULTIPLIER: f64 = 0.2408467;
}

impl Planet for Venus {
    const EARTH_YEAR_MULTIPLIER: f64 = 0.61519726;
}

impl Planet for Earth {
    const EARTH_YEAR_MULTIPLIER: f64 = 1.0;
}

impl Planet for Mars {
    const EARTH_YEAR_MULTIPLIER: f64 = 1.8808158;
}

impl Planet for Jupiter {
    const EARTH_YEAR_MULTIPLIER: f64 = 11.862615;
}

impl Planet for Saturn {
    const EARTH_YEAR_MULTIPLIER: f64 = 29.447498;
}

impl Planet for Uranus {
    const EARTH_YEAR_MULTIPLIER: f64 = 84.016846;
}

impl Planet for Neptune {
    const EARTH_YEAR_MULTIPLIER: f64 = 164.79132;
}
