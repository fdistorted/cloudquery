// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	{{range .MockImports}}
  "{{.}}"
  {{end}}
	"google.golang.org/api/option"
)

func create{{.SubService | ToCamel}}() (*client.Services, error) {
	var item {{.Service}}.{{.StructName}}
	if err := faker.FakeData(&item); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/projects/testProject/global/{{.SubService}}", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &{{.Service}}.{{.MockListStruct}}List{
			Items: []*{{.Service}}.{{.StructName}}{&item},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	svc, err := {{.Service}}.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		{{.Service|ToCamel}}: svc,
	}, nil
}

func Test{{.SubService | ToCamel}}(t *testing.T) {
	client.MockTestHelper(t,  {{.SubService | ToCamel}}(), create{{.SubService | ToCamel}}, client.TestOptions{})
}
