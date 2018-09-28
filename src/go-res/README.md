# Go Resource Packager

`go-res` is a minimalistic resource packager for creating battery-included Go applications.  It is very simple comparing to the popular [go-bindata](https://github.com/go-bindata/go-bindata) utility: `go-bindata` has nearly 100 funcs scattered in dozens of Go source file while `go-res` has only one source file with 5 funcs, of which only 2 are exported!

## The Idea

`go-res` provides a simple way to append all files in a specified directory (and all its sub-directories) at the end of the Go executable as tar.gz data, which in turn can be extracted on-demand.  In another word, it is a backpack where the content must be taken out when use.

## The APIs

### Pack

    func Pack(root string) error

Pack collect all files under directory `root` and its sub-directories, append them as tar.gz data at the end of the running application, then add a signature at the end to make the Pack action idempotent -- you can pack any directory many times, only the last operation's result is kept, all previous ones are discarded.

### Extract

    func Extract(path string, policy ExtractPolicy) error

Extract extracts embeded resources to the `path` specified. `policy` is used to control content overwriting behavior:

|policy  |logic  |
|-- |-- |
|`NoOverwrite`|if a file exists at destination location, it will _not_ be overwritten|
|`OverwriteIfNewer`|_only_ overwrite a file if the one in resource pack is _newer_|
|`AlwaysOverwrite`|_always_ overwrite file at destination with the on in resource pack|
|`Verbatim`|if `path` exists, _remove_ it with _all_ its contents, then extract resources to (newly created) `path`|

**CAUTION**： be careful when using the `Verbatim` policy, as it will remove the path **completely**, if you specify a wrong path, data loss will occur. As a protection, the `Extract()` function prohibits empty or root path `("/")`.

## The Use Case

Usually `go-res` is used at the beginning of the application's main function:

    func main() {
        pack := flag.String("pack", "", "pack specified directory as attached resources")
        webroot := flag.String("wwwroot", "../webroot", "root directory for resources")
        flag.Parse()
        if *pack != nil {
            assert(res.Pack(*pack))
            return
        }
        res.Extract(*webroot, res.OverwriteIfNewer)
    }

Resources are extracted on the launch of application. Afterwards, the application just use resources as normal files.

## The Pros & Cons

### Pros

* simple
* no runtime performance penalty
* debugging friendly, changes can be made after application is deployed

### Cons

* not compatible with executable packers such as UPX. 
* not thoroughly tested, comparing to [go-bindata](https://github.com/go-bindata/go-bindata) etc.