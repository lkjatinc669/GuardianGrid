use serde::{Serialize, Deserialize};
use serde_json::Value;

#[derive(Serialize, Deserialize, Debug)]
pub struct ApiResponse {
    pub error: bool,
    pub message: String,
    pub data: Value, 
}