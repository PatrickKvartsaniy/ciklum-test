package main

import (
	"fmt"
	"net/http"
	"runtime"
)

// PrintMemUsage prints RAM usage
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// RenderPage - index page rendering
func RenderPage(w http.ResponseWriter) {
	fmt.Fprint(w, tmpl)
}

const tmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <title>CSV Reader</title>
</head>
<body>
    <section id="cover">
        <div id="cover-caption">
            <div id="container" class="container">
                <div class="row text-white">
                    <div class="col-sm-10 offset-sm-1 text-center">
                        <h1 class="display-3">CSV Reader</h1>
                        <div class="info-form">
                            <form class="form-inline justify-content-center"  action="/" method="POST" enctype="multipart/form-data">
                                <div class="form-group">
                                    <input type="file" class="form-control" name="uploadfile" >
                                </div>
                                <button type="submit" class="btn btn-success ">Read</button>
                            </form>
                        </div>
                        <br>
                    </div>
                </div>
            </div>
        </div>
    </section>
</body>
</html>
`
