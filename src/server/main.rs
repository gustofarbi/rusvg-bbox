use anyhow::Context;
use serde::{Deserialize, Serialize};
use serde::ser::SerializeStruct;
use usvg::{NodeExt, TreeParsing};
use warp::Filter;

#[tokio::main]
async fn main() {
    env_logger::Builder::new()
        .filter_level(log::LevelFilter::Info)
        .target(env_logger::Target::Stdout)
        .init();

    let bbox = warp::post()
        .and(warp::path("bbox"))
        .and(warp::body::json())
        .map(|request: Request| {
            let content = request.content;
            let bbox = bbox(content);

            return if let Ok(bbox) = bbox {
                let bbox: Response = bbox.into();
                warp::reply::with_status(
                    warp::reply::json(&bbox),
                    warp::http::StatusCode::OK,
                )
            } else {
                let response = ErrorResponse {
                    reason: bbox.unwrap_err(),
                };

                warp::reply::with_status(
                    warp::reply::json(&response),
                    warp::http::StatusCode::INTERNAL_SERVER_ERROR,
                )
            };
        });

    let health = warp::get()
        .and(warp::path("health"))
        .map(|| "ok");

    let routes = bbox.or(health);

    log::info!("Starting server on port 8080");

    warp::serve(routes)
        .run(([0, 0, 0, 0], 8080))
        .await;
}

#[derive(Deserialize)]
struct Request {
    // the actual svg content
    content: String,
}

#[derive(Serialize)]
struct Response {
    x: f64,
    y: f64,
    width: f64,
    height: f64,
}

struct ErrorResponse {
    reason: anyhow::Error,
}

impl Serialize for ErrorResponse {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
        where S: serde::Serializer {
        let mut state = serializer.serialize_struct("ErrorResponse", 1)?;
        state.serialize_field("reason", &self.reason.to_string())?;
        state.end()
    }
}

impl From<usvg::Rect> for Response {
    fn from(rect: usvg::Rect) -> Self {
        Response {
            x: rect.x(),
            y: rect.y(),
            width: rect.width(),
            height: rect.height(),
        }
    }
}

fn bbox(content: String) -> anyhow::Result<usvg::Rect> {
    let mut tree = usvg::Tree::from_str(
        content.as_str(),
        &usvg::Options::default(),
    );

    if tree.is_err() {
        // try to wrap the content in an svg tag
        tree = usvg::Tree::from_str(
            wrap_content(content.clone()).as_str(),
            &usvg::Options::default(),
        );
    }

    let tree = tree.map_err(|e| anyhow::anyhow!(
        "failed to parse {}: {:?}", content, e))?;

    let bbox = tree.root.calculate_bbox()
        .context("failed to calculate bbox")?;

    let rect = bbox.to_rect();

    Ok(rect.context("failed to convert bbox to rect")?)
}

fn wrap_content(content: String) -> String {
    format!("<svg xmlns=\"http://www.w3.org/2000/svg\">{}</svg>", content)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_bbox() {
        // heart.svg,571.6644,122.1925,47.2223,40.7656
        let content = std::fs::read_to_string("testdata/heart.svg").unwrap();
        let bbox = bbox(content).unwrap();

        assert_eq!(float_4_decimals(bbox.x()), 571.6644);
        assert_eq!(float_4_decimals(bbox.y()), 122.1925);
        assert_eq!(float_4_decimals(bbox.width()), 47.2223);
        assert_eq!(float_4_decimals(bbox.height()), 40.7656);
    }

    fn float_4_decimals(f: f64) -> f64 {
        (f * 10000.0).round() / 10000.0
    }
}
