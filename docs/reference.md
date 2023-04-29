# Usage reference

## How it works

The tool just converts all existing MD files into HTML with some default simple template.\
It does not support custom theming or templates, as well as other customization features.

The layout should be the following: a single `README.md` file as entry point and possibly other MD files in another dir, e.g.:

```
.
├── README.md
└── docs/
    ├── about.md
    ├── reference.md
    ├── gettingstarted.md
    ├── logo.png
    ├── screenshot.jpg
    └── ...
```

## Options

To keep it simple, there are only couple of options.

| Argument Name   | Description                                                                                                                                         |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--dir` or `-d` | source directory with Markdown documentation (`./docs` by default).                                                                                 |
| `--out` or `-o` | output directory with HTML documentation (`./docs_gen` by default).                                                                                 |
| `--tpl` or `-t` | HTML template as a string, should contain `{{header}}` placeholder for header, `{{nav}}` for navigation block and `{{main}}` for Markdown content`. |
| `-i`            | the name of index Markdown README file (`README.md` by default).                                                                                    |
 

