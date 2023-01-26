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
    const ORBIT_MULTIPLIER: f64;

    fn years_during(d: &Duration) -> f64 {
        (d.seconds as f64) / (Self::ORBIT_MULTIPLIER * (31557600 as f64))
    }
}

macro_rules! planet {
    ($struct_name:ident, $orbit_multiplier:expr) => {
        pub struct $struct_name;
        impl Planet for $struct_name {
            const ORBIT_MULTIPLIER: f64 = $orbit_multiplier;
        }
    };
}

planet!(Mercury, 0.2408467);
planet!(Venus, 0.61519726);
planet!(Earth, 1.0);
planet!(Mars, 1.8808158);
planet!(Jupiter, 11.862615);
planet!(Saturn, 29.447498);
planet!(Uranus, 84.016846);
planet!(Neptune, 164.79132);