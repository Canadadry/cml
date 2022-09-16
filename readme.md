# Canadadry Markup Language (CML)

Why another json,yml,toml, ... format. Well, firstly for fun and because I can. Secondly I found those language either too verbose, or to complex to edit.

I toke a look at [hashicorp](https://github.com/hashicorp/hcl) which simplify but add some feature I dont want.

Here my target (starting from hashicorp example):

Let's start with this json :

```
{
	"io_mode": "async",
	"service": {
		"http": {
			"web_proxy": {
				"listen_addr": "127.0.0.1:8080",
				"process": {
					"main": {
						"command": ["/usr/local/bin/awesome-app", "server"]
					},
					"mgmt": {
						"command": ["/usr/local/bin/awesome-app", "mgmt"]
					}
				}
			}
		}
	}
}
```

Ok so mine look more like json but with some adjustment :

 - Why json key are a string value ? For me this is to allow multiline key or space inside the key. I think most of the case, keys are always snake_case so I 'll enforce on cml key must be snake_case (plus number but cannot start by)
 - Why the `:` to split keys and values ? It does not allow readability nor does it help to parse the file. After a key there is always a value. So no need for the `:` even more the `:` are now forbidden!! 😱
 - Why the `,` ? It cannot be after the last key-value or are element so it mess with git diff and is a root of a lot of invalid json. Does it help to parse json ? I don't think so, remove and forbidden too!
 - What about the `{` and the `}`, if we remove them we need to use indentation to detect block. I dont find this solution elegant or easier to read and even harder to write. Did I need to use tab ? Or can I mix tab and space ? My IDE show me a Tab is two space but the parser think a tab is Four space.... Too much issue I dont want to deal with this. But we can still improve. Why using this strange char ? Why not a more common and easy to type  like `()` ?
 - What about the `[]` we cannot remove them! If we remove them how to know if we are reading an array or an object, easy an object some kind or array of key-value so it must start we a key where the array start with a value.
 - Lastly I want to add comment in my config file, it's a config file, it will be read be human, we should be able to annotate it.

So in the end we have this :

```
io_mode "async"
service (
	http (
		web_proxy (
			listen_addr "127.0.0.1:8080"
			process (
				main (
					command ("/usr/local/bin/awesome-app" "server")
				)
				mgmt (
					command ("/usr/local/bin/awesome-app" "mgmt")
				)
			)
		)
	)
)
```

But it can be minified in

```
io_mode"async"service(http(web_proxy(listen_addr"127.0.0.1:8080"process(main(command("/usr/local/bin/awesome-app""server"))mgmt(command("/usr/local/bin/awesome-app""mgmt"))))))
```

Why allow this short version ? Well it's only a side effect, we dont really care about white space and line return so with or without it it will work the same.