use std::sync::Mutex;

pub struct AppState {
    pub ip_address: Mutex<Option<String>>,
}