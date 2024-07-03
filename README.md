# tat
Image to ascii art converter

```
Tat, short for tet-a-tet is an image to ascii art converter.

Usage:
  tat [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  img         Convert a given input image to ascii art.

Flags:
  -h, --help   help for tat

Use "tat [command] --help" for more information about a command.
```

## Example
```
go run . img ./testdata/gopher.png -x 60 -y 60
```
![example](assets/example.png)
