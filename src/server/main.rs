use anyhow::Context;
use serde::{Deserialize, Serialize};
use usvg::{NodeExt, TreeParsing};
use warp::Filter;

#[tokio::main]
async fn main() {
    let bbox = warp::post()
        .and(warp::path("bbox"))
        .and(warp::body::json())
        .map(|request: Request| {
            let content = request.content;
            let bbox = bbox(content);

            match bbox {
                Ok(bbox) => {
                    let response: Response = bbox.into();
                    warp::reply::json(&response)
                }
                Err(e) => warp::reply::json(&e.to_string()),
            }
        });

    let health = warp::get()
        .and(warp::path("health"))
        .map(|| "ok");

    let routes = bbox.or(health);

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

fn bbox(path: String) -> anyhow::Result<usvg::Rect> {
    let content = std::fs::read_to_string(path.clone())?;
    let mut tree = usvg::Tree::from_str(content.as_str(), &usvg::Options::default());

    if tree.is_err() {
        // try to wrap the content in an svg tag
        tree = usvg::Tree::from_str(wrap_content(content).as_str(), &usvg::Options::default());
    }

    let tree = tree.map_err(|e| anyhow::anyhow!("failed to parse {}: {:?}", path, e))?;
    let bbox = tree.root.calculate_bbox().context("failed to calculate bbox")?;
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
        let bbox = bbox("testdata/heart.svg".to_string()).unwrap();
        assert_eq!(float_4_decimals(bbox.x()), 571.6644);
        assert_eq!(float_4_decimals(bbox.y()), 122.1925);
        assert_eq!(float_4_decimals(bbox.width()), 47.2223);
        assert_eq!(float_4_decimals(bbox.height()), 40.7656);
    }

    fn float_4_decimals(f: f64) -> f64 {
        (f * 10000.0).round() / 10000.0
    }
}
