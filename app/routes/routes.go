// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) LoginPost(
		inputEmail string,
		inputPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "inputEmail", inputEmail)
	revel.Unbind(args, "inputPassword", inputPassword)
	return revel.MainRouter.Reverse("App.LoginPost", args).URL
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).URL
}

func (_ tApp) ForgotPassword(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ForgotPassword", args).URL
}

func (_ tApp) ForgotPasswordPost(
		inputEmail string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "inputEmail", inputEmail)
	return revel.MainRouter.Reverse("App.ForgotPasswordPost", args).URL
}

func (_ tApp) Dashboard(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Dashboard", args).URL
}

func (_ tApp) Registration(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Registration", args).URL
}

func (_ tApp) RegPost(
		inputEmail string,
		Password string,
		fname string,
		lname string,
		sQuestion string,
		answer string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "inputEmail", inputEmail)
	revel.Unbind(args, "Password", Password)
	revel.Unbind(args, "fname", fname)
	revel.Unbind(args, "lname", lname)
	revel.Unbind(args, "sQuestion", sQuestion)
	revel.Unbind(args, "answer", answer)
	return revel.MainRouter.Reverse("App.RegPost", args).URL
}

func (_ tApp) Booking(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Booking", args).URL
}

func (_ tApp) BookingList(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.BookingList", args).URL
}

func (_ tApp) ViewReserved(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ViewReserved", args).URL
}

func (_ tApp) BookingReserveList(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.BookingReserveList", args).URL
}

func (_ tApp) BookingReservePost(
		flightChoice string,
		passportValues string,
		nameValues string,
		cardNum string,
		formType string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "flightChoice", flightChoice)
	revel.Unbind(args, "passportValues", passportValues)
	revel.Unbind(args, "nameValues", nameValues)
	revel.Unbind(args, "cardNum", cardNum)
	revel.Unbind(args, "formType", formType)
	return revel.MainRouter.Reverse("App.BookingReservePost", args).URL
}

func (_ tApp) BookingPost(
		flightChoice string,
		seatNumber string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "flightChoice", flightChoice)
	revel.Unbind(args, "seatNumber", seatNumber)
	return revel.MainRouter.Reverse("App.BookingPost", args).URL
}

func (_ tApp) BookedList(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.BookedList", args).URL
}

func (_ tApp) TicketList(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.TicketList", args).URL
}

func (_ tApp) ChangePass(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ChangePass", args).URL
}

func (_ tApp) ChangePassPost(
		cPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "cPassword", cPassword)
	return revel.MainRouter.Reverse("App.ChangePassPost", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


