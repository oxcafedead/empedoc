# Getting started

To convert your existing documentation, just install the tool using the latest binary:

```sh
wget https://github.com/oxcafedead/empedoc/releases/download/v0.1.2/empedoc-linux-amd64
chmod +x empedoc-linux-amd64
```

And then just run it from the project dir:

```sh
cd project/path
path/to/binary/empedoc-linux-amd64
```

or in Windows:
```
cd d:/project/path
d:/path/to/binary/empedoc-windows-amd64.exe
```

It will produce static HTML files and other required resources to `./docs_gen` dir in your project.\
You can open `README.md.html` HTML file there to see how it looks.

This tools allows only minimal customizations, preferring simplicity over extensibility.
You might want to look for available command line arguments in [reference](./reference.md) section.