[YouPlot](https://github.com/red-data-tools/YouPlot)

*Description*: A command line tool that draw plots on the terminal.

*Labels*: #Ruby #Charts

*Examples*:

```bash
$ curl -sL https://git.io/ISLANDScsv \
    | sort -nk2 -t, \
    | tail -n15 \
    | uplot bar -d, -t "Areas of the World's Major Landmasses"
$ seq 1 100
    | uplot line
$ echo -e "from numpy import random;" \
          "n = random.randn(10000);"  \
          "print('\\\n'.join(str(i) for i in n))" \
    | python \
    | uplot hist --nbins 20
```
