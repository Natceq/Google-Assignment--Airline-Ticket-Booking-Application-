// GENERATED CODE - DO NOT EDIT
// This file is the run file for Revel.
// It registers all the controllers and provides details for the Revel server engine to
// properly inject parameters directly into the action endpoints.
package run

import (
	"reflect"
	"github.com/revel/revel"
	_ "airlines/app"
	controllers "airlines/app/controllers"
	tests "airlines/tests"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	_ "github.com/revel/revel"
	_ "github.com/revel/revel/cache"
	"github.com/revel/revel/testing"
)

var (
	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

// Register and run the application
func Run(port int) {
	Register()
	revel.Run(port)
}

// Register all the controllers
func Register() {
	revel.AppLog.Info("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					20: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "LoginPost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "inputEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "inputPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ForgotPassword",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					133: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "ForgotPasswordPost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "inputEmail", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Dashboard",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					190: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Registration",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					202: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "RegPost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "inputEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "Password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "fname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "lname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "sQuestion", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "answer", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Booking",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					236: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "BookingList",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ViewReserved",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					310: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "BookingReserveList",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "BookingReservePost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "flightChoice", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "passportValues", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "nameValues", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "cardNum", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "formType", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "BookingPost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "flightChoice", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "seatNumber", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "BookedList",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "TicketList",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ChangePass",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					747: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "ChangePassPost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "cPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}
}
