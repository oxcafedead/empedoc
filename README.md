# Empedoc

![Empedocles](./docs/empedocles.png "Empedocles")

A fast and simple static documentation generator based on Markdown written in Go.

This tool generates static documentation based on Markdown sources in a fast, efficient, and lightweight manner.\
It requires no external dependencies and is designed to be both easy to use and quick to set up.

## When to use

You might need this tool if this is relevant:

- you prefer to [get your HTML docs](./docs/gettingstarted.md) rapidly 🔥
- you like a small runtime images (3Mb executable)
- zero dependencies is a plus for you (no runtime dependencies or any third-party libs, everything is already bundled in the executable)
- you don't need fancy features like interactive search, animations etc.
- you need a documentation with no JavaScript
- your documentation is a single README.md file in the project dir and multiple MD/static files in `./docs` or other dir.

## When not to use

You might might want to consider other static documentation generators, if:

- you need more features: search, macros processing, advanced templates/theming, partial customizations.
- you want more interactive docs, containing javascript, animations, etc.

## Contribution

This tool is written in Go. If you found a bug / want some really needed feature, please open the issue for it and we could discuss.

## Stability

This is just an initial implementation. This is not intended to be used in production.\
However, it might be suitable for small projects documentation generation.
It is used already in [https://github.com/oxcafedead/barcode-reader-emulator](https://github.com/oxcafedead/barcode-reader-emulator) and couple of other non-public projects.

Some breaking changes may be introduced in 0.x.x versions, however this will be done only as a last resort.