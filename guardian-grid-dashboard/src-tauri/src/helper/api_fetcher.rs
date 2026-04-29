use reqwest;
use serde_json::Value;
use crate::models::response_model::ApiResponse;

pub async fn get_response_from_api(url: String) -> ApiResponse {
    let client = reqwest::Client::new();
    
    // Attempt the fetch
    match client.get(url).send().await {
        Ok(response) => {
            // Attempt to parse JSON
            match response.json::<Value>().await {
                Ok(json) => ApiResponse {
                    error: false,
                    message: "Success".into(),
                    data: json,
                },
                Err(e) => ApiResponse {
                    error: true,
                    message: format!("JSON Parse Error: {}", e),
                    data: Value::Null,
                },
            }
        }
        Err(e) => ApiResponse {
            error: true,
            message: format!("Network Error: {}", e),
            data: Value::Null,
        },
    }
}