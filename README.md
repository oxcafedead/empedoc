# Empedoc

![Empedocles](./docs/assets/empedocles.png "Empedocles, generated by AI")

A fast and simple static documentation generator based on Markdown written in Go.

This tool generates static documentation based on Markdown sources in a fast, efficient, and lightweight manner.\
It requires no external dependencies and is designed to be both easy to use and quick to set up.

## When to use

You might need this tool if this is relevant:

- you prefer to [get your HTML docs](./docs/gettingstarted.md) rapidly 🔥
- you're keen on small runtime images (3Mb)
- zero dependencies is a plus for you (old-school standalone executable)
- you don't need fancy features like interactive search, animations etc.
- no JavaScript
- your documentation consists of pure Markdown files.

## When not to use

You might might want to consider other static documentation generators, if:

- you need more features: search, macros processing, advanced templates/theming, partial customizations
- you want more interactive docs, containing javascript, animations, etc.

## Contribution

This tool is written in Go. If you found a bug / want some really needed feature, please open the issue for it and we could discuss.

## Stability

This is just an initial implementation. This is not intended to be used in production.\
However, it might be suitable for small projects documentation generation.
It is used already in [https://github.com/oxcafedead/barcode-reader-emulator](https://github.com/oxcafedead/barcode-reader-emulator) and couple of other non-public projects.

Some breaking changes may be introduced in 0.x.x versions, however this will be done only as a last resort.
