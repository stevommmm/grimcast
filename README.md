# Grimcast

Hacky combination of `grim` for webcasting the desktop to a browser tab (for sharing in teams etc)

Requires a geometry string to be passed on stdin in the grim format `<x>,<y> <width>x<height>`. The easiest way to do this is to combine `slurp` with the call.

```bash
slurp | grimcast
```
