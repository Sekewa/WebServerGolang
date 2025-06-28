# WebServerGolang
This is a simple Golang Web Server.

Before talking about how to run the server note that if you want to use your website on this web server you'll need to put it in a directory named 'static'.

Here is a schema on where to put it :

``` bash
.
├──📂bin
├──📂cmd
│   └──📁web
├──📂example
│   ├──📁css
│   └── index.html
├──📂internal
│   ├──📁 handlers
│   └──📁 utils
├──📁static (👈here is where your site will be)
├──.gitignore
├──go.mod
├──Makefile
└──README.md
```

To run the make command you'll need golang version 1.24.4. 
Once the WebServerGolang file created you can run it, if called you can also put some arguments la '--p' or '-port' to specifie the port in this format : xxxx where x is a number between 0 to 9.
``` bash
> bin\WebServerGolang --p 9090
```

If instead you choose to use Docker, you'll just need to run those too command :

``` bash
docker build --tag your_tag_for_this .

docker run -dit --name your_name_for_this -p xxxx:xxxx your_tag_for_this
```

Don't forget that the right part of the port is the one you choose to expose and use. If you want to change the port used you'll need to modifie the Dockerfile at those line :

``` Dockerfile
EXPOSE xxxx

...

CMD ["--p","xxxx"]
```

This way your web server will be listening to the right port whos exposed by docker !