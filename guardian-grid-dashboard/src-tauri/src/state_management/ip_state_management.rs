use tauri::State;
use std::net::IpAddr;

use crate::state::app_state::AppState;

#[tauri::command]
pub fn set_ip(ip: String, state: State<AppState>) -> Result<(), String> {
    match ip.parse::<IpAddr>() {
        Ok(valid_ip) => {
            let mut stored_ip = state.ip_address.lock().unwrap();
            *stored_ip = Some(valid_ip.to_string());
            Ok(())
        }
        Err(_) => Err("Invalid IP address".into()),
    }
}

#[tauri::command]
pub fn get_ip(state: State<AppState>) -> Option<String> {
    let stored_ip = state.ip_address.lock().unwrap();
    stored_ip.clone()
}