package inner

import (
    "fmt"
    "log"
    "os/exec"
    "bytes"

    // "github.com/loungeup/go-loungeup/resutil"
)

// package inner

// import (
// 	"time"

// 	"github.com/jirenius/go-res/resprot"
// 	"github.com/nats-io/nats.go"
// 	"gitlab.com/loungeup/go-loungeup/resutil"

// 	wmill "github.com/windmill-labs/windmill-go-client"
// )

func main() (interface{}, error) {
    // Define the Git command and its arguments
    cmd := exec.Command("git", "status", "--short")

    // Set the working directory (optional)
    // cmd.Dir = "/path/to/your/git/repo"

    // Capture the output and error streams
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    // Run the command
    err := cmd.Run()
    if err != nil {
        // Handle errors, such as command not found or non-zero exit status
        log.Fatalf("Command failed: %v\nError output: %s", err, stderr.String())
    }

    // Print the output
    fmt.Println("Git status output:")
    fmt.Println(stdout.String())

    return nil, nil
}
