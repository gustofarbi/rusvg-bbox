use std::io::Write;
use std::path::Path;

use anyhow::Result;
use usvg::{NodeExt, TreeParsing};

fn main() {
    let files = std::fs::read_dir("testdata").unwrap();
    let mut results_file = std::fs::File::create("results-rust.csv").unwrap();
    results_file
        .write_all("file,x,y,w,h\n".as_bytes())
        .unwrap();
    for file in files {
        let file = file.unwrap();
        let path = file.path();
        if path.is_file() {
            match bbox(path.to_str().unwrap().to_string()) {
                Ok(tree) => {
                    let bbox = tree.root.calculate_bbox().unwrap();
                    results_file
                        .write_all(
                            format!(
                                "{},{:.4},{:.4},{:.4},{:.4}\n",
                                Path::new(path.to_str().unwrap())
                                    .file_name()
                                    .unwrap()
                                    .to_str()
                                    .unwrap(),
                                bbox.x(),
                                bbox.y(),
                                bbox.width(),
                                bbox.height(),
                            )
                                .as_bytes(),
                        )
                        .unwrap();
                }
                Err(e) => {
                    println!("{}: {:?}", path.to_str().unwrap(), e);
                }
            }
        }
    }
}

fn bbox(path: String) -> Result<usvg::Tree> {
    let content = std::fs::read_to_string(path.clone())?;
    let mut tree = usvg::Tree::from_str(content.as_str(), &usvg::Options::default());

    if tree.is_ok() {
        return Ok(tree.unwrap());
    }

    // try to wrap the content in an svg tag
    tree = usvg::Tree::from_str(wrap_content(content).as_str(), &usvg::Options::default());

    if tree.is_err() {
        return Err(anyhow::anyhow!("failed to parse {}", path));
    }

    Ok(tree.unwrap())
}

fn wrap_content(content: String) -> String {
    format!("<svg xmlns=\"http://www.w3.org/2000/svg\">{}</svg>", content)
}
