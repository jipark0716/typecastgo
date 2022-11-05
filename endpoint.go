package typecastgo

var GoogleapiKey = "AIzaSyBJN3ZYdzTmjyQJ-9TdpikbsZDT9JUAYFk"
var GoogleApiVersion = "1"
var TypecastGid = "f8ON1ZpeF0mDFjZTQlr9G"

var (
	EndpointGoogleIdentitytoolkit = "https://identitytoolkit.googleapis.com/v" + GoogleApiVersion + "/"
	EndpointGoogleSecure          = "https://securetoken.googleapis.com/v" + GoogleApiVersion + "/"

	EndpointGoogleSecurePassword = EndpointGoogleIdentitytoolkit + "accounts:signInWithPassword?key=" + GoogleapiKey
	EndpointGoogleSecureToken    = EndpointGoogleIdentitytoolkit + "accounts:signInWithCustomToken?key=" + GoogleapiKey
	EndpointRefreshToken         = EndpointGoogleSecure + "token?key=" + GoogleapiKey

	EndpointTypecastApi        = "https://typecast.ai/api/"
	EndpointTypecastNoredirect = func(url string) string {
		return url + "/no-redirect"
	}

	EndpointTypecastOauth     = EndpointTypecastApi + "auth-fb/custom-token"
	EndpointTypecastSpeakPost = EndpointTypecastApi + "speak/batch/post"
	EndpointTypecastSpeakGet  = EndpointTypecastApi + "speak/batch/get"

	EndpointTypecast = EndpointTypecastApi + "actor"
)
