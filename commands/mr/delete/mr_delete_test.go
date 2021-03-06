package delete

import (
	"github.com/profclems/glab/commands/cmdtest"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

// TODO: test by mocking the appropriate api function
func TestMain(m *testing.M) {
	cmdtest.InitTest(m, "mr_delete_test")
}

func Test_deleteMergeRequest(t *testing.T) {
	t.Parallel()
	repo := cmdtest.CopyTestRepo(t, "mr_delete_test")
	var cmd *exec.Cmd
	tests := []struct {
		name       string
		args       []string
		wantErr    bool
		assertFunc func(t *testing.T, out string)
	}{
		{
			name:    "delete",
			args:    []string{"0"},
			wantErr: true,
			assertFunc: func(t *testing.T, out string) {
				assert.Contains(t, out, "404 Not Found")
			},
		},
		{
			name:    "delete no args",
			wantErr: true,
			assertFunc: func(t *testing.T, out string) {
				assert.Contains(t, out, "accepts 1 arg(s), received 0")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd = exec.Command(cmdtest.GlabBinaryPath, tt.args...)
			cmd.Dir = repo
			b, err := cmd.CombinedOutput()
			if err != nil && !tt.wantErr {
				t.Log(string(b))
				t.Error(err)
			}
			out := string(b)
			t.Log(out)
		})
	}
}
