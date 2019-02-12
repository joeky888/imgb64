Convert image(s) to base64 html data uri.

Support png, gif and jpg.

File type detected by [fil](https://github.com/joeky888/fil)

### Usage

```html
$ imgb64 1.png 2.gif 3.jpg

<img src="data:image/png;base64,iVBORw0KGgoAA"/>
<img src="data:image/gif;base64,R0lGODlhkAGQA"/>
<img src="data:image/jpeg;base64,/9j//gAQTGF2"/>
```

### Install binary

https://github.com/joeky888/imgb64/releases

### Install from source

```sh
go get -u github.com/joeky888/imgb64
```

### Licence

Public domain
