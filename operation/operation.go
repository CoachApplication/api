package operation

type Operation interface {

	Exec(api_property.Properties) api_result.Result

}
