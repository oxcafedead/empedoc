# Getting started

To convert your existing documentation, just download the binary and run it in the project directory:

```sh
# Download
wget https://github.com/oxcafedead/empedoc/releases/download/v1.0.0/empedoc-linux-amd64
# Exec permissions
chmod +x empedoc-linux-amd64
# Run it!
./empedoc-linux-amd64
```

or in Windows:

```
# Download
Invoke-WebRequest -Uri "https://github.com/oxcafedead/empedoc/releases/download/v1.0.0/empedoc-windows-amd64.exe" -OutFile "empedoc-windows-amd64.exe"
# Exec permissions
Set-ItemProperty -Path ".\empedoc-windows-amd64.exe" -Name "IsReadOnly" -Value $false
Set-ItemProperty -Path ".\empedoc-windows-amd64.exe" -Name "Attributes" -Value "Normal"
# Run it!
.\empedoc-windows-amd64.exe
```

It will produce static HTML files and other required resources to `./docs_gen` dir in your project.\
You can open `index.html` HTML file there to see how it looks.

This tools allows only minimal customizations, preferring simplicity over extensibility.
You might want to look for available command line arguments in [reference](./reference.md) section.
