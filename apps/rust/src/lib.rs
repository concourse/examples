pub fn add_six(num: f64) -> f64 {
    num + 6.0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add_six_positive_integer() {
        assert_eq!(add_six(5.0), 11.0);
    }

    #[test]
    fn test_add_six_negative_integer() {
        assert_eq!(add_six(-10.0), -4.0);
    }

    #[test]
    fn test_add_six_zero() {
        assert_eq!(add_six(0.0), 6.0);
    }

    #[test]
    fn test_add_six_decimal() {
        assert_eq!(add_six(3.5), 9.5);
    }

    #[test]
    fn test_add_six_large_number() {
        assert_eq!(add_six(1_000_000.0), 1_000_006.0);
    }
}
