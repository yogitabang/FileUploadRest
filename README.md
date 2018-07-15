# FileUploadRest

To post/upload the file use the following
curl -X POST -H 'Content-Type: multipart/form-data' -F uploadfile=@filename http://127.0.0.1:9090/upload


To get/download the file sue the following
curl -X GET  http://127.0.0.1:9090/files/filename

Note:
test is the folder on server where all the files are stored
