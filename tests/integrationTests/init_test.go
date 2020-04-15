package integrationTests

import (
	"flag"
	"fmt"
	"github.com/amitdavidson234/go-uniq/pkg/utils"
	"github.com/amitdavidson234/go-uniq/tests/testUtils"
	"math/rand"
	"os"
	"testing"
)

var binaryName string

var update *bool

func TestMain(m *testing.M) {
	update = flag.Bool("update", false, "update golden files")
	binaryName = fmt.Sprintf("%v", rand.Int())

	statusCode := run(m) // Moving the logic a level down so the cleanFunc defers will be called before os.exit
	os.Exit(statusCode)
}

func run(m *testing.M) int {
	cleanFunc, err := testUtils.GetBinary(binaryName)
	if err != nil {
		utils.ExitWithError(err)
	}
	defer func() {
		err := cleanFunc()
		if err != nil {
			utils.ExitWithError(err)
		}
	}()

	return m.Run()
}
