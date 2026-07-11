package chainofresponsibility

import "fmt"

// AuthorizationHandler handles authorization requests
type AuthorizationHandler struct {
	BaseHandler
}

// Handle processes authorization requests
func (h *AuthorizationHandler) Handle(request *Request) {
	if request.Type == "authz" {
		fmt.Printf("AuthorizationHandler: Processing authorization request for data %s\n", request.Data)
		if request.Data == "admin_resource" {
			fmt.Println("AuthorizationHandler: Authorization successful")
		} else {
			fmt.Println("AuthorizationHandler: Authorization failed")
		}
	} else {
		h.HandleNext(request)
	}
}
