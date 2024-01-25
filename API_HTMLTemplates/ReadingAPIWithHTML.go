// This is a standalone program which gives you an idea on how you can read a directory which has json files and parse them one by one
// And send the parsed Json file to the API through a POST Request which has converted json file to base64 encoded format

package goprograms

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/upload", fileUploadHandler)

	port := ":9090"
	fmt.Printf("Server running on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

// You can use constant and embed HTML content inside like below.
// Using <script> you can embed any javascript code for alerts and other asynchronous call handling.
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>File Upload with Bootstrap</title>
    <!-- Include Bootstrap CSS from CDN -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
</head>
<body>

<div class="container mt-5">
    <h2 class="mb-4">File Upload</h2>
    <form action="/upload" method="post" enctype="multipart/form-data">
        <div class="form-group">
            <label for="file">Choose Json File(s) to upload:</label>
            <input type="file" class="form-control-file" name="file" id="file" multiple>
        </div>
        <button type="submit" class="btn btn-primary">Upload</button>
    </form>
</div>

<!-- Include Bootstrap JS and Popper.js from CDN -->
<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>

</body>
</html>
`

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	// Display the upload form on the homepage. You can add any bootstrapping or styling you need here.
	tmpl, err := template.New("index").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// fileUploadHandler function for parsing file one by one and to call the API
func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the file headers
	files := r.MultipartForm.File["file"]

	// Iterate through each file
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error opening file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Process file contents
		base64Data, err := processFileContents(file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing file: %s", err), http.StatusInternalServerError)
			return
		}

		// Call API with POST request
		apiURL := "https://your_api_here.com/somepage"
		err = postCallToAPI(apiURL, base64Data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error calling API: %s", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "File uploaded and processed successfully")
	}
}

// processFileContents this function parses the file and converts them to base64 json encoded data
func processFileContents(file io.Reader) (string, error) {
	// Read JSON file
	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, file)
	if err != nil {
		return "", err
	}

	// Convert JSON to base64
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

func postCallToAPI(apiURL, base64Data string) error {
	// Prepare POST request with base64 data
	requestBody := []byte(fmt.Sprintf(`{"data": "%s"}`, base64Data))
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Read and print response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("API Response: %s\n", string(responseBody))
	return nil
}
