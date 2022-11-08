package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
	"traceWithGoV1/repository"
)

var repo *repository.Repository

func HandleGetPerson(writer http.ResponseWriter, request *http.Request) {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	spanCtx, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header),
	)
	span := opentracing.GlobalTracer().StartSpan(
		"/getPerson",
		opentracing.ChildOf(spanCtx),
	)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(request.Context(), span)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	email := request.FormValue("email")
	person, err := repo.GetDataByEmail(ctx, email)
	if err != nil {
		span.SetTag("error", true)
		span.LogFields(otlog.Error(err))
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	span.LogKV(
		"email", email,
		"title", person.Name,
		"activation_code", person.ActivationCode,
	)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	parseFiles, _ := template.ParseFiles("http/templates/get_person.html")
	parseFiles.Execute(writer, "")

	if email != "" {
		_, _ = fmt.Fprintf(writer,
			"<br><br><center><font color=\"green\" size=\"6\"><b>Email from field is: <br>"+email+"</font></center>")
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
}
