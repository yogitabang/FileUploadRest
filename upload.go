package main
 
import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strconv"
    "mime/multipart"
)
 
func upload(w http.ResponseWriter, r *http.Request) {
 
	if r.Method == "POST" {
		// Post small file
		/*
		file, handler, err := r.FormFile("uploadfile")
	       	if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	 
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	 
		io.Copy(f, file)
		*/
		
		const _24K = (1 << 10) * 24 
		err := r.ParseMultipartForm(_24K) 
		if err != nil {  
			fmt.Println(http.StatusInternalServerError)  
			return  
		}  
		for _, fheaders := range r.MultipartForm.File {  
			for _, hdr := range fheaders {  
				//open uploaded  
				var infile multipart.File 
				infile, err := hdr.Open() 
				if nil != err {  
					fmt.Println(http.StatusInternalServerError)  
				 	return  
				}  
				//open destination  
				//var outfile *os.File
				outfile, err := os.Create("./test/" + hdr.Filename)  
				if nil != err {  
				 	fmt.Println(http.StatusInternalServerError)  
				 	return  
				}  
				//buffer copy  
				//var written int64 
				written, err := io.Copy(outfile, infile) 
				if nil != err {  
				 	fmt.Println(http.StatusInternalServerError)  
				 	return  
				}  
				w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written)))) 
				fmt.Println("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))) 
			}  
		}  
	 
	} else {
		fmt.Println("Unknown HTTP " + r.Method + "  Method")
	}
}
 
func main() {
	http.HandleFunc("/upload", upload)
	fs := http.FileServer(http.Dir("./test"))
	http.Handle("/files/", http.StripPrefix("/files", fs))
	http.ListenAndServe(":9090", nil) // setting listening port
}



