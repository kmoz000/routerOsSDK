// routerOs/routerOsDebugger/cmd.go
package RouterOsSDK
import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"github.com/spf13/cobra"
)
var (
    username   string =  "admin"
    password   string = ""
    address    string = "http://192.168.1.1/rest"
)
func example(){
	c, err := N(address, WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}))
	if err != nil {
		fmt.Printf("Error Debugging RouterOS device at %s\nError> \t %v", address, err.Error())
		return
	}
	if res.StatusCode() != 200 {
		fmt.Printf("Error Debugging RouterOS device at %s\nStatusCode> \t", address, res.StatusCode())
		return
	}
	fmt.Printf("results: %v", res)
}