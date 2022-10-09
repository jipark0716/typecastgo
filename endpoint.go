package typecast

var GoogleapiKey = "AIzaSyBJN3ZYdzTmjyQJ-9TdpikbsZDT9JUAYFk"
var GoogleApiVersion = "3"
var TypecastGid = "f8ON1ZpeF0mDFjZTQlr9G"

var (
	EndpointGoogleIdentitytoolkit = "https://www.googleapis.com/identitytoolkit/v" + GoogleApiVersion + "/relyingparty/"

	EndpointGoogleSecurePassword = EndpointGoogleIdentitytoolkit + "verifyPassword?key=" + GoogleapiKey
	EndpointGoogleSecureToken    = EndpointGoogleIdentitytoolkit + "verifyCustomToken?key=" + GoogleapiKey

	EndpointTypecastApi        = "https://typecast.ai/api/"
	EndpointTypecastNoredirect = func(url string) string {
		return url + "/no-redirect"
	}

	EndpointTypecastOauth     = EndpointTypecastApi + "auth-fb/custom-token"
	EndpointTypecastSpeakPost = EndpointTypecastApi + "speak/batch/post"
	EndpointTypecastSpeakGet  = EndpointTypecastApi + "speak/batch/get"
)
